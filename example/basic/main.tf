module "onebhk" {
  source = "./modules/onebhk"

  building_name = "basic building"
  floor_name = "ground floor"
}

module "twobhk" {
  source = "./modules/twobhk"

  building_name = "ambani's building"
  floor_name = "ground floor"
}
