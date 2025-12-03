// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryservicetrust

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DirectoryServiceTrustConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/directory_service_trust#directory_id DirectoryServiceTrust#directory_id}.
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/directory_service_trust#remote_domain_name DirectoryServiceTrust#remote_domain_name}.
	RemoteDomainName *string `field:"required" json:"remoteDomainName" yaml:"remoteDomainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/directory_service_trust#trust_direction DirectoryServiceTrust#trust_direction}.
	TrustDirection *string `field:"required" json:"trustDirection" yaml:"trustDirection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/directory_service_trust#trust_password DirectoryServiceTrust#trust_password}.
	TrustPassword *string `field:"required" json:"trustPassword" yaml:"trustPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/directory_service_trust#conditional_forwarder_ip_addrs DirectoryServiceTrust#conditional_forwarder_ip_addrs}.
	ConditionalForwarderIpAddrs *[]*string `field:"optional" json:"conditionalForwarderIpAddrs" yaml:"conditionalForwarderIpAddrs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/directory_service_trust#delete_associated_conditional_forwarder DirectoryServiceTrust#delete_associated_conditional_forwarder}.
	DeleteAssociatedConditionalForwarder interface{} `field:"optional" json:"deleteAssociatedConditionalForwarder" yaml:"deleteAssociatedConditionalForwarder"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/directory_service_trust#region DirectoryServiceTrust#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/directory_service_trust#selective_auth DirectoryServiceTrust#selective_auth}.
	SelectiveAuth *string `field:"optional" json:"selectiveAuth" yaml:"selectiveAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/directory_service_trust#trust_type DirectoryServiceTrust#trust_type}.
	TrustType *string `field:"optional" json:"trustType" yaml:"trustType"`
}

