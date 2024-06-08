// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package launchtemplate


type LaunchTemplateMaintenanceOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/launch_template#auto_recovery LaunchTemplate#auto_recovery}.
	AutoRecovery *string `field:"optional" json:"autoRecovery" yaml:"autoRecovery"`
}

