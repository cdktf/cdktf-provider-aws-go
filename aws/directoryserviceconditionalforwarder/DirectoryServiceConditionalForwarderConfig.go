// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryserviceconditionalforwarder

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DirectoryServiceConditionalForwarderConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/directory_service_conditional_forwarder#directory_id DirectoryServiceConditionalForwarder#directory_id}.
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/directory_service_conditional_forwarder#dns_ips DirectoryServiceConditionalForwarder#dns_ips}.
	DnsIps *[]*string `field:"required" json:"dnsIps" yaml:"dnsIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/directory_service_conditional_forwarder#remote_domain_name DirectoryServiceConditionalForwarder#remote_domain_name}.
	RemoteDomainName *string `field:"required" json:"remoteDomainName" yaml:"remoteDomainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/directory_service_conditional_forwarder#id DirectoryServiceConditionalForwarder#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

