package quicksightdataset


type QuicksightDataSetLogicalTableMapSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_set#data_set_arn QuicksightDataSet#data_set_arn}.
	DataSetArn *string `field:"optional" json:"dataSetArn" yaml:"dataSetArn"`
	// join_instruction block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_set#join_instruction QuicksightDataSet#join_instruction}
	JoinInstruction *QuicksightDataSetLogicalTableMapSourceJoinInstruction `field:"optional" json:"joinInstruction" yaml:"joinInstruction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_set#physical_table_id QuicksightDataSet#physical_table_id}.
	PhysicalTableId *string `field:"optional" json:"physicalTableId" yaml:"physicalTableId"`
}

