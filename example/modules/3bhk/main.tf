module "two_bhk" {
  source = "../2bhk"
  building_name = var.building_name
}

resource "dwarka_room" "bedroom_3" {
  building_id = module.two_bhk.building_id
  floor_id    = module.two_bhk.floor_id
  name        = "third bedroom"
  description = "from terraform"
  direction   = "south"
}
