// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdacapacityprovider


type LambdaCapacityProviderPermissionsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_capacity_provider#capacity_provider_operator_role_arn LambdaCapacityProvider#capacity_provider_operator_role_arn}.
	CapacityProviderOperatorRoleArn *string `field:"required" json:"capacityProviderOperatorRoleArn" yaml:"capacityProviderOperatorRoleArn"`
}

