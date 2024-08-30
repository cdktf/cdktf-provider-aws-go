// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint


type DmsEndpointMongodbSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/dms_endpoint#auth_mechanism DmsEndpoint#auth_mechanism}.
	AuthMechanism *string `field:"optional" json:"authMechanism" yaml:"authMechanism"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/dms_endpoint#auth_source DmsEndpoint#auth_source}.
	AuthSource *string `field:"optional" json:"authSource" yaml:"authSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/dms_endpoint#auth_type DmsEndpoint#auth_type}.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/dms_endpoint#docs_to_investigate DmsEndpoint#docs_to_investigate}.
	DocsToInvestigate *string `field:"optional" json:"docsToInvestigate" yaml:"docsToInvestigate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/dms_endpoint#extract_doc_id DmsEndpoint#extract_doc_id}.
	ExtractDocId *string `field:"optional" json:"extractDocId" yaml:"extractDocId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/dms_endpoint#nesting_level DmsEndpoint#nesting_level}.
	NestingLevel *string `field:"optional" json:"nestingLevel" yaml:"nestingLevel"`
}

