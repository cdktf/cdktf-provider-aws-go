package ecr


type EcrpublicRepositoryTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecrpublic_repository#delete EcrpublicRepository#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

