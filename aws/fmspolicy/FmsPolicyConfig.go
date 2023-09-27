// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fmspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FmsPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/fms_policy#exclude_resource_tags FmsPolicy#exclude_resource_tags}.
	ExcludeResourceTags interface{} `field:"required" json:"excludeResourceTags" yaml:"excludeResourceTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/fms_policy#name FmsPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// security_service_policy_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/fms_policy#security_service_policy_data FmsPolicy#security_service_policy_data}
	SecurityServicePolicyData *FmsPolicySecurityServicePolicyData `field:"required" json:"securityServicePolicyData" yaml:"securityServicePolicyData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/fms_policy#delete_all_policy_resources FmsPolicy#delete_all_policy_resources}.
	DeleteAllPolicyResources interface{} `field:"optional" json:"deleteAllPolicyResources" yaml:"deleteAllPolicyResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/fms_policy#delete_unused_fm_managed_resources FmsPolicy#delete_unused_fm_managed_resources}.
	DeleteUnusedFmManagedResources interface{} `field:"optional" json:"deleteUnusedFmManagedResources" yaml:"deleteUnusedFmManagedResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/fms_policy#description FmsPolicy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// exclude_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/fms_policy#exclude_map FmsPolicy#exclude_map}
	ExcludeMap *FmsPolicyExcludeMap `field:"optional" json:"excludeMap" yaml:"excludeMap"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/fms_policy#id FmsPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// include_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/fms_policy#include_map FmsPolicy#include_map}
	IncludeMap *FmsPolicyIncludeMap `field:"optional" json:"includeMap" yaml:"includeMap"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/fms_policy#remediation_enabled FmsPolicy#remediation_enabled}.
	RemediationEnabled interface{} `field:"optional" json:"remediationEnabled" yaml:"remediationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/fms_policy#resource_tags FmsPolicy#resource_tags}.
	ResourceTags *map[string]*string `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/fms_policy#resource_type FmsPolicy#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/fms_policy#resource_type_list FmsPolicy#resource_type_list}.
	ResourceTypeList *[]*string `field:"optional" json:"resourceTypeList" yaml:"resourceTypeList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/fms_policy#tags FmsPolicy#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/fms_policy#tags_all FmsPolicy#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

