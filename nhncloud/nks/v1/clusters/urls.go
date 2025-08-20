package clusters

import "github.com/gophercloud/gophercloud"

const resourcePath = "clusters"

func resourceURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func listURL(c *gophercloud.ServiceClient) string {
	return resourceURL(c)
}

func createURL(c *gophercloud.ServiceClient) string {
	return resourceURL(c)
}

func getURL(c *gophercloud.ServiceClient, clusterIDOrName string) string {
	return c.ServiceURL(resourcePath, clusterIDOrName)
}

func deleteURL(c *gophercloud.ServiceClient, clusterIDOrName string) string {
	return getURL(c, clusterIDOrName)
}

func resizeURL(c *gophercloud.ServiceClient, clusterIDOrName string) string {
	return c.ServiceURL(resourcePath, clusterIDOrName, "actions", "resize")
}
