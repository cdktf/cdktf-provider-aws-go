package cloudcontrolapiresource


type CloudcontrolapiResourceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/cloudcontrolapi_resource#create CloudcontrolapiResource#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/cloudcontrolapi_resource#delete CloudcontrolapiResource#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/cloudcontrolapi_resource#update CloudcontrolapiResource#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

