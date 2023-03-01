module "one" {
  source = "../onebhk"

  building_name = var.building_name
  floor_name = var.floor_name
}

resource "dwarka_room" "masterbedroom" {
  building_id = var.building_name
  floor_id    = var.floor_name
  name        = "master bedroom"
  description = "from terraform"
  direction   = "south"
}
