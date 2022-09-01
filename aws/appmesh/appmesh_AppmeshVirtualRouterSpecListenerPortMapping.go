package appmesh


type AppmeshVirtualRouterSpecListenerPortMapping struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_router#port AppmeshVirtualRouter#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_router#protocol AppmeshVirtualRouter#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

