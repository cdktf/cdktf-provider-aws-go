// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pinpointsmsvoicev2configurationset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Pinpointsmsvoicev2ConfigurationSetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.0/docs/resources/pinpointsmsvoicev2_configuration_set#name Pinpointsmsvoicev2ConfigurationSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.0/docs/resources/pinpointsmsvoicev2_configuration_set#default_message_type Pinpointsmsvoicev2ConfigurationSet#default_message_type}.
	DefaultMessageType *string `field:"optional" json:"defaultMessageType" yaml:"defaultMessageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.0/docs/resources/pinpointsmsvoicev2_configuration_set#default_sender_id Pinpointsmsvoicev2ConfigurationSet#default_sender_id}.
	DefaultSenderId *string `field:"optional" json:"defaultSenderId" yaml:"defaultSenderId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.0/docs/resources/pinpointsmsvoicev2_configuration_set#tags Pinpointsmsvoicev2ConfigurationSet#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

