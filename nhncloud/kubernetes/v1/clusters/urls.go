package clusters

import "github.com/gophercloud/gophercloud"

const resourcePath = "clusters"

func resourceURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func listURL(c *gophercloud.ServiceClient) string {
	return resourceURL(c)
}

func listDetailURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("clusters", "detail")
}

func createURL(c *gophercloud.ServiceClient) string {
	return resourceURL(c)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}

func resizeURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id, "actions", "resize")
}

func upgradeURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id, "actions", "upgrade")
}
