// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsreplicationinstance


type DmsReplicationInstanceKerberosAuthenticationSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_replication_instance#key_cache_secret_iam_arn DmsReplicationInstance#key_cache_secret_iam_arn}.
	KeyCacheSecretIamArn *string `field:"required" json:"keyCacheSecretIamArn" yaml:"keyCacheSecretIamArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_replication_instance#key_cache_secret_id DmsReplicationInstance#key_cache_secret_id}.
	KeyCacheSecretId *string `field:"required" json:"keyCacheSecretId" yaml:"keyCacheSecretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_replication_instance#krb5_file_contents DmsReplicationInstance#krb5_file_contents}.
	Krb5FileContents *string `field:"required" json:"krb5FileContents" yaml:"krb5FileContents"`
}

