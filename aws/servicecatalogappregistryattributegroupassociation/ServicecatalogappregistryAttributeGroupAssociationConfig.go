// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicecatalogappregistryattributegroupassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicecatalogappregistryAttributeGroupAssociationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// ID of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/servicecatalogappregistry_attribute_group_association#application_id ServicecatalogappregistryAttributeGroupAssociation#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// ID of the attribute group to associate with the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/servicecatalogappregistry_attribute_group_association#attribute_group_id ServicecatalogappregistryAttributeGroupAssociation#attribute_group_id}
	AttributeGroupId *string `field:"required" json:"attributeGroupId" yaml:"attributeGroupId"`
}

