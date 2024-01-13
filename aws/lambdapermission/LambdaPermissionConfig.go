// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdapermission

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LambdaPermissionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/lambda_permission#action LambdaPermission#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/lambda_permission#function_name LambdaPermission#function_name}.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/lambda_permission#principal LambdaPermission#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/lambda_permission#event_source_token LambdaPermission#event_source_token}.
	EventSourceToken *string `field:"optional" json:"eventSourceToken" yaml:"eventSourceToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/lambda_permission#function_url_auth_type LambdaPermission#function_url_auth_type}.
	FunctionUrlAuthType *string `field:"optional" json:"functionUrlAuthType" yaml:"functionUrlAuthType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/lambda_permission#id LambdaPermission#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/lambda_permission#principal_org_id LambdaPermission#principal_org_id}.
	PrincipalOrgId *string `field:"optional" json:"principalOrgId" yaml:"principalOrgId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/lambda_permission#qualifier LambdaPermission#qualifier}.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/lambda_permission#source_account LambdaPermission#source_account}.
	SourceAccount *string `field:"optional" json:"sourceAccount" yaml:"sourceAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/lambda_permission#source_arn LambdaPermission#source_arn}.
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/lambda_permission#statement_id LambdaPermission#statement_id}.
	StatementId *string `field:"optional" json:"statementId" yaml:"statementId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/lambda_permission#statement_id_prefix LambdaPermission#statement_id_prefix}.
	StatementIdPrefix *string `field:"optional" json:"statementIdPrefix" yaml:"statementIdPrefix"`
}

