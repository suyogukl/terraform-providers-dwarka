package dwarka

import (
	"context"
	"time"

	"terraform-provider-dwarka/client/dwarka"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRoom() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRoomCreate,
		ReadContext:   resourceRoomRead,
		UpdateContext: resourceRoomUpdate,
		DeleteContext: resourceRoomDelete,
		Schema: map[string]*schema.Schema{
			"building_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"floor_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"direction": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceRoomCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dwarka.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	buildingID := d.Get("building_id").(string)
	floorID := d.Get("floor_id").(string)

	room := dwarka.Room{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Direction:   d.Get("direction").(string),
	}

	roomID, err := c.CreateRoom(buildingID, floorID, room)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*roomID)
	d.Set("building_id", buildingID)
	d.Set("floor_id", floorID)

	resourceRoomRead(ctx, d, m)

	return diags
}

func resourceRoomRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dwarka.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	roomID := d.Id()
	buildingID := d.Get("building_id").(string)
	floorID := d.Get("floor_id").(string)

	room, err := c.GetRoom(buildingID, floorID, roomID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", room.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", room.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("direction", room.Direction); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceRoomUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dwarka.Client)

	roomID := d.Id()
	buildingID := d.Get("building_id").(string)
	floorID := d.Get("floor_id").(string)

	if d.HasChanges("direction", "description") {
		room := dwarka.Room{
			Name:        d.Id(),
			Direction:   d.Get("direction").(string),
			Description: d.Get("description").(string),
		}

		_, err := c.UpdateRoom(buildingID, floorID, roomID, room)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("building_id", buildingID)
		d.Set("floor_id", buildingID)
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceRoomRead(ctx, d, m)
}

func resourceRoomDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dwarka.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	roomID := d.Id()
	buildingID := d.Get("building_id").(string)
	floorID := d.Get("floor_id").(string)

	err := c.DeleteRoom(buildingID, floorID, roomID)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
