package sagemakermodel


type SagemakerModelPrimaryContainerImageConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_model#repository_access_mode SagemakerModel#repository_access_mode}.
	RepositoryAccessMode *string `field:"required" json:"repositoryAccessMode" yaml:"repositoryAccessMode"`
	// repository_auth_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_model#repository_auth_config SagemakerModel#repository_auth_config}
	RepositoryAuthConfig *SagemakerModelPrimaryContainerImageConfigRepositoryAuthConfig `field:"optional" json:"repositoryAuthConfig" yaml:"repositoryAuthConfig"`
}

