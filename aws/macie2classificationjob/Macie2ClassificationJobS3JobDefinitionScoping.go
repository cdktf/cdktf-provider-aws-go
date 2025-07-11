// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package macie2classificationjob


type Macie2ClassificationJobS3JobDefinitionScoping struct {
	// excludes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/macie2_classification_job#excludes Macie2ClassificationJob#excludes}
	Excludes *Macie2ClassificationJobS3JobDefinitionScopingExcludes `field:"optional" json:"excludes" yaml:"excludes"`
	// includes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/macie2_classification_job#includes Macie2ClassificationJob#includes}
	Includes *Macie2ClassificationJobS3JobDefinitionScopingIncludes `field:"optional" json:"includes" yaml:"includes"`
}

