package apprunnerservice


type ApprunnerServiceSourceConfigurationCodeRepository struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/apprunner_service#repository_url ApprunnerService#repository_url}.
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
	// source_code_version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/apprunner_service#source_code_version ApprunnerService#source_code_version}
	SourceCodeVersion *ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersion `field:"required" json:"sourceCodeVersion" yaml:"sourceCodeVersion"`
	// code_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/apprunner_service#code_configuration ApprunnerService#code_configuration}
	CodeConfiguration *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfiguration `field:"optional" json:"codeConfiguration" yaml:"codeConfiguration"`
}

