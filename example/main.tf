resource "dwarka_building" "main" {
  name        = "terraform"
  description = "from terraform"
  lat         = 13.0827
  lan         = 80.2707
}

resource "dwarka_floor" "first" {
  building_id = dwarka_building.main.id
  name        = "first floor"
  description = "from terraform"
  level       = 1
}

resource "dwarka_room" "one" {
  building_id = dwarka_building.main.id
  floor_id    = dwarka_floor.first.id
  name        = "main"
  direction   = "east"
  description = "from terraform"
}
