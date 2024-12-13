// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawssecretsmanagersecretversions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsSecretsmanagerSecretVersionsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/data-sources/secretsmanager_secret_versions#secret_id DataAwsSecretsmanagerSecretVersions#secret_id}.
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/data-sources/secretsmanager_secret_versions#include_deprecated DataAwsSecretsmanagerSecretVersions#include_deprecated}.
	IncludeDeprecated interface{} `field:"optional" json:"includeDeprecated" yaml:"includeDeprecated"`
}

