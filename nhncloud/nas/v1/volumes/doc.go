/*
Package volumes provides information and interaction with NAS storage volumes, interfaces and mirrors
through the NAS service API.

Exmaple to Create a Volume

	createOpts := volumes.CreateVolumeOpts{
		Name: "test-volume",
		SizeGb: 100,
		MountProtocol: &volumes.MountProtocolOpts{
			Protocol: "nfs",
		},
	}
	volume, err := volumes.CreateVolume(serviceClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Get a Volume

	volumeID := "c79a4468-d788-410c-bf79-9a8ef6354852"
	volume, err := volumes.GetVolume(serviceClient, volumeID).Extract()
	if err != nil {
		panic(err)
	}

Example to Update a Volume

	volumeID := "c79a4468-d788-410c-bf79-9a8ef6354852"
	updateOpts := volumes.UpdateVolumeOpts{
		Description: "test-volume",
	}
	volume, err := volumes.UpdateVolume(serviceClient, volumeID, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Volume

	volumeID := "c79a4468-d788-410c-bf79-9a8ef6354852"
	err := volumes.DeleteVolume(serviceClient, volumeID).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Create an Interface

	volumeID := "c79a4468-d788-410c-bf79-9a8ef6354852"
	createOpts := volumes.CreateInterfaceOpts{
		SubnetID: "84f1b61f-58c4-45bf-a8a9-2dafb9e5214d",
	}
	interface, err := volumes.CreateInterface(serviceClient, volumeID, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete an Interface

	volumeID := "c79a4468-d788-410c-bf79-9a8ef6354852"
	interfaceID := "84f1b61f-58c4-45bf-a8a9-2dafb9e5214d"
	err := volumes.DeleteInterface(serviceClient, volumeID, interfaceID).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Create a Mirror

	volumeID := "c79a4468-d788-410c-bf79-9a8ef6354852"
	createOpts := volumes.CreateMirrorOpts{
		DstRegion: "KR1",
		DstTenantID: "aa5436ab58144c768ca4e9d2e9f5c3b2",
		DstVolume: &volumes.DstVolumeOpts{
			Name: "test-volume",
			SizeGb: 100,
			MountProtocol: &volumes.MountProtocolOpts{
				Protocol: "nfs",
			},
		},
	}
	mirror, err := volumes.CreateMirror(serviceClient, volumeID, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Mirror

	volumeID := "c79a4468-d788-410c-bf79-9a8ef6354852"
	mirrorID := "84f1b61f-58c4-45bf-a8a9-2dafb9e5214d"
	err := volumes.DeleteMirror(serviceClient, volumeID, mirrorID).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Get a Mirror Stat

	volumeID := "c79a4468-d788-410c-bf79-9a8ef6354852"
	mirrorID := "84f1b61f-58c4-45bf-a8a9-2dafb9e5214d"
	stat, err := volumes.GetMirrorStat(serviceClient, volumeID, mirrorID).Extract()
	if err != nil {
		panic(err)
	}
*/

package volumes
