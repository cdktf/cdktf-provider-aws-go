// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdacapacityprovider


type LambdaCapacityProviderInstanceRequirements struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lambda_capacity_provider#allowed_instance_types LambdaCapacityProvider#allowed_instance_types}.
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lambda_capacity_provider#architectures LambdaCapacityProvider#architectures}.
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lambda_capacity_provider#excluded_instance_types LambdaCapacityProvider#excluded_instance_types}.
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
}

