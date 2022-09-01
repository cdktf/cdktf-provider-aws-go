package fms


type FmsPolicyExcludeMap struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy#account FmsPolicy#account}.
	Account *[]*string `field:"optional" json:"account" yaml:"account"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy#orgunit FmsPolicy#orgunit}.
	Orgunit *[]*string `field:"optional" json:"orgunit" yaml:"orgunit"`
}

