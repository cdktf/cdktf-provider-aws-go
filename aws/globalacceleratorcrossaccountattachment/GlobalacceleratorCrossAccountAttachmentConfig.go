// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package globalacceleratorcrossaccountattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlobalacceleratorCrossAccountAttachmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/globalaccelerator_cross_account_attachment#name GlobalacceleratorCrossAccountAttachment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/globalaccelerator_cross_account_attachment#principals GlobalacceleratorCrossAccountAttachment#principals}.
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/globalaccelerator_cross_account_attachment#resource GlobalacceleratorCrossAccountAttachment#resource}
	Resource interface{} `field:"optional" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/globalaccelerator_cross_account_attachment#tags GlobalacceleratorCrossAccountAttachment#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

