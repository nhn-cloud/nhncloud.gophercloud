package volumes

import (
	"github.com/gophercloud/gophercloud"
)

// Volume represents a NAS volume in NHN Cloud.
type Volume struct {
	// ID is the unique identifier of the NAS volume.
	ID string `json:"id"`

	// Name is the name of the NAS volume.
	Name string `json:"name"`

	// Status is the current status of the NAS volume.
	Status string `json:"status"`

	// Description is the description of the NAS volume.
	Description string `json:"description"`

	// SizeGb is the size of the NAS volume in GB.
	SizeGb int `json:"sizeGb"`

	// ProjectID is the proejct ID where the NAS volume is created.
	ProjectID string `json:"projectId"`

	// TenantID is the tenant ID where the NAS volume is created.
	TenantID string `json:"tenantId"`

	// ACL is the list of the IPs or CIDR blocks that allow read and write permissions.
	ACL []string `json:"acl"`

	// Encryption contains encryption configurations.
	Encryption *Encryption `json:"encryption"`

	// Interfaces contains the list of interfaces associated with the NAS volume.
	Interfaces []*Interface `json:"interfaces"`

	// Mirrors contains the list of mirrors associated with the NAS volume.
	Mirrors []*Mirror `json:"mirrors"`

	// MountProtocol contains mount protocol configurations.
	MountProtocol *MountProtocol `json:"mountProtocol"`

	// SnapshotPolicy contains snapshot policy configurations.
	SnapshotPolicy *SnapshotPolicy `json:"snapshotPolicy"`

	// CreatedAt is the date and time when the NAS volume was created.
	CreatedAt string `json:"createdAt"`

	// UpdatedAt is the date and time when the NAS volume was last updated.
	UpdatedAt string `json:"updatedAt"`
}

// Encryption represents encryption configuration.
type Encryption struct {
	// Enabled indicates whether encryption is enabled.
	Enabled bool `json:"enabled"`

	// Keys contains the list of NAS Volume encryption keys information.
	Keys []*EncryptionKey `json:"keys"`
}

// EncryptionKey represents encryption key information.
type EncryptionKey struct {
	// KeyID is the ID of the key used for encryption.
	KeyID string `json:"keyId"`

	// KeyStoreID is the ID of the key store used for encryption.
	KeyStoreID int `json:"keyStoreId"`
}

// Interface represents interface information.
type Interface struct {
	// ID is the unique identifier of the interface.
	ID string `json:"id"`

	// Path is the connection path that the instance will use when mounting.
	Path string `json:"path"`

	// Status is the current status of the interface.
	Status string `json:"status"`

	// SubnetID is the ID of the subnet associated with the interface.
	SubnetID string `json:"subnetId"`

	// TenantID is the tenant ID where the interface is created.
	TenantID string `json:"tenantId"`
}

// Mirror represents mirror information.
type Mirror struct {
	// ID is the unique identifier of the mirror.
	ID string `json:"id"`

	// Role is the replication role.
	Role string `json:"role"`

	// Status is the current status of the mirror.
	Status string `json:"status"`

	// Direction is the replication direction.
	Direction string `json:"direction"`

	// DirectionChangedAt is the date and time when the replication direction was changed.
	DirectionChangedAt string `json:"directionChangedAt"`

	// DstProjectID is the project ID where the target volume is created.
	DstProjectID string `json:"dstProjectId"`

	// DstRegion is the region where the target volume is created.
	DstRegion string `json:"dstRegion"`

	// DstTenantID is the tenant ID where the target volume is created.
	DstTenantID string `json:"dstTenantId"`

	// DstVolumeID is the ID of the target volume.
	DstVolumeID string `json:"dstVolumeId"`

	// DstVolumeName is the name of the target volume.
	DstVolumeName string `json:"dstVolumeName"`

	// SrcProjectID is the project ID where the source volume is created.
	SrcProjectID string `json:"srcProjectId"`

	// SrcRegion is the region where the source volume is created.
	SrcRegion string `json:"srcRegion"`

	// SrcTenantID is the tenant ID where the source volume is created.
	SrcTenantID string `json:"srcTenantId"`

	// SrcVolumeID is the ID of the source volume.
	SrcVolumeID string `json:"srcVolumeId"`

	// SrcVolumeName is the name of the source volume.
	SrcVolumeName string `json:"srcVolumeName"`

	// CreatedAt is the date and time when the mirror was created.
	CreatedAt string `json:"createdAt"`
}

