// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package neptunegraphgraph

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NeptunegraphGraphConfig struct {
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
	// The provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use for the graph.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/neptunegraph_graph#provisioned_memory NeptunegraphGraph#provisioned_memory}
	ProvisionedMemory *float64 `field:"required" json:"provisionedMemory" yaml:"provisionedMemory"`
	// A value that indicates whether the graph has deletion protection enabled.
	//
	// The graph can't be deleted when deletion protection is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/neptunegraph_graph#deletion_protection NeptunegraphGraph#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The graph name.
	//
	// For example: my-graph-1.
	// 								The name must contain from 1 to 63 letters, numbers, or hyphens,
	// 								and its first character must be a letter. It cannot end with a hyphen or contain two consecutive hyphens.
	// 								If you don't specify a graph name, a unique graph name is generated for you using the prefix graph-for,
	// 								followed by a combination of Stack Name and a UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/neptunegraph_graph#graph_name NeptunegraphGraph#graph_name}
	GraphName *string `field:"optional" json:"graphName" yaml:"graphName"`
	// Allows user to specify name prefix and have remainder of name automatically generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/neptunegraph_graph#graph_name_prefix NeptunegraphGraph#graph_name_prefix}
	GraphNamePrefix *string `field:"optional" json:"graphNamePrefix" yaml:"graphNamePrefix"`
	// Specifies a KMS key to use to encrypt data in the new graph.
	//
	// Value must be ARN of KMS Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/neptunegraph_graph#kms_key_identifier NeptunegraphGraph#kms_key_identifier}
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// Specifies whether or not the graph can be reachable over the internet.
	//
	// All access to graphs is IAM authenticated.
	// 								When the graph is publicly available, its domain name system (DNS) endpoint resolves to
	// 								the public IP address from the internet. When the graph isn't publicly available, you need
	// 								to create a PrivateGraphEndpoint in a given VPC to ensure the DNS name resolves to a private
	// 								IP address that is reachable from the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/neptunegraph_graph#public_connectivity NeptunegraphGraph#public_connectivity}
	PublicConnectivity interface{} `field:"optional" json:"publicConnectivity" yaml:"publicConnectivity"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/neptunegraph_graph#region NeptunegraphGraph#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The number of replicas in other AZs.  Value must be between 0 and 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/neptunegraph_graph#replica_count NeptunegraphGraph#replica_count}
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/neptunegraph_graph#tags NeptunegraphGraph#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/neptunegraph_graph#timeouts NeptunegraphGraph#timeouts}
	Timeouts *NeptunegraphGraphTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vector_search_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/neptunegraph_graph#vector_search_configuration NeptunegraphGraph#vector_search_configuration}
	VectorSearchConfiguration interface{} `field:"optional" json:"vectorSearchConfiguration" yaml:"vectorSearchConfiguration"`
}

