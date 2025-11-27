// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagepipeline

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImagebuilderImagePipelineConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/imagebuilder_image_pipeline#infrastructure_configuration_arn ImagebuilderImagePipeline#infrastructure_configuration_arn}.
	InfrastructureConfigurationArn *string `field:"required" json:"infrastructureConfigurationArn" yaml:"infrastructureConfigurationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/imagebuilder_image_pipeline#name ImagebuilderImagePipeline#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/imagebuilder_image_pipeline#container_recipe_arn ImagebuilderImagePipeline#container_recipe_arn}.
	ContainerRecipeArn *string `field:"optional" json:"containerRecipeArn" yaml:"containerRecipeArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/imagebuilder_image_pipeline#description ImagebuilderImagePipeline#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/imagebuilder_image_pipeline#distribution_configuration_arn ImagebuilderImagePipeline#distribution_configuration_arn}.
	DistributionConfigurationArn *string `field:"optional" json:"distributionConfigurationArn" yaml:"distributionConfigurationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/imagebuilder_image_pipeline#enhanced_image_metadata_enabled ImagebuilderImagePipeline#enhanced_image_metadata_enabled}.
	EnhancedImageMetadataEnabled interface{} `field:"optional" json:"enhancedImageMetadataEnabled" yaml:"enhancedImageMetadataEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/imagebuilder_image_pipeline#execution_role ImagebuilderImagePipeline#execution_role}.
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/imagebuilder_image_pipeline#id ImagebuilderImagePipeline#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/imagebuilder_image_pipeline#image_recipe_arn ImagebuilderImagePipeline#image_recipe_arn}.
	ImageRecipeArn *string `field:"optional" json:"imageRecipeArn" yaml:"imageRecipeArn"`
	// image_scanning_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/imagebuilder_image_pipeline#image_scanning_configuration ImagebuilderImagePipeline#image_scanning_configuration}
	ImageScanningConfiguration *ImagebuilderImagePipelineImageScanningConfiguration `field:"optional" json:"imageScanningConfiguration" yaml:"imageScanningConfiguration"`
	// image_tests_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/imagebuilder_image_pipeline#image_tests_configuration ImagebuilderImagePipeline#image_tests_configuration}
	ImageTestsConfiguration *ImagebuilderImagePipelineImageTestsConfiguration `field:"optional" json:"imageTestsConfiguration" yaml:"imageTestsConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/imagebuilder_image_pipeline#region ImagebuilderImagePipeline#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/imagebuilder_image_pipeline#schedule ImagebuilderImagePipeline#schedule}
	Schedule *ImagebuilderImagePipelineSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/imagebuilder_image_pipeline#status ImagebuilderImagePipeline#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/imagebuilder_image_pipeline#tags ImagebuilderImagePipeline#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/imagebuilder_image_pipeline#tags_all ImagebuilderImagePipeline#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// workflow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/imagebuilder_image_pipeline#workflow ImagebuilderImagePipeline#workflow}
	Workflow interface{} `field:"optional" json:"workflow" yaml:"workflow"`
}

