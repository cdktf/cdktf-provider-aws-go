// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkforce


type SagemakerWorkforceSourceIpConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/sagemaker_workforce#cidrs SagemakerWorkforce#cidrs}.
	Cidrs *[]*string `field:"required" json:"cidrs" yaml:"cidrs"`
}

