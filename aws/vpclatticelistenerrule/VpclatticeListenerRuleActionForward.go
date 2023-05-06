package vpclatticelistenerrule


type VpclatticeListenerRuleActionForward struct {
	// target_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/vpclattice_listener_rule#target_groups VpclatticeListenerRule#target_groups}
	TargetGroups interface{} `field:"required" json:"targetGroups" yaml:"targetGroups"`
}

