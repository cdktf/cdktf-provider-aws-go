// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package instance


type InstanceMaintenanceOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/instance#auto_recovery Instance#auto_recovery}.
	AutoRecovery *string `field:"optional" json:"autoRecovery" yaml:"autoRecovery"`
}

