// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotcacertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotCaCertificateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/iot_ca_certificate#active IotCaCertificate#active}.
	Active interface{} `field:"required" json:"active" yaml:"active"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/iot_ca_certificate#allow_auto_registration IotCaCertificate#allow_auto_registration}.
	AllowAutoRegistration interface{} `field:"required" json:"allowAutoRegistration" yaml:"allowAutoRegistration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/iot_ca_certificate#ca_certificate_pem IotCaCertificate#ca_certificate_pem}.
	CaCertificatePem *string `field:"required" json:"caCertificatePem" yaml:"caCertificatePem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/iot_ca_certificate#certificate_mode IotCaCertificate#certificate_mode}.
	CertificateMode *string `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/iot_ca_certificate#id IotCaCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/iot_ca_certificate#region IotCaCertificate#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// registration_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/iot_ca_certificate#registration_config IotCaCertificate#registration_config}
	RegistrationConfig *IotCaCertificateRegistrationConfig `field:"optional" json:"registrationConfig" yaml:"registrationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/iot_ca_certificate#tags IotCaCertificate#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/iot_ca_certificate#tags_all IotCaCertificate#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/iot_ca_certificate#verification_certificate_pem IotCaCertificate#verification_certificate_pem}.
	VerificationCertificatePem *string `field:"optional" json:"verificationCertificatePem" yaml:"verificationCertificatePem"`
}

