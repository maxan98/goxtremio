package goxtremio

import xms "github.com/emccode/goxtremio/api/v3"

type InitiatorGroup xms.InitiatorGroup
type NewInitiatorGroupOptions xms.PostInitiatorGroupsReq
type NewInitiatorGroupResult *xms.PostInitiatorGroupsResp
type ModifyInitiatorGroupOptions xms.PutInitiatorGroupsReq

//GetInitiatorGroup returns a specific initiator by name or ID
func (c *Client) GetInitiatorGroup(
	id string, name string) (*InitiatorGroup, error) {
	ig, err := c.api.GetInitiatorGroup(id, name)
	if err != nil {
		return nil, err
	}
	igg := InitiatorGroup(*ig.Content)
	return &igg, nil
}

//GetInitiatorGroups returns a list of initiators
func (c *Client) GetInitiatorGroups() (Refs, error) {
	ig, err := c.api.GetInitiatorGroups()
	if err != nil {
		return nil, err
	}
	return ig.InitiatorGroups, nil
}

//NewInitiatorGroup creates an IG
func (c *Client) NewInitiatorGroup(opts *NewInitiatorGroupOptions) (NewInitiatorGroupResult, error) {
	req := xms.PostInitiatorGroupsReq(*opts)
	res, err := c.api.PostInitiatorGroups(&req)
	if err != nil {
		return nil, err
	}

	nigr := NewInitiatorGroupResult(res)
	return nigr, nil
}

//Modify IG
func (c *Client) ModifyInitiatorGroup(id string, opts *ModifyInitiatorGroupOptions) (err error) {
	req := xms.PutInitiatorGroupsReq(*opts)
	return c.api.PutInitiatorGroup(id, &req)
}
