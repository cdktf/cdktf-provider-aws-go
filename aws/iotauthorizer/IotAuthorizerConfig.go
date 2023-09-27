// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotauthorizer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotAuthorizerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/iot_authorizer#authorizer_function_arn IotAuthorizer#authorizer_function_arn}.
	AuthorizerFunctionArn *string `field:"required" json:"authorizerFunctionArn" yaml:"authorizerFunctionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/iot_authorizer#name IotAuthorizer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/iot_authorizer#enable_caching_for_http IotAuthorizer#enable_caching_for_http}.
	EnableCachingForHttp interface{} `field:"optional" json:"enableCachingForHttp" yaml:"enableCachingForHttp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/iot_authorizer#id IotAuthorizer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/iot_authorizer#signing_disabled IotAuthorizer#signing_disabled}.
	SigningDisabled interface{} `field:"optional" json:"signingDisabled" yaml:"signingDisabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/iot_authorizer#status IotAuthorizer#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/iot_authorizer#token_key_name IotAuthorizer#token_key_name}.
	TokenKeyName *string `field:"optional" json:"tokenKeyName" yaml:"tokenKeyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/iot_authorizer#token_signing_public_keys IotAuthorizer#token_signing_public_keys}.
	TokenSigningPublicKeys *map[string]*string `field:"optional" json:"tokenSigningPublicKeys" yaml:"tokenSigningPublicKeys"`
}

