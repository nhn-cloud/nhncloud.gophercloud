//go:build acceptance || networking || loadbalancer || listeners
// +build acceptance networking loadbalancer listeners

package v2

import (
	"testing"

	"github.com/nhn/nhncloud.gophercloud/acceptance/clients"
	"github.com/nhn/nhncloud.gophercloud/acceptance/tools"
	"github.com/nhn/nhncloud.gophercloud/openstack/loadbalancer/v2/listeners"
)

func TestListenersList(t *testing.T) {
	client, err := clients.NewLoadBalancerV2Client()
	if err != nil {
		t.Fatalf("Unable to create a loadbalancer client: %v", err)
	}

	allPages, err := listeners.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list listeners: %v", err)
	}

	allListeners, err := listeners.ExtractListeners(allPages)
	if err != nil {
		t.Fatalf("Unable to extract listeners: %v", err)
	}

	for _, listener := range allListeners {
		tools.PrintResource(t, listener)
	}
}
