// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package macie2classificationjob


type Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndTagCriterion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/macie2_classification_job#comparator Macie2ClassificationJob#comparator}.
	Comparator *string `field:"optional" json:"comparator" yaml:"comparator"`
	// tag_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/macie2_classification_job#tag_values Macie2ClassificationJob#tag_values}
	TagValues interface{} `field:"optional" json:"tagValues" yaml:"tagValues"`
}

