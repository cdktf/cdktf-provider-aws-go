package evidentlylaunch


type EvidentlyLaunchScheduledSplitsConfig struct {
	// steps block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/evidently_launch#steps EvidentlyLaunch#steps}
	Steps interface{} `field:"required" json:"steps" yaml:"steps"`
}

