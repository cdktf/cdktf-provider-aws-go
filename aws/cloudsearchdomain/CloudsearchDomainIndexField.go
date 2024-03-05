// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudsearchdomain


type CloudsearchDomainIndexField struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudsearch_domain#name CloudsearchDomain#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudsearch_domain#type CloudsearchDomain#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudsearch_domain#analysis_scheme CloudsearchDomain#analysis_scheme}.
	AnalysisScheme *string `field:"optional" json:"analysisScheme" yaml:"analysisScheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudsearch_domain#default_value CloudsearchDomain#default_value}.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudsearch_domain#facet CloudsearchDomain#facet}.
	Facet interface{} `field:"optional" json:"facet" yaml:"facet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudsearch_domain#highlight CloudsearchDomain#highlight}.
	Highlight interface{} `field:"optional" json:"highlight" yaml:"highlight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudsearch_domain#return CloudsearchDomain#return}.
	Return interface{} `field:"optional" json:"return" yaml:"return"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudsearch_domain#search CloudsearchDomain#search}.
	Search interface{} `field:"optional" json:"search" yaml:"search"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudsearch_domain#sort CloudsearchDomain#sort}.
	Sort interface{} `field:"optional" json:"sort" yaml:"sort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudsearch_domain#source_fields CloudsearchDomain#source_fields}.
	SourceFields *string `field:"optional" json:"sourceFields" yaml:"sourceFields"`
}

