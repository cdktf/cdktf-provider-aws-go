package appmesh


type AppmeshGatewayRouteSpecHttp2RouteActionTargetVirtualService struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#virtual_service_name AppmeshGatewayRoute#virtual_service_name}.
	VirtualServiceName *string `field:"required" json:"virtualServiceName" yaml:"virtualServiceName"`
}

