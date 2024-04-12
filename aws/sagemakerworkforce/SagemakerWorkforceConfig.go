// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkforce

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerWorkforceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sagemaker_workforce#workforce_name SagemakerWorkforce#workforce_name}.
	WorkforceName *string `field:"required" json:"workforceName" yaml:"workforceName"`
	// cognito_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sagemaker_workforce#cognito_config SagemakerWorkforce#cognito_config}
	CognitoConfig *SagemakerWorkforceCognitoConfig `field:"optional" json:"cognitoConfig" yaml:"cognitoConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sagemaker_workforce#id SagemakerWorkforce#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// oidc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sagemaker_workforce#oidc_config SagemakerWorkforce#oidc_config}
	OidcConfig *SagemakerWorkforceOidcConfig `field:"optional" json:"oidcConfig" yaml:"oidcConfig"`
	// source_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sagemaker_workforce#source_ip_config SagemakerWorkforce#source_ip_config}
	SourceIpConfig *SagemakerWorkforceSourceIpConfig `field:"optional" json:"sourceIpConfig" yaml:"sourceIpConfig"`
	// workforce_vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/sagemaker_workforce#workforce_vpc_config SagemakerWorkforce#workforce_vpc_config}
	WorkforceVpcConfig *SagemakerWorkforceWorkforceVpcConfig `field:"optional" json:"workforceVpcConfig" yaml:"workforceVpcConfig"`
}

