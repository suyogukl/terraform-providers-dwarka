locals {
  apartments = yamldecode(file("${path.module}/apartments.yaml"))
  name = local.apartments["name"]
  no_of_1bhk_apartments = local.apartments["1bhk"]
  no_of_2bhk_apartments = local.apartments["2bhk"]
  no_of_3bhk_apartments = local.apartments["3bhk"]
}

module "one_bkh" {
  count = local.no_of_1bhk_apartments

  source = "../modules/1bhk"
  building_name = "${local.name}-one-bkh-${count.index}"
}

module "two_bkh" {
  count = local.no_of_2bhk_apartments

  source = "../modules/2bhk"
  building_name = "${local.name}-two-bhk-${count.index}"
}

module "three_bkh" {
  count = local.no_of_3bhk_apartments

  source = "../modules/3bhk"
  building_name = "${local.name}-three-bhk-${count.index}"
}
