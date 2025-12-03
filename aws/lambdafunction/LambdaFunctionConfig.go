// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdafunction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LambdaFunctionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#function_name LambdaFunction#function_name}.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#role LambdaFunction#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#architectures LambdaFunction#architectures}.
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// capacity_provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#capacity_provider_config LambdaFunction#capacity_provider_config}
	CapacityProviderConfig *LambdaFunctionCapacityProviderConfig `field:"optional" json:"capacityProviderConfig" yaml:"capacityProviderConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#code_signing_config_arn LambdaFunction#code_signing_config_arn}.
	CodeSigningConfigArn *string `field:"optional" json:"codeSigningConfigArn" yaml:"codeSigningConfigArn"`
	// dead_letter_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#dead_letter_config LambdaFunction#dead_letter_config}
	DeadLetterConfig *LambdaFunctionDeadLetterConfig `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#description LambdaFunction#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#environment LambdaFunction#environment}
	Environment *LambdaFunctionEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// ephemeral_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#ephemeral_storage LambdaFunction#ephemeral_storage}
	EphemeralStorage *LambdaFunctionEphemeralStorage `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#filename LambdaFunction#filename}.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#file_system_config LambdaFunction#file_system_config}
	FileSystemConfig *LambdaFunctionFileSystemConfig `field:"optional" json:"fileSystemConfig" yaml:"fileSystemConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#handler LambdaFunction#handler}.
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#id LambdaFunction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// image_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#image_config LambdaFunction#image_config}
	ImageConfig *LambdaFunctionImageConfig `field:"optional" json:"imageConfig" yaml:"imageConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#image_uri LambdaFunction#image_uri}.
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#kms_key_arn LambdaFunction#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#layers LambdaFunction#layers}.
	Layers *[]*string `field:"optional" json:"layers" yaml:"layers"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#logging_config LambdaFunction#logging_config}
	LoggingConfig *LambdaFunctionLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#memory_size LambdaFunction#memory_size}.
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#package_type LambdaFunction#package_type}.
	PackageType *string `field:"optional" json:"packageType" yaml:"packageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#publish LambdaFunction#publish}.
	Publish interface{} `field:"optional" json:"publish" yaml:"publish"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#publish_to LambdaFunction#publish_to}.
	PublishTo *string `field:"optional" json:"publishTo" yaml:"publishTo"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#region LambdaFunction#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#replacement_security_group_ids LambdaFunction#replacement_security_group_ids}.
	ReplacementSecurityGroupIds *[]*string `field:"optional" json:"replacementSecurityGroupIds" yaml:"replacementSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#replace_security_groups_on_destroy LambdaFunction#replace_security_groups_on_destroy}.
	ReplaceSecurityGroupsOnDestroy interface{} `field:"optional" json:"replaceSecurityGroupsOnDestroy" yaml:"replaceSecurityGroupsOnDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#reserved_concurrent_executions LambdaFunction#reserved_concurrent_executions}.
	ReservedConcurrentExecutions *float64 `field:"optional" json:"reservedConcurrentExecutions" yaml:"reservedConcurrentExecutions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#runtime LambdaFunction#runtime}.
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#s3_bucket LambdaFunction#s3_bucket}.
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#s3_key LambdaFunction#s3_key}.
	S3Key *string `field:"optional" json:"s3Key" yaml:"s3Key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#s3_object_version LambdaFunction#s3_object_version}.
	S3ObjectVersion *string `field:"optional" json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#skip_destroy LambdaFunction#skip_destroy}.
	SkipDestroy interface{} `field:"optional" json:"skipDestroy" yaml:"skipDestroy"`
	// snap_start block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#snap_start LambdaFunction#snap_start}
	SnapStart *LambdaFunctionSnapStart `field:"optional" json:"snapStart" yaml:"snapStart"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#source_code_hash LambdaFunction#source_code_hash}.
	SourceCodeHash *string `field:"optional" json:"sourceCodeHash" yaml:"sourceCodeHash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#source_kms_key_arn LambdaFunction#source_kms_key_arn}.
	SourceKmsKeyArn *string `field:"optional" json:"sourceKmsKeyArn" yaml:"sourceKmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#tags LambdaFunction#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#tags_all LambdaFunction#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// tenancy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#tenancy_config LambdaFunction#tenancy_config}
	TenancyConfig *LambdaFunctionTenancyConfig `field:"optional" json:"tenancyConfig" yaml:"tenancyConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#timeout LambdaFunction#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#timeouts LambdaFunction#timeouts}
	Timeouts *LambdaFunctionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// tracing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#tracing_config LambdaFunction#tracing_config}
	TracingConfig *LambdaFunctionTracingConfig `field:"optional" json:"tracingConfig" yaml:"tracingConfig"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#vpc_config LambdaFunction#vpc_config}
	VpcConfig *LambdaFunctionVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

