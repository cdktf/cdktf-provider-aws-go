package sfnstatemachine


type SfnStateMachineLoggingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/sfn_state_machine#include_execution_data SfnStateMachine#include_execution_data}.
	IncludeExecutionData interface{} `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/sfn_state_machine#level SfnStateMachine#level}.
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/sfn_state_machine#log_destination SfnStateMachine#log_destination}.
	LogDestination *string `field:"optional" json:"logDestination" yaml:"logDestination"`
}

