package cloudwatcheventtarget


type CloudwatchEventTargetInputTransformer struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudwatch_event_target#input_template CloudwatchEventTarget#input_template}.
	InputTemplate *string `field:"required" json:"inputTemplate" yaml:"inputTemplate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudwatch_event_target#input_paths CloudwatchEventTarget#input_paths}.
	InputPaths *map[string]*string `field:"optional" json:"inputPaths" yaml:"inputPaths"`
}