// MountProtocol represents mount protocol configuration.
type MountProtocol struct {
	// CifsAuthIDs is the list of CIFS authentication IDs.
	CifsAuthIDs []string `json:"cifsAuthIds"`

	// Protocol specifies the protocol when mounting the NAS volume.
	Protocol string `json:"protocol"`
}

// SnapshotPolicy represents snapshot policy configuration.
type SnapshotPolicy struct {
	// MaxScheduledCount is the maximum number of snaptshots that can be saved
	MaxScheduledCount *int `json:"maxScheduledCount"`

	// ReservePercent is the snapshot capacity ratio.
	ReservePercent int `json:"reservePercent"`

	// Schedule contains schedule configurations.
	Schedule *Schedule `json:"schedule"`
}

// Schedule represents schedule configuration.
type Schedule struct {
	// Time is the snapshot auto-create time.
	Time string `json:"time"`

	// TimeOffset is the time zone for snapshot auto-create.
	TimeOffset string `json:"timeOffset"`

	// Weekdays is the days of the week that snapshots are automatically created.
	Weekdays []int `json:"weekdays"`
}

type commonResult struct {
	gophercloud.Result
}

// Extract will get the Volume object out of the commonResult object.
func (r commonResult) Extract() (*Volume, error) {
	var s Volume
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "volume")
}

// CreateResult is the response from a Create operation.
type CreateResult struct {
	commonResult
}

// GetREsult contains the response body and error from a Get request.
type GetResult struct {
	commonResult
}

// UpdateResult is the response from an Update operation.
type UpdateResult struct {
	commonResult
}

// DeleteResult is the response from a Delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

// ConnectInterfaceResult is the response from a ConnectInterface operation.
type ConnectInterfaceResult struct {
	gophercloud.Result
}

// Extract will get the Interface object out of the ConnectInterfaceResult object.
func (r ConnectInterfaceResult) Extract() (*Interface, error) {
	var s Interface
	err := r.ExtractInto(&s)
	return &s, err
}

func (r ConnectInterfaceResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "interface")
}

// DeleteInterfaceResult is the response from a DeleteInterface operation.
type DeleteInterfaceResult struct {
	gophercloud.ErrResult
}

// SetReplicationResult is the response from a SetReplication operation.
type SetReplicationResult struct {
	gophercloud.Result
}

// Extract will get the Mirror object out of the SetReplicationResult object.
func (r SetReplicationResult) Extract() (*Mirror, error) {
	var s Mirror
	err := r.ExtractInto(&s)
	return &s, err
}

func (r SetReplicationResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "volumeMirror")
}

// DisableReplicationResult is the response from a DisableReplication operation.
type DisableReplicationResult struct {
	gophercloud.ErrResult
}

// MirrorStat represents mirror status information.
type MirrorStat struct {
	// Status is the current status of the mirror.
	Status string `json:"status"`
}

// GetReplicationStatResult is the response from a GetReplicationStat operation.
type GetReplicationStatResult struct {
	commonResult
}

// Extract will get the MirrorStat object out of the GetReplicationStatResult object.
func (r GetReplicationStatResult) Extract() (*MirrorStat, error) {
	var s MirrorStat
	err := r.ExtractInto(&s)
	return &s, err
}

func (r GetReplicationStatResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "volumeMirrorStat")
}
