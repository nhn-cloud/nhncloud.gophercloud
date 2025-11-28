package volumes

import (
	"github.com/gophercloud/gophercloud"
)

type CreateVolumeOptsBuilder interface {
	ToVolumeCreateMap() (map[string]interface{}, error)
}

type CreateVolumeOpts struct {
	Name           string              `json:"name" required:"true"`
	Description    string              `json:"description,omitempty"`
	SizeGb         int                 `json:"sizeGb" required:"true"`
	ACL            []string            `json:"acl,omitempty"`
	Encryption     *EncryptionOpts     `json:"encryption,omitempty"`
	MountProtocol  *MountProtocolOpts  `json:"mountProtocol" required:"true"`
	SnapshotPolicy *SnapshotPolicyOpts `json:"snapshotPolicy,omitempty"`
}

type EncryptionOpts struct {
	Enabled bool `json:"enabled,omitempty"`
}

type MountProtocolOpts struct {
	CifsAuthIDs []string `json:"cifsAuthIds,omitempty"`
	Protocol    string   `json:"protocol" required:"true"`
}

type SnapshotPolicyOpts struct {
	MaxScheduledCount int           `json:"maxScheduledCount,omitempty"`
	ReservePercent    int           `json:"reservePercent,omitempty"`
	Schedule          *ScheduleOpts `json:"schedule,omitempty"`
}

type ScheduleOpts struct {
	Time       string `json:"time"`
	TimeOffset string `json:"timeOffset"`
	Weekdays   []int  `json:"weekdays"`
}

func (opts CreateVolumeOpts) ToVolumeCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "volume")
}

func CreateVolume(client *gophercloud.ServiceClient, opts CreateVolumeOptsBuilder) (r CreateVolumeResult) {
	b, err := opts.ToVolumeCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	resp, err := client.Post(createVolumeURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

func GetVolume(client *gophercloud.ServiceClient, id string) (r GetVolumeResult) {
	resp, err := client.Get(getVolumeURL(client, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

type UpdateVolumeOptsBuilder interface {
	ToVolumeUpdateMap() (map[string]interface{}, error)
}

type UpdateVolumeOpts struct {
	Description    *string                   `json:"description,omitempty"`
	SizeGb         *int                      `json:"sizeGb,omitempty"`
	ACL            *[]string                 `json:"acl,omitempty"`
	MountProtocol  *UpdateMountProtocolOpts  `json:"mountProtocol,omitempty"`
	SnapshotPolicy *UpdateSnapshotPolicyOpts `json:"snapshotPolicy,omitempty"`
}

type UpdateMountProtocolOpts struct {
	CifsAuthIDs []string `json:"cifsAuthIds,omitempty"`
	Protocol    string   `json:"protocol,omitempty"`
}

type UpdateSnapshotPolicyOpts struct {
	MaxScheduledCount *int                `json:"maxScheduledCount"`
	ReservePercent    *int                `json:"reservePercent,omitempty"`
	Schedule          *UpdateScheduleOpts `json:"schedule"`
}

type UpdateScheduleOpts struct {
	Time       string `json:"time"`
	TimeOffset string `json:"timeOffset"`
	Weekdays   []int  `json:"weekdays"`
}

func (opts UpdateVolumeOpts) ToVolumeUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "volume")
}

func UpdateVolume(client *gophercloud.ServiceClient, id string, opts UpdateVolumeOptsBuilder) (r UpdateVolumeResult) {
	b, err := opts.ToVolumeUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Patch(updateVolumeURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

func DeleteVolume(client *gophercloud.ServiceClient, id string) (r DeleteVolumeResult) {
	resp, err := client.Delete(deleteVolumeURL(client, id), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

type CreateInterfaceOptsBuilder interface {
	ToInterfaceCreateMap() (map[string]interface{}, error)
}

type CreateInterfaceOpts struct {
	SubnetID string `json:"subnetId" required:"true"`
}

func (opts CreateInterfaceOpts) ToInterfaceCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "interface")
}

func CreateInterface(client *gophercloud.ServiceClient, volumeID string, opts CreateInterfaceOptsBuilder) (r CreateInterfaceResult) {
	b, err := opts.ToInterfaceCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(createInterfaceURL(client, volumeID), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

func DeleteInterface(client *gophercloud.ServiceClient, volumeID string, interfaceID string) (r DeleteInterfaceResult) {
	resp, err := client.Delete(deleteInterfaceURL(client, volumeID, interfaceID), &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

type CreateMirrorOptsBuilder interface {
	ToVolumeMirrorCreateMap() (map[string]interface{}, error)
}

type CreateMirrorOpts struct {
	DstRegion   string         `json:"dstRegion" required:"true"`
	DstTenantID string         `json:"dstTenantId" required:"true"`
	DstVolume   *DstVolumeOpts `json:"dstVolume" required:"true"`
}

type DstVolumeOpts struct {
	Name           string              `json:"name" required:"true"`
	Description    string              `json:"description,omitempty"`
	SizeGb         int                 `json:"sizeGb" required:"true"`
	ACL            []string            `json:"acl,omitempty"`
	Encryption     *EncryptionOpts     `json:"encryption,omitempty"`
	MountProtocol  *MountProtocolOpts  `json:"mountProtocol,omitempty"`
	SnapshotPolicy *SnapshotPolicyOpts `json:"snapshotPolicy,omitempty"`
}

func (opts CreateMirrorOpts) ToVolumeMirrorCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "volumeMirror")
}

func CreateMirror(client *gophercloud.ServiceClient, volumeID string, opts CreateMirrorOptsBuilder) (r CreateMirrorResult) {
	b, err := opts.ToVolumeMirrorCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(createMirrorURL(client, volumeID), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

func DeleteMirror(client *gophercloud.ServiceClient, volumeID string, mirrorID string) (r DeleteMirrorResult) {
	resp, err := client.Delete(deleteMirrorURL(client, volumeID, mirrorID), &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

func GetMirrorStat(client *gophercloud.ServiceClient, volumeID string, mirrorID string) (r GetMirrorStatResult) {
	resp, err := client.Get(getMirrorStatURL(client, volumeID, mirrorID), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
