/*
Package clusters provides the ability to manage NKS clusters
through the NHN Cloud API.

# Example of Listing NKS Clusters

This example will list all NKS clusters:

	allPages, err := clusters.List(client, nil).AllPages(context.TODO())
	if err != nil {
		panic(err)
	}

	allClusters, err := clusters.ExtractClusters(allPages)
	if err != nil {
		panic(err)
	}

	for _, cluster := range allClusters {
		fmt.Printf("Cluster: %+v\n", cluster)
	}

# Example of Getting an NKS Cluster

This example will retrieve an NKS cluster by its UUID:

	cluster, err := clusters.Get(client, "abcd1234-efgh-5678-ijkl-9012mnop3456").Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Cluster UUID: %s\n", cluster.UUID)

# Example of Creating an NKS Cluster

This example will create a new NKS cluster:

	createOpts := clusters.CreateOpts{
		Name:              "test-cluster",
		ClusterTemplateID: "iaas_console",
		FixedNetwork:      "65e08bba-a7d5-4014-b28f-8f4fdaac001e",
		FixedSubnet:       "3a50e104-14d2-47a5-97a0-8ec87af8599e",
		FlavorID:          "b71c2d4e-31e4-4d0e-ac2f-f057ec4b6d71",
		Keypair:           "key-alpha",
		NodeCount:         2,
		Labels: map[string]interface{}{
			"kube_tag":                     "v1.32.3",
			"availability_zone":            "kr2-pub-a",
			"boot_volume_size":             "50",
			"boot_volume_type":             "General HDD",
			"ca_enable":                    "false",
			"cert_manager_api":             "True",
			"external_network_id":          "751b8227-1234-5678-9349-dbf829d0aba5",
			"external_subnet_id_list":      "59ddc195-1234-5678-9693-f09880747dc6",
			"master_lb_floating_ip_enabled": "true",
			"node_image":                   "1a10bf47-2f28-1234-5678-e2dc43f61789",
		},
		Addons: []clusters.Addon{
			{
				Name:    "calico",
				Version: "v3.28.2-nks1",
				Options: map[string]interface{}{
					"mode": "ebpf",
				},
			},
			{
				Name:    "coredns",
				Version: "1.8.4-nks1",
			},
		},
	}

	cluster, err := clusters.Create(client, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created cluster: %+v\n", cluster)

# Example of Deleting an NKS Cluster

This example will delete an NKS cluster:

	err := clusters.Delete(client, "abcd1234-efgh-5678-ijkl-9012mnop3456").ExtractErr()
	if err != nil {
		panic(err)
	}

# Example of Resizing an NKS Cluster

This example will resize an NKS cluster:

	resizeOpts := clusters.ResizeOpts{
		NodeCount: 3,
		Options: &clusters.ResizeOptions{
			NodesToRemove: []string{},
		},
	}

	cluster, err := clusters.Resize(client, "abcd1234-efgh-5678-ijkl-9012mnop3456", resizeOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Resized cluster: %+v\n", cluster)
*/
package clusters
