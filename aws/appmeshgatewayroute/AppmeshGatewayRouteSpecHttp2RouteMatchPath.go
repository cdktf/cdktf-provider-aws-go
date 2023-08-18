package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttp2RouteMatchPath struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/appmesh_gateway_route#exact AppmeshGatewayRoute#exact}.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/appmesh_gateway_route#regex AppmeshGatewayRoute#regex}.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}

