// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowTaskConnectorOperator struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appflow_flow#amplitude AppflowFlow#amplitude}.
	Amplitude *string `field:"optional" json:"amplitude" yaml:"amplitude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appflow_flow#custom_connector AppflowFlow#custom_connector}.
	CustomConnector *string `field:"optional" json:"customConnector" yaml:"customConnector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appflow_flow#datadog AppflowFlow#datadog}.
	Datadog *string `field:"optional" json:"datadog" yaml:"datadog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appflow_flow#dynatrace AppflowFlow#dynatrace}.
	Dynatrace *string `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appflow_flow#google_analytics AppflowFlow#google_analytics}.
	GoogleAnalytics *string `field:"optional" json:"googleAnalytics" yaml:"googleAnalytics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appflow_flow#infor_nexus AppflowFlow#infor_nexus}.
	InforNexus *string `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appflow_flow#marketo AppflowFlow#marketo}.
	Marketo *string `field:"optional" json:"marketo" yaml:"marketo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appflow_flow#s3 AppflowFlow#s3}.
	S3 *string `field:"optional" json:"s3" yaml:"s3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appflow_flow#salesforce AppflowFlow#salesforce}.
	Salesforce *string `field:"optional" json:"salesforce" yaml:"salesforce"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appflow_flow#sapo_data AppflowFlow#sapo_data}.
	SapoData *string `field:"optional" json:"sapoData" yaml:"sapoData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appflow_flow#service_now AppflowFlow#service_now}.
	ServiceNow *string `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appflow_flow#singular AppflowFlow#singular}.
	Singular *string `field:"optional" json:"singular" yaml:"singular"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appflow_flow#slack AppflowFlow#slack}.
	Slack *string `field:"optional" json:"slack" yaml:"slack"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appflow_flow#trendmicro AppflowFlow#trendmicro}.
	Trendmicro *string `field:"optional" json:"trendmicro" yaml:"trendmicro"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appflow_flow#veeva AppflowFlow#veeva}.
	Veeva *string `field:"optional" json:"veeva" yaml:"veeva"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appflow_flow#zendesk AppflowFlow#zendesk}.
	Zendesk *string `field:"optional" json:"zendesk" yaml:"zendesk"`
}

