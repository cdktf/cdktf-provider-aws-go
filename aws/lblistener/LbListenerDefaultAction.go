// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistener


type LbListenerDefaultAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/lb_listener#type LbListener#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// authenticate_cognito block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/lb_listener#authenticate_cognito LbListener#authenticate_cognito}
	AuthenticateCognito *LbListenerDefaultActionAuthenticateCognito `field:"optional" json:"authenticateCognito" yaml:"authenticateCognito"`
	// authenticate_oidc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/lb_listener#authenticate_oidc LbListener#authenticate_oidc}
	AuthenticateOidc *LbListenerDefaultActionAuthenticateOidc `field:"optional" json:"authenticateOidc" yaml:"authenticateOidc"`
	// fixed_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/lb_listener#fixed_response LbListener#fixed_response}
	FixedResponse *LbListenerDefaultActionFixedResponse `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// forward block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/lb_listener#forward LbListener#forward}
	Forward *LbListenerDefaultActionForward `field:"optional" json:"forward" yaml:"forward"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/lb_listener#order LbListener#order}.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/lb_listener#redirect LbListener#redirect}
	Redirect *LbListenerDefaultActionRedirect `field:"optional" json:"redirect" yaml:"redirect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/lb_listener#target_group_arn LbListener#target_group_arn}.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

