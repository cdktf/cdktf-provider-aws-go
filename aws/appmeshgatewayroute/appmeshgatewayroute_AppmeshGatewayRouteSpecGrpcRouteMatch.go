package appmeshgatewayroute


type AppmeshGatewayRouteSpecGrpcRouteMatch struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#service_name AppmeshGatewayRoute#service_name}.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
}

