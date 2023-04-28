package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttp2RouteMatchHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appmesh_gateway_route#name AppmeshGatewayRoute#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appmesh_gateway_route#invert AppmeshGatewayRoute#invert}.
	Invert interface{} `field:"optional" json:"invert" yaml:"invert"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appmesh_gateway_route#match AppmeshGatewayRoute#match}
	Match *AppmeshGatewayRouteSpecHttp2RouteMatchHeaderMatch `field:"optional" json:"match" yaml:"match"`
}

