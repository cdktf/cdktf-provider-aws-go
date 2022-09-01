package ec2


type Ec2ClientVpnAuthorizationRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_client_vpn_authorization_rule#create Ec2ClientVpnAuthorizationRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_client_vpn_authorization_rule#delete Ec2ClientVpnAuthorizationRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

