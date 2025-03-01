package seaweedfs

import (
	"encoding/json"
	"fmt"
)

type DirAssignResp struct {
	Count     int    `json:"count"`
	FID       string `json:"fid"`
	URL       string `json:"url"`
	PublicURL string `json:"publicUrl"`
}

func (c *Client) DirAssign() (DirAssignResp, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("%s/dir/assign", c.config.MasterURL))
	if err != nil {
		return DirAssignResp{}, err
	}

	var dirAssignResp DirAssignResp
	err = json.NewDecoder(resp.Body).Decode(&dirAssignResp)
	if err != nil {
		return DirAssignResp{}, err
	}

	return dirAssignResp, nil
}

type DirLookupResp struct {
	Locations []Location `json:"locations"`
}

type Location struct {
	URL       string `json:"url"`
	PublicURL string `json:"publicUrl"`
}

func (c *Client) DirLookup(volumeID string) (DirLookupResp, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("%s/dir/lookup?volumeId=%s", c.config.MasterURL, volumeID))
	if err != nil {
		return DirLookupResp{}, err
	}

	var dirLookupResp DirLookupResp
	err = json.NewDecoder(resp.Body).Decode(&dirLookupResp)
	if err != nil {
		return DirLookupResp{}, err
	}

	return dirLookupResp, nil
}
