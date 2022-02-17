package dwarka

import (
	"context"
	"time"

	"terraform-provider-dwarka/client/dwarka"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFloor() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFloorCreate,
		ReadContext:   resourceFloorRead,
		UpdateContext: resourceFloorUpdate,
		DeleteContext: resourceFloorDelete,
		Schema: map[string]*schema.Schema{
			"building_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"level": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceFloorCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dwarka.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	buildingID := d.Get("building_id").(string)

	floor := dwarka.Floor{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Level:       d.Get("level").(int),
	}

	floorID, err := c.CreateFloor(buildingID, floor)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*floorID)
	d.Set("building_id", buildingID)

	resourceFloorRead(ctx, d, m)

	return diags
}

func resourceFloorRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dwarka.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	floorID := d.Id()
	buildingID := d.Get("building_id").(string)

	floor, err := c.GetFloor(buildingID, floorID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", floor.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", floor.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("level", floor.Level); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceFloorUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dwarka.Client)

	floorID := d.Id()
	buildingID := d.Get("building_id").(string)

	if d.HasChanges("name", "lat", "lan", "description") {
		floor := dwarka.Floor{
			Name:        d.Id(),
			Level:       d.Get("level").(int),
			Description: d.Get("description").(string),
		}

		_, err := c.UpdateFloor(buildingID, floorID, floor)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("building_id", buildingID)
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceFloorRead(ctx, d, m)
}

func resourceFloorDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dwarka.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	floorID := d.Id()
	buildingID := d.Get("building_id").(string)

	err := c.DeleteFloor(buildingID, floorID)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
