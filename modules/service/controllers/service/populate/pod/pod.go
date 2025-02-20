package pod

import (
	"github.com/rancher/rio/modules/service/controllers/service/populate/rbac"
	"github.com/rancher/rio/modules/service/controllers/service/populate/servicelabels"
	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	"github.com/rancher/rio/pkg/constants"
	"github.com/rancher/rio/pkg/constructors"
	"github.com/rancher/wrangler/pkg/objectset"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Populate(service *riov1.Service, os *objectset.ObjectSet) (v1.PodTemplateSpec, error) {
	pts := v1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels:      servicelabels.ServiceLabels(service),
			Annotations: servicelabels.ServiceAnnotations(service),
		},
	}

	podSpec := podSpec(service)
	Roles(service, os)

	PersistVolume(service, os)

	pts.Spec = podSpec
	return pts, nil
}

func Roles(service *riov1.Service, os *objectset.ObjectSet) {
	if err := rbac.Populate(service, os); err != nil {
		os.AddErr(err)
		return
	}
}

func PersistVolume(service *riov1.Service, os *objectset.ObjectSet) {
	var volumes []riov1.Volume
	for _, volume := range service.Spec.Volumes {
		volumes = append(volumes, volume)
	}

	for _, c := range service.Spec.Sidecars {
		for _, volume := range c.Volumes {
			volumes = append(volumes, volume)
		}
	}

	for _, v := range volumes {
		if v.Persistent && constants.DefaultStorageClass {
			pv := constructors.NewPersistentVolumeClaim(service.Namespace, v.Name, v1.PersistentVolumeClaim{
				Spec: v1.PersistentVolumeClaimSpec{
					AccessModes: []v1.PersistentVolumeAccessMode{
						v1.ReadWriteOnce,
					},
					Resources: v1.ResourceRequirements{
						Requests: v1.ResourceList{
							v1.ResourceStorage: resource.MustParse(constants.RegistryStorageSize),
						},
					},
				},
			})
			os.Add(pv)
		}
	}
}
