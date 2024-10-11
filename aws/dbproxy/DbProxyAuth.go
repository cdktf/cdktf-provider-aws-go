// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dbproxy


type DbProxyAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/db_proxy#auth_scheme DbProxy#auth_scheme}.
	AuthScheme *string `field:"optional" json:"authScheme" yaml:"authScheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/db_proxy#client_password_auth_type DbProxy#client_password_auth_type}.
	ClientPasswordAuthType *string `field:"optional" json:"clientPasswordAuthType" yaml:"clientPasswordAuthType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/db_proxy#description DbProxy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/db_proxy#iam_auth DbProxy#iam_auth}.
	IamAuth *string `field:"optional" json:"iamAuth" yaml:"iamAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/db_proxy#secret_arn DbProxy#secret_arn}.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/db_proxy#username DbProxy#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

