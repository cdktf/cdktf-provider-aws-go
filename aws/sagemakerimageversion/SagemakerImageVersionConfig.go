// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerimageversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerImageVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/sagemaker_image_version#base_image SagemakerImageVersion#base_image}.
	BaseImage *string `field:"required" json:"baseImage" yaml:"baseImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/sagemaker_image_version#image_name SagemakerImageVersion#image_name}.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/sagemaker_image_version#aliases SagemakerImageVersion#aliases}.
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/sagemaker_image_version#horovod SagemakerImageVersion#horovod}.
	Horovod interface{} `field:"optional" json:"horovod" yaml:"horovod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/sagemaker_image_version#id SagemakerImageVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/sagemaker_image_version#job_type SagemakerImageVersion#job_type}.
	JobType *string `field:"optional" json:"jobType" yaml:"jobType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/sagemaker_image_version#ml_framework SagemakerImageVersion#ml_framework}.
	MlFramework *string `field:"optional" json:"mlFramework" yaml:"mlFramework"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/sagemaker_image_version#processor SagemakerImageVersion#processor}.
	Processor *string `field:"optional" json:"processor" yaml:"processor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/sagemaker_image_version#programming_lang SagemakerImageVersion#programming_lang}.
	ProgrammingLang *string `field:"optional" json:"programmingLang" yaml:"programmingLang"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/sagemaker_image_version#region SagemakerImageVersion#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/sagemaker_image_version#release_notes SagemakerImageVersion#release_notes}.
	ReleaseNotes *string `field:"optional" json:"releaseNotes" yaml:"releaseNotes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/sagemaker_image_version#vendor_guidance SagemakerImageVersion#vendor_guidance}.
	VendorGuidance *string `field:"optional" json:"vendorGuidance" yaml:"vendorGuidance"`
}

