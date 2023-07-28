package evidentlylaunch


type EvidentlyLaunchScheduledSplitsConfig struct {
	// steps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/evidently_launch#steps EvidentlyLaunch#steps}
	Steps interface{} `field:"required" json:"steps" yaml:"steps"`
}

