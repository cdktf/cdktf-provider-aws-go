package appmeshvirtualnode


type AppmeshVirtualNodeSpecBackendVirtualService struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#virtual_service_name AppmeshVirtualNode#virtual_service_name}.
	VirtualServiceName *string `field:"required" json:"virtualServiceName" yaml:"virtualServiceName"`
	// client_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#client_policy AppmeshVirtualNode#client_policy}
	ClientPolicy *AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicy `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

