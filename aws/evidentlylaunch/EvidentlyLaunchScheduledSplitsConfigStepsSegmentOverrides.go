package evidentlylaunch


type EvidentlyLaunchScheduledSplitsConfigStepsSegmentOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/evidently_launch#evaluation_order EvidentlyLaunch#evaluation_order}.
	EvaluationOrder *float64 `field:"required" json:"evaluationOrder" yaml:"evaluationOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/evidently_launch#segment EvidentlyLaunch#segment}.
	Segment *string `field:"required" json:"segment" yaml:"segment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/evidently_launch#weights EvidentlyLaunch#weights}.
	Weights *map[string]*float64 `field:"required" json:"weights" yaml:"weights"`
}

