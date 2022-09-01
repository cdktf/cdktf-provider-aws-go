// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoData struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#object_path AppflowFlow#object_path}.
	ObjectPath *string `field:"required" json:"objectPath" yaml:"objectPath"`
	// error_handling_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#error_handling_config AppflowFlow#error_handling_config}
	ErrorHandlingConfig *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfig `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#id_field_names AppflowFlow#id_field_names}.
	IdFieldNames *[]*string `field:"optional" json:"idFieldNames" yaml:"idFieldNames"`
	// success_response_handling_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#success_response_handling_config AppflowFlow#success_response_handling_config}
	SuccessResponseHandlingConfig *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataSuccessResponseHandlingConfig `field:"optional" json:"successResponseHandlingConfig" yaml:"successResponseHandlingConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#write_operation_type AppflowFlow#write_operation_type}.
	WriteOperationType *string `field:"optional" json:"writeOperationType" yaml:"writeOperationType"`
}

