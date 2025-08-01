// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devopsgurunotificationchannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DevopsguruNotificationChannelConfig struct {
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
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/devopsguru_notification_channel#filters DevopsguruNotificationChannel#filters}
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/devopsguru_notification_channel#region DevopsguruNotificationChannel#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// sns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/devopsguru_notification_channel#sns DevopsguruNotificationChannel#sns}
	Sns interface{} `field:"optional" json:"sns" yaml:"sns"`
}

