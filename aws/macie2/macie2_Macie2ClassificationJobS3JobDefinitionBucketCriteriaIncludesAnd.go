package macie2


type Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAnd struct {
	// simple_criterion block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job#simple_criterion Macie2ClassificationJob#simple_criterion}
	SimpleCriterion *Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndSimpleCriterion `field:"optional" json:"simpleCriterion" yaml:"simpleCriterion"`
	// tag_criterion block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job#tag_criterion Macie2ClassificationJob#tag_criterion}
	TagCriterion *Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndTagCriterion `field:"optional" json:"tagCriterion" yaml:"tagCriterion"`
}

