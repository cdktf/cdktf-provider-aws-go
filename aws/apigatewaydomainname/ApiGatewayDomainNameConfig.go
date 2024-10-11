// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewaydomainname

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiGatewayDomainNameConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/api_gateway_domain_name#domain_name ApiGatewayDomainName#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/api_gateway_domain_name#certificate_arn ApiGatewayDomainName#certificate_arn}.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/api_gateway_domain_name#certificate_body ApiGatewayDomainName#certificate_body}.
	CertificateBody *string `field:"optional" json:"certificateBody" yaml:"certificateBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/api_gateway_domain_name#certificate_chain ApiGatewayDomainName#certificate_chain}.
	CertificateChain *string `field:"optional" json:"certificateChain" yaml:"certificateChain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/api_gateway_domain_name#certificate_name ApiGatewayDomainName#certificate_name}.
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/api_gateway_domain_name#certificate_private_key ApiGatewayDomainName#certificate_private_key}.
	CertificatePrivateKey *string `field:"optional" json:"certificatePrivateKey" yaml:"certificatePrivateKey"`
	// endpoint_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/api_gateway_domain_name#endpoint_configuration ApiGatewayDomainName#endpoint_configuration}
	EndpointConfiguration *ApiGatewayDomainNameEndpointConfiguration `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/api_gateway_domain_name#id ApiGatewayDomainName#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// mutual_tls_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/api_gateway_domain_name#mutual_tls_authentication ApiGatewayDomainName#mutual_tls_authentication}
	MutualTlsAuthentication *ApiGatewayDomainNameMutualTlsAuthentication `field:"optional" json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/api_gateway_domain_name#ownership_verification_certificate_arn ApiGatewayDomainName#ownership_verification_certificate_arn}.
	OwnershipVerificationCertificateArn *string `field:"optional" json:"ownershipVerificationCertificateArn" yaml:"ownershipVerificationCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/api_gateway_domain_name#regional_certificate_arn ApiGatewayDomainName#regional_certificate_arn}.
	RegionalCertificateArn *string `field:"optional" json:"regionalCertificateArn" yaml:"regionalCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/api_gateway_domain_name#regional_certificate_name ApiGatewayDomainName#regional_certificate_name}.
	RegionalCertificateName *string `field:"optional" json:"regionalCertificateName" yaml:"regionalCertificateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/api_gateway_domain_name#security_policy ApiGatewayDomainName#security_policy}.
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/api_gateway_domain_name#tags ApiGatewayDomainName#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/api_gateway_domain_name#tags_all ApiGatewayDomainName#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

