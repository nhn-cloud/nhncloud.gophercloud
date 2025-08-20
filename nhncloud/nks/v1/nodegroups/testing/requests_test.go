package testing

import (
	"context"
	"testing"

	"github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/nhn-cloud/nhncloud.gophercloud/nhncloud/nks/v1/nodegroups"
)

func TestListNodegroups(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()
	HandleNodegroupListSuccessfully(t)

	opts := nodegroups.ListOpts{}

	count := 0
	err := nodegroups.List(client.ServiceClient(), "cluster-uuid-1", opts).EachPage(context.TODO(), func(page interface{}) (bool, error) {
		count++
		actual, err := nodegroups.ExtractNodegroups(page.(nodegroups.NodegroupPage))
		if err != nil {
			t.Errorf("Failed to extract nodegroups: %v", err)
			return false, err
		}

		if len(actual) != 1 {
			t.Errorf("Expected 1 nodegroup, got %d", len(actual))
		}

		if actual[0].Name != "default-worker" {
			t.Errorf("Expected name 'default-worker', got '%s'", actual[0].Name)
		}

		return true, nil
	})

	if err != nil {
		t.Errorf("Failed to list nodegroups: %v", err)
	}

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetNodegroup(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()
	HandleNodegroupGetSuccessfully(t)

	actual, err := nodegroups.Get(client.ServiceClient(), "cluster-uuid-1", "nodegroup-uuid-1").Extract()
	if err != nil {
		t.Errorf("Failed to get nodegroup: %v", err)
	}

	if actual.Name != "default-worker" {
		t.Errorf("Expected name 'default-worker', got '%s'", actual.Name)
	}

	if actual.UUID != "nodegroup-uuid-1" {
		t.Errorf("Expected UUID 'nodegroup-uuid-1', got '%s'", actual.UUID)
	}
}

func TestCreateNodegroup(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()
	HandleNodegroupCreateSuccessfully(t)

	opts := nodegroups.CreateOpts{
		Name:      "test-nodegroup",
		NodeCount: 2,
		FlavorID:  "flavor-uuid-1",
		ImageID:   "image-uuid-1",
		Labels: map[string]interface{}{
			"availability_zone": "kr2-pub-a",
			"boot_volume_size":  "50",
			"boot_volume_type":  "General HDD",
		},
	}

	actual, err := nodegroups.Create(client.ServiceClient(), "cluster-uuid-1", opts).Extract()
	if err != nil {
		t.Errorf("Failed to create nodegroup: %v", err)
	}

	if actual.UUID != "nodegroup-uuid-2" {
		t.Errorf("Expected UUID 'nodegroup-uuid-2', got '%s'", actual.UUID)
	}
}

func TestDeleteNodegroup(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()
	HandleNodegroupDeleteSuccessfully(t)

	err := nodegroups.Delete(client.ServiceClient(), "cluster-uuid-1", "nodegroup-uuid-1").ExtractErr()
	if err != nil {
		t.Errorf("Failed to delete nodegroup: %v", err)
	}
}
