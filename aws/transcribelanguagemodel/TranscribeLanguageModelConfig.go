// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcribelanguagemodel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TranscribeLanguageModelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/transcribe_language_model#base_model_name TranscribeLanguageModel#base_model_name}.
	BaseModelName *string `field:"required" json:"baseModelName" yaml:"baseModelName"`
	// input_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/transcribe_language_model#input_data_config TranscribeLanguageModel#input_data_config}
	InputDataConfig *TranscribeLanguageModelInputDataConfig `field:"required" json:"inputDataConfig" yaml:"inputDataConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/transcribe_language_model#language_code TranscribeLanguageModel#language_code}.
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/transcribe_language_model#model_name TranscribeLanguageModel#model_name}.
	ModelName *string `field:"required" json:"modelName" yaml:"modelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/transcribe_language_model#id TranscribeLanguageModel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/transcribe_language_model#tags TranscribeLanguageModel#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/transcribe_language_model#tags_all TranscribeLanguageModel#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/transcribe_language_model#timeouts TranscribeLanguageModel#timeouts}
	Timeouts *TranscribeLanguageModelTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

