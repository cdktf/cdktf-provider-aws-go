// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package macie2classificationjob


type Macie2ClassificationJobS3JobDefinition struct {
	// bucket_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/macie2_classification_job#bucket_criteria Macie2ClassificationJob#bucket_criteria}
	BucketCriteria *Macie2ClassificationJobS3JobDefinitionBucketCriteria `field:"optional" json:"bucketCriteria" yaml:"bucketCriteria"`
	// bucket_definitions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/macie2_classification_job#bucket_definitions Macie2ClassificationJob#bucket_definitions}
	BucketDefinitions interface{} `field:"optional" json:"bucketDefinitions" yaml:"bucketDefinitions"`
	// scoping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/macie2_classification_job#scoping Macie2ClassificationJob#scoping}
	Scoping *Macie2ClassificationJobS3JobDefinitionScoping `field:"optional" json:"scoping" yaml:"scoping"`
}

