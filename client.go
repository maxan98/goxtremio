package goxtremio

import (
	"os"
	"strconv"

	xms "github.com/emccode/goxtremio/api/v3"
)

type Client3 struct {
	api *xms.XMS
}

type Client Client3

func NewClient() (*Client, error) {
	insecure, _ := strconv.ParseBool(os.Getenv("GOXTREMIO_INSECURE"))
	return NewClientWithArgs(
		os.Getenv("GOXTREMIO_ENDPOINT"),
		insecure,
		os.Getenv("GOXTREMIO_USERNAME"),
		os.Getenv("GOXTREMIO_PASSWORD"))
}

func NewClientWithArgs(
	endpoint string,
	insecure bool,
	username, password string) (*Client, error) {

	api, err := xms.New(endpoint, insecure, username, password)
	if err != nil {
		return nil, err
	}

	return &Client{api}, nil
}

//GetXms returns XMS info
func (c *Client) GetXms() (*xms.Xms, error) {
	xmsResp, err := c.api.GetXms()
	if err != nil {
		return nil, err
	}
	return xmsResp.Content, nil
}

// GetClusters return list of all clusters
func (c *Client) GetClusters() (Refs, error) {
	clusters, err := c.api.GetClusters()
	if err != nil {
		return nil, err
	}
	return clusters.Clusters, nil
}

// GetCluster return info of specific clusters
func (c *Client) GetCluster(id string) (*xms.Cluster, error) {
	xmsResp, err := c.api.GetCluster(id)
	if err != nil {
		return nil, err
	}
	return xmsResp.Content, nil
}
