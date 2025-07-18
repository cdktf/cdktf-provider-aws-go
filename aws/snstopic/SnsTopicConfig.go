// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package snstopic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SnsTopicConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#application_failure_feedback_role_arn SnsTopic#application_failure_feedback_role_arn}.
	ApplicationFailureFeedbackRoleArn *string `field:"optional" json:"applicationFailureFeedbackRoleArn" yaml:"applicationFailureFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#application_success_feedback_role_arn SnsTopic#application_success_feedback_role_arn}.
	ApplicationSuccessFeedbackRoleArn *string `field:"optional" json:"applicationSuccessFeedbackRoleArn" yaml:"applicationSuccessFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#application_success_feedback_sample_rate SnsTopic#application_success_feedback_sample_rate}.
	ApplicationSuccessFeedbackSampleRate *float64 `field:"optional" json:"applicationSuccessFeedbackSampleRate" yaml:"applicationSuccessFeedbackSampleRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#archive_policy SnsTopic#archive_policy}.
	ArchivePolicy *string `field:"optional" json:"archivePolicy" yaml:"archivePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#content_based_deduplication SnsTopic#content_based_deduplication}.
	ContentBasedDeduplication interface{} `field:"optional" json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#delivery_policy SnsTopic#delivery_policy}.
	DeliveryPolicy *string `field:"optional" json:"deliveryPolicy" yaml:"deliveryPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#display_name SnsTopic#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#fifo_throughput_scope SnsTopic#fifo_throughput_scope}.
	FifoThroughputScope *string `field:"optional" json:"fifoThroughputScope" yaml:"fifoThroughputScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#fifo_topic SnsTopic#fifo_topic}.
	FifoTopic interface{} `field:"optional" json:"fifoTopic" yaml:"fifoTopic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#firehose_failure_feedback_role_arn SnsTopic#firehose_failure_feedback_role_arn}.
	FirehoseFailureFeedbackRoleArn *string `field:"optional" json:"firehoseFailureFeedbackRoleArn" yaml:"firehoseFailureFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#firehose_success_feedback_role_arn SnsTopic#firehose_success_feedback_role_arn}.
	FirehoseSuccessFeedbackRoleArn *string `field:"optional" json:"firehoseSuccessFeedbackRoleArn" yaml:"firehoseSuccessFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#firehose_success_feedback_sample_rate SnsTopic#firehose_success_feedback_sample_rate}.
	FirehoseSuccessFeedbackSampleRate *float64 `field:"optional" json:"firehoseSuccessFeedbackSampleRate" yaml:"firehoseSuccessFeedbackSampleRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#http_failure_feedback_role_arn SnsTopic#http_failure_feedback_role_arn}.
	HttpFailureFeedbackRoleArn *string `field:"optional" json:"httpFailureFeedbackRoleArn" yaml:"httpFailureFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#http_success_feedback_role_arn SnsTopic#http_success_feedback_role_arn}.
	HttpSuccessFeedbackRoleArn *string `field:"optional" json:"httpSuccessFeedbackRoleArn" yaml:"httpSuccessFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#http_success_feedback_sample_rate SnsTopic#http_success_feedback_sample_rate}.
	HttpSuccessFeedbackSampleRate *float64 `field:"optional" json:"httpSuccessFeedbackSampleRate" yaml:"httpSuccessFeedbackSampleRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#id SnsTopic#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#kms_master_key_id SnsTopic#kms_master_key_id}.
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#lambda_failure_feedback_role_arn SnsTopic#lambda_failure_feedback_role_arn}.
	LambdaFailureFeedbackRoleArn *string `field:"optional" json:"lambdaFailureFeedbackRoleArn" yaml:"lambdaFailureFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#lambda_success_feedback_role_arn SnsTopic#lambda_success_feedback_role_arn}.
	LambdaSuccessFeedbackRoleArn *string `field:"optional" json:"lambdaSuccessFeedbackRoleArn" yaml:"lambdaSuccessFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#lambda_success_feedback_sample_rate SnsTopic#lambda_success_feedback_sample_rate}.
	LambdaSuccessFeedbackSampleRate *float64 `field:"optional" json:"lambdaSuccessFeedbackSampleRate" yaml:"lambdaSuccessFeedbackSampleRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#name SnsTopic#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#name_prefix SnsTopic#name_prefix}.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#policy SnsTopic#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#region SnsTopic#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#signature_version SnsTopic#signature_version}.
	SignatureVersion *float64 `field:"optional" json:"signatureVersion" yaml:"signatureVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#sqs_failure_feedback_role_arn SnsTopic#sqs_failure_feedback_role_arn}.
	SqsFailureFeedbackRoleArn *string `field:"optional" json:"sqsFailureFeedbackRoleArn" yaml:"sqsFailureFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#sqs_success_feedback_role_arn SnsTopic#sqs_success_feedback_role_arn}.
	SqsSuccessFeedbackRoleArn *string `field:"optional" json:"sqsSuccessFeedbackRoleArn" yaml:"sqsSuccessFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#sqs_success_feedback_sample_rate SnsTopic#sqs_success_feedback_sample_rate}.
	SqsSuccessFeedbackSampleRate *float64 `field:"optional" json:"sqsSuccessFeedbackSampleRate" yaml:"sqsSuccessFeedbackSampleRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#tags SnsTopic#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#tags_all SnsTopic#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sns_topic#tracing_config SnsTopic#tracing_config}.
	TracingConfig *string `field:"optional" json:"tracingConfig" yaml:"tracingConfig"`
}

