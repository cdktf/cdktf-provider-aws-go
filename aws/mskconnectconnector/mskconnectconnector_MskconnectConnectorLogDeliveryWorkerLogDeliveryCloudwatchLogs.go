package mskconnectconnector


type MskconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mskconnect_connector#enabled MskconnectConnector#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mskconnect_connector#log_group MskconnectConnector#log_group}.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

