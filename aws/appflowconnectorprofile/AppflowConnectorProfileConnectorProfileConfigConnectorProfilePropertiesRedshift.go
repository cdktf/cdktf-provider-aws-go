// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesRedshift struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/appflow_connector_profile#bucket_name AppflowConnectorProfile#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/appflow_connector_profile#role_arn AppflowConnectorProfile#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/appflow_connector_profile#bucket_prefix AppflowConnectorProfile#bucket_prefix}.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/appflow_connector_profile#cluster_identifier AppflowConnectorProfile#cluster_identifier}.
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/appflow_connector_profile#data_api_role_arn AppflowConnectorProfile#data_api_role_arn}.
	DataApiRoleArn *string `field:"optional" json:"dataApiRoleArn" yaml:"dataApiRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/appflow_connector_profile#database_name AppflowConnectorProfile#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/appflow_connector_profile#database_url AppflowConnectorProfile#database_url}.
	DatabaseUrl *string `field:"optional" json:"databaseUrl" yaml:"databaseUrl"`
}

