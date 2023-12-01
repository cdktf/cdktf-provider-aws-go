// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileProperties struct {
	// amplitude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#amplitude AppflowConnectorProfile#amplitude}
	Amplitude *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesAmplitude `field:"optional" json:"amplitude" yaml:"amplitude"`
	// custom_connector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#custom_connector AppflowConnectorProfile#custom_connector}
	CustomConnector *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnector `field:"optional" json:"customConnector" yaml:"customConnector"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#datadog AppflowConnectorProfile#datadog}
	Datadog *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDatadog `field:"optional" json:"datadog" yaml:"datadog"`
	// dynatrace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#dynatrace AppflowConnectorProfile#dynatrace}
	Dynatrace *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDynatrace `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// google_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#google_analytics AppflowConnectorProfile#google_analytics}
	GoogleAnalytics *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesGoogleAnalytics `field:"optional" json:"googleAnalytics" yaml:"googleAnalytics"`
	// honeycode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#honeycode AppflowConnectorProfile#honeycode}
	Honeycode *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesHoneycode `field:"optional" json:"honeycode" yaml:"honeycode"`
	// infor_nexus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#infor_nexus AppflowConnectorProfile#infor_nexus}
	InforNexus *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesInforNexus `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// marketo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#marketo AppflowConnectorProfile#marketo}
	Marketo *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesMarketo `field:"optional" json:"marketo" yaml:"marketo"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#redshift AppflowConnectorProfile#redshift}
	Redshift *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesRedshift `field:"optional" json:"redshift" yaml:"redshift"`
	// salesforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#salesforce AppflowConnectorProfile#salesforce}
	Salesforce *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSalesforce `field:"optional" json:"salesforce" yaml:"salesforce"`
	// sapo_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#sapo_data AppflowConnectorProfile#sapo_data}
	SapoData *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoData `field:"optional" json:"sapoData" yaml:"sapoData"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#service_now AppflowConnectorProfile#service_now}
	ServiceNow *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesServiceNow `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// singular block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#singular AppflowConnectorProfile#singular}
	Singular *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSingular `field:"optional" json:"singular" yaml:"singular"`
	// slack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#slack AppflowConnectorProfile#slack}
	Slack *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSlack `field:"optional" json:"slack" yaml:"slack"`
	// snowflake block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#snowflake AppflowConnectorProfile#snowflake}
	Snowflake *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflake `field:"optional" json:"snowflake" yaml:"snowflake"`
	// trendmicro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#trendmicro AppflowConnectorProfile#trendmicro}
	Trendmicro *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesTrendmicro `field:"optional" json:"trendmicro" yaml:"trendmicro"`
	// veeva block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#veeva AppflowConnectorProfile#veeva}
	Veeva *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesVeeva `field:"optional" json:"veeva" yaml:"veeva"`
	// zendesk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appflow_connector_profile#zendesk AppflowConnectorProfile#zendesk}
	Zendesk *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesZendesk `field:"optional" json:"zendesk" yaml:"zendesk"`
}

