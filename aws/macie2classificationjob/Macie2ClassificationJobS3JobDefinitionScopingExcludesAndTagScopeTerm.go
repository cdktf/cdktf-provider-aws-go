// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package macie2classificationjob


type Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTerm struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/macie2_classification_job#comparator Macie2ClassificationJob#comparator}.
	Comparator *string `field:"optional" json:"comparator" yaml:"comparator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/macie2_classification_job#key Macie2ClassificationJob#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// tag_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/macie2_classification_job#tag_values Macie2ClassificationJob#tag_values}
	TagValues interface{} `field:"optional" json:"tagValues" yaml:"tagValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/macie2_classification_job#target Macie2ClassificationJob#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
}

