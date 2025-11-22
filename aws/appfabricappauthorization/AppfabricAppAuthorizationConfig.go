// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appfabricappauthorization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppfabricAppAuthorizationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appfabric_app_authorization#app AppfabricAppAuthorization#app}.
	App *string `field:"required" json:"app" yaml:"app"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appfabric_app_authorization#app_bundle_arn AppfabricAppAuthorization#app_bundle_arn}.
	AppBundleArn *string `field:"required" json:"appBundleArn" yaml:"appBundleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appfabric_app_authorization#auth_type AppfabricAppAuthorization#auth_type}.
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appfabric_app_authorization#credential AppfabricAppAuthorization#credential}
	Credential interface{} `field:"optional" json:"credential" yaml:"credential"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appfabric_app_authorization#region AppfabricAppAuthorization#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appfabric_app_authorization#tags AppfabricAppAuthorization#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// tenant block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appfabric_app_authorization#tenant AppfabricAppAuthorization#tenant}
	Tenant interface{} `field:"optional" json:"tenant" yaml:"tenant"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appfabric_app_authorization#timeouts AppfabricAppAuthorization#timeouts}
	Timeouts *AppfabricAppAuthorizationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

