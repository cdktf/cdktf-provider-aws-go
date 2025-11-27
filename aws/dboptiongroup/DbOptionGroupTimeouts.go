// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dboptiongroup


type DbOptionGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/db_option_group#delete DbOptionGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

