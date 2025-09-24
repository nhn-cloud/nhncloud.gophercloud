package clusters

import (
	"encoding/json"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Cluster represents a Kubernetes cluster in NHN Cloud.
type Cluster struct {
	// APIAddress is the endpoint where the kubernetes API is exposed.
	APIAddress string `json:"api_address"`

	// COEVersion is the version of the Container Orchestration Engine.
	COEVersion string `json:"coe_version"`

	// ClusterTemplateID is the UUID of the cluster template.
	ClusterTemplateID string `json:"cluster_template_id"`

	// ContainerVersion is the version of the container runtime.
	ContainerVersion string `json:"container_version"`

	// CreateTimeout is the timeout for cluster creation in minutes.
	CreateTimeout int `json:"create_timeout"`

	// CreatedAt is the date and time when the resource was created.
	CreatedAt time.Time `json:"created_at"`

	// DiscoveryURL is the URL used for cluster node discovery.
	DiscoveryURL string `json:"discovery_url"`

	// DockerVolumeSize is the size of the volume to allocate to docker.
	DockerVolumeSize int `json:"docker_volume_size"`

	// FixedNetwork is the network where cluster gets fixed IP from.
	FixedNetwork string `json:"fixed_network"`

	// FixedSubnet is the subnet where cluster gets fixed IP from.
	FixedSubnet string `json:"fixed_subnet"`

	// FlavorID is the flavor ID used for the cluster worker nodes.
	FlavorID string `json:"flavor_id"`

	// FloatingIPEnabled indicates whether cluster nodes have floating IP.
	FloatingIPEnabled bool `json:"floating_ip_enabled"`

	// KeyPair is the name of the SSH keypair to configure in cluster servers.
	KeyPair string `json:"keypair"`

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

	// MasterAddresses contains a list of IP addresses of master nodes.
	MasterAddresses []string `json:"master_addresses"`

	// MasterCount is the number of master nodes.
	MasterCount int `json:"master_count"`

	// MasterFlavorID is the flavor ID used for the cluster master nodes.
	MasterFlavorID string `json:"master_flavor_id"`

	// Name is the cluster name.
	Name string `json:"name"`

	// NodeAddresses contains a list of IP addresses of worker nodes.
	NodeAddresses []string `json:"node_addresses"`

	// NodeCount is the number of worker nodes.
	NodeCount int `json:"node_count"`

	// ProjectID is the project ID where cluster is created.
	ProjectID string `json:"project_id"`

	// StackID is the UUID of the heat stack used to create the cluster.
	StackID string `json:"stack_id"`

	// Status indicates the status of the cluster.
	Status string `json:"status"`

	// StatusReason provides additional details about the status.
	StatusReason string `json:"status_reason"`

	// UpdatedAt is the date and time when the resource was updated.
	UpdatedAt time.Time `json:"updated_at"`

	// UserID is the user ID that created the cluster.
	UserID string `json:"user_id"`

	// UUID is the UUID of the cluster.
	UUID string `json:"uuid"`

	// NKS specific fields
	// APIEndpointIPACL contains API endpoint IP access control settings.
	APIEndpointIPACL *APIEndpointIPACL `json:"api_ep_ipacl,omitempty"`

	// Addons contains addon configurations.
	Addons []Addon `json:"addons,omitempty"`
}

// ClusterPage stores a single page of all Cluster results from a List call.
type ClusterPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a ClusterPage is empty.
func (page ClusterPage) IsEmpty() (bool, error) {
	if page.StatusCode == 204 {
		return true, nil
	}

	clusters, err := ExtractClusters(page)
	return len(clusters) == 0, err
}

// NextPageURL extracts the "next" link from the links section of the result.
func (page ClusterPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"clusters_links"`
	}
	err := page.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
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
	var cluster Cluster
	err := r.ExtractInto(&cluster)
	return &cluster, err
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

// CreateResult is the response from a Create operation.
type CreateResult struct {
	clusterResult
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	clusterResult
}

// UpdateResult is the response from an Update operation.
type UpdateResult struct {
	clusterResult
}

// DeleteResult is the response from a Delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

// ResizeResult is the response from a Resize operation.
type ResizeResult struct {
	clusterResult
}

// UpgradeResult is the response from an Upgrade operation.
type UpgradeResult struct {
	clusterResult
}
