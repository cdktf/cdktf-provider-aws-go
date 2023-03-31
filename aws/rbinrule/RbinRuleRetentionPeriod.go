package rbinrule


type RbinRuleRetentionPeriod struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/rbin_rule#retention_period_unit RbinRule#retention_period_unit}.
	RetentionPeriodUnit *string `field:"required" json:"retentionPeriodUnit" yaml:"retentionPeriodUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/rbin_rule#retention_period_value RbinRule#retention_period_value}.
	RetentionPeriodValue *float64 `field:"required" json:"retentionPeriodValue" yaml:"retentionPeriodValue"`
}

