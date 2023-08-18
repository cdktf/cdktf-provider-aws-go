package codeartifactrepository


type CodeartifactRepositoryUpstream struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/codeartifact_repository#repository_name CodeartifactRepository#repository_name}.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
}

