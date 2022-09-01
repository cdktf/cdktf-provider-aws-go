package batch


type BatchSchedulingPolicyFairSharePolicyShareDistribution struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_scheduling_policy#share_identifier BatchSchedulingPolicy#share_identifier}.
	ShareIdentifier *string `field:"required" json:"shareIdentifier" yaml:"shareIdentifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_scheduling_policy#weight_factor BatchSchedulingPolicy#weight_factor}.
	WeightFactor *float64 `field:"optional" json:"weightFactor" yaml:"weightFactor"`
}

