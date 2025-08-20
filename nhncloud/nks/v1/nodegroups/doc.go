/*
Package nodegroups provides the ability to manage NKS nodegroups
through the NHN Cloud API.

# Example of Listing NKS Nodegroups

This example will list all nodegroups in an NKS cluster:

	listOpts := nodegroups.ListOpts{}

	allPages, err := nodegroups.List(client, "abcd1234-efgh-5678-ijkl-9012mnop3456", listOpts).AllPages(context.TODO())
	if err != nil {
		panic(err)
	}

	allNodegroups, err := nodegroups.ExtractNodegroups(allPages)
	if err != nil {
		panic(err)
	}

	for _, nodegroup := range allNodegroups {
		fmt.Printf("Nodegroup: %+v\n", nodegroup)
	}

# Example of Getting an NKS Nodegroup

This example will retrieve an NKS nodegroup by its UUID:

	nodegroup, err := nodegroups.Get(client, "abcd1234-efgh-5678-ijkl-9012mnop3456", "nodegroup-uuid").Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Nodegroup UUID: %s\n", nodegroup.UUID)

# Example of Creating an NKS Nodegroup

This example will create a new NKS nodegroup:

	createOpts := nodegroups.CreateOpts{
		Name:      "add-nodegroup",
		NodeCount: 2,
		FlavorID:  "b71c2d4e-1234-5678-ac2f-f057ec4b6d71",
		ImageID:   "1a10bf47-1234-5678-982b-e2dc43f61789",
		Labels: map[string]interface{}{
			"availability_zone":            "kr2-pub-a",
			"boot_volume_size":             "50",
			"boot_volume_type":             "General HDD",
			"ca_enable":                    "true",
			"ca_max_node_count":            "9",
			"ca_min_node_count":            "1",
			"ca_scale_down_enable":         "true",
		},
	}

	nodegroup, err := nodegroups.Create(client, "abcd1234-efgh-5678-ijkl-9012mnop3456", createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created nodegroup: %+v\n", nodegroup)

# Example of Deleting an NKS Nodegroup

This example will delete an NKS nodegroup:

	err := nodegroups.Delete(client, "abcd1234-efgh-5678-ijkl-9012mnop3456", "nodegroup-uuid").ExtractErr()
	if err != nil {
		panic(err)
	}

# Example of Upgrading an NKS Nodegroup

This example will upgrade an NKS nodegroup:

	upgradeOpts := nodegroups.UpgradeOpts{
		Version: "v1.32.3",
		Options: &nodegroups.UpgradeOptions{
			NumBufferNodes:         1,
			NumMaxUnavailableNodes: 1,
		},
	}

	nodegroup, err := nodegroups.Upgrade(client, "abcd1234-efgh-5678-ijkl-9012mnop3456", "nodegroup-uuid", upgradeOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Upgraded nodegroup: %+v\n", nodegroup)
*/
package nodegroups
