package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttp2RouteMatchHostname struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appmesh_gateway_route#exact AppmeshGatewayRoute#exact}.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appmesh_gateway_route#suffix AppmeshGatewayRoute#suffix}.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

