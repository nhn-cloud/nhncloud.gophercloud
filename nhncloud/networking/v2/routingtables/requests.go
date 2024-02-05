package routingtables

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToRoutingtableListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the routingtable attributes you want to see returned. SortKey allows you to sort
// by a particular routingtable attribute. SortDir sets the direction, and is either
// `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	TenantID     string `q:"tenant_id"`
	ID           string `q:"id"`
	Name         string `q:"name"`
	DefaultTable *bool  `q:"default_table"`
	GatewayID    string `q:"gateway_id"`
	Distributed  *bool  `q:"distributed"`
	Detail       *bool  `q:"detail"`
	SortKey      string `q:"sort_key"`
	SortDir      string `q:"sort_dir"`
	Marker       string `q:"marker"`
	Limit        int    `q:"limit"`
}

// ToRoutingtableListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToRoutingtableListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// routingtables. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToRoutingtableListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return RoutingtablePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific routingtable based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := c.Get(getURL(c, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToRoutingtableCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options used to create a routingtable.
type CreateOpts struct {
	Name        string `json:"name"`
	VpcID       string `json:"vpc_id"`
	Distributed bool   `json:"distributed,omitempty"`
}

// ToRoutingtableCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToRoutingtableCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "routingtable")
}

// Create accepts a CreateOpts struct and creates a new routingtable using the values
// provided. This operation does not actually require a request body, i.e. the
// CreateOpts struct argument can be empty.
//
// The tenant ID that is contained in the URI is the tenant that creates the
// routingtable. An admin user, however, has the option of specifying another tenant
// ID in the CreateOpts struct.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToRoutingtableCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Post(createURL(c), b, &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToRoutingtableUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents options used to update a routingtable.
type UpdateOpts struct {
	Name        *string `json:"name,omitempty"`
	Distributed *bool   `json:"distributed,omitempty"`
}

// ToRoutingtableUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToRoutingtableUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "routingtable")
}

// Update accepts a UpdateOpts struct and updates an existing routingtable using the
// values provided. For more information, see the Create function.
func Update(c *gophercloud.ServiceClient, routingtableID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToRoutingtableUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	h, err := gophercloud.BuildHeaders(opts)
	if err != nil {
		r.Err = err
		return
	}

	resp, err := c.Put(updateURL(c, routingtableID), b, &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: h,
		OkCodes:     []int{200, 201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete accepts a unique ID and deletes the routingtable associated with it.
func Delete(c *gophercloud.ServiceClient, routingtableID string) (r DeleteResult) {
	resp, err := c.Delete(deleteURL(c, routingtableID), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
