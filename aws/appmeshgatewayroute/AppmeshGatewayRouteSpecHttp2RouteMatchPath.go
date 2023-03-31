package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttp2RouteMatchPath struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#exact AppmeshGatewayRoute#exact}.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#regex AppmeshGatewayRoute#regex}.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}

