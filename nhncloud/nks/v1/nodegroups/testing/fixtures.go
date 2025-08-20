package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

// NodegroupListBody contains the canned body of a nodegroup list response.
const NodegroupListBody = `
{
	"nodegroups": [
		{
			"uuid": "nodegroup-uuid-1",
			"name": "default-worker",
			"cluster_id": "cluster-uuid-1",
			"node_count": 3,
			"flavor_id": "flavor-uuid-1",
			"image_id": "image-uuid-1",
			"status": "CREATE_COMPLETE",
			"created_at": "2024-01-01T00:00:00Z",
			"updated_at": "2024-01-01T00:00:00Z",
			"version": "v1.32.3",
			"labels": {
				"availability_zone": "kr2-pub-a",
				"boot_volume_size": "50",
				"boot_volume_type": "General HDD"
			}
		}
	]
}
`

// NodegroupGetBody contains the canned body of a nodegroup get response.
const NodegroupGetBody = `
{
	"nodegroup": {
		"uuid": "nodegroup-uuid-1",
		"name": "default-worker",
		"cluster_id": "cluster-uuid-1",
		"node_count": 3,
		"flavor_id": "flavor-uuid-1",
		"image_id": "image-uuid-1",
		"status": "CREATE_COMPLETE",
		"created_at": "2024-01-01T00:00:00Z",
		"updated_at": "2024-01-01T00:00:00Z",
		"version": "v1.32.3",
		"labels": {
			"availability_zone": "kr2-pub-a",
			"boot_volume_size": "50",
			"boot_volume_type": "General HDD"
		}
	}
}
`

// NodegroupCreateBody contains the canned body of a nodegroup create response.
const NodegroupCreateBody = `
{
	"nodegroup": {
		"uuid": "nodegroup-uuid-2"
	}
}
`

// HandleNodegroupListSuccessfully sets up the test server to respond to a nodegroup List request.
func HandleNodegroupListSuccessfully(t *testing.T) {
	testhelper.Mux.HandleFunc("/clusters/cluster-uuid-1/nodegroups", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "GET")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, NodegroupListBody)
	})
}

// HandleNodegroupGetSuccessfully sets up the test server to respond to a nodegroup Get request.
func HandleNodegroupGetSuccessfully(t *testing.T) {
	testhelper.Mux.HandleFunc("/clusters/cluster-uuid-1/nodegroups/nodegroup-uuid-1", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "GET")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, NodegroupGetBody)
	})
}

// HandleNodegroupCreateSuccessfully sets up the test server to respond to a nodegroup Create request.
func HandleNodegroupCreateSuccessfully(t *testing.T) {
	testhelper.Mux.HandleFunc("/clusters/cluster-uuid-1/nodegroups", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "POST")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		testhelper.TestHeader(t, r, "Content-Type", "application/json")
		testhelper.TestHeader(t, r, "Accept", "application/json")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, NodegroupCreateBody)
	})
}

// HandleNodegroupDeleteSuccessfully sets up the test server to respond to a nodegroup Delete request.
func HandleNodegroupDeleteSuccessfully(t *testing.T) {
	testhelper.Mux.HandleFunc("/clusters/cluster-uuid-1/nodegroups/nodegroup-uuid-1", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "DELETE")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})
}
