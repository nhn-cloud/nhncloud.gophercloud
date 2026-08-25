package internetgateways

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// Extract interprets any commonResult as a InternetGateway.
func (r commonResult) Extract() (*InternetGateway, error) {
	var s InternetGateway
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "internetgateway")
}

// CreateResult represents the result of a create operation.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation.
type GetResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

// InternetGateway represents a InternetGateway resource.
type InternetGateway struct {
	// 인터넷 게이트웨이 ID
	ID string `json:"id,omitempty"`
	// 인터넷 게이트웨이 이름
	Name string `json:"name,omitempty"`
	// 인터넷 게이트웨이가 연결한 외부 네트워크의 ID
	ExternalNetworkID string `json:"external_network_id,omitempty"`
	// 라우팅 테이블과 연결되어 있을 때 인터넷 게이트웨이를 연결한 라우팅 테이블 ID
	RoutingtableID string `json:"routingtable_id,omitempty"`
	// 인터넷 게이트웨이의 상태.
	State string `json:"state,omitempty"`
	// 인터넷 게이트웨이 생성 시간(UTC)
	CreateTime string `json:"create_time,omitempty"`
	// 인터넷 게이트웨이가 속한 테넌트 ID
	TenantID string `json:"tenant_id,omitempty"`
	// 점검으로 인한 인터넷 게이트웨이 이동 시 처리 상태
	MigrateStatus string `json:"migrate_status,omitempty"`
	// 점검으로 인한 인터넷 게이트웨이 이동 중 오류 발생 시 오류 메시지
	MigrateError string `json:"migrate_error,omitempty"`
}

// InternetGatewayPage is the page returned by a pager when traversing over a
// collection of InternetGateway resources.
type InternetGatewayPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of InternetGateway has
// reached the end of a page and the pager seeks to traverse over a new one.
func (r InternetGatewayPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"internetgateways_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a InternetGatewayPage struct is empty.
func (r InternetGatewayPage) IsEmpty() (bool, error) {
	if r.StatusCode == 204 {
		return true, nil
	}
	is, err := ExtractInternetGateways(r)
	return len(is) == 0, err
}

// ExtractInternetGateways accepts a Page struct, specifically a InternetGatewayPage
// struct, and extracts the elements into a slice of InternetGateway structs.
func ExtractInternetGateways(r pagination.Page) ([]InternetGateway, error) {
	var s []InternetGateway
	err := ExtractInternetGatewaysInto(r, &s)
	return s, err
}

func ExtractInternetGatewaysInto(r pagination.Page, v interface{}) error {
	return r.(InternetGatewayPage).Result.ExtractIntoSlicePtr(v, "internetgateways")
}
