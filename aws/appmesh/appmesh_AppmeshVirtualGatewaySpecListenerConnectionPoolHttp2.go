package appmesh


type AppmeshVirtualGatewaySpecListenerConnectionPoolHttp2 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#max_requests AppmeshVirtualGateway#max_requests}.
	MaxRequests *float64 `field:"required" json:"maxRequests" yaml:"maxRequests"`
}

