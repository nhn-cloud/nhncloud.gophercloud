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

func connectInterfaceURL(c *gophercloud.ServiceClient, volumeID string) string {
	return c.ServiceURL("v1", "volumes", volumeID, "interfaces")
}

func deleteInterfaceURL(c *gophercloud.ServiceClient, volumeID string, interfaceID string) string {
	return c.ServiceURL("v1", "volumes", volumeID, "interfaces", interfaceID)
}

func setReplicationURL(c *gophercloud.ServiceClient, volumeID string) string {
	return c.ServiceURL("v1", "volumes", volumeID, "volume-mirrors")
}

func disableReplicationURL(c *gophercloud.ServiceClient, volumeID string, mirrorID string) string {
	return c.ServiceURL("v1", "volumes", volumeID, "volume-mirrors", mirrorID)
}

func getReplicationStatURL(c *gophercloud.ServiceClient, volumeID string, mirrorID string) string {
	return c.ServiceURL("v1", "volumes", volumeID, "volume-mirrors", mirrorID, "stat")
}
