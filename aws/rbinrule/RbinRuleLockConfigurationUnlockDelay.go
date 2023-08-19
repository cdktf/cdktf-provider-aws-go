package rbinrule


type RbinRuleLockConfigurationUnlockDelay struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/rbin_rule#unlock_delay_unit RbinRule#unlock_delay_unit}.
	UnlockDelayUnit *string `field:"required" json:"unlockDelayUnit" yaml:"unlockDelayUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/rbin_rule#unlock_delay_value RbinRule#unlock_delay_value}.
	UnlockDelayValue *float64 `field:"required" json:"unlockDelayValue" yaml:"unlockDelayValue"`
}

