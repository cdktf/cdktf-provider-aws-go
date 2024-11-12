// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretsmanagersecret

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecretsmanagerSecretConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/secretsmanager_secret#description SecretsmanagerSecret#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/secretsmanager_secret#force_overwrite_replica_secret SecretsmanagerSecret#force_overwrite_replica_secret}.
	ForceOverwriteReplicaSecret interface{} `field:"optional" json:"forceOverwriteReplicaSecret" yaml:"forceOverwriteReplicaSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/secretsmanager_secret#id SecretsmanagerSecret#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/secretsmanager_secret#kms_key_id SecretsmanagerSecret#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/secretsmanager_secret#name SecretsmanagerSecret#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/secretsmanager_secret#name_prefix SecretsmanagerSecret#name_prefix}.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/secretsmanager_secret#policy SecretsmanagerSecret#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/secretsmanager_secret#recovery_window_in_days SecretsmanagerSecret#recovery_window_in_days}.
	RecoveryWindowInDays *float64 `field:"optional" json:"recoveryWindowInDays" yaml:"recoveryWindowInDays"`
	// replica block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/secretsmanager_secret#replica SecretsmanagerSecret#replica}
	Replica interface{} `field:"optional" json:"replica" yaml:"replica"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/secretsmanager_secret#tags SecretsmanagerSecret#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/secretsmanager_secret#tags_all SecretsmanagerSecret#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

