package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecListenerPortMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appmesh_virtual_gateway#port AppmeshVirtualGateway#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appmesh_virtual_gateway#protocol AppmeshVirtualGateway#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

