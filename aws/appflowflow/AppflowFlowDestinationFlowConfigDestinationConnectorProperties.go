// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowDestinationFlowConfigDestinationConnectorProperties struct {
	// custom_connector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/appflow_flow#custom_connector AppflowFlow#custom_connector}
	CustomConnector *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnector `field:"optional" json:"customConnector" yaml:"customConnector"`
	// customer_profiles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/appflow_flow#customer_profiles AppflowFlow#customer_profiles}
	CustomerProfiles *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfiles `field:"optional" json:"customerProfiles" yaml:"customerProfiles"`
	// event_bridge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/appflow_flow#event_bridge AppflowFlow#event_bridge}
	EventBridge *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridge `field:"optional" json:"eventBridge" yaml:"eventBridge"`
	// honeycode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/appflow_flow#honeycode AppflowFlow#honeycode}
	Honeycode *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycode `field:"optional" json:"honeycode" yaml:"honeycode"`
	// lookout_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/appflow_flow#lookout_metrics AppflowFlow#lookout_metrics}
	LookoutMetrics *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesLookoutMetrics `field:"optional" json:"lookoutMetrics" yaml:"lookoutMetrics"`
	// marketo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/appflow_flow#marketo AppflowFlow#marketo}
	Marketo *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketo `field:"optional" json:"marketo" yaml:"marketo"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/appflow_flow#redshift AppflowFlow#redshift}
	Redshift *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshift `field:"optional" json:"redshift" yaml:"redshift"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/appflow_flow#s3 AppflowFlow#s3}
	S3 *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3 `field:"optional" json:"s3" yaml:"s3"`
	// salesforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/appflow_flow#salesforce AppflowFlow#salesforce}
	Salesforce *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforce `field:"optional" json:"salesforce" yaml:"salesforce"`
	// sapo_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/appflow_flow#sapo_data AppflowFlow#sapo_data}
	SapoData *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoData `field:"optional" json:"sapoData" yaml:"sapoData"`
	// snowflake block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/appflow_flow#snowflake AppflowFlow#snowflake}
	Snowflake *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflake `field:"optional" json:"snowflake" yaml:"snowflake"`
	// upsolver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/appflow_flow#upsolver AppflowFlow#upsolver}
	Upsolver *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolver `field:"optional" json:"upsolver" yaml:"upsolver"`
	// zendesk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/appflow_flow#zendesk AppflowFlow#zendesk}
	Zendesk *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendesk `field:"optional" json:"zendesk" yaml:"zendesk"`
}

