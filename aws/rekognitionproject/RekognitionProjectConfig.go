// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rekognitionproject

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RekognitionProjectConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/rekognition_project#name RekognitionProject#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/rekognition_project#auto_update RekognitionProject#auto_update}.
	AutoUpdate *string `field:"optional" json:"autoUpdate" yaml:"autoUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/rekognition_project#feature RekognitionProject#feature}.
	Feature *string `field:"optional" json:"feature" yaml:"feature"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/rekognition_project#timeouts RekognitionProject#timeouts}
	Timeouts *RekognitionProjectTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

