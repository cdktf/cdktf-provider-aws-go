package macie2


type Macie2ClassificationJobS3JobDefinitionScoping struct {
	// excludes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job#excludes Macie2ClassificationJob#excludes}
	Excludes *Macie2ClassificationJobS3JobDefinitionScopingExcludes `field:"optional" json:"excludes" yaml:"excludes"`
	// includes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie2_classification_job#includes Macie2ClassificationJob#includes}
	Includes *Macie2ClassificationJobS3JobDefinitionScopingIncludes `field:"optional" json:"includes" yaml:"includes"`
}

