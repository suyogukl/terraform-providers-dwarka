resource "dwarka_building" "main" {
  name        = "terraform"
  description = "from terraform"
  lat         = 13.0827
  lan         = 80.2707
}

resource "dwarka_floor" "first" {
  building_id = dwarka_building.main.id
  name        = "first floor"
  description = "from terraform 123"
  level = 1
}
