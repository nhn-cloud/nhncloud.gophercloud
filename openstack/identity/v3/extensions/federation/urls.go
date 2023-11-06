package federation

import gophercloud "github.com/nhn/nhncloud.gophercloud"

const (
	rootPath     = "OS-FEDERATION"
	mappingsPath = "mappings"
)

func mappingsRootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath, mappingsPath)
}

func mappingsResourceURL(c *gophercloud.ServiceClient, mappingID string) string {
	return c.ServiceURL(rootPath, mappingsPath, mappingID)
}
