package guarddutyfilter


type GuarddutyFilterFindingCriteria struct {
	// criterion block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/guardduty_filter#criterion GuarddutyFilter#criterion}
	Criterion interface{} `field:"required" json:"criterion" yaml:"criterion"`
}

