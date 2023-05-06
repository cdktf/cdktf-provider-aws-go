package appmeshvirtualrouter


type AppmeshVirtualRouterSpecListenerPortMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/appmesh_virtual_router#port AppmeshVirtualRouter#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/appmesh_virtual_router#protocol AppmeshVirtualRouter#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

