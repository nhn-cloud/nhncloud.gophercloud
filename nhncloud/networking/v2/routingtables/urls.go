package routingtables

import "github.com/gophercloud/gophercloud"

func resourceURL(c *gophercloud.ServiceClient, id string, action ...string) string {
	return c.ServiceURL("routingtables", id)
}

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("routingtables")
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func createURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func attachGatewayURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("routingtables", id, "attach_gateway")
}

func detachGatewayURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("routingtables", id, "detach_gateway")
}
