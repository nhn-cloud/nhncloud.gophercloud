package nodegroups

import (
	"encoding/json"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Nodegroup represents an NKS nodegroup in NHN Cloud.
type Nodegroup struct {
	// UUID is the unique identifier for the nodegroup.
	UUID string `json:"uuid"`

	// Name is the name of the nodegroup.
	Name string `json:"name"`

	// ClusterID is the UUID of the cluster this nodegroup belongs to.
	ClusterID string `json:"cluster_id"`

	// NodeCount is the number of worker nodes in this nodegroup.
	NodeCount int `json:"node_count"`

	// FlavorID is the instance flavor UUID for worker nodes.
	FlavorID string `json:"flavor_id"`

	// ImageID is the image UUID for worker nodes.
	ImageID string `json:"image_id"`

	// Labels contains nodegroup configuration labels.
	Labels map[string]interface{} `json:"labels,omitempty"`

	// Status is the current status of the nodegroup.
	Status string `json:"status,omitempty"`

	// StatusReason provides additional information about the status.
	StatusReason string `json:"status_reason,omitempty"`

	// CreatedAt is the creation timestamp.
	CreatedAt string `json:"created_at,omitempty"`

	// UpdatedAt is the last update timestamp.
	UpdatedAt string `json:"updated_at,omitempty"`

	// Version is the Kubernetes version of the nodegroup.
	Version string `json:"version,omitempty"`
}

// NodegroupPage stores a single page of all Nodegroup results from a List call.
type NodegroupPage struct {
	pagination.SinglePageBase
}

// IsEmpty determines whether or not a NodegroupPage is empty.
func (page NodegroupPage) IsEmpty() (bool, error) {
	if page.StatusCode == 204 {
		return true, nil
	}

	nodegroups, err := ExtractNodegroups(page)
	return len(nodegroups) == 0, err
}

// ExtractNodegroups interprets a page of results as a slice of Nodegroups.
func ExtractNodegroups(r pagination.Page) ([]Nodegroup, error) {
	var s struct {
		Nodegroups []Nodegroup `json:"nodegroups"`
	}
	err := (r.(NodegroupPage)).ExtractInto(&s)
	return s.Nodegroups, err
}

type nodegroupResult struct {
	gophercloud.Result
}

// Extract will get the Nodegroup object out of the nodegroupResult object.
func (r nodegroupResult) Extract() (*Nodegroup, error) {
	var s struct {
		Nodegroup *Nodegroup `json:"nodegroup"`
	}
	err := r.ExtractInto(&s)
	return s.Nodegroup, err
}

// ExtractNodegroup extracts a nodegroup from a raw response.
func ExtractNodegroup(raw []byte) (*Nodegroup, error) {
	var response struct {
		Nodegroup Nodegroup `json:"nodegroup"`
	}
	err := json.Unmarshal(raw, &response)
	if err != nil {
		return nil, err
	}
	return &response.Nodegroup, nil
}

// ExtractNodegroupInto extracts a nodegroup into a provided interface.
func ExtractNodegroupInto(raw []byte, v interface{}) error {
	var response struct {
		Nodegroup interface{} `json:"nodegroup"`
	}
	response.Nodegroup = v
	return json.Unmarshal(raw, &response)
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	nodegroupResult
}

// CreateResult is the response from a Create operation.
type CreateResult struct {
	nodegroupResult
}

// DeleteResult is the response from a Delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

// UpgradeOptions contains options for nodegroup upgrade operation.
type UpgradeOptions struct {
	// NumBufferNodes is the number of buffer nodes.
	NumBufferNodes int `json:"num_buffer_nodes,omitempty"`

	// NumMaxUnavailableNodes is the maximum number of unavailable nodes.
	NumMaxUnavailableNodes int `json:"num_max_unavailable_nodes,omitempty"`
}

// UpgradeResult is the response from an Upgrade operation.
type UpgradeResult struct {
	nodegroupResult
}
