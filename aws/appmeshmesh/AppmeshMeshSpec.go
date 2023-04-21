package appmeshmesh


type AppmeshMeshSpec struct {
	// egress_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/appmesh_mesh#egress_filter AppmeshMesh#egress_filter}
	EgressFilter *AppmeshMeshSpecEgressFilter `field:"optional" json:"egressFilter" yaml:"egressFilter"`
}

