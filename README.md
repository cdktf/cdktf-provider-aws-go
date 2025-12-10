# CDKTF Go bindings for hashicorp/aws provider version 6.25.0

HashiCorp made the decision to stop publishing new versions of prebuilt [Terraform aws provider](https://registry.terraform.io/providers/hashicorp/aws/6.25.0) bindings for [CDK for Terraform](https://cdk.tf) on December 10, 2025. As such, this repository has been archived and is no longer supported in any way by HashiCorp. Previously-published versions of this prebuilt provider will still continue to be available on their respective package managers (e.g. npm, PyPi, Maven, NuGet), but these will not be compatible with new releases of `cdktf` past `0.21.0` and are no longer eligible for commercial support.

As a reminder, you can continue to use the `hashicorp/aws` provider in your CDK for Terraform (CDKTF) projects, even with newer versions of CDKTF, but you will need to generate the bindings locally. The easiest way to do so is to use the [`provider add` command](https://developer.hashicorp.com/terraform/cdktf/cli-reference/commands#provider-add), optionally with the `--force-local` flag enabled:

`cdktf provider add hashicorp/aws --force-local`

For more information and additional examples, check out our documentation on [generating provider bindings manually](https://cdk.tf/imports).

## Deprecated Package

The go package is generated into the [`github.com/cdktf/cdktf-provider-aws-go`](https://github.com/cdktf/cdktf-provider-aws-go) package.

`go get github.com/cdktf/cdktf-provider-aws-go/aws/<version>`

Where `<version>` is the version of the prebuilt provider you would like to use e.g. `v11`. The full module name can be found
within the [go.mod](https://github.com/cdktf/cdktf-provider-aws-go/blob/main/aws/go.mod#L1) file.

## Docs

Find auto-generated docs for this provider [here](https://github.com/cdktf/cdktf-provider-aws/blob/main/docs/API.go.md).

