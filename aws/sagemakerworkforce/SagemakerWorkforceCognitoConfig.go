// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkforce


type SagemakerWorkforceCognitoConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/sagemaker_workforce#client_id SagemakerWorkforce#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/sagemaker_workforce#user_pool SagemakerWorkforce#user_pool}.
	UserPool *string `field:"required" json:"userPool" yaml:"userPool"`
}

