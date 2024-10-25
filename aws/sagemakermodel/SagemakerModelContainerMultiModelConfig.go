// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel


type SagemakerModelContainerMultiModelConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/sagemaker_model#model_cache_setting SagemakerModel#model_cache_setting}.
	ModelCacheSetting *string `field:"optional" json:"modelCacheSetting" yaml:"modelCacheSetting"`
}

