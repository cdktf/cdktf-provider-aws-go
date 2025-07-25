// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchloganomalydetector

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudwatchLogAnomalyDetectorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/cloudwatch_log_anomaly_detector#enabled CloudwatchLogAnomalyDetector#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/cloudwatch_log_anomaly_detector#log_group_arn_list CloudwatchLogAnomalyDetector#log_group_arn_list}.
	LogGroupArnList *[]*string `field:"required" json:"logGroupArnList" yaml:"logGroupArnList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/cloudwatch_log_anomaly_detector#anomaly_visibility_time CloudwatchLogAnomalyDetector#anomaly_visibility_time}.
	AnomalyVisibilityTime *float64 `field:"optional" json:"anomalyVisibilityTime" yaml:"anomalyVisibilityTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/cloudwatch_log_anomaly_detector#detector_name CloudwatchLogAnomalyDetector#detector_name}.
	DetectorName *string `field:"optional" json:"detectorName" yaml:"detectorName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/cloudwatch_log_anomaly_detector#evaluation_frequency CloudwatchLogAnomalyDetector#evaluation_frequency}.
	EvaluationFrequency *string `field:"optional" json:"evaluationFrequency" yaml:"evaluationFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/cloudwatch_log_anomaly_detector#filter_pattern CloudwatchLogAnomalyDetector#filter_pattern}.
	FilterPattern *string `field:"optional" json:"filterPattern" yaml:"filterPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/cloudwatch_log_anomaly_detector#kms_key_id CloudwatchLogAnomalyDetector#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/cloudwatch_log_anomaly_detector#region CloudwatchLogAnomalyDetector#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/cloudwatch_log_anomaly_detector#tags CloudwatchLogAnomalyDetector#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

