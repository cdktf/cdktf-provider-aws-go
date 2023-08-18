package cloudwatcheventtarget


type CloudwatchEventTargetEcsTargetOrderedPlacementStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/cloudwatch_event_target#type CloudwatchEventTarget#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/cloudwatch_event_target#field CloudwatchEventTarget#field}.
	Field *string `field:"optional" json:"field" yaml:"field"`
}

