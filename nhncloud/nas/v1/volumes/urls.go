package volumes

import "github.com/gophercloud/gophercloud"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("v1", "volumes")
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("v1", "volumes", id)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("v1", "volumes", id)
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("v1", "volumes", id)
}

func connectInterfaceURL(c *gophercloud.ServiceClient, volmeID string) string {
	return c.ServiceURL("v1", "volumes", volmeID, "interfaces")
}

func deleteInterfaceURL(c *gophercloud.ServiceClient, volmeID string, interfaceID string) string {
	return c.ServiceURL("v1", "volumes", volmeID, "interfaces", interfaceID)
}

func setReplicationURL(c *gophercloud.ServiceClient, volmeID string) string {
	return c.ServiceURL("v1", "volumes", volmeID, "volume-mirrors")
}

func disableReplicationURL(c *gophercloud.ServiceClient, volmeID string, mirrorID string) string {
	return c.ServiceURL("v1", "volumes", volmeID, "volume-mirrors", mirrorID)
}

func getReplicationStatURL(c *gophercloud.ServiceClient, volmeID string, mirrorID string) string {
	return c.ServiceURL("v1", "volumes", volmeID, "volume-mirrors", mirrorID, "stat")
}
