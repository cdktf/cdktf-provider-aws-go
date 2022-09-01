package appmesh


type AppmeshGatewayRouteSpecHttpRouteActionRewritePrefix struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#default_prefix AppmeshGatewayRoute#default_prefix}.
	DefaultPrefix *string `field:"optional" json:"defaultPrefix" yaml:"defaultPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#value AppmeshGatewayRoute#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

