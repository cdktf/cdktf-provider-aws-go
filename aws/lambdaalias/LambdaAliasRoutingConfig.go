// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdaalias


type LambdaAliasRoutingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.50.0/docs/resources/lambda_alias#additional_version_weights LambdaAlias#additional_version_weights}.
	AdditionalVersionWeights *map[string]*float64 `field:"optional" json:"additionalVersionWeights" yaml:"additionalVersionWeights"`
}

