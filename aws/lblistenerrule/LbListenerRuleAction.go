package lblistenerrule


type LbListenerRuleAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lb_listener_rule#type LbListenerRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// authenticate_cognito block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lb_listener_rule#authenticate_cognito LbListenerRule#authenticate_cognito}
	AuthenticateCognito *LbListenerRuleActionAuthenticateCognito `field:"optional" json:"authenticateCognito" yaml:"authenticateCognito"`
	// authenticate_oidc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lb_listener_rule#authenticate_oidc LbListenerRule#authenticate_oidc}
	AuthenticateOidc *LbListenerRuleActionAuthenticateOidc `field:"optional" json:"authenticateOidc" yaml:"authenticateOidc"`
	// fixed_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lb_listener_rule#fixed_response LbListenerRule#fixed_response}
	FixedResponse *LbListenerRuleActionFixedResponse `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// forward block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lb_listener_rule#forward LbListenerRule#forward}
	Forward *LbListenerRuleActionForward `field:"optional" json:"forward" yaml:"forward"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lb_listener_rule#order LbListenerRule#order}.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lb_listener_rule#redirect LbListenerRule#redirect}
	Redirect *LbListenerRuleActionRedirect `field:"optional" json:"redirect" yaml:"redirect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lb_listener_rule#target_group_arn LbListenerRule#target_group_arn}.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

