package stage

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/aokoli/goutils"
	"github.com/rancher/rio/cli/cmd/edit/edit"
	"github.com/rancher/rio/cli/cmd/util"
	"github.com/rancher/rio/cli/pkg/clicontext"
	"github.com/rancher/rio/cli/pkg/types"
	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/rio/pkg/riofile/stringers"
	"github.com/rancher/rio/pkg/services"
	"github.com/rancher/wrangler/pkg/merr"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

type Stage struct {
	Image   string   `desc:"Runtime image (Docker image/OCI image)"`
	Edit    bool     `desc:"Edit the config to change the spec in new revision"`
	E_Env   []string `desc:"Set environment variables"`
	EnvFile []string `desc:"Read in a file of environment variables"`
}

func (r *Stage) Run(ctx *clicontext.CLIContext) error {
	if len(ctx.CLI.Args()) == 0 {
		return fmt.Errorf("must specify the service to update")
	}

	if len(ctx.CLI.Args()) > 2 {
		return fmt.Errorf("more than two arguments found")
	}

	serviceName := ctx.CLI.Args()[0]
	service, err := ctx.ByID(serviceName)
	if err != nil {
		return err
	}
	version := ""
	if len(ctx.CLI.Args()) == 2 {
		version = ctx.CLI.Args()[1]
	}
	if version == "" {
		var err error
		version, err = goutils.RandomNumeric(5)
		if err != nil {
			return fmt.Errorf("failed to generate random version, err: %v", err)
		}
		version = "v" + version
	}
	err = r.updatePromotedService(ctx, service, version)
	if err != nil {
		return err
	}
	err = r.updatePreviousServiceWeights(ctx, service, version) // this must come after stage call, if it fails old revisions shouldn't change
	if err != nil {
		return err
	}
	return nil
}

func (r *Stage) updatePromotedService(ctx *clicontext.CLIContext, service types.Resource, version string) error {
	if r.Edit {
		byteContent, err := json.Marshal(service.Object)
		if err != nil {
			return err
		}
		yamlBytes, err := yaml.JSONToYAML(byteContent)
		if err != nil {
			return err
		}
		yamlBytes = append(yamlBytes, []byte("/n")...)

		_, err = edit.Loop(nil, yamlBytes, func(content []byte) error {
			var obj *riov1.Service
			content = bytes.TrimSuffix(bytes.TrimSpace(content), []byte("/n"))
			if err := yaml.Unmarshal(content, &obj); err != nil {
				return err
			}
			svc := riov1.NewService(service.Namespace, service.Name, riov1.Service{
				Spec: obj.Spec,
			})
			app, _ := services.AppAndVersion(svc)
			svc.Name = app + "-" + version
			svc.Spec.Version = version
			if obj.Spec.App != "" {
				svc.Spec.App = obj.Spec.App
			} else {
				svc.Spec.App = app
			}
			svc.Spec.Weight = &[]int{0}[0]
			return ctx.Create(svc)
		})
		if err != nil {
			return err
		}
	} else {
		svc := service.Object.(*riov1.Service)
		app, _ := services.AppAndVersion(svc)
		spec := svc.Spec.DeepCopy()
		spec.Version = version
		spec.App = app
		spec.Weight = &[]int{0}[0]
		if ctx.CLI.String("image") != "" {
			spec.Image = ctx.CLI.String("image")
		}
		var err error
		spec.Env, err = r.mergeEnvVars(spec.Env)
		if err != nil {
			return err
		}
		stagedService := riov1.NewService(svc.Namespace, spec.App+"-"+version, riov1.Service{
			Spec: *spec,
			ObjectMeta: v1.ObjectMeta{
				Labels:      svc.Labels,
				Annotations: svc.Annotations,
			},
		})
		return ctx.Create(stagedService)
	}
	return nil
}

func (r *Stage) updatePreviousServiceWeights(ctx *clicontext.CLIContext, service types.Resource, version string) error {
	var allErrors []error
	svcs, err := util.ListAppServicesFromServiceName(ctx, service.Name)
	if err != nil {
		return err
	}
	var anyWeightSet bool
	for _, s := range svcs {
		if s.Spec.Version != version && s.Spec.Weight != nil && *s.Spec.Weight > 0 {
			anyWeightSet = true
			break
		}
	}
	if !anyWeightSet {
		for _, s := range svcs {
			if s.Spec.Version != version {
				err := ctx.UpdateResource(types.Resource{
					Namespace: s.Namespace,
					Name:      s.Name,
					App:       s.Spec.App,
					Version:   s.Spec.Version,
					Type:      types.ServiceType,
				}, func(obj runtime.Object) error {
					s := obj.(*riov1.Service)
					if s.Spec.Weight == nil {
						s.Spec.Weight = new(int)
					}
					*s.Spec.Weight = 100
					return nil
				})
				allErrors = append(allErrors, err)
			}
		}
		return merr.NewErrors(allErrors...)
	}
	return nil
}

// This keeps original and stage env vars in order and adds staged last, deletes any dups from original
func (r *Stage) mergeEnvVars(currEnvs []riov1.EnvVar) ([]riov1.EnvVar, error) {
	stageEnvs, err := stringers.ParseAllEnv(r.EnvFile, r.E_Env, true)
	if err != nil {
		return stageEnvs, err
	}
	if len(stageEnvs) == 0 {
		return currEnvs, nil
	}
	envMap := make(map[string]bool)
	for _, se := range stageEnvs {
		envMap[se.Name] = true
	}
	var orig []riov1.EnvVar
	for _, e := range currEnvs {
		if ok := envMap[e.Name]; !ok {
			orig = append(orig, e)
		}
	}
	return append(orig, stageEnvs...), nil
}
