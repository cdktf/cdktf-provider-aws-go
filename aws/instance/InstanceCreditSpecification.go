package instance


type InstanceCreditSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/instance#cpu_credits Instance#cpu_credits}.
	CpuCredits *string `field:"optional" json:"cpuCredits" yaml:"cpuCredits"`
}

