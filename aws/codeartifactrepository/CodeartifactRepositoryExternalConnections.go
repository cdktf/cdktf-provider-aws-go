package codeartifactrepository


type CodeartifactRepositoryExternalConnections struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codeartifact_repository#external_connection_name CodeartifactRepository#external_connection_name}.
	ExternalConnectionName *string `field:"required" json:"externalConnectionName" yaml:"externalConnectionName"`
}

