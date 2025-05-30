// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastictranscoderpipeline


type ElastictranscoderPipelineThumbnailConfigPermissions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/elastictranscoder_pipeline#access ElastictranscoderPipeline#access}.
	Access *[]*string `field:"optional" json:"access" yaml:"access"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/elastictranscoder_pipeline#grantee ElastictranscoderPipeline#grantee}.
	Grantee *string `field:"optional" json:"grantee" yaml:"grantee"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/elastictranscoder_pipeline#grantee_type ElastictranscoderPipeline#grantee_type}.
	GranteeType *string `field:"optional" json:"granteeType" yaml:"granteeType"`
}

