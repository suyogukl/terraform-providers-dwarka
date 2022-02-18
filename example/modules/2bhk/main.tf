module "one_bhk" {
  source        = "../1bhk"
  building_name = var.building_name
  floor_name    = var.floor_name
}

resource "dwarka_room" "bedroom_additional" {
  building_id = module.one_bhk.building_id
  floor_id    = module.one_bhk.floor_id
  name        = var.additional_bedroom_name
  description = "from terraform"
  direction   = "south"
}
