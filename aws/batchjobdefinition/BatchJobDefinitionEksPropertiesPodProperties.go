// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionEksPropertiesPodProperties struct {
	// containers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/batch_job_definition#containers BatchJobDefinition#containers}
	Containers *BatchJobDefinitionEksPropertiesPodPropertiesContainers `field:"required" json:"containers" yaml:"containers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/batch_job_definition#dns_policy BatchJobDefinition#dns_policy}.
	DnsPolicy *string `field:"optional" json:"dnsPolicy" yaml:"dnsPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/batch_job_definition#host_network BatchJobDefinition#host_network}.
	HostNetwork interface{} `field:"optional" json:"hostNetwork" yaml:"hostNetwork"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/batch_job_definition#metadata BatchJobDefinition#metadata}
	Metadata *BatchJobDefinitionEksPropertiesPodPropertiesMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/batch_job_definition#service_account_name BatchJobDefinition#service_account_name}.
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/batch_job_definition#volumes BatchJobDefinition#volumes}
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

