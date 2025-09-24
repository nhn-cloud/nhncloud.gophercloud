/*
Package clusters provides functionality to interact with Kubernetes clusters
through the Kubernetes service API.

This package is based on gophercloud's containerinfra clusters package and
extends it with NHN Cloud specific functionality.

Example to Create a Cluster

	createOpts := clusters.CreateOpts{
		ClusterTemplateID: "0562d357-8641-4759-8fed-8173f02c9633",
		Keypair:           "my-keypair",
		MasterCount:       1,
		Name:              "k8s-cluster",
		NodeCount:         1,
	}

	cluster, err := clusters.Create(kubernetesClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to List Clusters

	listOpts := clusters.ListOpts{
		Limit: 10,
	}

	allPages, err := clusters.List(kubernetesClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allClusters, err := clusters.ExtractClusters(allPages)
	if err != nil {
		panic(err)
	}

	for _, cluster := range allClusters {
		fmt.Printf("%+v\n", cluster)
	}

Example to Get a Cluster

	cluster, err := clusters.Get(kubernetesClient, "cluster-uuid").Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Cluster

	err := clusters.Delete(kubernetesClient, "cluster-uuid").ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package clusters
