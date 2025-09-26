package nodegroups

import (
	"encoding/json"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Nodegroup represents a Kubernetes nodegroup in NHN Cloud.
type Nodegroup struct {
	// ClusterID is the UUID of the cluster this nodegroup belongs to.
	ClusterID string `json:"cluster_id"`

	// CreatedAt is the date and time when the resource was created.
	CreatedAt time.Time `json:"created_at"`

	// DockerVolumeSize is the size of the volume to allocate to docker.
	DockerVolumeSize int `json:"docker_volume_size"`

	// FlavorID is the flavor ID used for the nodegroup.
	FlavorID string `json:"flavor_id"`

	// ID is the UUID of the nodegroup.
	ID int `json:"id"`

	// ImageID is the image ID used for the nodegroup.
	ImageID string `json:"image_id"`

	// IsDefault indicates whether this is the default nodegroup.
	IsDefault bool `json:"is_default"`

	// Labels is a set of key=value pairs.
	Labels map[string]string `json:"labels"`

	// LabelsAdded contains labels that were added compared to cluster template.
	LabelsAdded map[string]string `json:"labels_added"`

	// LabelsOverridden contains labels that were overridden from cluster template.
	LabelsOverridden map[string]string `json:"labels_overridden"`

	// LabelsSkipped contains labels that were skipped from cluster template.
	LabelsSkipped map[string]string `json:"labels_skipped"`

	// Links contains a list of links related to this resource.
	Links []gophercloud.Link `json:"links"`

	// MaxNodeCount is the maximum allowed number of nodes.
	MaxNodeCount int `json:"max_node_count"`

	// MinNodeCount is the minimum allowed number of nodes.
	MinNodeCount int `json:"min_node_count"`

	// Name is the nodegroup name.
	Name string `json:"name"`

	// NodeAddresses contains a list of IP addresses of nodes in this nodegroup.
	NodeAddresses []string `json:"node_addresses"`

	// NodeCount is the number of nodes in this nodegroup.
	NodeCount int `json:"node_count"`

	// ProjectID is the project ID where nodegroup is created.
	ProjectID string `json:"project_id"`

	// Role is the role of the nodegroup.
	Role string `json:"role"`

	// StackID is the UUID of the heat stack used to create the nodegroup.
	StackID string `json:"stack_id"`

	// Status indicates the status of the nodegroup.
	Status string `json:"status"`

	// StatusReason provides additional details about the status.
	StatusReason string `json:"status_reason"`

	// UpdatedAt is the date and time when the resource was updated.
	UpdatedAt time.Time `json:"updated_at"`

	// UUID is the UUID of the nodegroup.
	UUID string `json:"uuid"`

	// Version is the Kubernetes version of the nodegroup.
	Version string `json:"version"`
}

// NodegroupPage stores a single page of all Nodegroup results from a List call.
type NodegroupPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a NodegroupPage is empty.
func (page NodegroupPage) IsEmpty() (bool, error) {
	if page.StatusCode == 204 {
		return true, nil
	}

	nodegroups, err := ExtractNodegroups(page)
	return len(nodegroups) == 0, err
}

// NextPageURL extracts the "next" link from the links section of the result.
func (page NodegroupPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"nodegroups_links"`
	}
	err := page.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
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
	var nodegroup Nodegroup
	err := r.ExtractInto(&nodegroup)
	return &nodegroup, err
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

// CreateResult is the response from a Create operation.
type CreateResult struct {
	nodegroupResult
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	nodegroupResult
}

// UpdateResult is the response from an Update operation.
type UpdateResult struct {
	nodegroupResult
}

// DeleteResult is the response from a Delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

// UpgradeResult is the response from an Upgrade operation.
type UpgradeResult struct {
	nodegroupResult
}
