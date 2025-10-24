// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain


type OpensearchDomainAimlOptions struct {
	// natural_language_query_generation_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/opensearch_domain#natural_language_query_generation_options OpensearchDomain#natural_language_query_generation_options}
	NaturalLanguageQueryGenerationOptions *OpensearchDomainAimlOptionsNaturalLanguageQueryGenerationOptions `field:"optional" json:"naturalLanguageQueryGenerationOptions" yaml:"naturalLanguageQueryGenerationOptions"`
	// s3_vectors_engine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/opensearch_domain#s3_vectors_engine OpensearchDomain#s3_vectors_engine}
	S3VectorsEngine *OpensearchDomainAimlOptionsS3VectorsEngine `field:"optional" json:"s3VectorsEngine" yaml:"s3VectorsEngine"`
}

