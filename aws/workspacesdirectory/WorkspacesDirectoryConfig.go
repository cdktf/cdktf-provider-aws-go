// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspacesDirectoryConfig struct {
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
	// active_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/workspaces_directory#active_directory_config WorkspacesDirectory#active_directory_config}
	ActiveDirectoryConfig *WorkspacesDirectoryActiveDirectoryConfig `field:"optional" json:"activeDirectoryConfig" yaml:"activeDirectoryConfig"`
	// certificate_based_auth_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/workspaces_directory#certificate_based_auth_properties WorkspacesDirectory#certificate_based_auth_properties}
	CertificateBasedAuthProperties *WorkspacesDirectoryCertificateBasedAuthProperties `field:"optional" json:"certificateBasedAuthProperties" yaml:"certificateBasedAuthProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/workspaces_directory#directory_id WorkspacesDirectory#directory_id}.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/workspaces_directory#id WorkspacesDirectory#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/workspaces_directory#ip_group_ids WorkspacesDirectory#ip_group_ids}.
	IpGroupIds *[]*string `field:"optional" json:"ipGroupIds" yaml:"ipGroupIds"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/workspaces_directory#region WorkspacesDirectory#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// saml_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/workspaces_directory#saml_properties WorkspacesDirectory#saml_properties}
	SamlProperties *WorkspacesDirectorySamlProperties `field:"optional" json:"samlProperties" yaml:"samlProperties"`
	// self_service_permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/workspaces_directory#self_service_permissions WorkspacesDirectory#self_service_permissions}
	SelfServicePermissions *WorkspacesDirectorySelfServicePermissions `field:"optional" json:"selfServicePermissions" yaml:"selfServicePermissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/workspaces_directory#subnet_ids WorkspacesDirectory#subnet_ids}.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/workspaces_directory#tags WorkspacesDirectory#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/workspaces_directory#tags_all WorkspacesDirectory#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/workspaces_directory#user_identity_type WorkspacesDirectory#user_identity_type}.
	UserIdentityType *string `field:"optional" json:"userIdentityType" yaml:"userIdentityType"`
	// workspace_access_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/workspaces_directory#workspace_access_properties WorkspacesDirectory#workspace_access_properties}
	WorkspaceAccessProperties *WorkspacesDirectoryWorkspaceAccessProperties `field:"optional" json:"workspaceAccessProperties" yaml:"workspaceAccessProperties"`
	// workspace_creation_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/workspaces_directory#workspace_creation_properties WorkspacesDirectory#workspace_creation_properties}
	WorkspaceCreationProperties *WorkspacesDirectoryWorkspaceCreationProperties `field:"optional" json:"workspaceCreationProperties" yaml:"workspaceCreationProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/workspaces_directory#workspace_directory_description WorkspacesDirectory#workspace_directory_description}.
	WorkspaceDirectoryDescription *string `field:"optional" json:"workspaceDirectoryDescription" yaml:"workspaceDirectoryDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/workspaces_directory#workspace_directory_name WorkspacesDirectory#workspace_directory_name}.
	WorkspaceDirectoryName *string `field:"optional" json:"workspaceDirectoryName" yaml:"workspaceDirectoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/workspaces_directory#workspace_type WorkspacesDirectory#workspace_type}.
	WorkspaceType *string `field:"optional" json:"workspaceType" yaml:"workspaceType"`
}

