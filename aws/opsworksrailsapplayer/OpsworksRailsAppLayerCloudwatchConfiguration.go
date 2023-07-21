package opsworksrailsapplayer


type OpsworksRailsAppLayerCloudwatchConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/opsworks_rails_app_layer#enabled OpsworksRailsAppLayer#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// log_streams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/opsworks_rails_app_layer#log_streams OpsworksRailsAppLayer#log_streams}
	LogStreams interface{} `field:"optional" json:"logStreams" yaml:"logStreams"`
}

