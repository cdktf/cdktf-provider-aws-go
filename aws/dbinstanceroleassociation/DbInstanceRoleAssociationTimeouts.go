// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dbinstanceroleassociation


type DbInstanceRoleAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/db_instance_role_association#create DbInstanceRoleAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/db_instance_role_association#delete DbInstanceRoleAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

