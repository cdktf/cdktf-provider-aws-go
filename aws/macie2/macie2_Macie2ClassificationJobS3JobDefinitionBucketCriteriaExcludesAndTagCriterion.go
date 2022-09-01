package macie2


type Macie2ClassificationJobS3JobDefinitionBucketCriteriaExcludesAndTagCriterion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job#comparator Macie2ClassificationJob#comparator}.
	Comparator *string `field:"optional" json:"comparator" yaml:"comparator"`
	// tag_values block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job#tag_values Macie2ClassificationJob#tag_values}
	TagValues interface{} `field:"optional" json:"tagValues" yaml:"tagValues"`
}

