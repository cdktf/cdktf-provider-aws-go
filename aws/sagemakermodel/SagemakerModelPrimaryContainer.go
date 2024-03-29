// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel


type SagemakerModelPrimaryContainer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/sagemaker_model#container_hostname SagemakerModel#container_hostname}.
	ContainerHostname *string `field:"optional" json:"containerHostname" yaml:"containerHostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/sagemaker_model#environment SagemakerModel#environment}.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/sagemaker_model#image SagemakerModel#image}.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// image_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/sagemaker_model#image_config SagemakerModel#image_config}
	ImageConfig *SagemakerModelPrimaryContainerImageConfig `field:"optional" json:"imageConfig" yaml:"imageConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/sagemaker_model#mode SagemakerModel#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// model_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/sagemaker_model#model_data_source SagemakerModel#model_data_source}
	ModelDataSource *SagemakerModelPrimaryContainerModelDataSource `field:"optional" json:"modelDataSource" yaml:"modelDataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/sagemaker_model#model_data_url SagemakerModel#model_data_url}.
	ModelDataUrl *string `field:"optional" json:"modelDataUrl" yaml:"modelDataUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/sagemaker_model#model_package_name SagemakerModel#model_package_name}.
	ModelPackageName *string `field:"optional" json:"modelPackageName" yaml:"modelPackageName"`
}

