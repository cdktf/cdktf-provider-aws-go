package rbinrule


type RbinRuleLockConfiguration struct {
	// unlock_delay block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/rbin_rule#unlock_delay RbinRule#unlock_delay}
	UnlockDelay *RbinRuleLockConfigurationUnlockDelay `field:"required" json:"unlockDelay" yaml:"unlockDelay"`
}

