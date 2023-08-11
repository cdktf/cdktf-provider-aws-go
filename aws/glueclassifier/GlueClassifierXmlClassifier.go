package glueclassifier


type GlueClassifierXmlClassifier struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/glue_classifier#classification GlueClassifier#classification}.
	Classification *string `field:"required" json:"classification" yaml:"classification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/glue_classifier#row_tag GlueClassifier#row_tag}.
	RowTag *string `field:"required" json:"rowTag" yaml:"rowTag"`
}

