package volumes

import "github.com/gophercloud/gophercloud"

func createVolumeURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("v1", "volumes")
}

func getVolumeURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("v1", "volumes", id)
}

func updateVolumeURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("v1", "volumes", id)
}

func deleteVolumeURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("v1", "volumes", id)
}

func createInterfaceURL(c *gophercloud.ServiceClient, volmeID string) string {
	return c.ServiceURL("v1", "volumes", volmeID, "interfaces")
}

func deleteInterfaceURL(c *gophercloud.ServiceClient, volmeID string, interfaceID string) string {
	return c.ServiceURL("v1", "volumes", volmeID, "interfaces", interfaceID)
}

func createMirrorURL(c *gophercloud.ServiceClient, volmeID string) string {
	return c.ServiceURL("v1", "volumes", volmeID, "volume-mirrors")
}

func deleteMirrorURL(c *gophercloud.ServiceClient, volmeID string, mirrorID string) string {
	return c.ServiceURL("v1", "volumes", volmeID, "volume-mirrors", mirrorID)
}

func getMirrorStatURL(c *gophercloud.ServiceClient, volmeID string, mirrorID string) string {
	return c.ServiceURL("v1", "volumes", volmeID, "volume-mirrors", mirrorID, "stat")
}
