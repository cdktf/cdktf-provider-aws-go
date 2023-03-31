package rbinrule


type RbinRuleLockConfigurationUnlockDelay struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/rbin_rule#unlock_delay_unit RbinRule#unlock_delay_unit}.
	UnlockDelayUnit *string `field:"required" json:"unlockDelayUnit" yaml:"unlockDelayUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/rbin_rule#unlock_delay_value RbinRule#unlock_delay_value}.
	UnlockDelayValue *float64 `field:"required" json:"unlockDelayValue" yaml:"unlockDelayValue"`
}

