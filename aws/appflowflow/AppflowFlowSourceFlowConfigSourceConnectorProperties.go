// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowSourceFlowConfigSourceConnectorProperties struct {
	// amplitude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/appflow_flow#amplitude AppflowFlow#amplitude}
	Amplitude *AppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitude `field:"optional" json:"amplitude" yaml:"amplitude"`
	// custom_connector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/appflow_flow#custom_connector AppflowFlow#custom_connector}
	CustomConnector *AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnector `field:"optional" json:"customConnector" yaml:"customConnector"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/appflow_flow#datadog AppflowFlow#datadog}
	Datadog *AppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadog `field:"optional" json:"datadog" yaml:"datadog"`
	// dynatrace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/appflow_flow#dynatrace AppflowFlow#dynatrace}
	Dynatrace *AppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatrace `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// google_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/appflow_flow#google_analytics AppflowFlow#google_analytics}
	GoogleAnalytics *AppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalytics `field:"optional" json:"googleAnalytics" yaml:"googleAnalytics"`
	// infor_nexus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/appflow_flow#infor_nexus AppflowFlow#infor_nexus}
	InforNexus *AppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexus `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// marketo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/appflow_flow#marketo AppflowFlow#marketo}
	Marketo *AppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketo `field:"optional" json:"marketo" yaml:"marketo"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/appflow_flow#s3 AppflowFlow#s3}
	S3 *AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3 `field:"optional" json:"s3" yaml:"s3"`
	// salesforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/appflow_flow#salesforce AppflowFlow#salesforce}
	Salesforce *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforce `field:"optional" json:"salesforce" yaml:"salesforce"`
	// sapo_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/appflow_flow#sapo_data AppflowFlow#sapo_data}
	SapoData *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoData `field:"optional" json:"sapoData" yaml:"sapoData"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/appflow_flow#service_now AppflowFlow#service_now}
	ServiceNow *AppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNow `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// singular block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/appflow_flow#singular AppflowFlow#singular}
	Singular *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSingular `field:"optional" json:"singular" yaml:"singular"`
	// slack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/appflow_flow#slack AppflowFlow#slack}
	Slack *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSlack `field:"optional" json:"slack" yaml:"slack"`
	// trendmicro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/appflow_flow#trendmicro AppflowFlow#trendmicro}
	Trendmicro *AppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicro `field:"optional" json:"trendmicro" yaml:"trendmicro"`
	// veeva block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/appflow_flow#veeva AppflowFlow#veeva}
	Veeva *AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeeva `field:"optional" json:"veeva" yaml:"veeva"`
	// zendesk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/appflow_flow#zendesk AppflowFlow#zendesk}
	Zendesk *AppflowFlowSourceFlowConfigSourceConnectorPropertiesZendesk `field:"optional" json:"zendesk" yaml:"zendesk"`
}

