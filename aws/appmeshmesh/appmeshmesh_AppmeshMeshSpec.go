package appmeshmesh


type AppmeshMeshSpec struct {
	// egress_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_mesh#egress_filter AppmeshMesh#egress_filter}
	EgressFilter *AppmeshMeshSpecEgressFilter `field:"optional" json:"egressFilter" yaml:"egressFilter"`
}

