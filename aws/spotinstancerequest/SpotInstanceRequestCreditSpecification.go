package spotinstancerequest


type SpotInstanceRequestCreditSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/spot_instance_request#cpu_credits SpotInstanceRequest#cpu_credits}.
	CpuCredits *string `field:"optional" json:"cpuCredits" yaml:"cpuCredits"`
}

