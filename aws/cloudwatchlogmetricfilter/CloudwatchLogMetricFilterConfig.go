// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogmetricfilter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudwatchLogMetricFilterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudwatch_log_metric_filter#log_group_name CloudwatchLogMetricFilter#log_group_name}.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// metric_transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudwatch_log_metric_filter#metric_transformation CloudwatchLogMetricFilter#metric_transformation}
	MetricTransformation *CloudwatchLogMetricFilterMetricTransformation `field:"required" json:"metricTransformation" yaml:"metricTransformation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudwatch_log_metric_filter#name CloudwatchLogMetricFilter#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudwatch_log_metric_filter#pattern CloudwatchLogMetricFilter#pattern}.
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudwatch_log_metric_filter#apply_on_transformed_logs CloudwatchLogMetricFilter#apply_on_transformed_logs}.
	ApplyOnTransformedLogs interface{} `field:"optional" json:"applyOnTransformedLogs" yaml:"applyOnTransformedLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudwatch_log_metric_filter#id CloudwatchLogMetricFilter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudwatch_log_metric_filter#region CloudwatchLogMetricFilter#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

