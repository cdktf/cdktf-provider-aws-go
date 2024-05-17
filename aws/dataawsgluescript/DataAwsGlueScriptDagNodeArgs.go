// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsgluescript


type DataAwsGlueScriptDagNodeArgs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.50.0/docs/data-sources/glue_script#name DataAwsGlueScript#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.50.0/docs/data-sources/glue_script#value DataAwsGlueScript#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.50.0/docs/data-sources/glue_script#param DataAwsGlueScript#param}.
	Param interface{} `field:"optional" json:"param" yaml:"param"`
}

