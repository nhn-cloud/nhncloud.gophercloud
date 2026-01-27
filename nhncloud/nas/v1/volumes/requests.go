package volumes

import (
	"github.com/gophercloud/gophercloud"
)

// CreateOptsBuilder allows extensions to add additional parameters
// to the Create request.
type CreateOptsBuilder interface {
	ToVolumeCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options used to create a NAS volume.
type CreateOpts struct {
	// Name is the name of the NAS volume.
	Name string `json:"name" required:"true"`

	// Description is the description of the NAS volume.
	Description string `json:"description,omitempty"`

	// SizeGb is the size of the NAS volume in GB.
	SizeGb int `json:"sizeGb" required:"true"`

	// ACL is the list of the IPs or CIDR blocks that allow read and write permissions.
	ACL []string `json:"acl,omitempty"`

	// Encryption contains encryption configurations.
	Encryption *EncryptionOpts `json:"encryption,omitempty"`

	// Interfaces contains the list of interfaces associated with the NAS volume.
	Interfaces []*InterfaceOpts `json:"interfaces,omitempty"`

	// MountProtocol contains mount protocol configurations.
	MountProtocol *MountProtocolOpts `json:"mountProtocol" required:"true"`

	// SnapshotPolicy contains snapshot policy configurations.
	SnapshotPolicy *SnapshotPolicyOpts `json:"snapshotPolicy,omitempty"`
}

// InterfaceOpts represents interface configuration.
type InterfaceOpts struct {
	// SubnetID is the ID of the subnet associated with the NAS volume.
	SubnetID string `json:"subnetId,omitempty"`
}

// EncryptionOpts represents encryption configuration.
type EncryptionOpts struct {
	// Enabled indicates whether to enable encryption settings.
	Enabled bool `json:"enabled,omitempty"`
}

// MountProtocolOpts represents mount protocol configuration.
type MountProtocolOpts struct {
	// CifsAuthIDs is the lif of CIFS authentication IDs.
	CifsAuthIDs []string `json:"cifsAuthIds,omitempty"`

	// Protocol specifies the protocol when mounting the NAS volume.
	Protocol string `json:"protocol" required:"true"`
}

// SnapshotPolicyOpts represents snapshot policy configuration.
type SnapshotPolicyOpts struct {
	// MaxScheduledCount is the maximum number of snapshots that can be saved
	MaxScheduledCount *int `json:"maxScheduledCount"`

	// ReservePercent is the snapshot capacity ratio.
	ReservePercent int `json:"reservePercent,omitempty"`

	// Schedule contains schedule configurations.
	Schedule *ScheduleOpts `json:"schedule"`
}

// ScheduleOpts represents schedule configuration.
type ScheduleOpts struct {
	// Time is the snapshot auto-create time.
	Time string `json:"time"`

	// TimeOffset is the time zone for snapshot auto-create.
	TimeOffset string `json:"timeOffset"`

	// Weekdays is the days of the week that snapshots are automatically created.
	Weekdays []int `json:"weekdays"`
}

// ToVolumeCreateMap assembles a request body based on the contents of a CreateOpts.
func (opts CreateOpts) ToVolumeCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "volume")
}

// Create requests the creation of a new NAS volume.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToVolumeCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	resp, err := client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Get retrieves a specific NAS volume based on its unique ID.
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := client.Get(getURL(client, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToVolumeUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains the values used when updating a NAS volume.
type UpdateOpts struct {
	// Description is the description of the NAS volume.
	Description *string `json:"description,omitempty"`

	// SizeGb is the size of the NAS volume in GB.
	SizeGb *int `json:"sizeGb,omitempty"`

	// ACL is the list of the IPs or CIDR blocks that allow read and write permissions.
	ACL *[]string `json:"acl,omitempty"`

	// MountProtocol contains mount protocol configurations.
	MountProtocol *MountProtocolOpts `json:"mountProtocol,omitempty"`

	// SnapshotPolicy contains snapshot policy configurations.
	SnapshotPolicy *SnapshotPolicyOpts `json:"snapshotPolicy,omitempty"`
}

// ToVolumeUpdateMap assembles a request body based on the contents of an UpdateOpts.
func (opts UpdateOpts) ToVolumeUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "volume")
}

