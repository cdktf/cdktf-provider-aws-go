package identitystoreuser


type IdentitystoreUserEmails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/identitystore_user#primary IdentitystoreUser#primary}.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/identitystore_user#type IdentitystoreUser#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/identitystore_user#value IdentitystoreUser#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

