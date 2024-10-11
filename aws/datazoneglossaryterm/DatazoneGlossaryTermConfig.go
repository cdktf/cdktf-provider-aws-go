// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datazoneglossaryterm

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatazoneGlossaryTermConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/datazone_glossary_term#glossary_identifier DatazoneGlossaryTerm#glossary_identifier}.
	GlossaryIdentifier *string `field:"required" json:"glossaryIdentifier" yaml:"glossaryIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/datazone_glossary_term#name DatazoneGlossaryTerm#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/datazone_glossary_term#domain_identifier DatazoneGlossaryTerm#domain_identifier}.
	DomainIdentifier *string `field:"optional" json:"domainIdentifier" yaml:"domainIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/datazone_glossary_term#long_description DatazoneGlossaryTerm#long_description}.
	LongDescription *string `field:"optional" json:"longDescription" yaml:"longDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/datazone_glossary_term#short_description DatazoneGlossaryTerm#short_description}.
	ShortDescription *string `field:"optional" json:"shortDescription" yaml:"shortDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/datazone_glossary_term#status DatazoneGlossaryTerm#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// term_relations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/datazone_glossary_term#term_relations DatazoneGlossaryTerm#term_relations}
	TermRelations interface{} `field:"optional" json:"termRelations" yaml:"termRelations"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/datazone_glossary_term#timeouts DatazoneGlossaryTerm#timeouts}
	Timeouts *DatazoneGlossaryTermTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

