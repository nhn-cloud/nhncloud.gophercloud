package testing

import (
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/nhn-cloud/nhncloud.gophercloud/nhncloud/networking/v2/internetgateways"
)

func TestCreateInternetGateway(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleCreateInternetGatewaySuccessfully(t)

	res := internetgateways.Create(fake.ServiceClient(), internetgateways.CreateOpts{Name: "sample", ExternalNetworkID: "sample"})
	th.AssertNoErr(t, res.Err)

	// The golden request body is asserted by the mock handler (TestJSONRequest).
	_, err := res.Extract()
	th.AssertNoErr(t, err)
}

func TestGetInternetGateway(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleGetInternetGatewaySuccessfully(t)

	res := internetgateways.Get(fake.ServiceClient(), "71f8e19a-8af3-4d3f-9b0e-b45a547fd7a7")
	th.AssertNoErr(t, res.Err)

	_, err := res.Extract()
	th.AssertNoErr(t, err)
}

func TestListInternetGateway(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleListInternetGatewaySuccessfully(t)

	pager := internetgateways.List(fake.ServiceClient(), internetgateways.ListOpts{})
	pages, err := pager.AllPages()
	th.AssertNoErr(t, err)

	_, err = internetgateways.ExtractInternetGateways(pages)
	th.AssertNoErr(t, err)
}

func TestDeleteInternetGateway(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleDeleteInternetGatewaySuccessfully(t)

	res := internetgateways.Delete(fake.ServiceClient(), "71f8e19a-8af3-4d3f-9b0e-b45a547fd7a7")
	th.AssertNoErr(t, res.Err)
}
