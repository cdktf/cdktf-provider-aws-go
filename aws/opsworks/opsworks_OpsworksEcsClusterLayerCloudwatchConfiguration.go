package opsworks


type OpsworksEcsClusterLayerCloudwatchConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ecs_cluster_layer#enabled OpsworksEcsClusterLayer#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// log_streams block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_ecs_cluster_layer#log_streams OpsworksEcsClusterLayer#log_streams}
	LogStreams interface{} `field:"optional" json:"logStreams" yaml:"logStreams"`
}

