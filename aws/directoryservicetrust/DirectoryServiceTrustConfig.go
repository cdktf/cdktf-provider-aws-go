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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/directory_service_trust#directory_id DirectoryServiceTrust#directory_id}.
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/directory_service_trust#remote_domain_name DirectoryServiceTrust#remote_domain_name}.
	RemoteDomainName *string `field:"required" json:"remoteDomainName" yaml:"remoteDomainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/directory_service_trust#trust_direction DirectoryServiceTrust#trust_direction}.
	TrustDirection *string `field:"required" json:"trustDirection" yaml:"trustDirection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/directory_service_trust#trust_password DirectoryServiceTrust#trust_password}.
	TrustPassword *string `field:"required" json:"trustPassword" yaml:"trustPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/directory_service_trust#conditional_forwarder_ip_addrs DirectoryServiceTrust#conditional_forwarder_ip_addrs}.
	ConditionalForwarderIpAddrs *[]*string `field:"optional" json:"conditionalForwarderIpAddrs" yaml:"conditionalForwarderIpAddrs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/directory_service_trust#delete_associated_conditional_forwarder DirectoryServiceTrust#delete_associated_conditional_forwarder}.
	DeleteAssociatedConditionalForwarder interface{} `field:"optional" json:"deleteAssociatedConditionalForwarder" yaml:"deleteAssociatedConditionalForwarder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/directory_service_trust#selective_auth DirectoryServiceTrust#selective_auth}.
	SelectiveAuth *string `field:"optional" json:"selectiveAuth" yaml:"selectiveAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/directory_service_trust#trust_type DirectoryServiceTrust#trust_type}.
	TrustType *string `field:"optional" json:"trustType" yaml:"trustType"`
}

