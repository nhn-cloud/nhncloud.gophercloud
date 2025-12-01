package volumes

import (
	"github.com/gophercloud/gophercloud"
)

type Volume struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	Status         string          `json:"status"`
	Description    string          `json:"description"`
	SizeGb         int             `json:"sizeGb"`
	ProjectID      string          `json:"projectId"`
	TenantID       string          `json:"tenantId"`
	ACL            []string        `json:"acl"`
	Encryption     *Encryption     `json:"encryption"`
	Interfaces     []*Interface    `json:"interfaces"`
	Mirrors        []*Mirror       `json:"mirrors"`
	MountProtocol  *MountProtocol  `json:"mountProtocol"`
	SnapshotPolicy *SnapshotPolicy `json:"snapshotPolicy"`
	CreatedAt      string          `json:"createdAt"`
	UpdatedAt      string          `json:"updatedAt"`
}

type Encryption struct {
	Enabled bool             `json:"enabled"`
	Keys    []*EncryptionKey `json:"keys"`
}

type EncryptionKey struct {
	KeyID      string `json:"keyId"`
	KeyStoreID int    `json:"keyStoreId"`
}

type Interface struct {
	ID       string `json:"id"`
	Path     string `json:"path"`
	Status   string `json:"status"`
	SubnetID string `json:"subnetId"`
	TenantID string `json:"tenantId"`
}

type Mirror struct {
	ID                 string `json:"id"`
	Role               string `json:"role"`
	Status             string `json:"status"`
	Direction          string `json:"direction"`
	DirectionChangedAt string `json:"directionChangedAt"`
	DstProjectID       string `json:"dstProjectId"`
	DstRegion          string `json:"dstRegion"`
	DstTenantID        string `json:"dstTenantId"`
	DstVolumeID        string `json:"dstVolumeId"`
	DstVolumeName      string `json:"dstVolumeName"`
	SrcProjectID       string `json:"srcProjectId"`
	SrcRegion          string `json:"srcRegion"`
	SrcTenantID        string `json:"srcTenantId"`
	SrcVolumeID        string `json:"srcVolumeId"`
	SrcVolumeName      string `json:"srcVolumeName"`
	CreatedAt          string `json:"createdAt"`
}

type MountProtocol struct {
	CifsAuthIDs []string `json:"cifsAuthIds"`
	Protocol    string   `json:"protocol"`
}

type SnapshotPolicy struct {
	MaxScheduledCount *int      `json:"maxScheduledCount"`
	ReservePercent    int       `json:"reservePercent"`
	Schedule          *Schedule `json:"schedule"`
}

type Schedule struct {
	Time       string `json:"time"`
	TimeOffset string `json:"timeOffset"`
	Weekdays   []int  `json:"weekdays"`
}

type commonResult struct {
	gophercloud.Result
}

func (r commonResult) Extract() (*Volume, error) {
	var s Volume
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "volume")
}

type CreateResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type ConnectInterfaceResult struct {
	commonResult
}

func (r ConnectInterfaceResult) Extract() (*Interface, error) {
	var s Interface
	err := r.ExtractInto(&s)
	return &s, err
}

func (r ConnectInterfaceResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "interface")
}

type DeleteInterfaceResult struct {
	gophercloud.ErrResult
}

type SetReplicationResult struct {
	commonResult
}

func (r SetReplicationResult) Extract() (*Mirror, error) {
	var s Mirror
	err := r.ExtractInto(&s)
	return &s, err
}

func (r SetReplicationResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "volumeMirror")
}

type DisableReplicationResult struct {
	gophercloud.ErrResult
}

type MirrorStat struct {
	Status string `json:"status"`
}

type GetReplicationStatResult struct {
	commonResult
}

func (r GetReplicationStatResult) Extract() (*MirrorStat, error) {
	var s MirrorStat
	err := r.ExtractInto(&s)
	return &s, err
}

func (r GetReplicationStatResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "volumeMirrorStat")
}
