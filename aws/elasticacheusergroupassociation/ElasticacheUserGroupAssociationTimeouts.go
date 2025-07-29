// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticacheusergroupassociation


type ElasticacheUserGroupAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/elasticache_user_group_association#create ElasticacheUserGroupAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/elasticache_user_group_association#delete ElasticacheUserGroupAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

