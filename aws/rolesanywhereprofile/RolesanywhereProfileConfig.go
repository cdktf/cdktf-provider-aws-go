// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolesanywhereprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RolesanywhereProfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/rolesanywhere_profile#name RolesanywhereProfile#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/rolesanywhere_profile#role_arns RolesanywhereProfile#role_arns}.
	RoleArns *[]*string `field:"required" json:"roleArns" yaml:"roleArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/rolesanywhere_profile#duration_seconds RolesanywhereProfile#duration_seconds}.
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/rolesanywhere_profile#enabled RolesanywhereProfile#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/rolesanywhere_profile#id RolesanywhereProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/rolesanywhere_profile#managed_policy_arns RolesanywhereProfile#managed_policy_arns}.
	ManagedPolicyArns *[]*string `field:"optional" json:"managedPolicyArns" yaml:"managedPolicyArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/rolesanywhere_profile#require_instance_properties RolesanywhereProfile#require_instance_properties}.
	RequireInstanceProperties interface{} `field:"optional" json:"requireInstanceProperties" yaml:"requireInstanceProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/rolesanywhere_profile#session_policy RolesanywhereProfile#session_policy}.
	SessionPolicy *string `field:"optional" json:"sessionPolicy" yaml:"sessionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/rolesanywhere_profile#tags RolesanywhereProfile#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/rolesanywhere_profile#tags_all RolesanywhereProfile#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

