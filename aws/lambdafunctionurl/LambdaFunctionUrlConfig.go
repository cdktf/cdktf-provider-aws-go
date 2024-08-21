// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdafunctionurl

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LambdaFunctionUrlConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/lambda_function_url#authorization_type LambdaFunctionUrl#authorization_type}.
	AuthorizationType *string `field:"required" json:"authorizationType" yaml:"authorizationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/lambda_function_url#function_name LambdaFunctionUrl#function_name}.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// cors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/lambda_function_url#cors LambdaFunctionUrl#cors}
	Cors *LambdaFunctionUrlCors `field:"optional" json:"cors" yaml:"cors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/lambda_function_url#id LambdaFunctionUrl#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/lambda_function_url#invoke_mode LambdaFunctionUrl#invoke_mode}.
	InvokeMode *string `field:"optional" json:"invokeMode" yaml:"invokeMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/lambda_function_url#qualifier LambdaFunctionUrl#qualifier}.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/lambda_function_url#timeouts LambdaFunctionUrl#timeouts}
	Timeouts *LambdaFunctionUrlTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

