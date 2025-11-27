// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspaceswebidentityprovider

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspaceswebIdentityProviderConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/workspacesweb_identity_provider#identity_provider_details WorkspaceswebIdentityProvider#identity_provider_details}.
	IdentityProviderDetails *map[string]*string `field:"required" json:"identityProviderDetails" yaml:"identityProviderDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/workspacesweb_identity_provider#identity_provider_name WorkspaceswebIdentityProvider#identity_provider_name}.
	IdentityProviderName *string `field:"required" json:"identityProviderName" yaml:"identityProviderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/workspacesweb_identity_provider#identity_provider_type WorkspaceswebIdentityProvider#identity_provider_type}.
	IdentityProviderType *string `field:"required" json:"identityProviderType" yaml:"identityProviderType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/workspacesweb_identity_provider#portal_arn WorkspaceswebIdentityProvider#portal_arn}.
	PortalArn *string `field:"required" json:"portalArn" yaml:"portalArn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/workspacesweb_identity_provider#region WorkspaceswebIdentityProvider#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/workspacesweb_identity_provider#tags WorkspaceswebIdentityProvider#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

