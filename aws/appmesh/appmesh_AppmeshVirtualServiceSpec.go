package appmesh


type AppmeshVirtualServiceSpec struct {
	// provider block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_service#provider AppmeshVirtualService#provider}
	Provider *AppmeshVirtualServiceSpecProvider `field:"optional" json:"provider" yaml:"provider"`
}

