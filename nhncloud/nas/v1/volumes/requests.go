package volumes

import (
	"github.com/gophercloud/gophercloud"
)

type CreateOptsBuilder interface {
	ToVolumeCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
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

func (opts CreateOpts) ToVolumeCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "volume")
}

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

func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := client.Get(getURL(client, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

type UpdateOptsBuilder interface {
	ToVolumeUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
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

func (opts UpdateOpts) ToVolumeUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "volume")
}

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

func Delete(client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	resp, err := client.Delete(deleteURL(client, id), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

type ConnectInterfaceOptsBuilder interface {
	ToInterfaceConnectMap() (map[string]interface{}, error)
}

type ConnectInterfaceOpts struct {
	SubnetID string `json:"subnetId" required:"true"`
}

func (opts ConnectInterfaceOpts) ToInterfaceConnectMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "interface")
}

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

func DeleteInterface(client *gophercloud.ServiceClient, volumeID string, interfaceID string) (r DeleteInterfaceResult) {
	resp, err := client.Delete(deleteInterfaceURL(client, volumeID, interfaceID), &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

type SetMirrorOptsBuilder interface {
	ToReplicationSetMap() (map[string]interface{}, error)
}

type SetReplicationOpts struct {
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

func (opts SetReplicationOpts) ToReplicationSetMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "volumeMirror")
}

func SetReplication(client *gophercloud.ServiceClient, volumeID string, opts SetMirrorOptsBuilder) (r SetReplicationResult) {
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

func DisableReplication(client *gophercloud.ServiceClient, volumeID string, mirrorID string) (r DisableReplicationResult) {
	resp, err := client.Delete(disableReplicationURL(client, volumeID, mirrorID), &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

func GetReplicationStat(client *gophercloud.ServiceClient, volumeID string, mirrorID string) (r GetReplicationStatResult) {
	resp, err := client.Get(getReplicationStatURL(client, volumeID, mirrorID), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
