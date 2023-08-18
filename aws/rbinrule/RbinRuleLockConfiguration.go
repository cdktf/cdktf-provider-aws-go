package rbinrule


type RbinRuleLockConfiguration struct {
	// unlock_delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/rbin_rule#unlock_delay RbinRule#unlock_delay}
	UnlockDelay *RbinRuleLockConfigurationUnlockDelay `field:"required" json:"unlockDelay" yaml:"unlockDelay"`
}

