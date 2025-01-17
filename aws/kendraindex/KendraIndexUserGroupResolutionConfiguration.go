// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kendraindex


type KendraIndexUserGroupResolutionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/kendra_index#user_group_resolution_mode KendraIndex#user_group_resolution_mode}.
	UserGroupResolutionMode *string `field:"required" json:"userGroupResolutionMode" yaml:"userGroupResolutionMode"`
}

