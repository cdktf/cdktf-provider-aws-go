package codebuildproject


type CodebuildProjectEnvironmentRegistryCredential struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#credential CodebuildProject#credential}.
	Credential *string `field:"required" json:"credential" yaml:"credential"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#credential_provider CodebuildProject#credential_provider}.
	CredentialProvider *string `field:"required" json:"credentialProvider" yaml:"credentialProvider"`
}

