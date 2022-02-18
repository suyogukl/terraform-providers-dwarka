module "two_bhk" {
  source                  = "../2bhk"
  building_name           = var.building_name
  floor_name              = var.floor_name
  additional_bedroom_name = var.first_additional_bedroom_name
}

resource "dwarka_room" "bedroom_3" {
  building_id = module.two_bhk.building_id
  floor_id    = module.two_bhk.floor_id
  name        = var.second_additional_bedroom_name
  description = "from terraform"
  direction   = "south"
}
