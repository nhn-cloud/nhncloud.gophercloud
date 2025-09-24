/*
Package nodegroups provides functionality to interact with Kubernetes cluster
nodegroups through the Kubernetes service API.

This package is based on gophercloud's containerinfra nodegroups package and
extends it with NHN Cloud specific functionality.

Example to Create a Nodegroup

	createOpts := nodegroups.CreateOpts{
		Name:         "default-worker",
		FlavorID:     "m2.c2m4",
		ImageID:      "kubernetes-v1.23.3-ubuntu-18.04",
		NodeCount:    3,
	}

	nodegroup, err := nodegroups.Create(kubernetesClient, "cluster-id", createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to List Nodegroups

	listOpts := nodegroups.ListOpts{
		Limit: 10,
	}

	allPages, err := nodegroups.List(kubernetesClient, "cluster-id", listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allNodegroups, err := nodegroups.ExtractNodegroups(allPages)
	if err != nil {
		panic(err)
	}

	for _, nodegroup := range allNodegroups {
		fmt.Printf("%+v\n", nodegroup)
	}

Example to Get a Nodegroup

	nodegroup, err := nodegroups.Get(kubernetesClient, "cluster-id", "nodegroup-id").Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Nodegroup

	err := nodegroups.Delete(kubernetesClient, "cluster-id", "nodegroup-id").ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package nodegroups
