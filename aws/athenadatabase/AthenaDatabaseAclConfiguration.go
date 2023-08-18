package athenadatabase


type AthenaDatabaseAclConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/athena_database#s3_acl_option AthenaDatabase#s3_acl_option}.
	S3AclOption *string `field:"required" json:"s3AclOption" yaml:"s3AclOption"`
}

