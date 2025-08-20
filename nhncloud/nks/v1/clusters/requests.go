package clusters

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToClusterListQuery() (string, error)
}

// ListOpts enables listing Clusters based on specific attributes.
type ListOpts struct {
	// Currently no query parameters are supported for cluster listing
}

// ToClusterListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToClusterListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager that allows you to iterate over a collection of Clusters.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToClusterListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ClusterPage{pagination.SinglePageBase(r)}
	})
}

// Get retrieves a specific NKS cluster based on its unique ID or name.
// To extract the cluster object from the response, call the Extract method on the GetResult.
func Get(client *gophercloud.ServiceClient, clusterIDOrName string) (r GetResult) {
	resp, err := client.Get(getURL(client, clusterIDOrName), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToClusterCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies Cluster creation parameters.
type CreateOpts struct {
	// Name is the name of the cluster.
	Name string `json:"name" required:"true"`

	// ClusterTemplateID is the template ID used for the cluster.
	ClusterTemplateID string `json:"cluster_template_id" required:"true"`

	// FixedNetwork is the VPC network UUID.
	FixedNetwork string `json:"fixed_network" required:"true"`

	// FixedSubnet is the VPC subnet UUID.
	FixedSubnet string `json:"fixed_subnet" required:"true"`

	// FlavorID is the instance flavor UUID.
	FlavorID string `json:"flavor_id" required:"true"`

	// Keypair is the SSH keypair name.
	Keypair string `json:"keypair" required:"true"`

	// NodeCount is the number of worker nodes.
	NodeCount int `json:"node_count" required:"true"`

	// Labels contains cluster configuration labels.
	Labels map[string]interface{} `json:"labels,omitempty"`

	// APIEndpointIPACL contains API endpoint IP access control settings.
	APIEndpointIPACL *APIEndpointIPACL `json:"api_ep_ipacl,omitempty"`

	// Addons contains addon configurations.
	Addons []Addon `json:"addons,omitempty"`
}

// ToClusterCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToClusterCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create requests the creation of a new Cluster.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToClusterCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete requests the deletion of a Cluster.
func Delete(client *gophercloud.ServiceClient, clusterIDOrName string) (r DeleteResult) {
	resp, err := client.Delete(deleteURL(client, clusterIDOrName), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ResizeOptsBuilder allows extensions to add additional parameters to the
// Resize request.
type ResizeOptsBuilder interface {
	ToClusterResizeMap() (map[string]interface{}, error)
}

// ResizeOpts specifies Cluster resize parameters.
type ResizeOpts struct {
	// NodeCount is the new number of worker nodes.
	NodeCount int `json:"node_count" required:"true"`

	// Options contains resize options.
	Options *ResizeOptions `json:"options,omitempty"`
}

// ToClusterResizeMap constructs a request body from ResizeOpts.
func (opts ResizeOpts) ToClusterResizeMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Resize requests the resizing of a Cluster.
func Resize(client *gophercloud.ServiceClient, clusterIDOrName string, opts ResizeOptsBuilder) (r ResizeResult) {
	b, err := opts.ToClusterResizeMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(resizeURL(client, clusterIDOrName), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
