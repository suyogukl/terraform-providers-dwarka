package dwarka

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetRooms - Returns list of rooms
func (c *Client) GetRooms(buildingID, floorID string) ([]Room, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/buildings/%s/floors/%s/rooms", c.HostURL, buildingID, floorID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	rooms := []Room{}
	err = json.Unmarshal(body, &rooms)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

// GetRoom - Returns specific room
func (c *Client) GetRoom(buildingID, floorID, roomID string) (*Room, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/buildings/%s/floors/%s/rooms/%s", c.HostURL, buildingID, floorID, roomID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	room := Room{}
	err = json.Unmarshal(body, &room)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

// CreateRoom - Create new room
func (c *Client) CreateRoom(buildingID, floorID string, room Room) (*string, error) {
	rb, err := json.Marshal(room)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/buildings/%s/floors/%s/rooms", c.HostURL, buildingID, floorID), strings.NewReader(string(rb)))
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

func (c *Client) UpdateRoom(buildingID, floorID, roomID string, room Room) (*Room, error) {
	rb, err := json.Marshal(room)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/buildings/%s/floors/%s/rooms/%s", c.HostURL, buildingID, floorID, roomID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (c *Client) DeleteRoom(buildingID, floorID, roomID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/buildings/%s/floors/%s/rooms/%s", c.HostURL, buildingID, floorID, roomID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}
