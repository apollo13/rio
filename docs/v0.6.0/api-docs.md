<p>Packages:</p>
<ul>
<li>
<a href="#admin.rio.cattle.io%2fv1">admin.rio.cattle.io/v1</a>
</li>
<li>
<a href="#management.cattle.io%2fv3">management.cattle.io/v3</a>
</li>
<li>
<a href="#rio.cattle.io%2fv1">rio.cattle.io/v1</a>
</li>
</ul>
<h2 id="admin.rio.cattle.io/v1">admin.rio.cattle.io/v1</h2>
<p>
</p>
Resource Types:
<ul><li>
<a href="#admin.rio.cattle.io/v1.ClusterDomain">ClusterDomain</a>
</li><li>
<a href="#admin.rio.cattle.io/v1.PublicDomain">PublicDomain</a>
</li><li>
<a href="#admin.rio.cattle.io/v1.RioInfo">RioInfo</a>
</li><li>
<a href="#admin.rio.cattle.io/v1.SystemStack">SystemStack</a>
</li></ul>
<h3 id="ClusterDomain">ClusterDomain
</h3>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
admin.rio.cattle.io/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>ClusterDomain</code></td>
</tr>
<tr>
<td>
<code>metadata</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code></br>
<em>
<a href="#admin.rio.cattle.io/v1.ClusterDomainSpec">
ClusterDomainSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>secretName</code></br>
<em>
string
</em>
</td>
<td>
<p>SecretName holding the TLS certificate for this domain.  This is expected
to be a wildcard certificate</p>
</td>
</tr>
<tr>
<td>
<code>httpsPort</code></br>
<em>
int
</em>
</td>
<td>
<p>The public HTTPS port for the cluster domain</p>
</td>
</tr>
<tr>
<td>
<code>httpPort</code></br>
<em>
int
</em>
</td>
<td>
<p>The public HTTP port for the cluster domain</p>
</td>
</tr>
<tr>
<td>
<code>addresses</code></br>
<em>
<a href="#admin.rio.cattle.io/v1.Address">
[]Address
</a>
</em>
</td>
<td>
<p>The addresses assigned to the ClusterDomain by the provider</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code></br>
<em>
<a href="#admin.rio.cattle.io/v1.ClusterDomainStatus">
ClusterDomainStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="PublicDomain">PublicDomain
</h3>
<p>
<p>PublicDomain is a top-level resource to allow user to its own public domain for the services inside cluster. It can be pointed to
Router or Service. It is user&rsquo;s responsibility to setup a CNAME or A record to the clusterDomain or ingress IP.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
admin.rio.cattle.io/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>PublicDomain</code></td>
</tr>
<tr>
<td>
<code>metadata</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code></br>
<em>
<a href="#admin.rio.cattle.io/v1.PublicDomainSpec">
PublicDomainSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>secretName</code></br>
<em>
string
</em>
</td>
<td>
<p>SecretName holding the TLS certificate for this domain.</p>
</td>
</tr>
<tr>
<td>
<code>targetApp</code></br>
<em>
string
</em>
</td>
<td>
<p>Target App Name.  Only TargetAppName or TargetRouter can be set</p>
</td>
</tr>
<tr>
<td>
<code>targetRouter</code></br>
<em>
string
</em>
</td>
<td>
<p>Target Router Name.  Only TargetAppName or TargetRouter can be set</p>
</td>
</tr>
<tr>
<td>
<code>targetVersion</code></br>
<em>
string
</em>
</td>
<td>
<p>Target Version</p>
</td>
</tr>
<tr>
<td>
<code>targetNamespace</code></br>
<em>
string
</em>
</td>
<td>
<p>Target Service or Router Namespace</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code></br>
<em>
<a href="#admin.rio.cattle.io/v1.PublicDomainStatus">
PublicDomainStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="RioInfo">RioInfo
</h3>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
admin.rio.cattle.io/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>RioInfo</code></td>
</tr>
<tr>
<td>
<code>metadata</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>status</code></br>
<em>
<a href="#admin.rio.cattle.io/v1.RioInfoStatus">
RioInfoStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="SystemStack">SystemStack
</h3>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
admin.rio.cattle.io/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>SystemStack</code></td>
</tr>
<tr>
<td>
<code>metadata</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
</tbody>
</table>
<h3 id="Address">Address
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2fadmin.rio.cattle.io%2fv1.ClusterDomainSpec">ClusterDomainSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ip</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>hostname</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="ClusterDomainSpec">ClusterDomainSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2fadmin.rio.cattle.io%2fv1.ClusterDomain">ClusterDomain</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>secretName</code></br>
<em>
string
</em>
</td>
<td>
<p>SecretName holding the TLS certificate for this domain.  This is expected
to be a wildcard certificate</p>
</td>
</tr>
<tr>
<td>
<code>httpsPort</code></br>
<em>
int
</em>
</td>
<td>
<p>The public HTTPS port for the cluster domain</p>
</td>
</tr>
<tr>
<td>
<code>httpPort</code></br>
<em>
int
</em>
</td>
<td>
<p>The public HTTP port for the cluster domain</p>
</td>
</tr>
<tr>
<td>
<code>addresses</code></br>
<em>
<a href="#admin.rio.cattle.io/v1.Address">
[]Address
</a>
</em>
</td>
<td>
<p>The addresses assigned to the ClusterDomain by the provider</p>
</td>
</tr>
</tbody>
</table>
<h3 id="ClusterDomainStatus">ClusterDomainStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2fadmin.rio.cattle.io%2fv1.ClusterDomain">ClusterDomain</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>assignedSecretName</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>httpsSupported</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>conditions</code></br>
<em>
[]github.com/rancher/wrangler/pkg/genericcondition.GenericCondition
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="PublicDomainSpec">PublicDomainSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2fadmin.rio.cattle.io%2fv1.PublicDomain">PublicDomain</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>secretName</code></br>
<em>
string
</em>
</td>
<td>
<p>SecretName holding the TLS certificate for this domain.</p>
</td>
</tr>
<tr>
<td>
<code>targetApp</code></br>
<em>
string
</em>
</td>
<td>
<p>Target App Name.  Only TargetAppName or TargetRouter can be set</p>
</td>
</tr>
<tr>
<td>
<code>targetRouter</code></br>
<em>
string
</em>
</td>
<td>
<p>Target Router Name.  Only TargetAppName or TargetRouter can be set</p>
</td>
</tr>
<tr>
<td>
<code>targetVersion</code></br>
<em>
string
</em>
</td>
<td>
<p>Target Version</p>
</td>
</tr>
<tr>
<td>
<code>targetNamespace</code></br>
<em>
string
</em>
</td>
<td>
<p>Target Service or Router Namespace</p>
</td>
</tr>
</tbody>
</table>
<h3 id="PublicDomainStatus">PublicDomainStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2fadmin.rio.cattle.io%2fv1.PublicDomain">PublicDomain</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>httpsSupported</code></br>
<em>
bool
</em>
</td>
<td>
<p>Whether HTTP is supported in the Domain</p>
</td>
</tr>
<tr>
<td>
<code>assignedSecretName</code></br>
<em>
string
</em>
</td>
<td>
<p>Secret containing TLS cert for HTTPS</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code></br>
<em>
[]github.com/rancher/wrangler/pkg/genericcondition.GenericCondition
</em>
</td>
<td>
<p>Represents the latest available observations of a PublicDomain&rsquo;s current state.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="RioInfoStatus">RioInfoStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2fadmin.rio.cattle.io%2fv1.RioInfo">RioInfo</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>version</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>gitCommit</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>systemNamespace</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>conditions</code></br>
<em>
[]github.com/rancher/wrangler/pkg/genericcondition.GenericCondition
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ready</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>systemComponentReadyMap</code></br>
<em>
map[string]string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="management.cattle.io/v3">management.cattle.io/v3</h2>
<p>
</p>
Resource Types:
<ul><li>
<a href="#management.cattle.io/v3.Setting">Setting</a>
</li><li>
<a href="#management.cattle.io/v3.User">User</a>
</li></ul>
<h3 id="Setting">Setting
</h3>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
management.cattle.io/v3
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>Setting</code></td>
</tr>
<tr>
<td>
<code>metadata</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>value</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>default</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>customized</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>source</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="User">User
</h3>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
management.cattle.io/v3
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>User</code></td>
</tr>
<tr>
<td>
<code>metadata</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>displayName</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>description</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>username</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>password</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>mustChangePassword</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>principalIds</code></br>
<em>
[]string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>me</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>enabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>spec</code></br>
<em>
<a href="#management.cattle.io/v3.UserSpec">
UserSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code></br>
<em>
<a href="#management.cattle.io/v3.UserStatus">
UserStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="UserSpec">UserSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2fmanagement.cattle.io%2fv3.User">User</a>)
</p>
<p>
</p>
<h3 id="UserStatus">UserStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2fmanagement.cattle.io%2fv3.User">User</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>conditions</code></br>
<em>
[]github.com/rancher/wrangler/pkg/genericcondition.GenericCondition
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="rio.cattle.io/v1">rio.cattle.io/v1</h2>
<p>
</p>
Resource Types:
<ul><li>
<a href="#rio.cattle.io/v1.ExternalService">ExternalService</a>
</li><li>
<a href="#rio.cattle.io/v1.Router">Router</a>
</li><li>
<a href="#rio.cattle.io/v1.Service">Service</a>
</li><li>
<a href="#rio.cattle.io/v1.Stack">Stack</a>
</li></ul>
<h3 id="ExternalService">ExternalService
</h3>
<p>
<p>ExternalService creates a DNS record and route rules for any Service outside of the cluster, can be IPs or FQDN outside the mesh</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
rio.cattle.io/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>ExternalService</code></td>
</tr>
<tr>
<td>
<code>metadata</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code></br>
<em>
<a href="#rio.cattle.io/v1.ExternalServiceSpec">
ExternalServiceSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>ipAddresses</code></br>
<em>
[]string
</em>
</td>
<td>
<p>External service located outside mesh, represented by IPs</p>
</td>
</tr>
<tr>
<td>
<code>fqdn</code></br>
<em>
string
</em>
</td>
<td>
<p>External service located outside mesh, represented by DNS</p>
</td>
</tr>
<tr>
<td>
<code>targetApp</code></br>
<em>
string
</em>
</td>
<td>
<p>In-Mesh app in another namespace</p>
</td>
</tr>
<tr>
<td>
<code>targetVersion</code></br>
<em>
string
</em>
</td>
<td>
<p>In-Mesh version in another namespace</p>
</td>
</tr>
<tr>
<td>
<code>targetRouter</code></br>
<em>
string
</em>
</td>
<td>
<p>In-Mesh router in another namespace</p>
</td>
</tr>
<tr>
<td>
<code>targetServiceNamespace</code></br>
<em>
string
</em>
</td>
<td>
<p>Namespace of in-mesh service in another namespace</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code></br>
<em>
<a href="#rio.cattle.io/v1.ExternalServiceStatus">
ExternalServiceStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="Router">Router
</h3>
<p>
<p>Router is a top level resource to create L7 routing to different services. It will create VirtualService, ServiceEntry and DestinationRules</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
rio.cattle.io/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>Router</code></td>
</tr>
<tr>
<td>
<code>metadata</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code></br>
<em>
<a href="#rio.cattle.io/v1.RouterSpec">
RouterSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>routes</code></br>
<em>
<a href="#rio.cattle.io/v1.RouteSpec">
[]RouteSpec
</a>
</em>
</td>
<td>
<p>An ordered list of route rules for HTTP traffic. The first rule matching an incoming request is used.</p>
</td>
</tr>
<tr>
<td>
<code>internal</code></br>
<em>
bool
</em>
</td>
<td>
<p>By default all Routers are public and exposed outside of the cluster. Setting internal to true will
cause the Router to not be exposed</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code></br>
<em>
<a href="#rio.cattle.io/v1.RouterStatus">
RouterStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="Service">Service
</h3>
<p>
<p>Service acts as a top level resource for a container and its sidecars and routing resources.
Each service represents an individual revision, group by Spec.App(defaults to Service.Name), and Spec.Version(defaults to v0)</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
rio.cattle.io/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>Service</code></td>
</tr>
<tr>
<td>
<code>metadata</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code></br>
<em>
<a href="#rio.cattle.io/v1.ServiceSpec">
ServiceSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>PodConfig</code></br>
<em>
<a href="#rio.cattle.io/v1.PodConfig">
PodConfig
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>template</code></br>
<em>
bool
</em>
</td>
<td>
<p>Template this service is a template for new versions to be created base on changes
from the build.repo</p>
</td>
</tr>
<tr>
<td>
<code>stageOnly</code></br>
<em>
bool
</em>
</td>
<td>
<p>StageOnly whether to only stage services that are generated through template from build.repo</p>
</td>
</tr>
<tr>
<td>
<code>version</code></br>
<em>
string
</em>
</td>
<td>
<p>Version version of this service</p>
</td>
</tr>
<tr>
<td>
<code>app</code></br>
<em>
string
</em>
</td>
<td>
<p>App The exposed app name, if no value is set, then metadata.name of the Service is used</p>
</td>
</tr>
<tr>
<td>
<code>weight</code></br>
<em>
int
</em>
</td>
<td>
<p>Weight The weight among services with matching app field to determine how much traffic is load balanced
to this service.  If rollout is set, the weight become the target weight of the rollout.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code></br>
<em>
int
</em>
</td>
<td>
<p>Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1 in deployment.</p>
</td>
</tr>
<tr>
<td>
<code>maxUnavailable</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#intorstring-intstr-util">
k8s.io/apimachinery/pkg/util/intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The maximum number of pods that can be unavailable during the update.
Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).
Absolute number is calculated from percentage by rounding down.
This can not be 0 if MaxSurge is 0.
Defaults to 25%.
Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods
immediately when the rolling update starts. Once new pods are ready, old ReplicaSet
can be scaled down further, followed by scaling up the new ReplicaSet, ensuring
that the total number of pods available at all times during the update is at
least 70% of desired pods.</p>
</td>
</tr>
<tr>
<td>
<code>maxSurge</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#intorstring-intstr-util">
k8s.io/apimachinery/pkg/util/intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The maximum number of pods that can be scheduled above the desired number of
pods.
Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).
This can not be 0 if MaxUnavailable is 0.
Absolute number is calculated from percentage by rounding up.
Defaults to 25%.
Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when
the rolling update starts, such that the total number of old and new pods do not exceed
130% of desired pods. Once old pods have been killed,
new ReplicaSet can be scaled up further, ensuring that total number of pods running
at any time during the update is at most 130% of desired pods.</p>
</td>
</tr>
<tr>
<td>
<code>autoscale</code></br>
<em>
<a href="#rio.cattle.io/v1.AutoscaleConfig">
AutoscaleConfig
</a>
</em>
</td>
<td>
<p>Autoscale the replicas based on the amount of traffic received by this service</p>
</td>
</tr>
<tr>
<td>
<code>rollout</code></br>
<em>
<a href="#rio.cattle.io/v1.RolloutConfig">
RolloutConfig
</a>
</em>
</td>
<td>
<p>RolloutConfig controls how each service is allocated ComputedWeight</p>
</td>
</tr>
<tr>
<td>
<code>global</code></br>
<em>
bool
</em>
</td>
<td>
<p>Place one pod per node that matches the scheduling rules</p>
</td>
</tr>
<tr>
<td>
<code>serviceMesh</code></br>
<em>
bool
</em>
</td>
<td>
<p>Whether to disable Service mesh for Service. If true, no mesh sidecar will be deployed along with the Service</p>
</td>
</tr>
<tr>
<td>
<code>requestTimeoutSeconds</code></br>
<em>
int
</em>
</td>
<td>
<p>RequestTimeoutSeconds specify the timeout set on api gateway for each individual service</p>
</td>
</tr>
<tr>
<td>
<code>permissions</code></br>
<em>
<a href="#rio.cattle.io/v1.Permission">
[]Permission
</a>
</em>
</td>
<td>
<p>Permissions to the Services. It will create corresponding ServiceAccounts, Roles and RoleBinding.</p>
</td>
</tr>
<tr>
<td>
<code>globalPermissions</code></br>
<em>
<a href="#rio.cattle.io/v1.Permission">
[]Permission
</a>
</em>
</td>
<td>
<p>GlobalPermissions to the Services. It will create corresponding ServiceAccounts, ClusterRoles and ClusterRoleBinding.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code></br>
<em>
<a href="#rio.cattle.io/v1.ServiceStatus">
ServiceStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="Stack">Stack
</h3>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
rio.cattle.io/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>Stack</code></td>
</tr>
<tr>
<td>
<code>metadata</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code></br>
<em>
<a href="#rio.cattle.io/v1.StackSpec">
StackSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>build</code></br>
<em>
<a href="#rio.cattle.io/v1.StackBuild">
StackBuild
</a>
</em>
</td>
<td>
<p>Stack build parameters that watches git repo</p>
</td>
</tr>
<tr>
<td>
<code>permissions</code></br>
<em>
<a href="#rio.cattle.io/v1.Permission">
[]Permission
</a>
</em>
</td>
<td>
<p>Permissions used while deploying objects created by this stack</p>
</td>
</tr>
<tr>
<td>
<code>additionalGroupVersionKinds</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#groupversionkind-schema-runtime">
[]k8s.io/apimachinery/pkg/runtime/schema.GroupVersionKind
</a>
</em>
</td>
<td>
<p>Additional GVKs not in the rio.cattle.io that have the rio.cattle.io/stack label. These objects
are &ldquo;owned&rdquo; by this stack</p>
</td>
</tr>
<tr>
<td>
<code>answers</code></br>
<em>
map[string]string
</em>
</td>
<td>
<p>Stack answers</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code></br>
<em>
<a href="#rio.cattle.io/v1.StackStatus">
StackStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="AutoscaleConfig">AutoscaleConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.ServiceSpec">ServiceSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>concurrency</code></br>
<em>
int
</em>
</td>
<td>
<p>ContainerConcurrency specifies the maximum allowed in-flight (concurrent) requests per container of the Revision. Defaults to 0 which means unlimited concurrency.</p>
</td>
</tr>
<tr>
<td>
<code>minReplicas</code></br>
<em>
int32
</em>
</td>
<td>
<p>The minimal number of replicas Service can be scaled</p>
</td>
</tr>
<tr>
<td>
<code>maxReplicas</code></br>
<em>
int32
</em>
</td>
<td>
<p>The maximum number of replicas Service can be scaled</p>
</td>
</tr>
</tbody>
</table>
<h3 id="BuildRevision">BuildRevision
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.ServiceStatus">ServiceStatus</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>commits</code></br>
<em>
[]string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="Container">Container
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.NamedContainer">NamedContainer</a>, 
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.PodConfig">PodConfig</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>image</code></br>
<em>
string
</em>
</td>
<td>
<p>Docker image name. More info: <a href="https://kubernetes.io/docs/concepts/containers/images">https://kubernetes.io/docs/concepts/containers/images</a> This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.</p>
</td>
</tr>
<tr>
<td>
<code>build</code></br>
<em>
<a href="#rio.cattle.io/v1.ImageBuildSpec">
ImageBuildSpec
</a>
</em>
</td>
<td>
<p>ImageBuild specifies how to build this image</p>
</td>
</tr>
<tr>
<td>
<code>command</code></br>
<em>
[]string
</em>
</td>
<td>
<p>Entrypoint array. Not executed within a shell. The docker image&rsquo;s ENTRYPOINT is used if this is not provided.
Variable references $(VAR_NAME) are expanded using the container&rsquo;s environment. If a variable cannot be resolved, the reference in the input string will be unchanged.
The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not.
Cannot be updated. More info: <a href="https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell">https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell</a></p>
</td>
</tr>
<tr>
<td>
<code>args</code></br>
<em>
[]string
</em>
</td>
<td>
<p>Arguments to the entrypoint. The docker image&rsquo;s CMD is used if this is not provided.
Variable references $(VAR_NAME) are expanded using the container&rsquo;s environment.
If a variable cannot be resolved, the reference in the input string will be unchanged.
The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not.
Cannot be updated. More info: <a href="https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell">https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell</a></p>
</td>
</tr>
<tr>
<td>
<code>workingDir</code></br>
<em>
string
</em>
</td>
<td>
<p>Container&rsquo;s working directory. If not specified, the container runtime&rsquo;s default will be used, which might be configured in the container image. Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>ports</code></br>
<em>
<a href="#rio.cattle.io/v1.ContainerPort">
[]ContainerPort
</a>
</em>
</td>
<td>
<p>List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed.
Any port which is listening on the default &ldquo;0.0.0.0&rdquo; address inside a container will be accessible from the network. Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>env</code></br>
<em>
<a href="#rio.cattle.io/v1.EnvVar">
[]EnvVar
</a>
</em>
</td>
<td>
<p>List of environment variables to set in the container. Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>cpuMillis</code></br>
<em>
int64
</em>
</td>
<td>
<p>CPU, in cores</p>
</td>
</tr>
<tr>
<td>
<code>memoryBytes</code></br>
<em>
int64
</em>
</td>
<td>
<p>Memory, in bytes</p>
</td>
</tr>
<tr>
<td>
<code>secrets</code></br>
<em>
<a href="#rio.cattle.io/v1.DataMount">
[]DataMount
</a>
</em>
</td>
<td>
<p>Secrets Mounts</p>
</td>
</tr>
<tr>
<td>
<code>configs</code></br>
<em>
<a href="#rio.cattle.io/v1.DataMount">
[]DataMount
</a>
</em>
</td>
<td>
<p>Configmaps Mounts</p>
</td>
</tr>
<tr>
<td>
<code>livenessProbe</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#probe-v1-core">
Kubernetes core/v1.Probe
</a>
</em>
</td>
<td>
<p>Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: <a href="https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes">https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes</a></p>
</td>
</tr>
<tr>
<td>
<code>readinessProbe</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#probe-v1-core">
Kubernetes core/v1.Probe
</a>
</em>
</td>
<td>
<p>Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: <a href="https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes">https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes</a></p>
</td>
</tr>
<tr>
<td>
<code>imagePullPolicy</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#pullpolicy-v1-core">
Kubernetes core/v1.PullPolicy
</a>
</em>
</td>
<td>
<p>Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if tag is does not start with v[0-9] or [0-9], or IfNotPresent otherwise. Cannot be updated. More info: <a href="https://kubernetes.io/docs/concepts/containers/images#updating-images">https://kubernetes.io/docs/concepts/containers/images#updating-images</a></p>
</td>
</tr>
<tr>
<td>
<code>stdin</code></br>
<em>
bool
</em>
</td>
<td>
<p>Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.</p>
</td>
</tr>
<tr>
<td>
<code>stdinOnce</code></br>
<em>
bool
</em>
</td>
<td>
<p>Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions.
If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false</p>
</td>
</tr>
<tr>
<td>
<code>tty</code></br>
<em>
bool
</em>
</td>
<td>
<p>Whether this container should allocate a TTY for itself, also requires &lsquo;stdin&rsquo; to be true. Default is false.</p>
</td>
</tr>
<tr>
<td>
<code>volumes</code></br>
<em>
<a href="#rio.cattle.io/v1.Volume">
[]Volume
</a>
</em>
</td>
<td>
<p>Pod volumes to mount into the container&rsquo;s filesystem</p>
</td>
</tr>
<tr>
<td>
<code>ContainerSecurityContext</code></br>
<em>
<a href="#rio.cattle.io/v1.ContainerSecurityContext">
ContainerSecurityContext
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="ContainerPort">ContainerPort
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.Container">Container</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>expose</code></br>
<em>
bool
</em>
</td>
<td>
<p>Expose will make the port available outside the cluster. All http/https ports will be set to true by default
if Expose is nil.  All other protocols are set to false by default</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code></br>
<em>
<a href="#rio.cattle.io/v1.Protocol">
Protocol
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>port</code></br>
<em>
int32
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>targetPort</code></br>
<em>
int32
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>hostport</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="ContainerSecurityContext">ContainerSecurityContext
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.Container">Container</a>)
</p>
<p>
<p>ContainerSecurityContext holds pod-level security attributes and common container constants. Optional: Defaults to empty. See type description for default values of each field.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>runAsUser</code></br>
<em>
int64
</em>
</td>
<td>
<p>The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.
If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container</p>
</td>
</tr>
<tr>
<td>
<code>runAsGroup</code></br>
<em>
int64
</em>
</td>
<td>
<p>The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.
If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.</p>
</td>
</tr>
<tr>
<td>
<code>readOnlyRootFilesystem</code></br>
<em>
bool
</em>
</td>
<td>
<p>Whether this container has a read-only root filesystem. Default is false.</p>
</td>
</tr>
<tr>
<td>
<code>privileged</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Run container in privileged mode.
Processes in privileged containers are essentially equivalent to root on the host.
Defaults to false.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="DNS">DNS
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.PodConfig">PodConfig</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>policy</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#dnspolicy-v1-core">
Kubernetes core/v1.DNSPolicy
</a>
</em>
</td>
<td>
<p>Set DNS policy for the pod. Defaults to &ldquo;ClusterFirst&rdquo;. Valid values are &lsquo;ClusterFirstWithHostNet&rsquo;, &lsquo;ClusterFirst&rsquo;, &lsquo;Default&rsquo; or &lsquo;None&rsquo;.
DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy.
To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to &lsquo;ClusterFirstWithHostNet&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>nameservers</code></br>
<em>
[]string
</em>
</td>
<td>
<p>A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.</p>
</td>
</tr>
<tr>
<td>
<code>searches</code></br>
<em>
[]string
</em>
</td>
<td>
<p>A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.</p>
</td>
</tr>
<tr>
<td>
<code>options</code></br>
<em>
<a href="#rio.cattle.io/v1.PodDNSConfigOption">
[]PodDNSConfigOption
</a>
</em>
</td>
<td>
<p>A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy.
Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="DataMount">DataMount
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.Container">Container</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>target</code></br>
<em>
string
</em>
</td>
<td>
<p>The directory or file to mount the value to in the container</p>
</td>
</tr>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>The name of the ConfigMap or Secret to mount</p>
</td>
</tr>
<tr>
<td>
<code>key</code></br>
<em>
string
</em>
</td>
<td>
<p>The key in the data of the ConfigMap or Secret to mount to a file.
If Key is set the Target must be a file.  If key is set the target must be a directory and will
contain one file per key from the Secret/ConfigMap data field.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="Destination">Destination
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.RouteSpec">RouteSpec</a>, 
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.WeightedDestination">WeightedDestination</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>app</code></br>
<em>
string
</em>
</td>
<td>
<p>Destination Service</p>
</td>
</tr>
<tr>
<td>
<code>version</code></br>
<em>
string
</em>
</td>
<td>
<p>Destination Revision</p>
</td>
</tr>
<tr>
<td>
<code>port</code></br>
<em>
uint32
</em>
</td>
<td>
<p>Destination Port</p>
</td>
</tr>
</tbody>
</table>
<h3 id="EnvVar">EnvVar
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.Container">Container</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>value</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>secretName</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>configMapName</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>key</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="ExternalServiceSpec">ExternalServiceSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.ExternalService">ExternalService</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ipAddresses</code></br>
<em>
[]string
</em>
</td>
<td>
<p>External service located outside mesh, represented by IPs</p>
</td>
</tr>
<tr>
<td>
<code>fqdn</code></br>
<em>
string
</em>
</td>
<td>
<p>External service located outside mesh, represented by DNS</p>
</td>
</tr>
<tr>
<td>
<code>targetApp</code></br>
<em>
string
</em>
</td>
<td>
<p>In-Mesh app in another namespace</p>
</td>
</tr>
<tr>
<td>
<code>targetVersion</code></br>
<em>
string
</em>
</td>
<td>
<p>In-Mesh version in another namespace</p>
</td>
</tr>
<tr>
<td>
<code>targetRouter</code></br>
<em>
string
</em>
</td>
<td>
<p>In-Mesh router in another namespace</p>
</td>
</tr>
<tr>
<td>
<code>targetServiceNamespace</code></br>
<em>
string
</em>
</td>
<td>
<p>Namespace of in-mesh service in another namespace</p>
</td>
</tr>
</tbody>
</table>
<h3 id="ExternalServiceStatus">ExternalServiceStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.ExternalService">ExternalService</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>conditions</code></br>
<em>
[]github.com/rancher/wrangler/pkg/genericcondition.GenericCondition
</em>
</td>
<td>
<p>Represents the latest available observations of a ExternalService&rsquo;s current state.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="Fault">Fault
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.RouteSpec">RouteSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>percentage</code></br>
<em>
int
</em>
</td>
<td>
<p>Percentage of requests on which the delay will be injected.</p>
</td>
</tr>
<tr>
<td>
<code>delayMillis</code></br>
<em>
int
</em>
</td>
<td>
<p>REQUIRED. Add a fixed delay before forwarding the request. Units: milliseconds</p>
</td>
</tr>
<tr>
<td>
<code>abortHTTPStatus</code></br>
<em>
int
</em>
</td>
<td>
<p>Abort Http request attempts and return error codes back to downstream service, giving the impression that the upstream service is faulty.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="HeaderMatch">HeaderMatch
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.Match">Match</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>value</code></br>
<em>
<a href="#rio.cattle.io/v1.StringMatch">
StringMatch
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="HeaderOperations">HeaderOperations
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.RouteSpec">RouteSpec</a>)
</p>
<p>
<p>HeaderOperations Describes the header manipulations to apply</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>add</code></br>
<em>
<a href="#rio.cattle.io/v1.NameValue">
[]NameValue
</a>
</em>
</td>
<td>
<p>Append the given values to the headers specified by keys
(will create a comma-separated list of values)</p>
</td>
</tr>
<tr>
<td>
<code>set</code></br>
<em>
<a href="#rio.cattle.io/v1.NameValue">
[]NameValue
</a>
</em>
</td>
<td>
<p>Append the given values to the headers specified by keys
(will create a comma-separated list of values)</p>
</td>
</tr>
<tr>
<td>
<code>remove</code></br>
<em>
[]string
</em>
</td>
<td>
<p>Remove a the specified headers</p>
</td>
</tr>
</tbody>
</table>
<h3 id="ImageBuildSpec">ImageBuildSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.Container">Container</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>repo</code></br>
<em>
string
</em>
</td>
<td>
<p>Repository url</p>
</td>
</tr>
<tr>
<td>
<code>revision</code></br>
<em>
string
</em>
</td>
<td>
<p>Repo Revision. Can be a git commit or tag</p>
</td>
</tr>
<tr>
<td>
<code>branch</code></br>
<em>
string
</em>
</td>
<td>
<p>Repo Branch. If specified, a gitmodule will be created to watch the repo and creating new revision if new commit or tag is pushed.</p>
</td>
</tr>
<tr>
<td>
<code>dockerfile</code></br>
<em>
string
</em>
</td>
<td>
<p>Specify the name of the Dockerfile in the Repo. This is the full path relative to the repo root. Defaults to <code>Dockerfile</code>.</p>
</td>
</tr>
<tr>
<td>
<code>context</code></br>
<em>
string
</em>
</td>
<td>
<p>Specify build context. Defaults to &ldquo;.&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>args</code></br>
<em>
[]string
</em>
</td>
<td>
<p>Specify build args</p>
</td>
</tr>
<tr>
<td>
<code>template</code></br>
<em>
string
</em>
</td>
<td>
<p>Specify the build template. Defaults to <code>buildkit</code>.</p>
</td>
</tr>
<tr>
<td>
<code>webhookSecretName</code></br>
<em>
string
</em>
</td>
<td>
<p>Specify the github secret name. Used to create Github webhook, the secret key has to be <code>accessToken</code></p>
</td>
</tr>
<tr>
<td>
<code>cloneSecretName</code></br>
<em>
string
</em>
</td>
<td>
<p>Specify secret name for checking our git resources</p>
</td>
</tr>
<tr>
<td>
<code>pushRegistry</code></br>
<em>
string
</em>
</td>
<td>
<p>Specify custom registry to push the image instead of built-in one</p>
</td>
</tr>
<tr>
<td>
<code>pushRegistrySecretName</code></br>
<em>
string
</em>
</td>
<td>
<p>Specify secret for pushing to custom registry</p>
</td>
</tr>
<tr>
<td>
<code>imageName</code></br>
<em>
string
</em>
</td>
<td>
<p>Specify image name instead of the one generated from service name, format: $registry/$imageName:$revision</p>
</td>
</tr>
<tr>
<td>
<code>pr</code></br>
<em>
bool
</em>
</td>
<td>
<p>Whether to enable builds for pull requests</p>
</td>
</tr>
<tr>
<td>
<code>tag</code></br>
<em>
bool
</em>
</td>
<td>
<p>Whether to enable builds for tags</p>
</td>
</tr>
<tr>
<td>
<code>noCache</code></br>
<em>
bool
</em>
</td>
<td>
<p>Build image with no cache</p>
</td>
</tr>
<tr>
<td>
<code>timeout</code></br>
<em>
int
</em>
</td>
<td>
<p>TimeoutSeconds describes how long the build can run</p>
</td>
</tr>
</tbody>
</table>
<h3 id="Match">Match
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.RouteSpec">RouteSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>path</code></br>
<em>
<a href="#rio.cattle.io/v1.StringMatch">
StringMatch
</a>
</em>
</td>
<td>
<p>URI to match values are case-sensitive and formatted as follows:</p>
<p>exact: &ldquo;value&rdquo; for exact string match</p>
<p>prefix: &ldquo;value&rdquo; for prefix-based match</p>
<p>regex: &ldquo;value&rdquo; for ECMAscript style regex-based match</p>
</td>
</tr>
<tr>
<td>
<code>methods</code></br>
<em>
[]string
</em>
</td>
<td>
<p>HTTP Method values are case-sensitive and formatted as follows:</p>
<p>exact: &ldquo;value&rdquo; for exact string match</p>
<p>prefix: &ldquo;value&rdquo; for prefix-based match</p>
<p>regex: &ldquo;value&rdquo; for ECMAscript style regex-based match</p>
</td>
</tr>
<tr>
<td>
<code>headers</code></br>
<em>
<a href="#rio.cattle.io/v1.HeaderMatch">
[]HeaderMatch
</a>
</em>
</td>
<td>
<p>The header keys must be lowercase and use hyphen as the separator, e.g. x-request-id.</p>
<p>Header values are case-sensitive and formatted as follows:</p>
<p>exact: &ldquo;value&rdquo; for exact string match</p>
<p>prefix: &ldquo;value&rdquo; for prefix-based match</p>
<p>regex: &ldquo;value&rdquo; for ECMAscript style regex-based match</p>
<p>Note: The keys uri, scheme, method, and authority will be ignored.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="NameValue">NameValue
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.HeaderOperations">HeaderOperations</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>value</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="NamedContainer">NamedContainer
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.PodConfig">PodConfig</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>The name of the container</p>
</td>
</tr>
<tr>
<td>
<code>init</code></br>
<em>
bool
</em>
</td>
<td>
<p>List of initialization containers belonging to the pod.
Init containers are executed in order prior to containers being started.
If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy.
The name for an init container or normal container must be unique among all containers.
Init containers may not have Lifecycle actions, Readiness probes, or Liveness probes.
The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit for each resource type, and then using the max of of that value or the sum of the normal containers.
Limits are applied to init containers in a similar fashion. Init containers cannot currently be added or removed. Cannot be updated. More info: <a href="https://kubernetes.io/docs/concepts/workloads/pods/init-containers/">https://kubernetes.io/docs/concepts/workloads/pods/init-containers/</a></p>
</td>
</tr>
<tr>
<td>
<code>Container</code></br>
<em>
<a href="#rio.cattle.io/v1.Container">
Container
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="Permission">Permission
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.ServiceSpec">ServiceSpec</a>, 
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.StackSpec">StackSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>role</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>verbs</code></br>
<em>
[]string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>apiGroup</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>resource</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>url</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>resourceName</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="PodConfig">PodConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.ServiceSpec">ServiceSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>containers</code></br>
<em>
<a href="#rio.cattle.io/v1.NamedContainer">
[]NamedContainer
</a>
</em>
</td>
<td>
<p>List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>hostname</code></br>
<em>
string
</em>
</td>
<td>
<p>Specifies the hostname of the Pod If not specified, the pod&rsquo;s hostname will be set to a system-defined value.</p>
</td>
</tr>
<tr>
<td>
<code>hostAliases</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#hostalias-v1-core">
[]Kubernetes core/v1.HostAlias
</a>
</em>
</td>
<td>
<p>HostAliases is an optional list of hosts and IPs that will be injected into the pod&rsquo;s hosts file if specified. This is only valid for non-hostNetwork pods.</p>
</td>
</tr>
<tr>
<td>
<code>hostNetwork</code></br>
<em>
bool
</em>
</td>
<td>
<p>Host networking requested for this pod. Use the host&rsquo;s network namespace. If this option is set, the ports that will be used must be specified. Default to false.</p>
</td>
</tr>
<tr>
<td>
<code>imagePullSecrets</code></br>
<em>
[]string
</em>
</td>
<td>
<p>Image pull secret</p>
</td>
</tr>
<tr>
<td>
<code>volumeTemplates</code></br>
<em>
<a href="#rio.cattle.io/v1.VolumeTemplate">
[]VolumeTemplate
</a>
</em>
</td>
<td>
<p>Volumes to create per replica</p>
</td>
</tr>
<tr>
<td>
<code>dns</code></br>
<em>
<a href="#rio.cattle.io/v1.DNS">
DNS
</a>
</em>
</td>
<td>
<p>DNS settings for this Pod</p>
</td>
</tr>
<tr>
<td>
<code>Affinity</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#affinity-v1-core">
Kubernetes core/v1.Affinity
</a>
</em>
</td>
<td>
<p>
(Members of <code>Affinity</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>Container</code></br>
<em>
<a href="#rio.cattle.io/v1.Container">
Container
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="PodDNSConfigOption">PodDNSConfigOption
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.DNS">DNS</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>value</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="Protocol">Protocol
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.ContainerPort">ContainerPort</a>)
</p>
<p>
</p>
<h3 id="Question">Question
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.TemplateMeta">TemplateMeta</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>variable</code></br>
<em>
string
</em>
</td>
<td>
<p>The variable name to reference using ${&hellip;} syntax</p>
</td>
</tr>
<tr>
<td>
<code>label</code></br>
<em>
string
</em>
</td>
<td>
<p>A friend name for the question</p>
</td>
</tr>
<tr>
<td>
<code>description</code></br>
<em>
string
</em>
</td>
<td>
<p>A longer description of the question</p>
</td>
</tr>
<tr>
<td>
<code>type</code></br>
<em>
string
</em>
</td>
<td>
<p>The field type: string, int, bool, enum. default is string</p>
</td>
</tr>
<tr>
<td>
<code>required</code></br>
<em>
bool
</em>
</td>
<td>
<p>The answer can not be blank</p>
</td>
</tr>
<tr>
<td>
<code>default</code></br>
<em>
string
</em>
</td>
<td>
<p>Default value of the answer if not specified by the user</p>
</td>
</tr>
<tr>
<td>
<code>group</code></br>
<em>
string
</em>
</td>
<td>
<p>Group the question with questions in the same group (Most used by UI)</p>
</td>
</tr>
<tr>
<td>
<code>minLength</code></br>
<em>
int
</em>
</td>
<td>
<p>Minimum length of the answer</p>
</td>
</tr>
<tr>
<td>
<code>maxLength</code></br>
<em>
int
</em>
</td>
<td>
<p>Maximum length of the answer</p>
</td>
</tr>
<tr>
<td>
<code>min</code></br>
<em>
int
</em>
</td>
<td>
<p>Minimum value of an int answer</p>
</td>
</tr>
<tr>
<td>
<code>max</code></br>
<em>
int
</em>
</td>
<td>
<p>Maximum value of an int answer</p>
</td>
</tr>
<tr>
<td>
<code>options</code></br>
<em>
[]string
</em>
</td>
<td>
<p>An array of valid answers for type enum questions</p>
</td>
</tr>
<tr>
<td>
<code>validChars</code></br>
<em>
string
</em>
</td>
<td>
<p>Answer must be composed of only these characters</p>
</td>
</tr>
<tr>
<td>
<code>invalidChars</code></br>
<em>
string
</em>
</td>
<td>
<p>Answer must not have any of these characters</p>
</td>
</tr>
<tr>
<td>
<code>subquestions</code></br>
<em>
<a href="#rio.cattle.io/v1.SubQuestion">
[]SubQuestion
</a>
</em>
</td>
<td>
<p>A list of questions that are considered child questions</p>
</td>
</tr>
<tr>
<td>
<code>showIf</code></br>
<em>
string
</em>
</td>
<td>
<p>Ask question only if this evaluates to true, more info on syntax below</p>
</td>
</tr>
<tr>
<td>
<code>showSubquestionIf</code></br>
<em>
string
</em>
</td>
<td>
<p>Ask subquestions if this evaluates to true</p>
</td>
</tr>
</tbody>
</table>
<h3 id="Redirect">Redirect
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.RouteSpec">RouteSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>host</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>path</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>prefix</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>toHTTPS</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="Retry">Retry
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.RouteSpec">RouteSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>timeoutSeconds</code></br>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>attempts</code></br>
<em>
int
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="Rewrite">Rewrite
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.RouteSpec">RouteSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>host</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>path</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="RolloutConfig">RolloutConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.ServiceSpec">ServiceSpec</a>)
</p>
<p>
<p>RolloutConfig specifies the configuration when promoting a new revision</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>increment</code></br>
<em>
int
</em>
</td>
<td>
<p>Increment Value each Rollout can scale up or down</p>
</td>
</tr>
<tr>
<td>
<code>intervalSeconds</code></br>
<em>
int
</em>
</td>
<td>
<p>Interval between each Rollout in seconds</p>
</td>
</tr>
<tr>
<td>
<code>pause</code></br>
<em>
bool
</em>
</td>
<td>
<p>Pause if true the rollout will stop in place until set to false.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="RouteSpec">RouteSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.RouterSpec">RouterSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>match</code></br>
<em>
<a href="#rio.cattle.io/v1.Match">
Match
</a>
</em>
</td>
<td>
<p>Match conditions to be satisfied for the rule to be activated. All conditions inside a single match block have AND semantics, while the list of match blocks have OR semantics.
The rule is matched if any one of the match blocks succeed.</p>
</td>
</tr>
<tr>
<td>
<code>to</code></br>
<em>
<a href="#rio.cattle.io/v1.WeightedDestination">
[]WeightedDestination
</a>
</em>
</td>
<td>
<p>A http rule can either redirect or forward (default) traffic. The forwarding target can be one of several versions of a service (see glossary in beginning of document).
Weights associated with the service version determine the proportion of traffic it receives.</p>
</td>
</tr>
<tr>
<td>
<code>redirect</code></br>
<em>
<a href="#rio.cattle.io/v1.Redirect">
Redirect
</a>
</em>
</td>
<td>
<p>A http rule can either redirect or forward (default) traffic. If traffic passthrough option is specified in the rule, route/redirect will be ignored.
The redirect primitive can be used to send a HTTP 301 redirect to a different URI or Authority.</p>
</td>
</tr>
<tr>
<td>
<code>rewrite</code></br>
<em>
<a href="#rio.cattle.io/v1.Rewrite">
Rewrite
</a>
</em>
</td>
<td>
<p>Rewrite HTTP URIs and Authority headers. Rewrite cannot be used with Redirect primitive. Rewrite will be performed before forwarding.</p>
</td>
</tr>
<tr>
<td>
<code>retry</code></br>
<em>
<a href="#rio.cattle.io/v1.Retry">
Retry
</a>
</em>
</td>
<td>
<p>Retries specifies the retry logic for each route</p>
</td>
</tr>
<tr>
<td>
<code>headers</code></br>
<em>
<a href="#rio.cattle.io/v1.HeaderOperations">
HeaderOperations
</a>
</em>
</td>
<td>
<p>Header manipulation rules</p>
</td>
</tr>
<tr>
<td>
<code>fault</code></br>
<em>
<a href="#rio.cattle.io/v1.Fault">
Fault
</a>
</em>
</td>
<td>
<p>Fault injection policy to apply on HTTP traffic at the client side. Note that timeouts or retries will not be enabled when faults are enabled on the client side.</p>
</td>
</tr>
<tr>
<td>
<code>mirror</code></br>
<em>
<a href="#rio.cattle.io/v1.Destination">
Destination
</a>
</em>
</td>
<td>
<p>Mirror HTTP traffic to a another destination in addition to forwarding the requests to the intended destination.
Mirrored traffic is on a best effort basis where the sidecar/gateway will not wait for the mirrored cluster to respond before returning the response from the original destination.
Statistics will be generated for the mirrored destination.</p>
</td>
</tr>
<tr>
<td>
<code>timeoutSeconds</code></br>
<em>
int
</em>
</td>
<td>
<p>TimeoutSeconds specifies timeout setting for each route</p>
</td>
</tr>
</tbody>
</table>
<h3 id="RouterSpec">RouterSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.Router">Router</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>routes</code></br>
<em>
<a href="#rio.cattle.io/v1.RouteSpec">
[]RouteSpec
</a>
</em>
</td>
<td>
<p>An ordered list of route rules for HTTP traffic. The first rule matching an incoming request is used.</p>
</td>
</tr>
<tr>
<td>
<code>internal</code></br>
<em>
bool
</em>
</td>
<td>
<p>By default all Routers are public and exposed outside of the cluster. Setting internal to true will
cause the Router to not be exposed</p>
</td>
</tr>
</tbody>
</table>
<h3 id="RouterStatus">RouterStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.Router">Router</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>endpoints</code></br>
<em>
[]string
</em>
</td>
<td>
<p>The endpoint to access the router</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code></br>
<em>
[]github.com/rancher/wrangler/pkg/genericcondition.GenericCondition
</em>
</td>
<td>
<p>Represents the latest available observations of a PublicDomain&rsquo;s current state.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="ScaleStatus">ScaleStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.ServiceStatus">ServiceStatus</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>unavailable</code></br>
<em>
int
</em>
</td>
<td>
<p>Total number of unavailable pods targeted by this deployment. This is the total number of pods that are still required for the deployment to have 100% available capacity.
They may either be pods that are running but not yet available or pods that still have not been created.</p>
</td>
</tr>
<tr>
<td>
<code>available</code></br>
<em>
int
</em>
</td>
<td>
<p>Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="ServiceSpec">ServiceSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.Service">Service</a>)
</p>
<p>
<p>ServiceSpec represents spec for Service</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>PodConfig</code></br>
<em>
<a href="#rio.cattle.io/v1.PodConfig">
PodConfig
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>template</code></br>
<em>
bool
</em>
</td>
<td>
<p>Template this service is a template for new versions to be created base on changes
from the build.repo</p>
</td>
</tr>
<tr>
<td>
<code>stageOnly</code></br>
<em>
bool
</em>
</td>
<td>
<p>StageOnly whether to only stage services that are generated through template from build.repo</p>
</td>
</tr>
<tr>
<td>
<code>version</code></br>
<em>
string
</em>
</td>
<td>
<p>Version version of this service</p>
</td>
</tr>
<tr>
<td>
<code>app</code></br>
<em>
string
</em>
</td>
<td>
<p>App The exposed app name, if no value is set, then metadata.name of the Service is used</p>
</td>
</tr>
<tr>
<td>
<code>weight</code></br>
<em>
int
</em>
</td>
<td>
<p>Weight The weight among services with matching app field to determine how much traffic is load balanced
to this service.  If rollout is set, the weight become the target weight of the rollout.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code></br>
<em>
int
</em>
</td>
<td>
<p>Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1 in deployment.</p>
</td>
</tr>
<tr>
<td>
<code>maxUnavailable</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#intorstring-intstr-util">
k8s.io/apimachinery/pkg/util/intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The maximum number of pods that can be unavailable during the update.
Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).
Absolute number is calculated from percentage by rounding down.
This can not be 0 if MaxSurge is 0.
Defaults to 25%.
Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods
immediately when the rolling update starts. Once new pods are ready, old ReplicaSet
can be scaled down further, followed by scaling up the new ReplicaSet, ensuring
that the total number of pods available at all times during the update is at
least 70% of desired pods.</p>
</td>
</tr>
<tr>
<td>
<code>maxSurge</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#intorstring-intstr-util">
k8s.io/apimachinery/pkg/util/intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The maximum number of pods that can be scheduled above the desired number of
pods.
Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).
This can not be 0 if MaxUnavailable is 0.
Absolute number is calculated from percentage by rounding up.
Defaults to 25%.
Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when
the rolling update starts, such that the total number of old and new pods do not exceed
130% of desired pods. Once old pods have been killed,
new ReplicaSet can be scaled up further, ensuring that total number of pods running
at any time during the update is at most 130% of desired pods.</p>
</td>
</tr>
<tr>
<td>
<code>autoscale</code></br>
<em>
<a href="#rio.cattle.io/v1.AutoscaleConfig">
AutoscaleConfig
</a>
</em>
</td>
<td>
<p>Autoscale the replicas based on the amount of traffic received by this service</p>
</td>
</tr>
<tr>
<td>
<code>rollout</code></br>
<em>
<a href="#rio.cattle.io/v1.RolloutConfig">
RolloutConfig
</a>
</em>
</td>
<td>
<p>RolloutConfig controls how each service is allocated ComputedWeight</p>
</td>
</tr>
<tr>
<td>
<code>global</code></br>
<em>
bool
</em>
</td>
<td>
<p>Place one pod per node that matches the scheduling rules</p>
</td>
</tr>
<tr>
<td>
<code>serviceMesh</code></br>
<em>
bool
</em>
</td>
<td>
<p>Whether to disable Service mesh for Service. If true, no mesh sidecar will be deployed along with the Service</p>
</td>
</tr>
<tr>
<td>
<code>requestTimeoutSeconds</code></br>
<em>
int
</em>
</td>
<td>
<p>RequestTimeoutSeconds specify the timeout set on api gateway for each individual service</p>
</td>
</tr>
<tr>
<td>
<code>permissions</code></br>
<em>
<a href="#rio.cattle.io/v1.Permission">
[]Permission
</a>
</em>
</td>
<td>
<p>Permissions to the Services. It will create corresponding ServiceAccounts, Roles and RoleBinding.</p>
</td>
</tr>
<tr>
<td>
<code>globalPermissions</code></br>
<em>
<a href="#rio.cattle.io/v1.Permission">
[]Permission
</a>
</em>
</td>
<td>
<p>GlobalPermissions to the Services. It will create corresponding ServiceAccounts, ClusterRoles and ClusterRoleBinding.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="ServiceStatus">ServiceStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.Service">Service</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>deploymentReady</code></br>
<em>
bool
</em>
</td>
<td>
<p>DeploymentReady for ready status on deployment</p>
</td>
</tr>
<tr>
<td>
<code>scaleStatus</code></br>
<em>
<a href="#rio.cattle.io/v1.ScaleStatus">
ScaleStatus
</a>
</em>
</td>
<td>
<p>ScaleStatus for the Service</p>
</td>
</tr>
<tr>
<td>
<code>computedApp</code></br>
<em>
string
</em>
</td>
<td>
<p>ComputedApp is the calculated value of Spec.App if not set</p>
</td>
</tr>
<tr>
<td>
<code>computedVersion</code></br>
<em>
string
</em>
</td>
<td>
<p>ComputedVersion is the calculated value of Spec.Version if not set</p>
</td>
</tr>
<tr>
<td>
<code>computedReplicas</code></br>
<em>
int
</em>
</td>
<td>
<p>ComputedReplicas is calculated from autoscaling component to make sure pod has the desired load</p>
</td>
</tr>
<tr>
<td>
<code>computedWeight</code></br>
<em>
int
</em>
</td>
<td>
<p>ComputedWeight is the weight calculated from the rollout revision</p>
</td>
</tr>
<tr>
<td>
<code>containerRevision</code></br>
<em>
<a href="#rio.cattle.io/v1.BuildRevision">
map[string]github.com/rancher/rio/pkg/apis/rio.cattle.io/v1.BuildRevision
</a>
</em>
</td>
<td>
<p>ContainerRevision are populated from builds to store commits for each repo</p>
</td>
</tr>
<tr>
<td>
<code>generatedServices</code></br>
<em>
map[string]bool
</em>
</td>
<td>
<p>GeneratedServices contains all the service names are generated from build template</p>
</td>
</tr>
<tr>
<td>
<code>gitCommits</code></br>
<em>
[]string
</em>
</td>
<td>
<p>GitCommits contains all git commits that triggers template update</p>
</td>
</tr>
<tr>
<td>
<code>shouldGenerate</code></br>
<em>
string
</em>
</td>
<td>
<p>ShouldGenerate contains the serviceName that should be generated on the next controller run</p>
</td>
</tr>
<tr>
<td>
<code>shouldClean</code></br>
<em>
map[string]bool
</em>
</td>
<td>
<p>ShouldClean contains all the services that are generated from template but should be cleaned up.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code></br>
<em>
[]github.com/rancher/wrangler/pkg/genericcondition.GenericCondition
</em>
</td>
<td>
<p>Represents the latest available observations of a deployment&rsquo;s current state.</p>
</td>
</tr>
<tr>
<td>
<code>endpoints</code></br>
<em>
[]string
</em>
</td>
<td>
<p>The Endpoints to access this version directly</p>
</td>
</tr>
<tr>
<td>
<code>appEndpoints</code></br>
<em>
[]string
</em>
</td>
<td>
<p>The Endpoints to access this service as part of an app</p>
</td>
</tr>
<tr>
<td>
<code>buildLogToken</code></br>
<em>
string
</em>
</td>
<td>
<p>log token to access build log</p>
</td>
</tr>
<tr>
<td>
<code>watch</code></br>
<em>
bool
</em>
</td>
<td>
<p>Watch represents if a service should creates git watcher to watch git changes</p>
</td>
</tr>
</tbody>
</table>
<h3 id="StackBuild">StackBuild
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.StackSpec">StackSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>repo</code></br>
<em>
string
</em>
</td>
<td>
<p>Git repo url</p>
</td>
</tr>
<tr>
<td>
<code>branch</code></br>
<em>
string
</em>
</td>
<td>
<p>Git branch</p>
</td>
</tr>
<tr>
<td>
<code>revision</code></br>
<em>
string
</em>
</td>
<td>
<p>Git revision</p>
</td>
</tr>
<tr>
<td>
<code>cloneSecretName</code></br>
<em>
string
</em>
</td>
<td>
<p>Git secret name for repository</p>
</td>
</tr>
<tr>
<td>
<code>rioFile</code></br>
<em>
string
</em>
</td>
<td>
<p>Specify the name of the Riofile in the Repo. This is the full path relative to the repo root. Defaults to <code>Riofile</code>.</p>
</td>
</tr>
<tr>
<td>
<code>webhookSecretName</code></br>
<em>
string
</em>
</td>
<td>
<p>Specify the github secret name. Used to create Github webhook, the secret key has to be <code>accessToken</code></p>
</td>
</tr>
</tbody>
</table>
<h3 id="StackSpec">StackSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.Stack">Stack</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>build</code></br>
<em>
<a href="#rio.cattle.io/v1.StackBuild">
StackBuild
</a>
</em>
</td>
<td>
<p>Stack build parameters that watches git repo</p>
</td>
</tr>
<tr>
<td>
<code>permissions</code></br>
<em>
<a href="#rio.cattle.io/v1.Permission">
[]Permission
</a>
</em>
</td>
<td>
<p>Permissions used while deploying objects created by this stack</p>
</td>
</tr>
<tr>
<td>
<code>additionalGroupVersionKinds</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#groupversionkind-schema-runtime">
[]k8s.io/apimachinery/pkg/runtime/schema.GroupVersionKind
</a>
</em>
</td>
<td>
<p>Additional GVKs not in the rio.cattle.io that have the rio.cattle.io/stack label. These objects
are &ldquo;owned&rdquo; by this stack</p>
</td>
</tr>
<tr>
<td>
<code>answers</code></br>
<em>
map[string]string
</em>
</td>
<td>
<p>Stack answers</p>
</td>
</tr>
</tbody>
</table>
<h3 id="StackStatus">StackStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.Stack">Stack</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>revision</code></br>
<em>
string
</em>
</td>
<td>
<p>Observed commit for the build</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code></br>
<em>
[]github.com/rancher/wrangler/pkg/genericcondition.GenericCondition
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="StringMatch">StringMatch
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.HeaderMatch">HeaderMatch</a>, 
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.Match">Match</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>exact</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>prefix</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>regexp</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="SubQuestion">SubQuestion
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.Question">Question</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>variable</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>label</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>description</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>type</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>required</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>default</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>group</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>minLength</code></br>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>maxLength</code></br>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>min</code></br>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>max</code></br>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>options</code></br>
<em>
[]string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>validChars</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>invalidChars</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>showIf</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="TemplateMeta">TemplateMeta
</h3>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>version</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>iconUrl</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>readme</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>questions</code></br>
<em>
<a href="#rio.cattle.io/v1.Question">
[]Question
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>goTemplate</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>envSubst</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="Volume">Volume
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.Container">Container</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name is the name of the volume. If multiple Volumes in the same pod share the same name they
will be the same underlying storage. If persistent is set to true Name is required and will be
used to reference a PersistentVolumeClaim in the current namespace.</p>
<p>If Name matches the name of a VolumeTemplate on this service then the VolumeTemplate will be used as the
source of the volume.</p>
</td>
</tr>
<tr>
<td>
<code>path</code></br>
<em>
string
</em>
</td>
<td>
<p>That path within the container to mount the volume to</p>
</td>
</tr>
<tr>
<td>
<code>hostpath</code></br>
<em>
string
</em>
</td>
<td>
<p>That path on the host to mount into this container</p>
</td>
</tr>
<tr>
<td>
<code>hostPathType</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#hostpathtype-v1-core">
Kubernetes core/v1.HostPathType
</a>
</em>
</td>
<td>
<p>The</p>
</td>
</tr>
<tr>
<td>
<code>persistent</code></br>
<em>
bool
</em>
</td>
<td>
<p>If Persistent is true then this volume refers to a PersistentVolumeClaim in this namespace. The
Name field is used to reference PersistentVolumeClaim.  If the Name of this Volume matches a VolumeTemplate
then Persistent is assumed to be true</p>
</td>
</tr>
</tbody>
</table>
<h3 id="VolumeTemplate">VolumeTemplate
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.PodConfig">PodConfig</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>labels</code></br>
<em>
map[string]string
</em>
</td>
<td>
<p>Labels to be applied to the created PVC</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code></br>
<em>
map[string]string
</em>
</td>
<td>
<p>Annotations to be applied to the created PVC</p>
</td>
</tr>
<tr>
<td>
<code>Name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name of the VolumeTemplate. A volume entry will use this name to refer to the created volume</p>
</td>
</tr>
<tr>
<td>
<code>accessModes</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#persistentvolumeaccessmode-v1-core">
[]Kubernetes core/v1.PersistentVolumeAccessMode
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>AccessModes contains the desired access modes the volume should have.
More info: <a href="https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1">https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1</a></p>
</td>
</tr>
<tr>
<td>
<code>storage</code></br>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Resources represents the minimum resources the volume should have.
More info: <a href="https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources">https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources</a></p>
</td>
</tr>
<tr>
<td>
<code>storageClassName</code></br>
<em>
string
</em>
</td>
<td>
<p>Name of the StorageClass required by the claim.
More info: <a href="https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1">https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1</a></p>
</td>
</tr>
<tr>
<td>
<code>volumeMode</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#persistentvolumemode-v1-core">
Kubernetes core/v1.PersistentVolumeMode
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>volumeMode defines what type of volume is required by the claim.
Value of Filesystem is implied when not included in claim spec.
This is a beta feature.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="WeightedDestination">WeightedDestination
</h3>
<p>
(<em>Appears on:</em>
<a href="#github.com%2francher%2frio%2fpkg%2fapis%2frio.cattle.io%2fv1.RouteSpec">RouteSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Destination</code></br>
<em>
<a href="#rio.cattle.io/v1.Destination">
Destination
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>weight</code></br>
<em>
int
</em>
</td>
<td>
<p>Weight for the Destination</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<p><em>
Generated with <code>gen-crd-api-reference-docs</code>
on git commit <code>a593a33c</code>.
</em></p>
