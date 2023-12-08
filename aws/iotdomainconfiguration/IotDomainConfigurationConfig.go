// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotdomainconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotDomainConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/iot_domain_configuration#name IotDomainConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// authorizer_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/iot_domain_configuration#authorizer_config IotDomainConfiguration#authorizer_config}
	AuthorizerConfig *IotDomainConfigurationAuthorizerConfig `field:"optional" json:"authorizerConfig" yaml:"authorizerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/iot_domain_configuration#domain_name IotDomainConfiguration#domain_name}.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/iot_domain_configuration#id IotDomainConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/iot_domain_configuration#server_certificate_arns IotDomainConfiguration#server_certificate_arns}.
	ServerCertificateArns *[]*string `field:"optional" json:"serverCertificateArns" yaml:"serverCertificateArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/iot_domain_configuration#service_type IotDomainConfiguration#service_type}.
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/iot_domain_configuration#status IotDomainConfiguration#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/iot_domain_configuration#tags IotDomainConfiguration#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/iot_domain_configuration#tags_all IotDomainConfiguration#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// tls_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/iot_domain_configuration#tls_config IotDomainConfiguration#tls_config}
	TlsConfig *IotDomainConfigurationTlsConfig `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/iot_domain_configuration#validation_certificate_arn IotDomainConfiguration#validation_certificate_arn}.
	ValidationCertificateArn *string `field:"optional" json:"validationCertificateArn" yaml:"validationCertificateArn"`
}

