// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spotinstancerequest


type SpotInstanceRequestMaintenanceOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_instance_request#auto_recovery SpotInstanceRequest#auto_recovery}.
	AutoRecovery *string `field:"optional" json:"autoRecovery" yaml:"autoRecovery"`
}

