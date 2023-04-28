package opsworksapplication


type OpsworksApplicationSslConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#certificate OpsworksApplication#certificate}.
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#private_key OpsworksApplication#private_key}.
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#chain OpsworksApplication#chain}.
	Chain *string `field:"optional" json:"chain" yaml:"chain"`
}

