package cloudwatcheventtarget


type CloudwatchEventTargetRunCommandTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/cloudwatch_event_target#key CloudwatchEventTarget#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/cloudwatch_event_target#values CloudwatchEventTarget#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

