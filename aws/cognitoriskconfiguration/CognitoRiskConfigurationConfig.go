// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitoriskconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CognitoRiskConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/cognito_risk_configuration#user_pool_id CognitoRiskConfiguration#user_pool_id}.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// account_takeover_risk_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/cognito_risk_configuration#account_takeover_risk_configuration CognitoRiskConfiguration#account_takeover_risk_configuration}
	AccountTakeoverRiskConfiguration *CognitoRiskConfigurationAccountTakeoverRiskConfiguration `field:"optional" json:"accountTakeoverRiskConfiguration" yaml:"accountTakeoverRiskConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/cognito_risk_configuration#client_id CognitoRiskConfiguration#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// compromised_credentials_risk_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/cognito_risk_configuration#compromised_credentials_risk_configuration CognitoRiskConfiguration#compromised_credentials_risk_configuration}
	CompromisedCredentialsRiskConfiguration *CognitoRiskConfigurationCompromisedCredentialsRiskConfiguration `field:"optional" json:"compromisedCredentialsRiskConfiguration" yaml:"compromisedCredentialsRiskConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/cognito_risk_configuration#id CognitoRiskConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// risk_exception_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/cognito_risk_configuration#risk_exception_configuration CognitoRiskConfiguration#risk_exception_configuration}
	RiskExceptionConfiguration *CognitoRiskConfigurationRiskExceptionConfiguration `field:"optional" json:"riskExceptionConfiguration" yaml:"riskExceptionConfiguration"`
}

