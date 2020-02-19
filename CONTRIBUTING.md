# Contributing

Contributions to this project are welcome.

## Requirements
- Terraform v0.11+
- Go 1.13.x (to build the provider plugin)

## Building

Clone this repository to `$GOPATH/src/github.com/IBM/terraform-provider-compose`:

```
mkdir -p $GOPATH/src/github.com/IBM; cd $GOPATH/src/github.com/IBM
git clone git@github.com:IBM/terraform-provider-compose.git
```

Enter the provider directory and build the provider:

```
cd $GOPATH/src/github.com/IBM/terraform-provider-compose
make build
```

## Tests

In order to run tests, make sure you have the `COMPOSE_API_TOKEN` environment variable set,
and run `make testacc`. These tests will create real resources, and _will cost money to run_.
