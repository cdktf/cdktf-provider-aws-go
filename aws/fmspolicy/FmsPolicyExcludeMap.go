package fmspolicy


type FmsPolicyExcludeMap struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/fms_policy#account FmsPolicy#account}.
	Account *[]*string `field:"optional" json:"account" yaml:"account"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/fms_policy#orgunit FmsPolicy#orgunit}.
	Orgunit *[]*string `field:"optional" json:"orgunit" yaml:"orgunit"`
}

