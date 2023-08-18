package imagebuildercontainerrecipe

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImagebuilderContainerRecipeConfig struct {
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
	// component block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/imagebuilder_container_recipe#component ImagebuilderContainerRecipe#component}
	Component interface{} `field:"required" json:"component" yaml:"component"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/imagebuilder_container_recipe#container_type ImagebuilderContainerRecipe#container_type}.
	ContainerType *string `field:"required" json:"containerType" yaml:"containerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/imagebuilder_container_recipe#name ImagebuilderContainerRecipe#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/imagebuilder_container_recipe#parent_image ImagebuilderContainerRecipe#parent_image}.
	ParentImage *string `field:"required" json:"parentImage" yaml:"parentImage"`
	// target_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/imagebuilder_container_recipe#target_repository ImagebuilderContainerRecipe#target_repository}
	TargetRepository *ImagebuilderContainerRecipeTargetRepository `field:"required" json:"targetRepository" yaml:"targetRepository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/imagebuilder_container_recipe#version ImagebuilderContainerRecipe#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/imagebuilder_container_recipe#description ImagebuilderContainerRecipe#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/imagebuilder_container_recipe#dockerfile_template_data ImagebuilderContainerRecipe#dockerfile_template_data}.
	DockerfileTemplateData *string `field:"optional" json:"dockerfileTemplateData" yaml:"dockerfileTemplateData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/imagebuilder_container_recipe#dockerfile_template_uri ImagebuilderContainerRecipe#dockerfile_template_uri}.
	DockerfileTemplateUri *string `field:"optional" json:"dockerfileTemplateUri" yaml:"dockerfileTemplateUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/imagebuilder_container_recipe#id ImagebuilderContainerRecipe#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// instance_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/imagebuilder_container_recipe#instance_configuration ImagebuilderContainerRecipe#instance_configuration}
	InstanceConfiguration *ImagebuilderContainerRecipeInstanceConfiguration `field:"optional" json:"instanceConfiguration" yaml:"instanceConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/imagebuilder_container_recipe#kms_key_id ImagebuilderContainerRecipe#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/imagebuilder_container_recipe#platform_override ImagebuilderContainerRecipe#platform_override}.
	PlatformOverride *string `field:"optional" json:"platformOverride" yaml:"platformOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/imagebuilder_container_recipe#tags ImagebuilderContainerRecipe#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/imagebuilder_container_recipe#tags_all ImagebuilderContainerRecipe#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/imagebuilder_container_recipe#working_directory ImagebuilderContainerRecipe#working_directory}.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

