package guarddutyfilter


type GuarddutyFilterFindingCriteria struct {
	// criterion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/guardduty_filter#criterion GuarddutyFilter#criterion}
	Criterion interface{} `field:"required" json:"criterion" yaml:"criterion"`
}

