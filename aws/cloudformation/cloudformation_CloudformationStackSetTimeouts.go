package cloudformation


type CloudformationStackSetTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudformation_stack_set#update CloudformationStackSet#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

