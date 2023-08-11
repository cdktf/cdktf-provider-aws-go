package sfnalias


type SfnAliasRoutingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/sfn_alias#state_machine_version_arn SfnAlias#state_machine_version_arn}.
	StateMachineVersionArn *string `field:"required" json:"stateMachineVersionArn" yaml:"stateMachineVersionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/sfn_alias#weight SfnAlias#weight}.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

