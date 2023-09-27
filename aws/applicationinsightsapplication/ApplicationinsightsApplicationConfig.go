// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationinsightsApplicationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/applicationinsights_application#resource_group_name ApplicationinsightsApplication#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/applicationinsights_application#auto_config_enabled ApplicationinsightsApplication#auto_config_enabled}.
	AutoConfigEnabled interface{} `field:"optional" json:"autoConfigEnabled" yaml:"autoConfigEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/applicationinsights_application#auto_create ApplicationinsightsApplication#auto_create}.
	AutoCreate interface{} `field:"optional" json:"autoCreate" yaml:"autoCreate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/applicationinsights_application#cwe_monitor_enabled ApplicationinsightsApplication#cwe_monitor_enabled}.
	CweMonitorEnabled interface{} `field:"optional" json:"cweMonitorEnabled" yaml:"cweMonitorEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/applicationinsights_application#grouping_type ApplicationinsightsApplication#grouping_type}.
	GroupingType *string `field:"optional" json:"groupingType" yaml:"groupingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/applicationinsights_application#id ApplicationinsightsApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/applicationinsights_application#ops_center_enabled ApplicationinsightsApplication#ops_center_enabled}.
	OpsCenterEnabled interface{} `field:"optional" json:"opsCenterEnabled" yaml:"opsCenterEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/applicationinsights_application#ops_item_sns_topic_arn ApplicationinsightsApplication#ops_item_sns_topic_arn}.
	OpsItemSnsTopicArn *string `field:"optional" json:"opsItemSnsTopicArn" yaml:"opsItemSnsTopicArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/applicationinsights_application#tags ApplicationinsightsApplication#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/applicationinsights_application#tags_all ApplicationinsightsApplication#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

