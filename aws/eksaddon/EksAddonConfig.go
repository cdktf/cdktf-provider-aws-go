// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eksaddon

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EksAddonConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_addon#addon_name EksAddon#addon_name}.
	AddonName *string `field:"required" json:"addonName" yaml:"addonName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_addon#cluster_name EksAddon#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_addon#addon_version EksAddon#addon_version}.
	AddonVersion *string `field:"optional" json:"addonVersion" yaml:"addonVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_addon#configuration_values EksAddon#configuration_values}.
	ConfigurationValues *string `field:"optional" json:"configurationValues" yaml:"configurationValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_addon#id EksAddon#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_addon#preserve EksAddon#preserve}.
	Preserve interface{} `field:"optional" json:"preserve" yaml:"preserve"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_addon#resolve_conflicts EksAddon#resolve_conflicts}.
	ResolveConflicts *string `field:"optional" json:"resolveConflicts" yaml:"resolveConflicts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_addon#resolve_conflicts_on_create EksAddon#resolve_conflicts_on_create}.
	ResolveConflictsOnCreate *string `field:"optional" json:"resolveConflictsOnCreate" yaml:"resolveConflictsOnCreate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_addon#resolve_conflicts_on_update EksAddon#resolve_conflicts_on_update}.
	ResolveConflictsOnUpdate *string `field:"optional" json:"resolveConflictsOnUpdate" yaml:"resolveConflictsOnUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_addon#service_account_role_arn EksAddon#service_account_role_arn}.
	ServiceAccountRoleArn *string `field:"optional" json:"serviceAccountRoleArn" yaml:"serviceAccountRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_addon#tags EksAddon#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_addon#tags_all EksAddon#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/eks_addon#timeouts EksAddon#timeouts}
	Timeouts *EksAddonTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

