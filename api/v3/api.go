package apiv3

import (
	"errors"
	"os"
	"strconv"
)

var debug bool

var ErrIdOrNameNotSpecified = errors.New("Either an ID or Name must be specified")

func init() {
	debug, _ = strconv.ParseBool(os.Getenv("GOXTREMIO_DEBUG"))
}

type Ref struct {
	Href    string `json:"href"`
	Name    string `json:"name"`
	SysName string `json:"sys-name"`
}

type getVolumesResp struct {
	Volumes []*Ref `json:"volumes"`
	Links   []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetVolumes() (resp *getVolumesResp, err error) {
	err = xms.query("GET", "/types/volumes", "", nil, nil, &resp)
	return resp, err
}

type Volume struct {
	UnalignedIoAlerts       string          `json:"unaligned-io-alerts"`
	SmallIoAlerts           string          `json:"small-io-alerts"`
	SmallIops               string          `json:"small-iops"`
	DestSnapList            []interface{}   `json:"dest-snap-list"`
	VolID                   []interface{}   `json:"vol-id"`
	ObjSeverity             string          `json:"obj-severity"`
	NumOfLunMappings        int             `json:"num-of-lun-mappings"`
	UnalignedRdBw           string          `json:"unaligned-rd-bw"`
	NumOfDestSnaps          int             `json:"num-of-dest-snaps"`
	Iops                    string          `json:"iops"`
	AccNumOfSmallWr         string          `json:"acc-num-of-small-wr"`
	AlignmentOffset         int             `json:"alignment-offset"`
	LbSize                  int             `json:"lb-size"`
	LogicalSpaceInUse       string          `json:"logical-space-in-use"`
	UnalignedIoRatioLevel   string          `json:"unaligned-io-ratio-level"`
	AccNumOfRd              string          `json:"acc-num-of-rd"`
	Index                   int             `json:"index"`
	SmallRdBw               string          `json:"small-rd-bw"`
	NaaName                 string          `json:"naa-name"`
	AccSizeOfWr             string          `json:"acc-size-of-wr"`
	AccNumOfSmallRd         string          `json:"acc-num-of-small-rd"`
	UnalignedRdIops         string          `json:"unaligned-rd-iops"`
	WrLatency               string          `json:"wr-latency"`
	SnapgrpID               []interface{}   `json:"snapgrp-id"`
	AncestorVolID           []interface{}   `json:"ancestor-vol-id"`
	VaaiTpAlerts            string          `json:"vaai-tp-alerts"`
	CreationTime            string          `json:"creation-time"`
	RdBw                    string          `json:"rd-bw"`
	XmsID                   []interface{}   `json:"xms-id"`
	Compressible            string          `json:"compressible"`
	SmallWrBw               string          `json:"small-wr-bw"`
	AccNumOfUnalignedRd     string          `json:"acc-num-of-unaligned-rd"`
	LuName                  string          `json:"lu-name"`
	UnalignedIops           string          `json:"unaligned-iops"`
	UnalignedBw             string          `json:"unaligned-bw"`
	Bw                      string          `json:"bw"`
	SmallIoRatioLevel       string          `json:"small-io-ratio-level"`
	LunMappingList          [][]interface{} `json:"lun-mapping-list"`
	VolSize                 string          `json:"vol-size"`
	WrIops                  string          `json:"wr-iops"`
	SysID                   []interface{}   `json:"sys-id"`
	AvgLatency              string          `json:"avg-latency"`
	RdLatency               string          `json:"rd-latency"`
	SmallWrIops             string          `json:"small-wr-iops"`
	SmallBw                 string          `json:"small-bw"`
	Name                    string          `json:"name"`
	AccNumOfUnalignedWr     string          `json:"acc-num-of-unaligned-wr"`
	UnalignedWrIops         string          `json:"unaligned-wr-iops"`
	AccNumOfWr              string          `json:"acc-num-of-wr"`
	SmallIoRatio            string          `json:"small-io-ratio"`
	AccSizeOfRd             string          `json:"acc-size-of-rd"`
	UnalignedWrBw           string          `json:"unaligned-wr-bw"`
	SmallRdIops             string          `json:"small-rd-iops"`
	UnalignedIoRatio        string          `json:"unaligned-io-ratio"`
	RdIops                  string          `json:"rd-iops"`
	WrBw                    string          `json:"wr-bw"`
	GUID                    string          `json:"guid"`
	LastRefreshTime         string          `json:"last-refresh-time"`
	VolType                 string          `json:"vol-type"`
	CreatedByExternalClient string          `json:"created-by-external-client,omitempty"`
	QoSPolicyID             []interface{}   `json:"qos-policy-id,omitempty"`
}
type getVolumeResp struct {
	Content *Volume `json:"content"`
	Links   []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetVolume(id string, name string) (resp *getVolumeResp, err error) {
	if id == "" && name == "" {
		return nil, ErrIdOrNameNotSpecified
	}

	if id != "" && name != "" {
		err = xms.query("GET", "/types/volumes", "", map[string]string{"name": name, "cluster-name": id}, nil, &resp)
	} else {
		err = xms.query("GET", "/types/volumes", id, map[string]string{"name": name}, nil, &resp)
	}
	return resp, err
}

type getLunMapsResp struct {
	LunMaps []*Ref `json:"lun-maps"`
	Links   []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetLunMaps(filter string) (resp *getLunMapsResp, err error) {
	err = xms.query("GET", "/types/lun-maps", "", map[string]string{"filter": filter}, nil, &resp)
	return resp, err
}

type LunMap struct {
	TgName       string        `json:"tg-name"`
	IgIndex      int           `json:"ig-index"`
	XmsID        []interface{} `json:"xms-id"`
	MappingIndex int           `json:"mapping-index"`
	ObjSeverity  string        `json:"obj-severity"`
	TgIndex      int           `json:"tg-index"`
	Lun          int           `json:"lun"`
	IgName       string        `json:"ig-name"`
	VolIndex     int           `json:"vol-index"`
	VolName      string        `json:"vol-name"`
	MappingID    []interface{} `json:"mapping-id"`
}

type getLunMapResp struct {
	Content *LunMap `json:"content"`
	Links   []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetLunMap(id string, name string) (resp *getLunMapResp, err error) {
	if id == "" && name == "" {
		return nil, ErrIdOrNameNotSpecified
	}

	err = xms.query("GET", "/types/lun-maps", id, map[string]string{"name": name}, nil, &resp)
	return resp, err
}

type getInitiatorsResp struct {
	Initiators []*Ref `json:"initiators"`
	Links      []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetInitiators() (resp *getInitiatorsResp, err error) {
	err = xms.query("GET", "/types/initiators", "", nil, nil, &resp)
	return resp, err
}

type Initiator struct {
	SmallIops                           string        `json:"small-iops"`
	WrLatency                           string        `json:"wr-latency"`
	ChapDiscoveryInitiatorPassword      string        `json:"chap-discovery-initiator-password"`
	ObjSeverity                         string        `json:"obj-severity"`
	RdBw                                string        `json:"rd-bw"`
	UnalignedRdBw                       string        `json:"unaligned-rd-bw"`
	ChapDiscoveryInitiatorUserName      string        `json:"chap-discovery-initiator-user-name"`
	ChapDiscoveryClusterUserName        string        `json:"chap-discovery-cluster-user-name"`
	Iops                                string        `json:"iops"`
	NumOfConnTars                       int           `json:"num-of-conn-tars"`
	AccNumOfSmallWr                     string        `json:"acc-num-of-small-wr"`
	ChapAuthenticationInitiatorPassword string        `json:"chap-authentication-initiator-password"`
	AccNumOfRd                          string        `json:"acc-num-of-rd"`
	Index                               int           `json:"index"`
	PortAddress                         string        `json:"port-address"`
	SmallRdBw                           string        `json:"small-rd-bw"`
	ChapAuthenticationInitiatorUserName string        `json:"chap-authentication-initiator-user-name"`
	IgID                                []interface{} `json:"ig-id"`
	SmallBw                             string        `json:"small-bw"`
	AccSizeOfWr                         string        `json:"acc-size-of-wr"`
	AccNumOfSmallRd                     string        `json:"acc-num-of-small-rd"`
	UnalignedRdIops                     string        `json:"unaligned-rd-iops"`
	ChapDiscoveryClusterPassword        string        `json:"chap-discovery-cluster-password"`
	ChapAuthenticationClusterPassword   string        `json:"chap-authentication-cluster-password"`
	XmsID                               []interface{} `json:"xms-id"`
	UnalignedWrIops                     string        `json:"unaligned-wr-iops"`
	AccNumOfUnalignedRd                 string        `json:"acc-num-of-unaligned-rd"`
	SmallWrBw                           string        `json:"small-wr-bw"`
	UnalignedIops                       string        `json:"unaligned-iops"`
	UnalignedBw                         string        `json:"unaligned-bw"`
	Bw                                  string        `json:"bw"`
	WrIops                              string        `json:"wr-iops"`
	AvgLatency                          string        `json:"avg-latency"`
	RdLatency                           string        `json:"rd-latency"`
	SmallWrIops                         string        `json:"small-wr-iops"`
	ChapAuthenticationClusterUserName   string        `json:"chap-authentication-cluster-user-name"`
	Name                                string        `json:"name"`
	AccNumOfUnalignedWr                 string        `json:"acc-num-of-unaligned-wr"`
	AccNumOfWr                          string        `json:"acc-num-of-wr"`
	InitiatorID                         []interface{} `json:"initiator-id"`
	AccSizeOfRd                         string        `json:"acc-size-of-rd"`
	UnalignedWrBw                       string        `json:"unaligned-wr-bw"`
	SmallRdIops                         string        `json:"small-rd-iops"`
	InitiatorConnState                  string        `json:"initiator-conn-state"`
	RdIops                              string        `json:"rd-iops"`
	WrBw                                string        `json:"wr-bw"`
	PortType                            string        `json:"port-type"`
}

type getInitiatorResp struct {
	Content *Initiator `json:"content"`
	Links   []Link     `json:"links"`
}

type Link struct {
	GUID string `json:"guid"`
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

func (xms *XMS) GetInitiator(id string, name string) (resp *getInitiatorResp, err error) {
	if id == "" && name == "" {
		return nil, ErrIdOrNameNotSpecified
	}

	err = xms.query("GET", "/types/initiators", id, map[string]string{"name": name}, nil, &resp)
	return resp, err
}

type getInitiatorGroupsResp struct {
	InitiatorGroups []*Ref `json:"initiator-groups"`
	Links           []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetInitiatorGroups() (resp *getInitiatorGroupsResp, err error) {
	err = xms.query("GET", "/types/initiator-groups", "", nil, nil, &resp)
	return resp, err
}

type InitiatorGroup struct {
	SmallIops           string        `json:"small-iops"`
	NumOfInitiators     int           `json:"num-of-initiators"`
	ObjSeverity         string        `json:"obj-severity"`
	RdBw                string        `json:"rd-bw"`
	UnalignedRdBw       string        `json:"unaligned-rd-bw"`
	Iops                string        `json:"iops"`
	AccNumOfSmallWr     string        `json:"acc-num-of-small-wr"`
	AccNumOfRd          string        `json:"acc-num-of-rd"`
	Index               int           `json:"index"`
	SmallRdBw           string        `json:"small-rd-bw"`
	IgID                []interface{} `json:"ig-id"`
	AccSizeOfWr         string        `json:"acc-size-of-wr"`
	AccNumOfSmallRd     string        `json:"acc-num-of-small-rd"`
	UnalignedRdIops     string        `json:"unaligned-rd-iops"`
	NumOfVols           int           `json:"num-of-vols"`
	XmsID               []interface{} `json:"xms-id"`
	UnalignedWrIops     string        `json:"unaligned-wr-iops"`
	AccNumOfUnalignedRd string        `json:"acc-num-of-unaligned-rd"`
	SmallWrBw           string        `json:"small-wr-bw"`
	UnalignedIops       string        `json:"unaligned-iops"`
	UnalignedBw         string        `json:"unaligned-bw"`
	SmallRdIops         string        `json:"small-rd-iops"`
	WrIops              string        `json:"wr-iops"`
	SmallWrIops         string        `json:"small-wr-iops"`
	SmallBw             string        `json:"small-bw"`
	Name                string        `json:"name"`
	AccNumOfUnalignedWr string        `json:"acc-num-of-unaligned-wr"`
	AccNumOfWr          string        `json:"acc-num-of-wr"`
	AccSizeOfRd         string        `json:"acc-size-of-rd"`
	UnalignedWrBw       string        `json:"unaligned-wr-bw"`
	Bw                  string        `json:"bw"`
	RdIops              string        `json:"rd-iops"`
	WrBw                string        `json:"wr-bw"`
}

type getInitiatorGroupResp struct {
	Content *InitiatorGroup `json:"content"`
	Links   []Link          `json:"links"`
}

func (xms *XMS) GetInitiatorGroup(id string, name string) (resp *getInitiatorGroupResp, err error) {
	if id == "" && name == "" {
		return nil, ErrIdOrNameNotSpecified
	}

	err = xms.query("GET", "/types/initiator-groups", id, map[string]string{"name": name}, nil, &resp)
	return resp, err
}

type InitiatorConnectivity struct {
	Index               int             `json:"index"`
	PortAddress         string          `json:"port-address"`
	Name                string          `json:"name"`
	TargetIDLlist       [][]interface{} `json:"target-list"`
	NumConnectedTragets int             `json:"num-of-conn-tars"`
	SysID               []interface{}   `json:"sys-id"`
	PortType            string          `json:"port-type"`
}

type getInitiatorsConnectivityResp struct {
	Content []*InitiatorConnectivity `json:"content"`
	Links   []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetInitiatorsConnectivity() (resp *getInitiatorsConnectivityResp, err error) {
	err = xms.query("GET", "/types/initiators-connectivity", "", nil, nil, &resp)
	return resp, err
}

//figure out how to get initiator names - from events

type getVolumeFoldersResp struct {
	Folders []*Ref `json:"folders"`
	Links   []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetVolumeFolders() (resp *getVolumeFoldersResp, err error) {
	err = xms.query("GET", "/types/volume-folders", "", nil, nil, &resp)
	return resp, err
}

type VolumeFolder struct {
	SmallIops           string          `json:"small-iops"`
	ObjSeverity         string          `json:"obj-severity"`
	RdBw                string          `json:"rd-bw"`
	UnalignedRdBw       string          `json:"unaligned-rd-bw"`
	Iops                string          `json:"iops"`
	AccNumOfSmallWr     string          `json:"acc-num-of-small-wr"`
	ParentFolderID      []interface{}   `json:"parent-folder-id"`
	AccNumOfRd          string          `json:"acc-num-of-rd"`
	Index               int             `json:"index"`
	SmallRdBw           string          `json:"small-rd-bw"`
	NumOfDirectObjs     int             `json:"num-of-direct-objs"`
	AccSizeOfWr         string          `json:"acc-size-of-wr"`
	NumOfSubfolders     int             `json:"num-of-subfolders"`
	AccNumOfSmallRd     string          `json:"acc-num-of-small-rd"`
	UnalignedRdIops     string          `json:"unaligned-rd-iops"`
	EmptyVolBlocks      string          `json:"empty-vol-blocks"`
	NumOfVols           int             `json:"num-of-vols"`
	XmsID               []interface{}   `json:"xms-id"`
	SubfolderList       []interface{}   `json:"subfolder-list"`
	NonEmptyVolBlocks   string          `json:"non-empty-vol-blocks"`
	UnalignedWrIops     string          `json:"unaligned-wr-iops"`
	AccNumOfUnalignedRd string          `json:"acc-num-of-unaligned-rd"`
	FolderID            []interface{}   `json:"folder-id"`
	SmallWrBw           string          `json:"small-wr-bw"`
	UnalignedIops       string          `json:"unaligned-iops"`
	UnalignedBw         string          `json:"unaligned-bw"`
	Bw                  string          `json:"bw"`
	VolSize             string          `json:"vol-size"`
	WrIops              string          `json:"wr-iops"`
	DirectList          [][]interface{} `json:"direct-list"`
	SmallWrIops         string          `json:"small-wr-iops"`
	SmallBw             string          `json:"small-bw"`
	Name                string          `json:"name"`
	AccNumOfUnalignedWr string          `json:"acc-num-of-unaligned-wr"`
	AccNumOfWr          string          `json:"acc-num-of-wr"`
	AccSizeOfRd         string          `json:"acc-size-of-rd"`
	Caption             string          `json:"caption"`
	UnalignedWrBw       string          `json:"unaligned-wr-bw"`
	SmallRdIops         string          `json:"small-rd-iops"`
	RdIops              string          `json:"rd-iops"`
	WrBw                string          `json:"wr-bw"`
}

type getVolumeFolderResp struct {
	Content *VolumeFolder `json:"content"`
	Links   []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetVolumeFolder(id string, name string) (resp *getVolumeFolderResp, err error) {
	if id == "" && name == "" {
		return nil, ErrIdOrNameNotSpecified
	}

	err = xms.query("GET", "/types/volume-folders", id, map[string]string{"name": name}, nil, &resp)
	return resp, err
}

type getIGFoldersResp struct {
	Folders []*Ref `json:"folders"`
	Links   []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetIGFolders() (resp *getIGFoldersResp, err error) {
	err = xms.query("GET", "/types/ig-folders", "", nil, nil, &resp)
	return resp, err
}

type IGFolder struct {
	SmallIops           string          `json:"small-iops"`
	ObjSeverity         string          `json:"obj-severity"`
	RdBw                string          `json:"rd-bw"`
	UnalignedRdBw       string          `json:"unaligned-rd-bw"`
	Iops                string          `json:"iops"`
	AccNumOfSmallWr     string          `json:"acc-num-of-small-wr"`
	ParentFolderID      []interface{}   `json:"parent-folder-id"`
	AccNumOfRd          string          `json:"acc-num-of-rd"`
	Index               int             `json:"index"`
	SmallRdBw           string          `json:"small-rd-bw"`
	NumOfDirectObjs     int             `json:"num-of-direct-objs"`
	AccSizeOfWr         string          `json:"acc-size-of-wr"`
	NumOfSubfolders     int             `json:"num-of-subfolders"`
	AccNumOfSmallRd     string          `json:"acc-num-of-small-rd"`
	UnalignedRdIops     string          `json:"unaligned-rd-iops"`
	XmsID               []interface{}   `json:"xms-id"`
	SubfolderList       []interface{}   `json:"subfolder-list"`
	UnalignedWrIops     string          `json:"unaligned-wr-iops"`
	AccNumOfUnalignedRd string          `json:"acc-num-of-unaligned-rd"`
	FolderID            []interface{}   `json:"folder-id"`
	SmallWrBw           string          `json:"small-wr-bw"`
	UnalignedIops       string          `json:"unaligned-iops"`
	UnalignedBw         string          `json:"unaligned-bw"`
	Bw                  string          `json:"bw"`
	WrIops              string          `json:"wr-iops"`
	DirectList          [][]interface{} `json:"direct-list"`
	SmallWrIops         string          `json:"small-wr-iops"`
	RdIops              string          `json:"rd-iops"`
	Name                string          `json:"name"`
	AccNumOfUnalignedWr string          `json:"acc-num-of-unaligned-wr"`
	AccNumOfWr          string          `json:"acc-num-of-wr"`
	AccSizeOfRd         string          `json:"acc-size-of-rd"`
	Caption             string          `json:"caption"`
	UnalignedWrBw       string          `json:"unaligned-wr-bw"`
	SmallRdIops         string          `json:"small-rd-iops"`
	SmallBw             string          `json:"small-bw"`
	WrBw                string          `json:"wr-bw"`
}

type getIGFolderResp struct {
	Content *IGFolder `json:"content"`
	Links   []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetIGFolder(id string, name string) (resp *getIGFolderResp, err error) {
	if id == "" && name == "" {
		return nil, ErrIdOrNameNotSpecified
	}

	err = xms.query("GET", "/types/ig-folders", id, map[string]string{"name": name}, nil, &resp)
	return resp, err
}

type getISCSIPortalsResp struct {
	ISCSIPortals []*Ref `json:"iscsi-portals"`
	Links        []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetISCSIPortals() (resp *getISCSIPortalsResp, err error) {
	err = xms.query("GET", "/types/iscsi-portals", "", nil, nil, &resp)
	return resp, err
}

type ISCSIPortal struct {
	Index       int           `json:"index"`
	PortAddress string        `json:"port-address"`
	XmsID       []interface{} `json:"xms-id"`
	Name        string        `json:"name"`
	ObjSeverity string        `json:"obj-severity"`
	Certainty   string        `json:"certainty"`
	Vlan        int           `json:"vlan"`
	TarID       []interface{} `json:"tar-id"`
	IPPort      int           `json:"ip-port"`
	IPAddr      string        `json:"ip-addr"`
	GUID        string        `json:"guid"`
}

type getISCSIPortalResp struct {
	Content *ISCSIPortal `json:"content"`
	Links   []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetISCSIPortal(id string, name string) (resp *getISCSIPortalResp, err error) {
	if id == "" && name == "" {
		return nil, ErrIdOrNameNotSpecified
	}

	err = xms.query("GET", "/types/iscsi-portals", id, map[string]string{"name": name}, nil, &resp)
	return resp, err
}

type Xms struct {
	Version string `json:"version"`
}

type getXmsResp struct {
	Content *Xms `json:"content"`
	Links   []struct {
		GUID string `json:"guid"`
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetXms() (resp *getXmsResp, err error) {
	err = xms.query("GET", "/types/xms", "1", nil, nil, &resp)
	return resp, err
}

type getClustersResp struct {
	Clusters []*Ref `json:"clusters"`
	Links    []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetClusters() (resp *getClustersResp, err error) {
	err = xms.query("GET", "/types/clusters", "", nil, nil, &resp)
	return resp, err
}

type Cluster struct {
	FreeUdSsdSpace string `json:"free-ud-ssd-space"`
	Name           string `json:"name"`
}

type getClusterResp struct {
	Content *Cluster `json:"content"`
	Links   []struct {
		GUID string `json:"guid"`
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetCluster(id string) (resp *getClusterResp, err error) {
	err = xms.query("GET", "/types/clusters", id, nil, nil, &resp)
	return resp, err
}

type getTargetsResp struct {
	Targets []*Ref `json:"targets"`
	Links   []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetTargets() (resp *getTargetsResp, err error) {
	err = xms.query("GET", "/types/targets", "", nil, nil, &resp)
	return resp, err
}

type Target struct {
	PortType    string          `json:"port-type"`
	PortAddress string          `json:"port-address"`
	SysName     string          `json:"sys-name"`
	PortalList  [][]interface{} `json:"portal-list"`
}

type getTargetResp struct {
	Content *Target `json:"content"`
	Links   []struct {
		GUID string `json:"guid"`
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetTarget(id string) (resp *getTargetResp, err error) {
	err = xms.query("GET", "/types/targets", id, nil, nil, &resp)
	return resp, err
}

type getEventsResp struct {
	Events []*Event `json:"events"`
	Links  []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

type Event struct {
	EntityDetails  string `json:"entity_details"`
	Severity       string `json:"severity"`
	Classification string `json:"classification"`
	Timestamp      string `json:"timestamp"`
	Entity         string `json:"entity"`
	EventCode      string `json:"event_code"`
	ID             int    `json:"id"`
	Description    string `json:"description"`
}

func (xms *XMS) GetEvents(severity string) (resp *getEventsResp, err error) {
	err = xms.query("GET", "/types/events", "", map[string]string{"severity": severity}, nil, &resp)
	return resp, err
}

type getSnapshotsResp struct {
	Snapshots []*Ref `json:"snapshots"`
	Links     []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetSnapshots() (resp *getSnapshotsResp, err error) {
	err = xms.query("GET", "/types/snapshots", "", nil, nil, &resp)
	return resp, err
}

type Snapshot struct {
	UnalignedIoAlerts     string        `json:"unaligned-io-alerts"`
	SmallIoAlerts         string        `json:"small-io-alerts"`
	SmallIops             string        `json:"small-iops"`
	DestSnapList          []interface{} `json:"dest-snap-list"`
	VolID                 []interface{} `json:"vol-id"`
	ObjSeverity           string        `json:"obj-severity"`
	NumOfLunMappings      int           `json:"num-of-lun-mappings"`
	UnalignedRdBw         string        `json:"unaligned-rd-bw"`
	NumOfDestSnaps        int           `json:"num-of-dest-snaps"`
	Iops                  string        `json:"iops"`
	AccNumOfSmallWr       string        `json:"acc-num-of-small-wr"`
	AlignmentOffset       int           `json:"alignment-offset"`
	LbSize                int           `json:"lb-size"`
	LogicalSpaceInUse     string        `json:"logical-space-in-use"`
	UnalignedIoRatioLevel string        `json:"unaligned-io-ratio-level"`
	AccNumOfRd            string        `json:"acc-num-of-rd"`
	Index                 int           `json:"index"`
	SmallRdBw             string        `json:"small-rd-bw"`
	NaaName               string        `json:"naa-name"`
	AccSizeOfWr           string        `json:"acc-size-of-wr"`
	AccNumOfSmallRd       string        `json:"acc-num-of-small-rd"`
	UnalignedRdIops       string        `json:"unaligned-rd-iops"`
	WrLatency             string        `json:"wr-latency"`
	SnapgrpID             []interface{} `json:"snapgrp-id"`
	AncestorVolID         []interface{} `json:"ancestor-vol-id"`
	VaaiTpAlerts          string        `json:"vaai-tp-alerts"`
	CreationTime          string        `json:"creation-time"`
	RdBw                  string        `json:"rd-bw"`
	XmsID                 []interface{} `json:"xms-id"`
	Compressible          string        `json:"compressible"`
	SmallWrBw             string        `json:"small-wr-bw"`
	AccNumOfUnalignedRd   string        `json:"acc-num-of-unaligned-rd"`
	LuName                string        `json:"lu-name"`
	UnalignedIops         string        `json:"unaligned-iops"`
	UnalignedBw           string        `json:"unaligned-bw"`
	Bw                    string        `json:"bw"`
	SmallIoRatioLevel     string        `json:"small-io-ratio-level"`
	LunMappingList        []interface{} `json:"lun-mapping-list"`
	VolSize               string        `json:"vol-size"`
	WrIops                string        `json:"wr-iops"`
	SysID                 []interface{} `json:"sys-id"`
	AvgLatency            string        `json:"avg-latency"`
	RdLatency             string        `json:"rd-latency"`
	SmallWrIops           string        `json:"small-wr-iops"`
	SmallBw               string        `json:"small-bw"`
	Name                  string        `json:"name"`
	AccNumOfUnalignedWr   string        `json:"acc-num-of-unaligned-wr"`
	UnalignedWrIops       string        `json:"unaligned-wr-iops"`
	AccNumOfWr            string        `json:"acc-num-of-wr"`
	SmallIoRatio          string        `json:"small-io-ratio"`
	AccSizeOfRd           string        `json:"acc-size-of-rd"`
	UnalignedWrBw         string        `json:"unaligned-wr-bw"`
	SmallRdIops           string        `json:"small-rd-iops"`
	UnalignedIoRatio      string        `json:"unaligned-io-ratio"`
	RdIops                string        `json:"rd-iops"`
	WrBw                  string        `json:"wr-bw"`
}

type getSnapshotResp struct {
	Content *Snapshot `json:"content"`
	Links   []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) GetSnapshot(id string, name string) (resp *getSnapshotResp, err error) {
	if id == "" && name == "" {
		return nil, ErrIdOrNameNotSpecified
	}

	err = xms.query("GET", "/types/snapshots", id, map[string]string{"name": name}, nil, &resp)
	return resp, err
}

type UserAccount struct {
	Index int `json:"index"`
}

type getUserAccountResp struct {
	Content *UserAccount `json:"content"`
	Links   []Link       `json:"links"`
}

func (xms *XMS) GetUserAccount(id string, name string) (resp *getUserAccountResp, err error) {
	if id == "" && name == "" {
		return nil, ErrIdOrNameNotSpecified
	}

	err = xms.query("GET", "/types/user-accounts", id, map[string]string{"name": name}, nil, &resp)
	return resp, err
}

type PostUserAccountReq struct {
	Name     string `json:"usr-name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type PostUserAccountResp struct {
	Links []Link `json:"links"`
}

func (xms *XMS) PostUserAccount(req *PostUserAccountReq) (resp *PostUserAccountResp, err error) {
	err = xms.query("POST", "/types/user-accounts", "", nil, req, &resp)
	return resp, err
}

type PostQoSPolicyReq struct {
	LimitType       string      `json:"limit-type"`
	MaxBw           string      `json:"max-bw,omitempty"`
	SysID           interface{} `json:"cluster-id,omitempty"`
	QosPolicyName   string      `json:"qos-policy-name"`
	MaxIOPS         int         `json:"max-iops,omitempty"`
	BurstPercentage string      `json:"burst-percentage,omitempty"`
	IoSize          string      `json:"io-size,omitempty"`
}

type PostQoSPolicyResp struct {
	Links []Link `json:"links"`
}

func (xms *XMS) PostQoSPolicy(req *PostQoSPolicyReq) (resp *PostQoSPolicyResp, err error) {
	err = xms.query("POST", "/types/qos-policies", "", nil, req, &resp)
	return resp, err
}

type PostVolumesReq struct {
	AlignmentOffset         int         `json:"alignment-offset,omitempty"`
	LbSize                  int         `json:"lb-size,omitempty"`
	SysID                   interface{} `json:"cluster-id,omitempty"`
	VolName                 string      `json:"vol-name,omitempty"`
	VolSize                 int         `json:"vol-size"`
	ParentFolderID          string      `json:"parent-folder-id,omitempty"`
	CreatedByExternalClient string      `json:"created-by-external-client,omitempty"`
	QoSPolicyID             interface{} `json:"qos-policy-id,omitempty"`
}

type PostVolumesResp struct {
	Links []Link `json:"links"`
}

func (xms *XMS) PostVolumes(req *PostVolumesReq) (resp *PostVolumesResp, err error) {
	err = xms.query("POST", "/types/volumes", "", nil, req, &resp)
	return resp, err
}

func (xms *XMS) DeleteVolumes(id string, name string) (err error) {
	err = xms.query("DELETE", "/types/volumes", id, map[string]string{"name": name}, nil, nil)
	return err
}

type PutVolumeReq struct {
	VolName         string `json:"vol-name,omitempty"`
	VolSize         int    `json:"vol-size,omitempty"`
	QoSPolicyID     int    `json:"qos-policy-id,omitempty"`
	RemoveQoSPolicy string `json:"remove-qos-policy,omitempty"`
}

func (xms *XMS) PutVolume(id string, name string, req *PutVolumeReq) (err error) {
	if id == "" && name == "" {
		return ErrIdOrNameNotSpecified
	}

	err = xms.query("PUT", "/types/volumes", id, map[string]string{"name": name}, req, nil)
	return err
}

type Item struct {
	GUID  string `json:"guid"`
	Href  string `json:"href"`
	Index int    `json:"index"`
	Name  string `json:"name"`
}

type VolumeCreateCopyReq struct {
	FromVolumeList          [][]interface{} `json:"from-volume-list"`
	SysID                   interface{}     `json:"cluster-id,omitempty"`
	CreatedByExternalClient string          `json:"created-by-external-client,omitempty"`
	ManagementLocked        string          `json:"management-locked,omitempty"`
	NewVolSuffix            string          `json:"new-vol-suffix,omitempty"`
	NewSnapshotSetName      string          `json:"new-snapshot-set-name,omitempty"`
}

type VolumeCreateCopyResponseContent struct {
	Volumes     []Item `json:"volumes"`
	SnapshotSet Item   `json:"snapshot-set"`
}

type VolumeCreateCopyResp struct {
	Content *VolumeCreateCopyResponseContent `json:"content"`
	Links   []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) VolumeCreateRepurposeCopy(req *VolumeCreateCopyReq) (resp *VolumeCreateCopyResp, err error) {
	err = xms.query("PUT", "/commands/volumes/create-repurpose-copy", "", nil, req, &resp)
	return resp, err
}

func (xms *XMS) VolumeCreateProtectionCopy(req *VolumeCreateCopyReq) (resp *VolumeCreateCopyResp, err error) {
	err = xms.query("PUT", "/commands/volumes/create-protection-copy", "", nil, req, &resp)
	return resp, err
}

type PostLunMapsReq struct {
	SysID   interface{} `json:"cluster-id,omitempty"`
	VolID   int         `json:"vol-id,omitempty"`
	VolName string      `json:"vol-name,omitempty"`
	IgID    interface{} `json:"ig-id"`
	Lun     int         `json:"lun,omitempty"`
	TgID    int         `json:"tg-id,omitempty"`
	TgName  string      `json:"tg-name,omitempty"`
}

type PostLunMapsResp struct {
	Links []Link `json:"links"`
}

func (xms *XMS) PostLunMaps(req *PostLunMapsReq) (resp *PostLunMapsResp, err error) {
	err = xms.query("POST", "/types/lun-maps", "", nil, req, &resp)
	return resp, err
}

func (xms *XMS) DeleteLunMaps(id string, name string) (err error) {
	err = xms.query("DELETE", "/types/lun-maps", id, map[string]string{"name": name}, nil, nil)
	return err
}

type PostInitiatorsReq struct {
	SysID                             interface{} `json:"cluster-id,omitempty"`
	IgID                              int         `json:"ig-id,omitempty"`
	IgName                            string      `json:"ig-name,omitempty"`
	InitiatorName                     string      `json:"initiator-name"`
	PortAddress                       string      `json:"port-address"`
	InitiatorAuthenticationPassword   string      `json:"initiator-authentication-password,omitempty"`
	InitiatorAuthenticationUserName   string      `json:"initiator-authentication-user-name,omitempty"`
	InitiatorDiscoveryPassword        string      `json:"initiator-discovery-password,omitempty"`
	InitiatorDiscoveryUserName        string      `json:"initiator-discovery-user-name,omitempty"`
	InitiatorChapDiscoveryCredentials string      `json:"initiator-chap-discovery-credentials,omitempty"`
}

type PostInitiatorsResp struct {
	Links []Link `json:"links"`
}

func (xms *XMS) PostInitiators(req *PostInitiatorsReq) (resp *PostInitiatorsResp, err error) {
	err = xms.query("POST", "/types/initiators", "", nil, req, &resp)
	return resp, err
}

func (xms *XMS) DeleteInitiators(id string, name string) (err error) {
	err = xms.query("DELETE", "/types/initiators", id, map[string]string{"name": name}, nil, nil)
	return err
}

type PostInitiatorGroupsReq struct {
	SysID                   interface{} `json:"cluster-id,omitempty"`
	IgName                  string      `json:"ig-name,omitempty"`
	QoSEnabledMode          string      `json:"qos-enabled-mode,omitempty"`
	QoSPolicyID             string      `json:"qos-policy-id,omitempty"`
	CreatedByExternalClient string      `json:"created-by-external-client,omitempty"`
}

type PostInitiatorGroupsResp struct {
	Links []Link `json:"links"`
}

func (xms *XMS) PostInitiatorGroups(req *PostInitiatorGroupsReq) (resp *PostInitiatorGroupsResp, err error) {
	err = xms.query("POST", "/types/initiator-groups", "", nil, req, &resp)
	return resp, err
}

type PutInitiatorGroupsReq struct {
	IgName         string `json:"new-name,omitempty"`
	QoSEnabledMode string `json:"qos-enabled-mode,omitempty"`
	QoSPolicyID    string `json:"qos-policy-id,omitempty"`
}

func (xms *XMS) PutInitiatorGroup(id string, req *PutInitiatorGroupsReq) (err error) {
	err = xms.query("PUT", "/types/initiator-groups", id, nil, req, nil)
	return err
}

type SnapListItem struct {
	AncestorVolID string `json:"ancestor-vol-id"`
	SnapVolName   string `json:"snap-vol-name"`
}

type PostSnapshotsReq struct {
	SnapList []*SnapListItem `json:"snap-list"`
	FolderID string          `json:"folder-id"`
}

type PostSnapshotsResp struct {
	Links []Link `json:"links"`
}

func (xms *XMS) PostSnapshots(req *PostSnapshotsReq) (resp *PostSnapshotsResp, err error) {
	err = xms.query("POST", "/types/snapshots", "", nil, req, &resp)
	return resp, err
}

func (xms *XMS) DeleteSnapshots(id string, name string) (err error) {
	err = xms.query("DELETE", "/types/snapshots", id, map[string]string{"name": name}, nil, nil)
	return err
}
