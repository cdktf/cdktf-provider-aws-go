package athenadatabase


type AthenaDatabaseAclConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/athena_database#s3_acl_option AthenaDatabase#s3_acl_option}.
	S3AclOption *string `field:"required" json:"s3AclOption" yaml:"s3AclOption"`
}