// Update allows a NAS volume to be updated.
func Update(client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToVolumeUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Patch(updateURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete deletes a specific NAS volume based on its unique ID.
func Delete(client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	resp, err := client.Delete(deleteURL(client, id), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ConnectInterfaceOptsBuilder allows extensions to add additional parameters to the
// ConnectInterface request.
type ConnectInterfaceOptsBuilder interface {
	ToInterfaceConnectMap() (map[string]interface{}, error)
}

// ConnectInterfaceOpts contains the values used when setting the interface of a NAS volume.
type ConnectInterfaceOpts struct {
	//SubnetID is the ID of subnet associated with the NAS volume.
	SubnetID string `json:"subnetId" required:"true"`
}

// ToInterfaceConnectMap assembles a request body based on the contents of a ConnectInterfaceOpts.
func (opts ConnectInterfaceOpts) ToInterfaceConnectMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "interface")
}

// ConnectInterface set t the interface of a NAS volume.
func ConnectInterface(client *gophercloud.ServiceClient, volumeID string, opts ConnectInterfaceOptsBuilder) (r ConnectInterfaceResult) {
	b, err := opts.ToInterfaceConnectMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(connectInterfaceURL(client, volumeID), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// DeleteInterface deletes the interface of a NAS volume.
func DeleteInterface(client *gophercloud.ServiceClient, volumeID string, interfaceID string) (r DeleteInterfaceResult) {
	resp, err := client.Delete(deleteInterfaceURL(client, volumeID, interfaceID), &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// SetReplicationOptsBuilder allows extensions to add additional parameters to the
// SetReplication request.
type SetReplicationOptsBuilder interface {
	ToReplicationSetMap() (map[string]interface{}, error)
}

// SetReplicationOpts contains the values used when setting the replication of a NAS volume.
type SetReplicationOpts struct {
	// DstRegion is the region of the target volume.
	DstRegion string `json:"dstRegion" required:"true"`

	// DstTenantID is the tenant ID of the target volume.
	DstTenantID string `json:"dstTenantId" required:"true"`

	// DstVolume contains the target volume configurations.
	DstVolume *DstVolumeOpts `json:"dstVolume" required:"true"`
}

// DstVolumeOpts contains the target volume configurations.
type DstVolumeOpts struct {
	// Name is the name of the target volume.
	Name string `json:"name" required:"true"`

	// Description is the description of the target volume.
	Description string `json:"description,omitempty"`

	// SizeGb is the size of the target volume in GB.
	SizeGb int `json:"sizeGb" required:"true"`

	// ACL is the list of the IPs or CIDR blocks that allow read and write permissions.
	ACL []string `json:"acl,omitempty"`

	// Encryption contains encryption configurations.
	Encryption *EncryptionOpts `json:"encryption,omitempty"`

	// MountProtocol contains mount protocol configurations.
	MountProtocol *MountProtocolOpts `json:"mountProtocol,omitempty"`

	// SnapshotPolicy contains snapshot policy configurations.
	SnapshotPolicy *SnapshotPolicyOpts `json:"snapshotPolicy,omitempty"`
}

// ToReplicationSetMap assembles a request body based on the contents of a SetReplicationOpts.
func (opts SetReplicationOpts) ToReplicationSetMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "volumeMirror")
}

// SetReplication sets the replication of a NAS volume.
func SetReplication(client *gophercloud.ServiceClient, volumeID string, opts SetReplicationOptsBuilder) (r SetReplicationResult) {
	b, err := opts.ToReplicationSetMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(setReplicationURL(client, volumeID), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// DisableReplication disables the replication of a NAS volume.
func DisableReplication(client *gophercloud.ServiceClient, volumeID string, mirrorID string) (r DisableReplicationResult) {
	resp, err := client.Delete(disableReplicationURL(client, volumeID, mirrorID), &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// GetReplicationStat retrieves the replication status of a NAS volume.
func GetReplicationStat(client *gophercloud.ServiceClient, volumeID string, mirrorID string) (r GetReplicationStatResult) {
	resp, err := client.Get(getReplicationStatURL(client, volumeID, mirrorID), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
