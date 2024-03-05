// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdaprovisionedconcurrencyconfig


type LambdaProvisionedConcurrencyConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/lambda_provisioned_concurrency_config#create LambdaProvisionedConcurrencyConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/lambda_provisioned_concurrency_config#update LambdaProvisionedConcurrencyConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

