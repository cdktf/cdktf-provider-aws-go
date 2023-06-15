package quicksightdataset


type QuicksightDataSetLogicalTableMapSourceJoinInstruction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/quicksight_data_set#left_operand QuicksightDataSet#left_operand}.
	LeftOperand *string `field:"required" json:"leftOperand" yaml:"leftOperand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/quicksight_data_set#on_clause QuicksightDataSet#on_clause}.
	OnClause *string `field:"required" json:"onClause" yaml:"onClause"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/quicksight_data_set#right_operand QuicksightDataSet#right_operand}.
	RightOperand *string `field:"required" json:"rightOperand" yaml:"rightOperand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/quicksight_data_set#type QuicksightDataSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// left_join_key_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/quicksight_data_set#left_join_key_properties QuicksightDataSet#left_join_key_properties}
	LeftJoinKeyProperties *QuicksightDataSetLogicalTableMapSourceJoinInstructionLeftJoinKeyProperties `field:"optional" json:"leftJoinKeyProperties" yaml:"leftJoinKeyProperties"`
	// right_join_key_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/quicksight_data_set#right_join_key_properties QuicksightDataSet#right_join_key_properties}
	RightJoinKeyProperties *QuicksightDataSetLogicalTableMapSourceJoinInstructionRightJoinKeyProperties `field:"optional" json:"rightJoinKeyProperties" yaml:"rightJoinKeyProperties"`
}

