package appflowflow


type AppflowFlowTriggerConfigTriggerProperties struct {
	// scheduled block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/appflow_flow#scheduled AppflowFlow#scheduled}
	Scheduled *AppflowFlowTriggerConfigTriggerPropertiesScheduled `field:"optional" json:"scheduled" yaml:"scheduled"`
}

