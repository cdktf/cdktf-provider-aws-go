// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type MedialiveInputSecurityGroupWhitelistRules struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_input_security_group#cidr MedialiveInputSecurityGroup#cidr}.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
}

