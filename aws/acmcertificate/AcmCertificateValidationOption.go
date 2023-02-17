package acmcertificate


type AcmCertificateValidationOption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate#domain_name AcmCertificate#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate#validation_domain AcmCertificate#validation_domain}.
	ValidationDomain *string `field:"required" json:"validationDomain" yaml:"validationDomain"`
}

