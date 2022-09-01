package cloudformation


type CloudformationStackSetAutoDeployment struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudformation_stack_set#enabled CloudformationStackSet#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudformation_stack_set#retain_stacks_on_account_removal CloudformationStackSet#retain_stacks_on_account_removal}.
	RetainStacksOnAccountRemoval interface{} `field:"optional" json:"retainStacksOnAccountRemoval" yaml:"retainStacksOnAccountRemoval"`
}

