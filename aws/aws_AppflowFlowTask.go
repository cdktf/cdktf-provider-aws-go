// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowFlowTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#source_fields AppflowFlow#source_fields}.
	SourceFields *[]*string `field:"required" json:"sourceFields" yaml:"sourceFields"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#task_type AppflowFlow#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// connector_operator block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#connector_operator AppflowFlow#connector_operator}
	ConnectorOperator interface{} `field:"optional" json:"connectorOperator" yaml:"connectorOperator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#destination_field AppflowFlow#destination_field}.
	DestinationField *string `field:"optional" json:"destinationField" yaml:"destinationField"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#task_properties AppflowFlow#task_properties}.
	TaskProperties *map[string]*string `field:"optional" json:"taskProperties" yaml:"taskProperties"`
}

