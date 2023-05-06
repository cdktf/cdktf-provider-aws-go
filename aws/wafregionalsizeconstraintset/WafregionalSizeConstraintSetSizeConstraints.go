package wafregionalsizeconstraintset


type WafregionalSizeConstraintSetSizeConstraints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/wafregional_size_constraint_set#comparison_operator WafregionalSizeConstraintSet#comparison_operator}.
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/wafregional_size_constraint_set#field_to_match WafregionalSizeConstraintSet#field_to_match}
	FieldToMatch *WafregionalSizeConstraintSetSizeConstraintsFieldToMatch `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/wafregional_size_constraint_set#size WafregionalSizeConstraintSet#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/wafregional_size_constraint_set#text_transformation WafregionalSizeConstraintSet#text_transformation}.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
}

