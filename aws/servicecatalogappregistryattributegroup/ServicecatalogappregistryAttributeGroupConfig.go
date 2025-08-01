// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicecatalogappregistryattributegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicecatalogappregistryAttributeGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/servicecatalogappregistry_attribute_group#attributes ServicecatalogappregistryAttributeGroup#attributes}.
	Attributes *string `field:"required" json:"attributes" yaml:"attributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/servicecatalogappregistry_attribute_group#name ServicecatalogappregistryAttributeGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/servicecatalogappregistry_attribute_group#description ServicecatalogappregistryAttributeGroup#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/servicecatalogappregistry_attribute_group#region ServicecatalogappregistryAttributeGroup#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/servicecatalogappregistry_attribute_group#tags ServicecatalogappregistryAttributeGroup#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

