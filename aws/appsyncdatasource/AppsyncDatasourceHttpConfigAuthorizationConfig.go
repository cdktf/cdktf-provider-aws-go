// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource


type AppsyncDatasourceHttpConfigAuthorizationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/appsync_datasource#authorization_type AppsyncDatasource#authorization_type}.
	AuthorizationType *string `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// aws_iam_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/appsync_datasource#aws_iam_config AppsyncDatasource#aws_iam_config}
	AwsIamConfig *AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig `field:"optional" json:"awsIamConfig" yaml:"awsIamConfig"`
}

