terraform {
  required_providers {
    xtratusipam = {
      version = "0.1"
      source  = "xtratuscloud/xtratusipam"
    }
  }
}

locals {
  ipam_url_dev   = "https://fna-we-c-ipam-01.azurewebsites.net"
  ipam_apiId_dev = "fb09120f-6fc4-4d82-91d8-69a47d73779e"
}

provider "xtratusipam" {
  api_url = local.ipam_url_dev
  token   = data.external.get_access_token.result.accessToken
}

##get an access token to storage resources
data "external" "get_access_token" {
  program = ["az", "account", "get-access-token", "--resource", "api://${local.ipam_apiId_dev}"]
}
