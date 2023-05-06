package macie2classificationjob


type Macie2ClassificationJobS3JobDefinitionBucketCriteriaExcludesAnd struct {
	// simple_criterion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/macie2_classification_job#simple_criterion Macie2ClassificationJob#simple_criterion}
	SimpleCriterion *Macie2ClassificationJobS3JobDefinitionBucketCriteriaExcludesAndSimpleCriterion `field:"optional" json:"simpleCriterion" yaml:"simpleCriterion"`
	// tag_criterion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/macie2_classification_job#tag_criterion Macie2ClassificationJob#tag_criterion}
	TagCriterion *Macie2ClassificationJobS3JobDefinitionBucketCriteriaExcludesAndTagCriterion `field:"optional" json:"tagCriterion" yaml:"tagCriterion"`
}

