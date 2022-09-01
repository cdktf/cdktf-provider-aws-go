// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AccountAlternateContactTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/account_alternate_contact#create AccountAlternateContact#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/account_alternate_contact#delete AccountAlternateContact#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/account_alternate_contact#update AccountAlternateContact#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

