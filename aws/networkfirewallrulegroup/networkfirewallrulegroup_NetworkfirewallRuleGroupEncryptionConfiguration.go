package networkfirewallrulegroup


type NetworkfirewallRuleGroupEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_rule_group#type NetworkfirewallRuleGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_rule_group#key_id NetworkfirewallRuleGroup#key_id}.
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

