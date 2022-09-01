package athena


type AthenaWorkgroupConfigurationResultConfigurationAclConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/athena_workgroup#s3_acl_option AthenaWorkgroup#s3_acl_option}.
	S3AclOption *string `field:"required" json:"s3AclOption" yaml:"s3AclOption"`
}

