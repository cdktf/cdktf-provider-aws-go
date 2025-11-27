// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchoutboundconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpensearchOutboundConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/opensearch_outbound_connection#connection_alias OpensearchOutboundConnection#connection_alias}.
	ConnectionAlias *string `field:"required" json:"connectionAlias" yaml:"connectionAlias"`
	// local_domain_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/opensearch_outbound_connection#local_domain_info OpensearchOutboundConnection#local_domain_info}
	LocalDomainInfo *OpensearchOutboundConnectionLocalDomainInfo `field:"required" json:"localDomainInfo" yaml:"localDomainInfo"`
	// remote_domain_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/opensearch_outbound_connection#remote_domain_info OpensearchOutboundConnection#remote_domain_info}
	RemoteDomainInfo *OpensearchOutboundConnectionRemoteDomainInfo `field:"required" json:"remoteDomainInfo" yaml:"remoteDomainInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/opensearch_outbound_connection#accept_connection OpensearchOutboundConnection#accept_connection}.
	AcceptConnection interface{} `field:"optional" json:"acceptConnection" yaml:"acceptConnection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/opensearch_outbound_connection#connection_mode OpensearchOutboundConnection#connection_mode}.
	ConnectionMode *string `field:"optional" json:"connectionMode" yaml:"connectionMode"`
	// connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/opensearch_outbound_connection#connection_properties OpensearchOutboundConnection#connection_properties}
	ConnectionProperties *OpensearchOutboundConnectionConnectionProperties `field:"optional" json:"connectionProperties" yaml:"connectionProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/opensearch_outbound_connection#id OpensearchOutboundConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/opensearch_outbound_connection#region OpensearchOutboundConnection#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/opensearch_outbound_connection#timeouts OpensearchOutboundConnection#timeouts}
	Timeouts *OpensearchOutboundConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

