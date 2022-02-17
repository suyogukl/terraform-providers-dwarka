package dwarka

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetBuildings - Returns list of buildings
func (c *Client) GetBuildings() ([]Building, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/buildings", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	buildings := []Building{}
	err = json.Unmarshal(body, &buildings)
	if err != nil {
		return nil, err
	}

	return buildings, nil
}

// GetBuilding - Returns specific building
func (c *Client) GetBuilding(buildingID string) (*Building, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/buildings/%s", c.HostURL, buildingID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	building := Building{}
	err = json.Unmarshal(body, &building)
	if err != nil {
		return nil, err
	}

	return &building, nil
}

// CreateBuilding - Create new building
func (c *Client) CreateBuilding(building Building) (*string, error) {
	rb, err := json.Marshal(building)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/buildings", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	result := map[string]string{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	id := result["id"]
	return &id, nil
}

func (c *Client) UpdateBuilding(buildingID string, building Building) (*Building, error) {
	rb, err := json.Marshal(building)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/buildings/%s", c.HostURL, buildingID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return &building, nil
}

func (c *Client) DeleteBuilding(buildingID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/buildings/%s", c.HostURL, buildingID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}
