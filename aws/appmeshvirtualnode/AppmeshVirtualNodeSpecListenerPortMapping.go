package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerPortMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/appmesh_virtual_node#port AppmeshVirtualNode#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/appmesh_virtual_node#protocol AppmeshVirtualNode#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

