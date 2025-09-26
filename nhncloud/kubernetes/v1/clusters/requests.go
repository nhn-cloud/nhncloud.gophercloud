package clusters

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/containerinfra/v1/clusters"
	"github.com/gophercloud/gophercloud/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToClustersCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options used to create a kubernetes cluster.
type CreateOpts struct {
	// ClusterTemplateID is the UUID of the cluster template.
	ClusterTemplateID string `json:"cluster_template_id" required:"true"`

	// CreateTimeout is the timeout for cluster creation.
	CreateTimeout *int `json:"create_timeout,omitempty"`

	// DiscoveryURL is the URL used for cluster node discovery.
	DiscoveryURL string `json:"discovery_url,omitempty"`

	// DockerVolumeSize is the size of a volume to allocate to docker for
	// container/image storage.
	DockerVolumeSize *int `json:"docker_volume_size,omitempty"`

	// FlavorID is the nova flavor ID to use when launching the cluster.
	FlavorID string `json:"flavor_id,omitempty"`

	// FloatingIPEnabled indicates whether created cluster should create IP
	// floating IP for every node or not.
	FloatingIPEnabled *bool `json:"floating_ip_enabled,omitempty"`

	// FixedNetwork is the network the cluster gets fixed IP from.
	FixedNetwork string `json:"fixed_network,omitempty"`

	// FixedSubnet is the subnet the cluster gets fixed IP from.
	FixedSubnet string `json:"fixed_subnet,omitempty"`

	// Keypair is the name of the SSH keypair to configure in the cluster
	// servers.
	Keypair string `json:"keypair,omitempty"`

	// Labels is an arbitrary key=value pair.
	Labels map[string]string `json:"labels,omitempty"`

	// MasterCount is the number of master nodes for the cluster.
	MasterCount *int `json:"master_count,omitempty"`

	// MasterFlavorID is the nova flavor ID to use when launching the master
	// nodes.
	MasterFlavorID string `json:"master_flavor_id,omitempty"`

	// MergeLabels indicates whether we want to merge cluster template labels
	// into cluster labels when creating cluster.
	MergeLabels *bool `json:"merge_labels,omitempty"`

	// Name is the name of the cluster.
	Name string `json:"name,omitempty"`

	// NodeCount is the number of worker nodes for the cluster.
	NodeCount *int `json:"node_count,omitempty"`

	// APIEndpointIPACL contains API endpoint IP access control settings.
	APIEndpointIPACL *APIEndpointIPACL `json:"api_ep_ipacl,omitempty"`

	// Addons contains addon configurations.
	Addons []Addon `json:"addons,omitempty"`
}

// APIEndpointIPACL represents IP access control configuration.
type APIEndpointIPACL struct {
	// Enable indicates whether IP access control is enabled.
	Enable string `json:"enable,omitempty"`

	// Action specifies the access control action (ALLOW/DENY).
	Action string `json:"action,omitempty"`

	// IPACLTargets contains the list of IP access control targets.
	IPACLTargets []IPACLTarget `json:"ipacl_targets,omitempty"`
}

// IPACLTarget represents an IP access control target.
type IPACLTarget struct {
	// CidrAddress is the IP address or CIDR range.
	CidrAddress string `json:"cidr_address"`

	// Description is the description for this IP access control target.
	Description string `json:"description,omitempty"`
}

// Addon represents an addon configuration.
type Addon struct {
	// Name is the addon name.
	Name string `json:"name"`

	// Version is the addon version.
	Version string `json:"version"`

	// Options contains addon-specific options.
	Options map[string]interface{} `json:"options,omitempty"`
}

// ToClustersCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToClustersCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create requests the creation of a new kubernetes cluster.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToClustersCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Get retrieves a specific kubernetes cluster based on its unique ID.
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := client.Get(getURL(client, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToClustersListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API.
type ListOpts struct {
	// Limit is the maximum number of clusters to return.
	Limit int `q:"limit"`

	// Marker is the UUID of the last cluster on the previous page.
	Marker string `q:"marker"`

	// SortDir allows to choose the direction of sorting.
	SortDir string `q:"sort_dir"`

	// SortKey allows to sort by one of the cluster attributes.
	SortKey string `q:"sort_key"`
}

// ToClustersListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToClustersListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// kubernetes clusters.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToClustersListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return ClusterPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// ListDetail returns a Pager which allows you to iterate over a collection of
// clusters with detailed information.
// It accepts a ListOptsBuilder, which allows you to sort the returned
// collection for greater efficiency.
func ListDetail(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listDetailURL(c)
	if opts != nil {
		query, err := opts.ToClustersListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return ClusterPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

type UpdateOp string

const (
	AddOp     UpdateOp = "add"
	RemoveOp  UpdateOp = "remove"
	ReplaceOp UpdateOp = "replace"
)

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToClustersUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains the values used when updating a kubernetes cluster.
type UpdateOpts struct {
	// Op is the operation to be performed on the cluster attribute.
	Op clusters.UpdateOp `json:"op" required:"true"`

	// Path is the attribute path of the cluster.
	Path string `json:"path" required:"true"`

	// Value is the new value of the specified attribute.
	Value interface{} `json:"value,omitempty"`
}

// ToClustersUpdateMap assembles a request body based on the contents of an
// UpdateOpts.
func (opts UpdateOpts) ToClustersUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Update allows clusters to be updated.
func Update(client *gophercloud.ServiceClient, id string, opts []UpdateOptsBuilder) (r UpdateResult) {
	var body []map[string]interface{}
	for _, opt := range opts {
		b, err := opt.ToClustersUpdateMap()
		if err != nil {
			r.Err = err
			return r
		}
		body = append(body, b)
	}
	resp, err := client.Patch(updateURL(client, id), body, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete deletes the specified kubernetes cluster ID.
func Delete(client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	resp, err := client.Delete(deleteURL(client, id), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ResizeOptsBuilder allows extensions to add additional parameters to the
// Resize request.
type ResizeOptsBuilder interface {
	ToClustersResizeMap() (map[string]interface{}, error)
}

// ResizeOpts contains the values used when resizing a kubernetes cluster.
type ResizeOpts struct {
	// NodeCount is the target number of worker nodes for the cluster.
	NodeCount *int `json:"node_count,omitempty"`

	// NodeGroup is the target nodegroup for the cluster.
	NodeGroup string `json:"nodegroup,omitempty"`

	// NodesToRemove contains the list of node UUIDs to remove.
	NodesToRemove []string `json:"nodes_to_remove,omitempty"`
}

// ToClustersResizeMap assembles a request body based on the contents of a
// ResizeOpts.
func (opts ResizeOpts) ToClustersResizeMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Resize allows a cluster to be resized.
func Resize(client *gophercloud.ServiceClient, id string, opts ResizeOptsBuilder) (r ResizeResult) {
	b, err := opts.ToClustersResizeMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(resizeURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpgradeOptsBuilder allows extensions to add additional parameters to the
// Upgrade request.
type UpgradeOptsBuilder interface {
	ToClustersUpgradeMap() (map[string]interface{}, error)
}

// UpgradeOpts contains the values used when upgrading a kubernetes cluster.
type UpgradeOpts struct {
	// ClusterTemplate is the new cluster template ID to upgrade to.
	ClusterTemplate string `json:"cluster_template,omitempty"`
}

// ToClustersUpgradeMap assembles a request body based on the contents of an
// UpgradeOpts.
func (opts UpgradeOpts) ToClustersUpgradeMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Upgrade allows a cluster to be upgraded.
func Upgrade(client *gophercloud.ServiceClient, id string, opts UpgradeOptsBuilder) (r UpgradeResult) {
	b, err := opts.ToClustersUpgradeMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(upgradeURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
