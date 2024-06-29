// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflake struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/appflow_connector_profile#bucket_name AppflowConnectorProfile#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/appflow_connector_profile#stage AppflowConnectorProfile#stage}.
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/appflow_connector_profile#warehouse AppflowConnectorProfile#warehouse}.
	Warehouse *string `field:"required" json:"warehouse" yaml:"warehouse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/appflow_connector_profile#account_name AppflowConnectorProfile#account_name}.
	AccountName *string `field:"optional" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/appflow_connector_profile#bucket_prefix AppflowConnectorProfile#bucket_prefix}.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/appflow_connector_profile#private_link_service_name AppflowConnectorProfile#private_link_service_name}.
	PrivateLinkServiceName *string `field:"optional" json:"privateLinkServiceName" yaml:"privateLinkServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/appflow_connector_profile#region AppflowConnectorProfile#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

