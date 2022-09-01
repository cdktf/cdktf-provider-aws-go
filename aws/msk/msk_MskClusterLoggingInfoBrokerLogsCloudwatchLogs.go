package msk


type MskClusterLoggingInfoBrokerLogsCloudwatchLogs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/msk_cluster#enabled MskCluster#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/msk_cluster#log_group MskCluster#log_group}.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

