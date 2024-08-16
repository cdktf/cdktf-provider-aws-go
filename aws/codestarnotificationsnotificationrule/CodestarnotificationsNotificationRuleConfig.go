// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codestarnotificationsnotificationrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodestarnotificationsNotificationRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/codestarnotifications_notification_rule#detail_type CodestarnotificationsNotificationRule#detail_type}.
	DetailType *string `field:"required" json:"detailType" yaml:"detailType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/codestarnotifications_notification_rule#event_type_ids CodestarnotificationsNotificationRule#event_type_ids}.
	EventTypeIds *[]*string `field:"required" json:"eventTypeIds" yaml:"eventTypeIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/codestarnotifications_notification_rule#name CodestarnotificationsNotificationRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/codestarnotifications_notification_rule#resource CodestarnotificationsNotificationRule#resource}.
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/codestarnotifications_notification_rule#id CodestarnotificationsNotificationRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/codestarnotifications_notification_rule#status CodestarnotificationsNotificationRule#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/codestarnotifications_notification_rule#tags CodestarnotificationsNotificationRule#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/codestarnotifications_notification_rule#tags_all CodestarnotificationsNotificationRule#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/codestarnotifications_notification_rule#target CodestarnotificationsNotificationRule#target}
	Target interface{} `field:"optional" json:"target" yaml:"target"`
}

