package wafv2webaclloggingconfiguration


type Wafv2WebAclLoggingConfigurationLoggingFilterFilterCondition struct {
	// action_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/wafv2_web_acl_logging_configuration#action_condition Wafv2WebAclLoggingConfiguration#action_condition}
	ActionCondition *Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionActionCondition `field:"optional" json:"actionCondition" yaml:"actionCondition"`
	// label_name_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/wafv2_web_acl_logging_configuration#label_name_condition Wafv2WebAclLoggingConfiguration#label_name_condition}
	LabelNameCondition *Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameCondition `field:"optional" json:"labelNameCondition" yaml:"labelNameCondition"`
}

