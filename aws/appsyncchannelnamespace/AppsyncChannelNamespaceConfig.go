// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncchannelnamespace

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsyncChannelNamespaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appsync_channel_namespace#api_id AppsyncChannelNamespace#api_id}.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appsync_channel_namespace#name AppsyncChannelNamespace#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appsync_channel_namespace#code_handlers AppsyncChannelNamespace#code_handlers}.
	CodeHandlers *string `field:"optional" json:"codeHandlers" yaml:"codeHandlers"`
	// handler_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appsync_channel_namespace#handler_configs AppsyncChannelNamespace#handler_configs}
	HandlerConfigs interface{} `field:"optional" json:"handlerConfigs" yaml:"handlerConfigs"`
	// publish_auth_mode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appsync_channel_namespace#publish_auth_mode AppsyncChannelNamespace#publish_auth_mode}
	PublishAuthMode interface{} `field:"optional" json:"publishAuthMode" yaml:"publishAuthMode"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appsync_channel_namespace#region AppsyncChannelNamespace#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// subscribe_auth_mode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appsync_channel_namespace#subscribe_auth_mode AppsyncChannelNamespace#subscribe_auth_mode}
	SubscribeAuthMode interface{} `field:"optional" json:"subscribeAuthMode" yaml:"subscribeAuthMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appsync_channel_namespace#tags AppsyncChannelNamespace#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

