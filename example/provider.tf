terraform {
  required_providers {
    dwarka = {
      version = "0.1.1"
      source  = "github.com/jskswamy/dwarka"
    }
  }
}

provider "dwarka" {
  host = "http://localhost:1410"
}
