module "one_bhk" {
  source = "../1bhk"
  building_name = var.building_name
}

resource "dwarka_room" "bedroom_2" {
  building_id = module.one_bhk.building_id
  floor_id    = module.one_bhk.floor_id
  name        = "second bedroom"
  description = "from terraform"
  direction   = "south"
}
