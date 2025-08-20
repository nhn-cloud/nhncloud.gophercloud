package nodegroups

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToNodegroupListQuery() (string, error)
}

// ListOpts enables listing Nodegroups based on specific attributes.
type ListOpts struct {
	// Currently no query parameters are supported for nodegroup listing
}

// ToNodegroupListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToNodegroupListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager that allows you to iterate over a collection of Nodegroups.
func List(client *gophercloud.ServiceClient, clusterIDOrName string, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client, clusterIDOrName)
	if opts != nil {
		query, err := opts.ToNodegroupListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return NodegroupPage{pagination.SinglePageBase(r)}
	})
}

// Get retrieves a specific NKS nodegroup based on its unique ID.
// clusterIDOrName can be either the cluster UUID or name.
// To extract the nodegroup object from the response, call the Extract method on the GetResult.
func Get(client *gophercloud.ServiceClient, clusterIDOrName, nodegroupIDOrName string) (r GetResult) {
	resp, err := client.Get(getURL(client, clusterIDOrName, nodegroupIDOrName), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToNodegroupCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies Nodegroup creation parameters.
type CreateOpts struct {
	// Name is the name of the nodegroup.
	Name string `json:"name" required:"true"`

	// NodeCount is the number of worker nodes.
	NodeCount int `json:"node_count" required:"true"`

	// FlavorID is the instance flavor UUID for worker nodes.
	FlavorID string `json:"flavor_id" required:"true"`

	// ImageID is the image UUID for worker nodes.
	ImageID string `json:"image_id" required:"true"`

	// Labels contains nodegroup configuration labels.
	Labels map[string]interface{} `json:"labels,omitempty"`
}

// ToNodegroupCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToNodegroupCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create requests the creation of a new Nodegroup.
// clusterIDOrName can be either the cluster UUID or name.
func Create(client *gophercloud.ServiceClient, clusterIDOrName string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToNodegroupCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(createURL(client, clusterIDOrName), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete requests the deletion of a Nodegroup.
// clusterIDOrName can be either the cluster UUID or name.
func Delete(client *gophercloud.ServiceClient, clusterIDOrName, nodegroupIDOrName string) (r DeleteResult) {
	resp, err := client.Delete(deleteURL(client, clusterIDOrName, nodegroupIDOrName), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpgradeOptsBuilder allows extensions to add additional parameters to the
// Upgrade request.
type UpgradeOptsBuilder interface {
	ToNodegroupUpgradeMap() (map[string]interface{}, error)
}

// UpgradeOpts specifies Nodegroup upgrade parameters.
type UpgradeOpts struct {
	// Version is the Kubernetes version to upgrade to.
	Version string `json:"version" required:"true"`

	// Options contains upgrade options.
	Options *UpgradeOptions `json:"options,omitempty"`
}

// ToNodegroupUpgradeMap constructs a request body from UpgradeOpts.
func (opts UpgradeOpts) ToNodegroupUpgradeMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Upgrade requests the upgrade of a Nodegroup.
// clusterIDOrName can be either the cluster UUID or name.
func Upgrade(client *gophercloud.ServiceClient, clusterIDOrName, nodegroupIDOrName string, opts UpgradeOptsBuilder) (r UpgradeResult) {
	b, err := opts.ToNodegroupUpgradeMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(upgradeURL(client, clusterIDOrName, nodegroupIDOrName), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
