package appmesh


type AppmeshGatewayRouteSpecHttpRouteMatchHostname struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#exact AppmeshGatewayRoute#exact}.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#suffix AppmeshGatewayRoute#suffix}.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

