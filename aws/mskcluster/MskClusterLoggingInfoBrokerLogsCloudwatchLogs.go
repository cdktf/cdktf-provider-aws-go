package mskcluster


type MskClusterLoggingInfoBrokerLogsCloudwatchLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/msk_cluster#enabled MskCluster#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/msk_cluster#log_group MskCluster#log_group}.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

