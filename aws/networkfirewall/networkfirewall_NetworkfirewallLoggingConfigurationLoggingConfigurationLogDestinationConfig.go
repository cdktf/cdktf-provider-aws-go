package networkfirewall


type NetworkfirewallLoggingConfigurationLoggingConfigurationLogDestinationConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_logging_configuration#log_destination NetworkfirewallLoggingConfiguration#log_destination}.
	LogDestination *map[string]*string `field:"required" json:"logDestination" yaml:"logDestination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_logging_configuration#log_destination_type NetworkfirewallLoggingConfiguration#log_destination_type}.
	LogDestinationType *string `field:"required" json:"logDestinationType" yaml:"logDestinationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_logging_configuration#log_type NetworkfirewallLoggingConfiguration#log_type}.
	LogType *string `field:"required" json:"logType" yaml:"logType"`
}

