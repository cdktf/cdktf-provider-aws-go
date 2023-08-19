package inspector2organizationconfiguration


type Inspector2OrganizationConfigurationAutoEnable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/inspector2_organization_configuration#ec2 Inspector2OrganizationConfiguration#ec2}.
	Ec2 interface{} `field:"required" json:"ec2" yaml:"ec2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/inspector2_organization_configuration#ecr Inspector2OrganizationConfiguration#ecr}.
	Ecr interface{} `field:"required" json:"ecr" yaml:"ecr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/inspector2_organization_configuration#lambda Inspector2OrganizationConfiguration#lambda}.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
}

