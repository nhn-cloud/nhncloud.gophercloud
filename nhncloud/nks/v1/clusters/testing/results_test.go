package testing

import (
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"

	"github.com/nhncloud/nhncloud.gophercloud/nhncloud/kubernetes/v1/clusters"
)

func TestExtractCluster(t *testing.T) {
	cluster, err := clusters.ExtractCluster([]byte(GetClusterResponse))
	th.AssertNoErr(t, err)

	th.AssertEquals(t, "abcd1234-efgh-5678-ijkl-9012mnop3456", cluster.UUID)
}

func TestExtractClusterInto(t *testing.T) {
	var cluster clusters.Cluster
	err := clusters.ExtractClusterInto([]byte(GetClusterResponse), &cluster)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, "abcd1234-efgh-5678-ijkl-9012mnop3456", cluster.UUID)
}
