// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appconfigextension


type AppconfigExtensionActionPoint struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/appconfig_extension#action AppconfigExtension#action}
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/appconfig_extension#point AppconfigExtension#point}.
	Point *string `field:"required" json:"point" yaml:"point"`
}

