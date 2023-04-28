package macies3bucketassociation


type MacieS3BucketAssociationClassificationType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/macie_s3_bucket_association#continuous MacieS3BucketAssociation#continuous}.
	Continuous *string `field:"optional" json:"continuous" yaml:"continuous"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/macie_s3_bucket_association#one_time MacieS3BucketAssociation#one_time}.
	OneTime *string `field:"optional" json:"oneTime" yaml:"oneTime"`
}

