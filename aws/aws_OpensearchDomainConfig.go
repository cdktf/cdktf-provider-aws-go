// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpensearchDomainConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#domain_name OpensearchDomain#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#access_policies OpensearchDomain#access_policies}.
	AccessPolicies *string `field:"optional" json:"accessPolicies" yaml:"accessPolicies"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#advanced_options OpensearchDomain#advanced_options}.
	AdvancedOptions *map[string]*string `field:"optional" json:"advancedOptions" yaml:"advancedOptions"`
	// advanced_security_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#advanced_security_options OpensearchDomain#advanced_security_options}
	AdvancedSecurityOptions *OpensearchDomainAdvancedSecurityOptions `field:"optional" json:"advancedSecurityOptions" yaml:"advancedSecurityOptions"`
	// auto_tune_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#auto_tune_options OpensearchDomain#auto_tune_options}
	AutoTuneOptions *OpensearchDomainAutoTuneOptions `field:"optional" json:"autoTuneOptions" yaml:"autoTuneOptions"`
	// cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#cluster_config OpensearchDomain#cluster_config}
	ClusterConfig *OpensearchDomainClusterConfig `field:"optional" json:"clusterConfig" yaml:"clusterConfig"`
	// cognito_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#cognito_options OpensearchDomain#cognito_options}
	CognitoOptions *OpensearchDomainCognitoOptions `field:"optional" json:"cognitoOptions" yaml:"cognitoOptions"`
	// domain_endpoint_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#domain_endpoint_options OpensearchDomain#domain_endpoint_options}
	DomainEndpointOptions *OpensearchDomainDomainEndpointOptions `field:"optional" json:"domainEndpointOptions" yaml:"domainEndpointOptions"`
	// ebs_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#ebs_options OpensearchDomain#ebs_options}
	EbsOptions *OpensearchDomainEbsOptions `field:"optional" json:"ebsOptions" yaml:"ebsOptions"`
	// encrypt_at_rest block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#encrypt_at_rest OpensearchDomain#encrypt_at_rest}
	EncryptAtRest *OpensearchDomainEncryptAtRest `field:"optional" json:"encryptAtRest" yaml:"encryptAtRest"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#engine_version OpensearchDomain#engine_version}.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#id OpensearchDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// log_publishing_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#log_publishing_options OpensearchDomain#log_publishing_options}
	LogPublishingOptions interface{} `field:"optional" json:"logPublishingOptions" yaml:"logPublishingOptions"`
	// node_to_node_encryption block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#node_to_node_encryption OpensearchDomain#node_to_node_encryption}
	NodeToNodeEncryption *OpensearchDomainNodeToNodeEncryption `field:"optional" json:"nodeToNodeEncryption" yaml:"nodeToNodeEncryption"`
	// snapshot_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#snapshot_options OpensearchDomain#snapshot_options}
	SnapshotOptions *OpensearchDomainSnapshotOptions `field:"optional" json:"snapshotOptions" yaml:"snapshotOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#tags OpensearchDomain#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#tags_all OpensearchDomain#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#timeouts OpensearchDomain#timeouts}
	Timeouts *OpensearchDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vpc_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#vpc_options OpensearchDomain#vpc_options}
	VpcOptions *OpensearchDomainVpcOptions `field:"optional" json:"vpcOptions" yaml:"vpcOptions"`
}

