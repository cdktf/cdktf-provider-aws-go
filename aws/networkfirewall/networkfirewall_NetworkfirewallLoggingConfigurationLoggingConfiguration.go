package networkfirewall


type NetworkfirewallLoggingConfigurationLoggingConfiguration struct {
	// log_destination_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_logging_configuration#log_destination_config NetworkfirewallLoggingConfiguration#log_destination_config}
	LogDestinationConfig interface{} `field:"required" json:"logDestinationConfig" yaml:"logDestinationConfig"`
}

