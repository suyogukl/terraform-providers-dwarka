module "apartments" {
  source                = "../modules/apartments"
  name                  = var.name
  no_of_1bhk_apartments = var.no_of_1bhk_apartments
  no_of_2bhk_apartments = var.no_of_2bhk_apartments
  no_of_3bhk_apartments = var.no_of_3bhk_apartments
}
