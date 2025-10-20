package nodegroups

import "github.com/gophercloud/gophercloud"

const resourcePath = "clusters"
const nodegroupsPath = "nodegroups"

func resourceURL(c *gophercloud.ServiceClient, clusterID string) string {
	return c.ServiceURL(resourcePath, clusterID, nodegroupsPath)
}

func listURL(c *gophercloud.ServiceClient, clusterID string) string {
	return resourceURL(c, clusterID)
}

func createURL(c *gophercloud.ServiceClient, clusterID string) string {
	return resourceURL(c, clusterID)
}

func getURL(c *gophercloud.ServiceClient, clusterID, nodegroupID string) string {
	return c.ServiceURL(resourcePath, clusterID, nodegroupsPath, nodegroupID)
}

func updateURL(c *gophercloud.ServiceClient, clusterID, nodegroupID string) string {
	return c.ServiceURL(resourcePath, clusterID, nodegroupsPath, nodegroupID)
}

func deleteURL(c *gophercloud.ServiceClient, clusterID, nodegroupID string) string {
	return c.ServiceURL(resourcePath, clusterID, nodegroupsPath, nodegroupID)
}
func upgradeURL(c *gophercloud.ServiceClient, clusterID, nodegroupID string) string {
	return c.ServiceURL(resourcePath, clusterID, nodegroupsPath, nodegroupID, "upgrade")
}
