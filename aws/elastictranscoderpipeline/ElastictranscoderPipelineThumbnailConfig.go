// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastictranscoderpipeline


type ElastictranscoderPipelineThumbnailConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/elastictranscoder_pipeline#bucket ElastictranscoderPipeline#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/elastictranscoder_pipeline#storage_class ElastictranscoderPipeline#storage_class}.
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

