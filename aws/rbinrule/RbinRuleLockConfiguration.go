package rbinrule


type RbinRuleLockConfiguration struct {
	// unlock_delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/rbin_rule#unlock_delay RbinRule#unlock_delay}
	UnlockDelay *RbinRuleLockConfigurationUnlockDelay `field:"required" json:"unlockDelay" yaml:"unlockDelay"`
}

