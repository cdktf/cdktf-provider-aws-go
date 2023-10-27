// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LambdaEventSourceMappingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#function_name LambdaEventSourceMapping#function_name}.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// amazon_managed_kafka_event_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#amazon_managed_kafka_event_source_config LambdaEventSourceMapping#amazon_managed_kafka_event_source_config}
	AmazonManagedKafkaEventSourceConfig *LambdaEventSourceMappingAmazonManagedKafkaEventSourceConfig `field:"optional" json:"amazonManagedKafkaEventSourceConfig" yaml:"amazonManagedKafkaEventSourceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#batch_size LambdaEventSourceMapping#batch_size}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#bisect_batch_on_function_error LambdaEventSourceMapping#bisect_batch_on_function_error}.
	BisectBatchOnFunctionError interface{} `field:"optional" json:"bisectBatchOnFunctionError" yaml:"bisectBatchOnFunctionError"`
	// destination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#destination_config LambdaEventSourceMapping#destination_config}
	DestinationConfig *LambdaEventSourceMappingDestinationConfig `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// document_db_event_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#document_db_event_source_config LambdaEventSourceMapping#document_db_event_source_config}
	DocumentDbEventSourceConfig *LambdaEventSourceMappingDocumentDbEventSourceConfig `field:"optional" json:"documentDbEventSourceConfig" yaml:"documentDbEventSourceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#enabled LambdaEventSourceMapping#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#event_source_arn LambdaEventSourceMapping#event_source_arn}.
	EventSourceArn *string `field:"optional" json:"eventSourceArn" yaml:"eventSourceArn"`
	// filter_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#filter_criteria LambdaEventSourceMapping#filter_criteria}
	FilterCriteria *LambdaEventSourceMappingFilterCriteria `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#function_response_types LambdaEventSourceMapping#function_response_types}.
	FunctionResponseTypes *[]*string `field:"optional" json:"functionResponseTypes" yaml:"functionResponseTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#id LambdaEventSourceMapping#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#maximum_batching_window_in_seconds LambdaEventSourceMapping#maximum_batching_window_in_seconds}.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#maximum_record_age_in_seconds LambdaEventSourceMapping#maximum_record_age_in_seconds}.
	MaximumRecordAgeInSeconds *float64 `field:"optional" json:"maximumRecordAgeInSeconds" yaml:"maximumRecordAgeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#maximum_retry_attempts LambdaEventSourceMapping#maximum_retry_attempts}.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#parallelization_factor LambdaEventSourceMapping#parallelization_factor}.
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#queues LambdaEventSourceMapping#queues}.
	Queues *[]*string `field:"optional" json:"queues" yaml:"queues"`
	// scaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#scaling_config LambdaEventSourceMapping#scaling_config}
	ScalingConfig *LambdaEventSourceMappingScalingConfig `field:"optional" json:"scalingConfig" yaml:"scalingConfig"`
	// self_managed_event_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#self_managed_event_source LambdaEventSourceMapping#self_managed_event_source}
	SelfManagedEventSource *LambdaEventSourceMappingSelfManagedEventSource `field:"optional" json:"selfManagedEventSource" yaml:"selfManagedEventSource"`
	// self_managed_kafka_event_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#self_managed_kafka_event_source_config LambdaEventSourceMapping#self_managed_kafka_event_source_config}
	SelfManagedKafkaEventSourceConfig *LambdaEventSourceMappingSelfManagedKafkaEventSourceConfig `field:"optional" json:"selfManagedKafkaEventSourceConfig" yaml:"selfManagedKafkaEventSourceConfig"`
	// source_access_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#source_access_configuration LambdaEventSourceMapping#source_access_configuration}
	SourceAccessConfiguration interface{} `field:"optional" json:"sourceAccessConfiguration" yaml:"sourceAccessConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#starting_position LambdaEventSourceMapping#starting_position}.
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#starting_position_timestamp LambdaEventSourceMapping#starting_position_timestamp}.
	StartingPositionTimestamp *string `field:"optional" json:"startingPositionTimestamp" yaml:"startingPositionTimestamp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#topics LambdaEventSourceMapping#topics}.
	Topics *[]*string `field:"optional" json:"topics" yaml:"topics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lambda_event_source_mapping#tumbling_window_in_seconds LambdaEventSourceMapping#tumbling_window_in_seconds}.
	TumblingWindowInSeconds *float64 `field:"optional" json:"tumblingWindowInSeconds" yaml:"tumblingWindowInSeconds"`
}

