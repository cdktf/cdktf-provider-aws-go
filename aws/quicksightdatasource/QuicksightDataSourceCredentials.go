// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceCredentials struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/quicksight_data_source#copy_source_arn QuicksightDataSource#copy_source_arn}.
	CopySourceArn *string `field:"optional" json:"copySourceArn" yaml:"copySourceArn"`
	// credential_pair block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/quicksight_data_source#credential_pair QuicksightDataSource#credential_pair}
	CredentialPair *QuicksightDataSourceCredentialsCredentialPair `field:"optional" json:"credentialPair" yaml:"credentialPair"`
}

