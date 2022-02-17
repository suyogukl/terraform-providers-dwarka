locals {
  buildings = jsondecode(file("${path.module}/buildings.json"))
}

resource "dwarka_building" "buildings_via_count" {
  count = length(local.buildings)

  name        = local.buildings[count.index]["name"]
  description = "from terraform"
  lat         = 13.0827
  lan         = 80.2707
}

resource "dwarka_building" "buildings_via_loop" {
  for_each = {for building in local.buildings: building["name"] => building}

  name        = each.key
  description = "from terraform"
  lat         = 13.0827
  lan         = 80.2707
}
