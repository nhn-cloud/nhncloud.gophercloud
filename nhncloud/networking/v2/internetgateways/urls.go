package internetgateways

import "github.com/gophercloud/gophercloud"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("internetgateways")
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("internetgateways", id)
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

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
