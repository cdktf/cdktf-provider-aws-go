package sagemakerdomain


type SagemakerDomainRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_domain#home_efs_file_system SagemakerDomain#home_efs_file_system}.
	HomeEfsFileSystem *string `field:"optional" json:"homeEfsFileSystem" yaml:"homeEfsFileSystem"`
}

