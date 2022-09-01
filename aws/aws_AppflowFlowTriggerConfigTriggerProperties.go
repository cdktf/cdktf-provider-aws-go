// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowFlowTriggerConfigTriggerProperties struct {
	// scheduled block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#scheduled AppflowFlow#scheduled}
	Scheduled *AppflowFlowTriggerConfigTriggerPropertiesScheduled `field:"optional" json:"scheduled" yaml:"scheduled"`
}

