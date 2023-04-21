package ivschatloggingconfiguration


type IvschatLoggingConfigurationDestinationConfigurationCloudwatchLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/ivschat_logging_configuration#log_group_name IvschatLoggingConfiguration#log_group_name}.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
}

