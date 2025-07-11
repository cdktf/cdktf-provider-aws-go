// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package neptunegraphgraph


type NeptunegraphGraphVectorSearchConfiguration struct {
	// Specifies the number of dimensions for vector embeddings.  Value must be between 1 and 65,535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/neptunegraph_graph#vector_search_dimension NeptunegraphGraph#vector_search_dimension}
	VectorSearchDimension *float64 `field:"optional" json:"vectorSearchDimension" yaml:"vectorSearchDimension"`
}

