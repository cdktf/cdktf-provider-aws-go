// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grafanaworkspacesamlconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GrafanaWorkspaceSamlConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/grafana_workspace_saml_configuration#editor_role_values GrafanaWorkspaceSamlConfiguration#editor_role_values}.
	EditorRoleValues *[]*string `field:"required" json:"editorRoleValues" yaml:"editorRoleValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/grafana_workspace_saml_configuration#workspace_id GrafanaWorkspaceSamlConfiguration#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/grafana_workspace_saml_configuration#admin_role_values GrafanaWorkspaceSamlConfiguration#admin_role_values}.
	AdminRoleValues *[]*string `field:"optional" json:"adminRoleValues" yaml:"adminRoleValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/grafana_workspace_saml_configuration#allowed_organizations GrafanaWorkspaceSamlConfiguration#allowed_organizations}.
	AllowedOrganizations *[]*string `field:"optional" json:"allowedOrganizations" yaml:"allowedOrganizations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/grafana_workspace_saml_configuration#email_assertion GrafanaWorkspaceSamlConfiguration#email_assertion}.
	EmailAssertion *string `field:"optional" json:"emailAssertion" yaml:"emailAssertion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/grafana_workspace_saml_configuration#groups_assertion GrafanaWorkspaceSamlConfiguration#groups_assertion}.
	GroupsAssertion *string `field:"optional" json:"groupsAssertion" yaml:"groupsAssertion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/grafana_workspace_saml_configuration#id GrafanaWorkspaceSamlConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/grafana_workspace_saml_configuration#idp_metadata_url GrafanaWorkspaceSamlConfiguration#idp_metadata_url}.
	IdpMetadataUrl *string `field:"optional" json:"idpMetadataUrl" yaml:"idpMetadataUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/grafana_workspace_saml_configuration#idp_metadata_xml GrafanaWorkspaceSamlConfiguration#idp_metadata_xml}.
	IdpMetadataXml *string `field:"optional" json:"idpMetadataXml" yaml:"idpMetadataXml"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/grafana_workspace_saml_configuration#login_assertion GrafanaWorkspaceSamlConfiguration#login_assertion}.
	LoginAssertion *string `field:"optional" json:"loginAssertion" yaml:"loginAssertion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/grafana_workspace_saml_configuration#login_validity_duration GrafanaWorkspaceSamlConfiguration#login_validity_duration}.
	LoginValidityDuration *float64 `field:"optional" json:"loginValidityDuration" yaml:"loginValidityDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/grafana_workspace_saml_configuration#name_assertion GrafanaWorkspaceSamlConfiguration#name_assertion}.
	NameAssertion *string `field:"optional" json:"nameAssertion" yaml:"nameAssertion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/grafana_workspace_saml_configuration#org_assertion GrafanaWorkspaceSamlConfiguration#org_assertion}.
	OrgAssertion *string `field:"optional" json:"orgAssertion" yaml:"orgAssertion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/grafana_workspace_saml_configuration#role_assertion GrafanaWorkspaceSamlConfiguration#role_assertion}.
	RoleAssertion *string `field:"optional" json:"roleAssertion" yaml:"roleAssertion"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/grafana_workspace_saml_configuration#timeouts GrafanaWorkspaceSamlConfiguration#timeouts}
	Timeouts *GrafanaWorkspaceSamlConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

