module "one_bkh" {
  count = var.no_of_1bhk_apartments

  source        = "../modules/1bhk"
  building_name = "${var.name}-one-bkh-${count.index}"
}

module "two_bkh" {
  count = var.no_of_2bhk_apartments

  source        = "../modules/2bhk"
  building_name = "${var.name}-two-bhk-${count.index}"
}

module "three_bkh" {
  count = var.no_of_3bhk_apartments

  source        = "../modules/3bhk"
  building_name = "${var.name}-three-bhk-${count.index}"
}
