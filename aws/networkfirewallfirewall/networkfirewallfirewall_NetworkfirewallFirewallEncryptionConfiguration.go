package networkfirewallfirewall


type NetworkfirewallFirewallEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_firewall#type NetworkfirewallFirewall#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_firewall#key_id NetworkfirewallFirewall#key_id}.
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

