package internetgateways

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToInternetGatewayListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through the API.
type ListOpts struct {
	// 조회할 인터넷 게이트웨이가 속한 테넌트
	TenantID string `q:"tenant_id"`
	// 조회할 인터넷 게이트웨이 ID
	ID string `q:"id"`
	// 조회할 인터넷 게이트웨이 이름
	Name string `q:"name"`
	// 조회할 인터넷 게이트웨이가 연결한 외부 네트워크의 ID
	ExternalNetworkID string `q:"external_network_id"`
	// 조회할 인터넷 게이트웨이를 연결한 라우팅 테이블의 ID
	RoutingtableID string `q:"routingtable_id"`
}

// ToInternetGatewayListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToInternetGatewayListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of InternetGateway.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToInternetGatewayListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return InternetGatewayPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific InternetGateway based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := c.Get(getURL(c, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToInternetGatewayCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options used to create a InternetGateway.
type CreateOpts struct {
	// 인터넷 게이트웨이 이름
	Name string `json:"name" required:"true"`
	// 인터넷 게이트웨이가 연결할 외부 네트워크 ID
	ExternalNetworkID string `json:"external_network_id" required:"true"`
}

// ToInternetGatewayCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToInternetGatewayCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "internetgateway")
}

// Create accepts a CreateOpts struct and creates a new InternetGateway.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToInternetGatewayCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Post(createURL(c), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201, 202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete accepts a unique ID and deletes the InternetGateway associated with it.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	resp, err := c.Delete(deleteURL(c, id), &gophercloud.RequestOpts{
		OkCodes: []int{202, 204},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
