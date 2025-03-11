// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ceanomalysubscription

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CeAnomalySubscriptionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/ce_anomaly_subscription#frequency CeAnomalySubscription#frequency}.
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/ce_anomaly_subscription#monitor_arn_list CeAnomalySubscription#monitor_arn_list}.
	MonitorArnList *[]*string `field:"required" json:"monitorArnList" yaml:"monitorArnList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/ce_anomaly_subscription#name CeAnomalySubscription#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// subscriber block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/ce_anomaly_subscription#subscriber CeAnomalySubscription#subscriber}
	Subscriber interface{} `field:"required" json:"subscriber" yaml:"subscriber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/ce_anomaly_subscription#account_id CeAnomalySubscription#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/ce_anomaly_subscription#id CeAnomalySubscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/ce_anomaly_subscription#tags CeAnomalySubscription#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/ce_anomaly_subscription#tags_all CeAnomalySubscription#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// threshold_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/ce_anomaly_subscription#threshold_expression CeAnomalySubscription#threshold_expression}
	ThresholdExpression *CeAnomalySubscriptionThresholdExpression `field:"optional" json:"thresholdExpression" yaml:"thresholdExpression"`
}

