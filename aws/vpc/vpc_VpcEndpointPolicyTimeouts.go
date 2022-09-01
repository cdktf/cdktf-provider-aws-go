package vpc


type VpcEndpointPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/vpc_endpoint_policy#create VpcEndpointPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/vpc_endpoint_policy#delete VpcEndpointPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

