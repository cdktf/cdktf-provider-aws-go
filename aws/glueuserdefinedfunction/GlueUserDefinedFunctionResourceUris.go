// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package glueuserdefinedfunction


type GlueUserDefinedFunctionResourceUris struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/glue_user_defined_function#resource_type GlueUserDefinedFunction#resource_type}.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/glue_user_defined_function#uri GlueUserDefinedFunction#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

