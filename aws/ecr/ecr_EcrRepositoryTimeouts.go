package ecr


type EcrRepositoryTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecr_repository#delete EcrRepository#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

