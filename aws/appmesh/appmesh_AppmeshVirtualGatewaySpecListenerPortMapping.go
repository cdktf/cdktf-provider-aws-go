package appmesh


type AppmeshVirtualGatewaySpecListenerPortMapping struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#port AppmeshVirtualGateway#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#protocol AppmeshVirtualGateway#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

