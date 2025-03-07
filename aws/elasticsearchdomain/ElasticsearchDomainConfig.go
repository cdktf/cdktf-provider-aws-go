// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticsearchdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElasticsearchDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#domain_name ElasticsearchDomain#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#access_policies ElasticsearchDomain#access_policies}.
	AccessPolicies *string `field:"optional" json:"accessPolicies" yaml:"accessPolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#advanced_options ElasticsearchDomain#advanced_options}.
	AdvancedOptions *map[string]*string `field:"optional" json:"advancedOptions" yaml:"advancedOptions"`
	// advanced_security_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#advanced_security_options ElasticsearchDomain#advanced_security_options}
	AdvancedSecurityOptions *ElasticsearchDomainAdvancedSecurityOptions `field:"optional" json:"advancedSecurityOptions" yaml:"advancedSecurityOptions"`
	// auto_tune_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#auto_tune_options ElasticsearchDomain#auto_tune_options}
	AutoTuneOptions *ElasticsearchDomainAutoTuneOptions `field:"optional" json:"autoTuneOptions" yaml:"autoTuneOptions"`
	// cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#cluster_config ElasticsearchDomain#cluster_config}
	ClusterConfig *ElasticsearchDomainClusterConfig `field:"optional" json:"clusterConfig" yaml:"clusterConfig"`
	// cognito_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#cognito_options ElasticsearchDomain#cognito_options}
	CognitoOptions *ElasticsearchDomainCognitoOptions `field:"optional" json:"cognitoOptions" yaml:"cognitoOptions"`
	// domain_endpoint_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#domain_endpoint_options ElasticsearchDomain#domain_endpoint_options}
	DomainEndpointOptions *ElasticsearchDomainDomainEndpointOptions `field:"optional" json:"domainEndpointOptions" yaml:"domainEndpointOptions"`
	// ebs_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#ebs_options ElasticsearchDomain#ebs_options}
	EbsOptions *ElasticsearchDomainEbsOptions `field:"optional" json:"ebsOptions" yaml:"ebsOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#elasticsearch_version ElasticsearchDomain#elasticsearch_version}.
	ElasticsearchVersion *string `field:"optional" json:"elasticsearchVersion" yaml:"elasticsearchVersion"`
	// encrypt_at_rest block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#encrypt_at_rest ElasticsearchDomain#encrypt_at_rest}
	EncryptAtRest *ElasticsearchDomainEncryptAtRest `field:"optional" json:"encryptAtRest" yaml:"encryptAtRest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#id ElasticsearchDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// log_publishing_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#log_publishing_options ElasticsearchDomain#log_publishing_options}
	LogPublishingOptions interface{} `field:"optional" json:"logPublishingOptions" yaml:"logPublishingOptions"`
	// node_to_node_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#node_to_node_encryption ElasticsearchDomain#node_to_node_encryption}
	NodeToNodeEncryption *ElasticsearchDomainNodeToNodeEncryption `field:"optional" json:"nodeToNodeEncryption" yaml:"nodeToNodeEncryption"`
	// snapshot_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#snapshot_options ElasticsearchDomain#snapshot_options}
	SnapshotOptions *ElasticsearchDomainSnapshotOptions `field:"optional" json:"snapshotOptions" yaml:"snapshotOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#tags ElasticsearchDomain#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#tags_all ElasticsearchDomain#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#timeouts ElasticsearchDomain#timeouts}
	Timeouts *ElasticsearchDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vpc_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticsearch_domain#vpc_options ElasticsearchDomain#vpc_options}
	VpcOptions *ElasticsearchDomainVpcOptions `field:"optional" json:"vpcOptions" yaml:"vpcOptions"`
}

