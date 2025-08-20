package clusters

import (
	"encoding/json"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Cluster represents an NKS cluster in NHN Cloud.
type Cluster struct {
	// UUID is the unique identifier for the NKS cluster.
	UUID string `json:"uuid"`

	// Name is the name of the cluster.
	Name string `json:"name"`

	// ClusterTemplateID is the template ID used for the cluster.
	ClusterTemplateID string `json:"cluster_template_id"`

	// FixedNetwork is the VPC network UUID.
	FixedNetwork string `json:"fixed_network"`

	// FixedSubnet is the VPC subnet UUID.
	FixedSubnet string `json:"fixed_subnet"`

	// FlavorID is the instance flavor UUID.
	FlavorID string `json:"flavor_id"`

	// Keypair is the SSH keypair name.
	Keypair string `json:"keypair"`

	// NodeCount is the number of worker nodes.
	NodeCount int `json:"node_count"`

	// Labels contains cluster configuration labels.
	Labels map[string]interface{} `json:"labels,omitempty"`

	// APIEndpointIPACL contains API endpoint IP access control settings.
	APIEndpointIPACL *APIEndpointIPACL `json:"api_ep_ipacl,omitempty"`

	// Addons contains addon configurations.
	Addons []Addon `json:"addons,omitempty"`

	// Status is the current status of the cluster.
	Status string `json:"status,omitempty"`

	// StatusReason provides additional information about the status.
	StatusReason string `json:"status_reason,omitempty"`

	// CreatedAt is the creation timestamp.
	CreatedAt string `json:"created_at,omitempty"`

	// UpdatedAt is the last update timestamp.
	UpdatedAt string `json:"updated_at,omitempty"`
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

// ClusterPage stores a single page of all Cluster results from a List call.
type ClusterPage struct {
	pagination.SinglePageBase
}

// IsEmpty determines whether or not a ClusterPage is empty.
func (page ClusterPage) IsEmpty() (bool, error) {
	if page.StatusCode == 204 {
		return true, nil
	}

	clusters, err := ExtractClusters(page)
	return len(clusters) == 0, err
}

// ExtractClusters interprets a page of results as a slice of Clusters.
func ExtractClusters(r pagination.Page) ([]Cluster, error) {
	var s struct {
		Clusters []Cluster `json:"clusters"`
	}
	err := (r.(ClusterPage)).ExtractInto(&s)
	return s.Clusters, err
}

type clusterResult struct {
	gophercloud.Result
}

// Extract will get the Cluster object out of the clusterResult object.
func (r clusterResult) Extract() (*Cluster, error) {
	var s struct {
		Cluster *Cluster `json:"cluster"`
	}
	err := r.ExtractInto(&s)
	return s.Cluster, err
}

// ExtractCluster extracts a cluster from a raw response.
func ExtractCluster(raw []byte) (*Cluster, error) {
	var response struct {
		Cluster Cluster `json:"cluster"`
	}
	err := json.Unmarshal(raw, &response)
	if err != nil {
		return nil, err
	}
	return &response.Cluster, nil
}

// ExtractClusterInto extracts a cluster into a provided interface.
func ExtractClusterInto(raw []byte, v interface{}) error {
	var response struct {
		Cluster interface{} `json:"cluster"`
	}
	response.Cluster = v
	return json.Unmarshal(raw, &response)
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	clusterResult
}

// CreateResult is the response from a Create operation.
type CreateResult struct {
	clusterResult
}

// DeleteResult is the response from a Delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

// ResizeOptions contains options for cluster resize operation.
type ResizeOptions struct {
	// NodesToRemove contains the list of node UUIDs to remove.
	NodesToRemove []string `json:"nodes_to_remove,omitempty"`
}

// ResizeResult is the response from a Resize operation.
type ResizeResult struct {
	clusterResult
}
