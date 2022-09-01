package macie2


type Macie2ClassificationJobS3JobDefinitionBucketCriteria struct {
	// excludes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job#excludes Macie2ClassificationJob#excludes}
	Excludes *Macie2ClassificationJobS3JobDefinitionBucketCriteriaExcludes `field:"optional" json:"excludes" yaml:"excludes"`
	// includes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job#includes Macie2ClassificationJob#includes}
	Includes *Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludes `field:"optional" json:"includes" yaml:"includes"`
}

