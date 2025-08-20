package nodegroups

import "github.com/gophercloud/gophercloud"

const clustersPath = "clusters"
const resourcePath = "nodegroups"

func resourceURL(c *gophercloud.ServiceClient, clusterIDOrName string) string {
	return c.ServiceURL(clustersPath, clusterIDOrName, resourcePath)
}

func listURL(c *gophercloud.ServiceClient, clusterIDOrName string) string {
	return resourceURL(c, clusterIDOrName)
}

func createURL(c *gophercloud.ServiceClient, clusterIDOrName string) string {
	return resourceURL(c, clusterIDOrName)
}

func getURL(c *gophercloud.ServiceClient, clusterIDOrName, nodegroupIDOrName string) string {
	return c.ServiceURL(clustersPath, clusterIDOrName, resourcePath, nodegroupIDOrName)
}

func deleteURL(c *gophercloud.ServiceClient, clusterIDOrName, nodegroupIDOrName string) string {
	return getURL(c, clusterIDOrName, nodegroupIDOrName)
}

func upgradeURL(c *gophercloud.ServiceClient, clusterIDOrName, nodegroupIDOrName string) string {
	return c.ServiceURL(clustersPath, clusterIDOrName, resourcePath, nodegroupIDOrName, "upgrade")
}
