package routingtables

import (
	"encoding/json"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

func (r commonResult) Extract() (*Routingtable, error) {
	var s Routingtable
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "routingtable")
}

type CreateResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type Routingtable struct {
	// Name Routing table name
	Name string `json:"name"`

	// TenantID Tenant ID to which routing table is included
	TenantID string `json:"tenant_id"`

	// ID Routing table ID
	ID string `json:"id"`

	// DefaultTable Whether routing table is default table
	DefaultTable bool `json:"default_table"`

	// Distributed Routing method of routing table to query
	Distributed bool `json:"distributed"`

	// GatewayID The ID of the internet gateway when the gateway is associated with routing table
	GatewayID string `json:"gateway_id"`

	// State Status of the routing table. Currently, only the available status exists
	State string `json:"state"`

	// CreateTime Routing table creation time
	CreateTime string `json:"create_time"`

	// UpdatedAt and CreatedAt contain ISO-8601 timestamps of when the state of the
	// network last changed, and when it was created.
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}

type RoutingtableDetail struct {
	// Name Routing table name
	Name string `json:"name"`

	// TenantID Tenant ID to which routing table is included
	TenantID string `json:"tenant_id"`

	// Routingtable Routing table ID
	ID string `json:"id"`

	// DefaultTable Whether routing table is default table
	DefaultTable bool `json:"default_table"`

	// Distributed Routing method of routing table to query
	Distributed bool `json:"distributed"`

	// GatewayID The ID of the internet gateway when the gateway is associated with routing table
	GatewayID string `json:"gateway_id"`

	// State Status of the routing table. Currently, only the available status exists
	State string `json:"state"`

	// VPCs List of VPC information strings to which routing table belongs
	VPCs []string `json:"vpcs,omitempty"`

	// Subnets List of subnet information strings associated with routing table
	Subnets []string `json:"subnets,omitempty"`

	// Routes List of route information objects set to routing table
	Routes []Route `json:"routes,omitempty"`

	// CreateTime Routing table creation time
	CreateTime string `json:"create_time"`

	// UpdatedAt and CreatedAt contain ISO-8601 timestamps of when the state of the
	// network last changed, and when it was created.
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}

type Route struct {
	TenantID       string `json:"tenant_id,omitempty"`
	Mask           int    `json:"mask,omitempty"`
	Gateway        string `json:"gateway,omitempty"`
	GatewayID      string `json:"gateway_id,omitempty"`
	RoutingtableID string `json:"routingtable_id,omitempty"`
	CIDR           string `json:"cidr,omitempty"`
	ID             string `json:"id,omitempty"`
}

func (r *Routingtable) UnmarshalJSON(b []byte) error {
	type tmp Routingtable

	// Support for older neutron time format
	var s1 struct {
		tmp
		CreatedAt gophercloud.JSONRFC3339NoZ `json:"created_at"`
		UpdatedAt gophercloud.JSONRFC3339NoZ `json:"updated_at"`
	}

	err := json.Unmarshal(b, &s1)
	if err == nil {
		*r = Routingtable(s1.tmp)
		r.CreatedAt = time.Time(s1.CreatedAt)
		r.UpdatedAt = time.Time(s1.UpdatedAt)

		return nil
	}

	// Support for newer neutron time format
	var s2 struct {
		tmp
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	err = json.Unmarshal(b, &s2)
	if err != nil {
		return err
	}

	*r = Routingtable(s2.tmp)
	r.CreatedAt = time.Time(s2.CreatedAt)
	r.UpdatedAt = time.Time(s2.UpdatedAt)

	return nil
}

// RoutingtablePage is the page returned by a pager when traversing over a
// collection of networks.
type RoutingtablePage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of networks has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r RoutingtablePage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"networks_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a RoutingtablePage struct is empty.
func (r RoutingtablePage) IsEmpty() (bool, error) {
	if r.StatusCode == 204 {
		return true, nil
	}

	is, err := ExtractRoutingtables(r)
	return len(is) == 0, err
}

// ExtractRoutingtables accepts a Page struct, specifically a RoutingtablePage struct,
// and extracts the elements into a slice of Routingtable structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractRoutingtables(r pagination.Page) ([]Routingtable, error) {
	var s []Routingtable
	err := ExtractRoutingtablesInto(r, &s)
	return s, err
}

func ExtractRoutingtablesInto(r pagination.Page, v interface{}) error {
	return r.(RoutingtablePage).Result.ExtractIntoSlicePtr(v, "routingtables")
}
