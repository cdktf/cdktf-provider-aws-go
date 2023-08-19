package acmcertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AcmCertificateConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/acm_certificate#certificate_authority_arn AcmCertificate#certificate_authority_arn}.
	CertificateAuthorityArn *string `field:"optional" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/acm_certificate#certificate_body AcmCertificate#certificate_body}.
	CertificateBody *string `field:"optional" json:"certificateBody" yaml:"certificateBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/acm_certificate#certificate_chain AcmCertificate#certificate_chain}.
	CertificateChain *string `field:"optional" json:"certificateChain" yaml:"certificateChain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/acm_certificate#domain_name AcmCertificate#domain_name}.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/acm_certificate#early_renewal_duration AcmCertificate#early_renewal_duration}.
	EarlyRenewalDuration *string `field:"optional" json:"earlyRenewalDuration" yaml:"earlyRenewalDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/acm_certificate#id AcmCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/acm_certificate#key_algorithm AcmCertificate#key_algorithm}.
	KeyAlgorithm *string `field:"optional" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/acm_certificate#options AcmCertificate#options}
	Options *AcmCertificateOptions `field:"optional" json:"options" yaml:"options"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/acm_certificate#private_key AcmCertificate#private_key}.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/acm_certificate#subject_alternative_names AcmCertificate#subject_alternative_names}.
	SubjectAlternativeNames *[]*string `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/acm_certificate#tags AcmCertificate#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/acm_certificate#tags_all AcmCertificate#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/acm_certificate#validation_method AcmCertificate#validation_method}.
	ValidationMethod *string `field:"optional" json:"validationMethod" yaml:"validationMethod"`
	// validation_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/acm_certificate#validation_option AcmCertificate#validation_option}
	ValidationOption interface{} `field:"optional" json:"validationOption" yaml:"validationOption"`
}

