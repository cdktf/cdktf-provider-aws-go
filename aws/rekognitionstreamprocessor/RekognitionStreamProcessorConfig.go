// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rekognitionstreamprocessor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RekognitionStreamProcessorConfig struct {
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
	// An identifier you assign to the stream processor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/rekognition_stream_processor#name RekognitionStreamProcessor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Number (ARN) of the IAM role that allows access to the stream processor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/rekognition_stream_processor#role_arn RekognitionStreamProcessor#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// data_sharing_preference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/rekognition_stream_processor#data_sharing_preference RekognitionStreamProcessor#data_sharing_preference}
	DataSharingPreference interface{} `field:"optional" json:"dataSharingPreference" yaml:"dataSharingPreference"`
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/rekognition_stream_processor#input RekognitionStreamProcessor#input}
	Input interface{} `field:"optional" json:"input" yaml:"input"`
	// The identifier for your AWS Key Management Service key (AWS KMS key).
	//
	// You can supply the Amazon Resource Name (ARN) of your KMS key, the ID of your KMS key, an alias for your KMS key, or an alias ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/rekognition_stream_processor#kms_key_id RekognitionStreamProcessor#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// notification_channel block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/rekognition_stream_processor#notification_channel RekognitionStreamProcessor#notification_channel}
	NotificationChannel interface{} `field:"optional" json:"notificationChannel" yaml:"notificationChannel"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/rekognition_stream_processor#output RekognitionStreamProcessor#output}
	Output interface{} `field:"optional" json:"output" yaml:"output"`
	// regions_of_interest block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/rekognition_stream_processor#regions_of_interest RekognitionStreamProcessor#regions_of_interest}
	RegionsOfInterest interface{} `field:"optional" json:"regionsOfInterest" yaml:"regionsOfInterest"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/rekognition_stream_processor#settings RekognitionStreamProcessor#settings}
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/rekognition_stream_processor#tags RekognitionStreamProcessor#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/rekognition_stream_processor#timeouts RekognitionStreamProcessor#timeouts}
	Timeouts *RekognitionStreamProcessorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

