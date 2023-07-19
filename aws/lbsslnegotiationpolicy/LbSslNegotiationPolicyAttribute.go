package lbsslnegotiationpolicy


type LbSslNegotiationPolicyAttribute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/lb_ssl_negotiation_policy#name LbSslNegotiationPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/lb_ssl_negotiation_policy#value LbSslNegotiationPolicy#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

