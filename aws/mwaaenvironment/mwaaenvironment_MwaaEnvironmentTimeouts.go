package mwaaenvironment


type MwaaEnvironmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#create MwaaEnvironment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#delete MwaaEnvironment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mwaa_environment#update MwaaEnvironment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

