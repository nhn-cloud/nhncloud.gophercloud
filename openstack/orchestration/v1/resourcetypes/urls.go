package resourcetypes

import gophercloud "github.com/nhn/nhncloud.gophercloud"

const (
	resTypesPath = "resource_types"
)

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resTypesPath)
}

func getSchemaURL(c *gophercloud.ServiceClient, resourceType string) string {
	return c.ServiceURL(resTypesPath, resourceType)
}

func generateTemplateURL(c *gophercloud.ServiceClient, resourceType string) string {
	return c.ServiceURL(resTypesPath, resourceType, "template")
}
