// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type MedialiveInputSecurityGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_input_security_group#create MedialiveInputSecurityGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_input_security_group#delete MedialiveInputSecurityGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_input_security_group#update MedialiveInputSecurityGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

