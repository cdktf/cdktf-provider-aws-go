package vpclatticelistener


type VpclatticeListenerDefaultActionForwardTargetGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/vpclattice_listener#target_group_identifier VpclatticeListener#target_group_identifier}.
	TargetGroupIdentifier *string `field:"optional" json:"targetGroupIdentifier" yaml:"targetGroupIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/vpclattice_listener#weight VpclatticeListener#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

