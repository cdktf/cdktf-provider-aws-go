package appflowflow


type AppflowFlowTriggerConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#trigger_type AppflowFlow#trigger_type}.
	TriggerType *string `field:"required" json:"triggerType" yaml:"triggerType"`
	// trigger_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#trigger_properties AppflowFlow#trigger_properties}
	TriggerProperties *AppflowFlowTriggerConfigTriggerProperties `field:"optional" json:"triggerProperties" yaml:"triggerProperties"`
}

