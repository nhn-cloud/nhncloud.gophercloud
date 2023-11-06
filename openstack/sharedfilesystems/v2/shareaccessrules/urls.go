package shareaccessrules

import (
	"fmt"

	gophercloud "github.com/nhn/nhncloud.gophercloud"
)

const shareAccessRulesEndpoint = "share-access-rules"

func getURL(c *gophercloud.ServiceClient, accessID string) string {
	return c.ServiceURL(shareAccessRulesEndpoint, accessID)
}

func listURL(c *gophercloud.ServiceClient, shareID string) string {
	return fmt.Sprintf("%s?share_id=%s", c.ServiceURL(shareAccessRulesEndpoint), shareID)
}
