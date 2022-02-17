resource "dwarka_building" "main" {
  name        = "basic building"
  description = "from terraform"
  lat         = 13.0827
  lan         = 80.2707
}

resource "dwarka_floor" "ground" {
  building_id = dwarka_building.main.id
  name        = "ground floor"
  description = "from terraform"
  level       = 1
}

resource "dwarka_room" "hall" {
  building_id = dwarka_building.main.id
  floor_id    = dwarka_floor.ground.id
  name        = "main room"
  description = "from terraform"
  direction   = "south"
}
