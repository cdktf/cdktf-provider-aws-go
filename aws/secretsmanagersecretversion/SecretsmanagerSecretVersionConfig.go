// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretsmanagersecretversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecretsmanagerSecretVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/secretsmanager_secret_version#secret_id SecretsmanagerSecretVersion#secret_id}.
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/secretsmanager_secret_version#id SecretsmanagerSecretVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/secretsmanager_secret_version#secret_binary SecretsmanagerSecretVersion#secret_binary}.
	SecretBinary *string `field:"optional" json:"secretBinary" yaml:"secretBinary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/secretsmanager_secret_version#secret_string SecretsmanagerSecretVersion#secret_string}.
	SecretString *string `field:"optional" json:"secretString" yaml:"secretString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/secretsmanager_secret_version#version_stages SecretsmanagerSecretVersion#version_stages}.
	VersionStages *[]*string `field:"optional" json:"versionStages" yaml:"versionStages"`
}

