// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallloggingconfiguration


type NetworkfirewallLoggingConfigurationLoggingConfiguration struct {
	// log_destination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/networkfirewall_logging_configuration#log_destination_config NetworkfirewallLoggingConfiguration#log_destination_config}
	LogDestinationConfig interface{} `field:"required" json:"logDestinationConfig" yaml:"logDestinationConfig"`
}

