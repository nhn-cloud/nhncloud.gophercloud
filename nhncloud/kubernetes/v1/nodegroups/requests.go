package nodegroups

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/containerinfra/v1/nodegroups"
	"github.com/gophercloud/gophercloud/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToNodegroupCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options used to create a kubernetes nodegroup.
type CreateOpts struct {
	// DockerVolumeSize is the size of a volume to allocate to docker for
	// container/image storage.
	DockerVolumeSize *int `json:"docker_volume_size,omitempty"`

	// FlavorID is the nova flavor ID to use when launching the nodegroup.
	FlavorID string `json:"flavor_id,omitempty"`

	// ImageID is the glance image ID used to boot the nodegroup.
	ImageID string `json:"image_id,omitempty"`

	// Labels is an arbitrary key=value pair.
	Labels map[string]string `json:"labels,omitempty"`

	// MaxNodeCount is the maximum allowed number of nodes in the nodegroup.
	MaxNodeCount *int `json:"max_node_count,omitempty"`

	// MergeLabels indicates whether we want to merge cluster template labels
	// into nodegroup labels when creating nodegroup.
	MergeLabels *bool `json:"merge_labels,omitempty"`

	// MinNodeCount is the minimum allowed number of nodes in the nodegroup.
	MinNodeCount int `json:"min_node_count,omitempty"`

	// Name is the name of the nodegroup.
	Name string `json:"name" required:"true"`

	// NodeCount is the number of nodes in the nodegroup.
	NodeCount *int `json:"node_count,omitempty"`

	// Role is the role of the nodegroup (e.g., "worker", "master").
	Role string `json:"role,omitempty"`
}

// ToNodegroupCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToNodegroupCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create requests the creation of a new kubernetes nodegroup.
func Create(client *gophercloud.ServiceClient, clusterID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToNodegroupCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(createURL(client, clusterID), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Get retrieves a specific kubernetes nodegroup based on its unique ID.
func Get(client *gophercloud.ServiceClient, clusterID, nodegroupID string) (r GetResult) {
	resp, err := client.Get(getURL(client, clusterID, nodegroupID), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToNodegroupListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API.
type ListOpts struct {
	// Limit is the maximum number of nodegroups to return.
	Limit int `q:"limit"`

	// Marker is the UUID of the last nodegroup on the previous page.
	Marker string `q:"marker"`

	// SortDir allows to choose the direction of sorting.
	SortDir string `q:"sort_dir"`

	// SortKey allows to sort by one of the nodegroup attributes.
	SortKey string `q:"sort_key"`

	// Role allows to filter nodegroups by role.
	Role string `q:"role"`
}

// ToNodegroupListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToNodegroupListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// kubernetes nodegroups.
func List(c *gophercloud.ServiceClient, clusterID string, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c, clusterID)
	if opts != nil {
		query, err := opts.ToNodegroupListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return NodegroupPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToNodegroupUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains the values used when updating a kubernetes nodegroup.
type UpdateOpts struct {
	// Op is the operation to be performed on the nodegroup attribute.
	Op nodegroups.UpdateOp `json:"op" required:"true"`

	// Path is the attribute path of the nodegroup.
	Path string `json:"path" required:"true"`

	// Value is the new value of the specified attribute.
	Value interface{} `json:"value,omitempty"`
}

// ToNodegroupUpdateMap assembles a request body based on the contents of an
// UpdateOpts.
func (opts UpdateOpts) ToNodegroupUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Update allows nodegroups to be updated.
func Update(client *gophercloud.ServiceClient, clusterID, nodegroupID string, opts []UpdateOptsBuilder) (r UpdateResult) {
	var body []map[string]interface{}
	for _, opt := range opts {
		b, err := opt.ToNodegroupUpdateMap()
		if err != nil {
			r.Err = err
			return r
		}
		body = append(body, b)
	}
	resp, err := client.Patch(updateURL(client, clusterID, nodegroupID), body, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete deletes the specified kubernetes nodegroup ID.
func Delete(client *gophercloud.ServiceClient, clusterID, nodegroupID string) (r DeleteResult) {
	resp, err := client.Delete(deleteURL(client, clusterID, nodegroupID), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
