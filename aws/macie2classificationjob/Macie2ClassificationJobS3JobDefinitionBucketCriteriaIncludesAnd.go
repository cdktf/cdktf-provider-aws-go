// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package macie2classificationjob


type Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAnd struct {
	// simple_criterion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/macie2_classification_job#simple_criterion Macie2ClassificationJob#simple_criterion}
	SimpleCriterion *Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndSimpleCriterion `field:"optional" json:"simpleCriterion" yaml:"simpleCriterion"`
	// tag_criterion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/macie2_classification_job#tag_criterion Macie2ClassificationJob#tag_criterion}
	TagCriterion *Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndTagCriterion `field:"optional" json:"tagCriterion" yaml:"tagCriterion"`
}

