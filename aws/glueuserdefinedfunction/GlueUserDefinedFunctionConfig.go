// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package glueuserdefinedfunction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlueUserDefinedFunctionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/glue_user_defined_function#class_name GlueUserDefinedFunction#class_name}.
	ClassName *string `field:"required" json:"className" yaml:"className"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/glue_user_defined_function#database_name GlueUserDefinedFunction#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/glue_user_defined_function#name GlueUserDefinedFunction#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/glue_user_defined_function#owner_name GlueUserDefinedFunction#owner_name}.
	OwnerName *string `field:"required" json:"ownerName" yaml:"ownerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/glue_user_defined_function#owner_type GlueUserDefinedFunction#owner_type}.
	OwnerType *string `field:"required" json:"ownerType" yaml:"ownerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/glue_user_defined_function#catalog_id GlueUserDefinedFunction#catalog_id}.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/glue_user_defined_function#id GlueUserDefinedFunction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// resource_uris block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/glue_user_defined_function#resource_uris GlueUserDefinedFunction#resource_uris}
	ResourceUris interface{} `field:"optional" json:"resourceUris" yaml:"resourceUris"`
}

