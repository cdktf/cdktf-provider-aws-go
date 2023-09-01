// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appconfigextension


type AppconfigExtensionActionPointAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/appconfig_extension#name AppconfigExtension#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/appconfig_extension#role_arn AppconfigExtension#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/appconfig_extension#uri AppconfigExtension#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/appconfig_extension#description AppconfigExtension#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

