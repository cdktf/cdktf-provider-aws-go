// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclAssociationConfigRequestBody struct {
	// api_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.54.1/docs/resources/wafv2_web_acl#api_gateway Wafv2WebAcl#api_gateway}
	ApiGateway interface{} `field:"optional" json:"apiGateway" yaml:"apiGateway"`
	// app_runner_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.54.1/docs/resources/wafv2_web_acl#app_runner_service Wafv2WebAcl#app_runner_service}
	AppRunnerService interface{} `field:"optional" json:"appRunnerService" yaml:"appRunnerService"`
	// cloudfront block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.54.1/docs/resources/wafv2_web_acl#cloudfront Wafv2WebAcl#cloudfront}
	Cloudfront interface{} `field:"optional" json:"cloudfront" yaml:"cloudfront"`
	// cognito_user_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.54.1/docs/resources/wafv2_web_acl#cognito_user_pool Wafv2WebAcl#cognito_user_pool}
	CognitoUserPool interface{} `field:"optional" json:"cognitoUserPool" yaml:"cognitoUserPool"`
	// verified_access_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.54.1/docs/resources/wafv2_web_acl#verified_access_instance Wafv2WebAcl#verified_access_instance}
	VerifiedAccessInstance interface{} `field:"optional" json:"verifiedAccessInstance" yaml:"verifiedAccessInstance"`
}

