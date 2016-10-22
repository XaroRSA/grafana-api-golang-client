package gapi

import (
	"io/ioutil"
	"encoding/json"
	"errors"
)

type Datasource struct {
	ID int `json:"id"`
	OrgID int `json:"orgId"`
	Name string `json:"name"`
	Type string `json:"type"`
	TypeLogoURL string `json:"typeLogoUrl"`
	Access string `json:"access"`
	URL string `json:"url"`
	Password string `json:"password"`
	User string `json:"user"`
	Database string `json:"database"`
	BasicAuth bool `json:"basicAuth"`
	BasicAuthUser string `json:"basicAuthUser"`
	BasicAuthPassword string `json:"basicAuthPassword"`
	WithCredentials bool `json:"withCredentials"`
	IsDefault bool `json:"isDefault"`
}

func (c *Client) Datasources_all() ([]Datasource, error) {
	datasources := make([]Datasource, 0)
	req, err := c.newRequest("GET", "/api/datasources", nil)
	if err != nil {
		return datasources, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return datasources, err
	}
	if resp.StatusCode != 200 {
		return datasources, errors.New(resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return datasources, err
	}
	err = json.Unmarshal(data, &datasources)
	if err != nil {
		return datasources, err
	}
	return datasources, err
}