package goxtremio

import xms "github.com/emccode/goxtremio/api/v3"

type Volume *xms.Volume
type NewVolumeOptions xms.PostVolumesReq
type NewVolumeResult *xms.PostVolumesResp
type CloneVolumeOptions xms.VolumeCreateCopyReq
type CloneVolumeResult *xms.VolumeCreateCopyResponseContent
type ModifyVolumeOptions xms.PutVolumeReq

// GetVolume returns a specific volume by name or ID
func (c *Client) GetVolume(id string, name string) (Volume, error) {
	volume, err := c.api.GetVolume(id, name)
	if err != nil {
		return nil, err
	}
	return volume.Content, nil
}

// GetVolumes returns a list of volumes
func (c *Client) GetVolumes() (Refs, error) {
	volumes, err := c.api.GetVolumes()
	if err != nil {
		return nil, err
	}
	return volumes.Volumes, nil
}

// Constructs a new Volume instance
func VolumeCtor() Volume {
	return &xms.Volume{}
}

// Constructs a new Volume instance
func VolumeCtorNameIndex(name string, index int) Volume {
	return &xms.Volume{Name: name, Index: index}
}

// NewVolume creates a volume
func (c *Client) NewVolume(opts *NewVolumeOptions) (NewVolumeResult, error) {
	req := xms.PostVolumesReq(*opts)
	return c.api.PostVolumes(&req)
}

// CloneVolume creates a volume from an existing volume
func (c *Client) CloneVolume(opts *CloneVolumeOptions) (CloneVolumeResult, error) {
	req := xms.VolumeCreateCopyReq(*opts)
	resp, err := c.api.VolumeCreateRepurposeCopy(&req)
	if err != nil {
		return nil, err
	}
	return resp.Content, nil
}

// DeleteVolume deletes a volume
func (c *Client) DeleteVolume(id string, name string) error {
	return c.api.DeleteVolumes(id, name)
}

// ModifyVolume update an existing volume
func (c *Client) ModifyVolume(id string, name string, opts *ModifyVolumeOptions) error {
	req := xms.PutVolumeReq(*opts)
	return c.api.PutVolume(id, name, &req)
}
