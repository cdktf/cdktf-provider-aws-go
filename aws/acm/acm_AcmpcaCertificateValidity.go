package acm


type AcmpcaCertificateValidity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate#type AcmpcaCertificate#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate#value AcmpcaCertificate#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

