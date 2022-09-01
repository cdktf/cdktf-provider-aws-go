package macie


type MacieS3BucketAssociationClassificationType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie_s3_bucket_association#continuous MacieS3BucketAssociation#continuous}.
	Continuous *string `field:"optional" json:"continuous" yaml:"continuous"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie_s3_bucket_association#one_time MacieS3BucketAssociation#one_time}.
	OneTime *string `field:"optional" json:"oneTime" yaml:"oneTime"`
}

