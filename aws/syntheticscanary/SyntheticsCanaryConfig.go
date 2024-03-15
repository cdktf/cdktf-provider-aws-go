// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticscanary

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsCanaryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#artifact_s3_location SyntheticsCanary#artifact_s3_location}.
	ArtifactS3Location *string `field:"required" json:"artifactS3Location" yaml:"artifactS3Location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#execution_role_arn SyntheticsCanary#execution_role_arn}.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#handler SyntheticsCanary#handler}.
	Handler *string `field:"required" json:"handler" yaml:"handler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#name SyntheticsCanary#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#runtime_version SyntheticsCanary#runtime_version}.
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#schedule SyntheticsCanary#schedule}
	Schedule *SyntheticsCanarySchedule `field:"required" json:"schedule" yaml:"schedule"`
	// artifact_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#artifact_config SyntheticsCanary#artifact_config}
	ArtifactConfig *SyntheticsCanaryArtifactConfig `field:"optional" json:"artifactConfig" yaml:"artifactConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#delete_lambda SyntheticsCanary#delete_lambda}.
	DeleteLambda interface{} `field:"optional" json:"deleteLambda" yaml:"deleteLambda"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#failure_retention_period SyntheticsCanary#failure_retention_period}.
	FailureRetentionPeriod *float64 `field:"optional" json:"failureRetentionPeriod" yaml:"failureRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#id SyntheticsCanary#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// run_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#run_config SyntheticsCanary#run_config}
	RunConfig *SyntheticsCanaryRunConfig `field:"optional" json:"runConfig" yaml:"runConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#s3_bucket SyntheticsCanary#s3_bucket}.
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#s3_key SyntheticsCanary#s3_key}.
	S3Key *string `field:"optional" json:"s3Key" yaml:"s3Key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#s3_version SyntheticsCanary#s3_version}.
	S3Version *string `field:"optional" json:"s3Version" yaml:"s3Version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#start_canary SyntheticsCanary#start_canary}.
	StartCanary interface{} `field:"optional" json:"startCanary" yaml:"startCanary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#success_retention_period SyntheticsCanary#success_retention_period}.
	SuccessRetentionPeriod *float64 `field:"optional" json:"successRetentionPeriod" yaml:"successRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#tags SyntheticsCanary#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#tags_all SyntheticsCanary#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#vpc_config SyntheticsCanary#vpc_config}
	VpcConfig *SyntheticsCanaryVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/synthetics_canary#zip_file SyntheticsCanary#zip_file}.
	ZipFile *string `field:"optional" json:"zipFile" yaml:"zipFile"`
}

