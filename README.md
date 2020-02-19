# terraform-provider-compose
[![Go Report Card](https://goreportcard.com/badge/github.com/ibm/terraform-provider-compose)](https://goreportcard.com/report/github.com/ibm/terraform-provider-compose) [![Build Status](https://travis-ci.com/reevejd/terraform-provider-compose.svg?branch=master)](https://travis-ci.com/IBM/terraform-provider-compose)

## Installation:

Download a [release for your operating system](https://github.com/IBM/terraform-provider-compose/releases) and place it in your [third-party plugins directory](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).

## Usage

Export your compose api key as an environment variable:
```
export COMPOSE_API_KEY="my compose api key"
```

or directly in your terraform configuration file (not recommended):

```
provider "compose" {
  api_token = "my compose api key"
}
```
### Example
```
provider "compose" {}

data "compose_account" "account" {}

resource "compose_deployment" "postgres" {
  name       = "mypostgres"
  account_id = "${data.compose_account.account.id}"
  datacenter = "aws:us-east-1"
  type       = "postgresql"
  units      = 1
  version = "9.6.16"
}

provider "postgresql" {
  host     = "${compose_deployment.deployment.connection_details.0.host}"
  port     = "${compose_deployment.deployment.connection_details.0.port}"
  database = "${compose_deployment.deployment.connection_details.0.database}"
  username = "${compose_deployment.deployment.connection_details.0.admin_username}"
  password = "${compose_deployment.deployment.connection_details.0.admin_password}"
}

resource "postgresql_database" "test" {
  name = "test"
}
```

## Contributing

For information on how to build and develop this project, see [CONTRIBUTING.md](CONTRIBUTING.md).
