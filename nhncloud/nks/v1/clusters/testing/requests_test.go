package testing

import (
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"

	"github.com/nhn-cloud/nhncloud.gophercloud/nhncloud/nks/v1/clusters"
)

func TestGetCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/clusters/abcd1234-efgh-5678-ijkl-9012mnop3456", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(GetClusterResponse))
	})

	client, err := openstack.NewComputeV2(client.ServiceClient(), gophercloud.EndpointOpts{})
	th.AssertNoErr(t, err)

	cluster, err := clusters.Get(client, "abcd1234-efgh-5678-ijkl-9012mnop3456").Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, "abcd1234-efgh-5678-ijkl-9012mnop3456", cluster.UUID)
}
