package vpclatticelistenerrule


type VpclatticeListenerRuleActionForwardTargetGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/vpclattice_listener_rule#target_group_identifier VpclatticeListenerRule#target_group_identifier}.
	TargetGroupIdentifier *string `field:"required" json:"targetGroupIdentifier" yaml:"targetGroupIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/vpclattice_listener_rule#weight VpclatticeListenerRule#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

