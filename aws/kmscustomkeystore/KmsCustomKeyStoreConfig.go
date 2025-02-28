// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmscustomkeystore

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KmsCustomKeyStoreConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/kms_custom_key_store#custom_key_store_name KmsCustomKeyStore#custom_key_store_name}.
	CustomKeyStoreName *string `field:"required" json:"customKeyStoreName" yaml:"customKeyStoreName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/kms_custom_key_store#cloud_hsm_cluster_id KmsCustomKeyStore#cloud_hsm_cluster_id}.
	CloudHsmClusterId *string `field:"optional" json:"cloudHsmClusterId" yaml:"cloudHsmClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/kms_custom_key_store#custom_key_store_type KmsCustomKeyStore#custom_key_store_type}.
	CustomKeyStoreType *string `field:"optional" json:"customKeyStoreType" yaml:"customKeyStoreType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/kms_custom_key_store#id KmsCustomKeyStore#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/kms_custom_key_store#key_store_password KmsCustomKeyStore#key_store_password}.
	KeyStorePassword *string `field:"optional" json:"keyStorePassword" yaml:"keyStorePassword"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/kms_custom_key_store#timeouts KmsCustomKeyStore#timeouts}
	Timeouts *KmsCustomKeyStoreTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/kms_custom_key_store#trust_anchor_certificate KmsCustomKeyStore#trust_anchor_certificate}.
	TrustAnchorCertificate *string `field:"optional" json:"trustAnchorCertificate" yaml:"trustAnchorCertificate"`
	// xks_proxy_authentication_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/kms_custom_key_store#xks_proxy_authentication_credential KmsCustomKeyStore#xks_proxy_authentication_credential}
	XksProxyAuthenticationCredential *KmsCustomKeyStoreXksProxyAuthenticationCredential `field:"optional" json:"xksProxyAuthenticationCredential" yaml:"xksProxyAuthenticationCredential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/kms_custom_key_store#xks_proxy_connectivity KmsCustomKeyStore#xks_proxy_connectivity}.
	XksProxyConnectivity *string `field:"optional" json:"xksProxyConnectivity" yaml:"xksProxyConnectivity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/kms_custom_key_store#xks_proxy_uri_endpoint KmsCustomKeyStore#xks_proxy_uri_endpoint}.
	XksProxyUriEndpoint *string `field:"optional" json:"xksProxyUriEndpoint" yaml:"xksProxyUriEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/kms_custom_key_store#xks_proxy_uri_path KmsCustomKeyStore#xks_proxy_uri_path}.
	XksProxyUriPath *string `field:"optional" json:"xksProxyUriPath" yaml:"xksProxyUriPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/kms_custom_key_store#xks_proxy_vpc_endpoint_service_name KmsCustomKeyStore#xks_proxy_vpc_endpoint_service_name}.
	XksProxyVpcEndpointServiceName *string `field:"optional" json:"xksProxyVpcEndpointServiceName" yaml:"xksProxyVpcEndpointServiceName"`
}

