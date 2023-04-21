package codebuildproject


type CodebuildProjectEnvironment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/codebuild_project#compute_type CodebuildProject#compute_type}.
	ComputeType *string `field:"required" json:"computeType" yaml:"computeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/codebuild_project#image CodebuildProject#image}.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/codebuild_project#type CodebuildProject#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/codebuild_project#certificate CodebuildProject#certificate}.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// environment_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/codebuild_project#environment_variable CodebuildProject#environment_variable}
	EnvironmentVariable interface{} `field:"optional" json:"environmentVariable" yaml:"environmentVariable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/codebuild_project#image_pull_credentials_type CodebuildProject#image_pull_credentials_type}.
	ImagePullCredentialsType *string `field:"optional" json:"imagePullCredentialsType" yaml:"imagePullCredentialsType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/codebuild_project#privileged_mode CodebuildProject#privileged_mode}.
	PrivilegedMode interface{} `field:"optional" json:"privilegedMode" yaml:"privilegedMode"`
	// registry_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/codebuild_project#registry_credential CodebuildProject#registry_credential}
	RegistryCredential *CodebuildProjectEnvironmentRegistryCredential `field:"optional" json:"registryCredential" yaml:"registryCredential"`
}

