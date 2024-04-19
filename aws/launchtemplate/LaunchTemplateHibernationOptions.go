// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package launchtemplate


type LaunchTemplateHibernationOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/launch_template#configured LaunchTemplate#configured}.
	Configured interface{} `field:"required" json:"configured" yaml:"configured"`
}

