// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqsqueue

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SqsQueueConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#content_based_deduplication SqsQueue#content_based_deduplication}.
	ContentBasedDeduplication interface{} `field:"optional" json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#deduplication_scope SqsQueue#deduplication_scope}.
	DeduplicationScope *string `field:"optional" json:"deduplicationScope" yaml:"deduplicationScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#delay_seconds SqsQueue#delay_seconds}.
	DelaySeconds *float64 `field:"optional" json:"delaySeconds" yaml:"delaySeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#fifo_queue SqsQueue#fifo_queue}.
	FifoQueue interface{} `field:"optional" json:"fifoQueue" yaml:"fifoQueue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#fifo_throughput_limit SqsQueue#fifo_throughput_limit}.
	FifoThroughputLimit *string `field:"optional" json:"fifoThroughputLimit" yaml:"fifoThroughputLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#id SqsQueue#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#kms_data_key_reuse_period_seconds SqsQueue#kms_data_key_reuse_period_seconds}.
	KmsDataKeyReusePeriodSeconds *float64 `field:"optional" json:"kmsDataKeyReusePeriodSeconds" yaml:"kmsDataKeyReusePeriodSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#kms_master_key_id SqsQueue#kms_master_key_id}.
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#max_message_size SqsQueue#max_message_size}.
	MaxMessageSize *float64 `field:"optional" json:"maxMessageSize" yaml:"maxMessageSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#message_retention_seconds SqsQueue#message_retention_seconds}.
	MessageRetentionSeconds *float64 `field:"optional" json:"messageRetentionSeconds" yaml:"messageRetentionSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#name SqsQueue#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#name_prefix SqsQueue#name_prefix}.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#policy SqsQueue#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#receive_wait_time_seconds SqsQueue#receive_wait_time_seconds}.
	ReceiveWaitTimeSeconds *float64 `field:"optional" json:"receiveWaitTimeSeconds" yaml:"receiveWaitTimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#redrive_allow_policy SqsQueue#redrive_allow_policy}.
	RedriveAllowPolicy *string `field:"optional" json:"redriveAllowPolicy" yaml:"redriveAllowPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#redrive_policy SqsQueue#redrive_policy}.
	RedrivePolicy *string `field:"optional" json:"redrivePolicy" yaml:"redrivePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#sqs_managed_sse_enabled SqsQueue#sqs_managed_sse_enabled}.
	SqsManagedSseEnabled interface{} `field:"optional" json:"sqsManagedSseEnabled" yaml:"sqsManagedSseEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#tags SqsQueue#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#tags_all SqsQueue#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sqs_queue#visibility_timeout_seconds SqsQueue#visibility_timeout_seconds}.
	VisibilityTimeoutSeconds *float64 `field:"optional" json:"visibilityTimeoutSeconds" yaml:"visibilityTimeoutSeconds"`
}

