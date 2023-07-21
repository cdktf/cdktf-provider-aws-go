package loadbalancerpolicy


type LoadBalancerPolicyPolicyAttribute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/load_balancer_policy#name LoadBalancerPolicy#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/load_balancer_policy#value LoadBalancerPolicy#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

