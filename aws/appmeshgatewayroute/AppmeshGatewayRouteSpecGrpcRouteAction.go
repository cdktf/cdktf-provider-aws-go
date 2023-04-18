package appmeshgatewayroute


type AppmeshGatewayRouteSpecGrpcRouteAction struct {
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/appmesh_gateway_route#target AppmeshGatewayRoute#target}
	Target *AppmeshGatewayRouteSpecGrpcRouteActionTarget `field:"required" json:"target" yaml:"target"`
}

