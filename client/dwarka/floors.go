package dwarka

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetFloors - Returns list of floors
func (c *Client) GetFloors(buildingID string) ([]Floor, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/buildings/%s/floors", c.HostURL, buildingID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	floors := []Floor{}
	err = json.Unmarshal(body, &floors)
	if err != nil {
		return nil, err
	}

	return floors, nil
}

// GetFloor - Returns specific floor
func (c *Client) GetFloor(buildingID string, floorID string) (*Floor, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/buildings/%s/floors/%s", c.HostURL, buildingID, floorID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	floor := Floor{}
	err = json.Unmarshal(body, &floor)
	if err != nil {
		return nil, err
	}

	return &floor, nil
}

// CreateFloor - Create new floor
func (c *Client) CreateFloor(buildingID string, floor Floor) (*string, error) {
	rb, err := json.Marshal(floor)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/buildings/%s/floors", c.HostURL, buildingID), strings.NewReader(string(rb)))
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

func (c *Client) UpdateFloor(buildingID string, floorID string, floor Floor) (*Floor, error) {
	rb, err := json.Marshal(floor)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/buildings/%s/floors/%s", c.HostURL, buildingID, floorID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return &floor, nil
}

func (c *Client) DeleteFloor(buildingID, floorID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/buildings/%s/floors/%s", c.HostURL, buildingID, floorID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}
