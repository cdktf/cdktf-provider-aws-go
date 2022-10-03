package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerPortMapping struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#port AppmeshVirtualNode#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#protocol AppmeshVirtualNode#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

