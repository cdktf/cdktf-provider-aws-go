package appmeshroute


type AppmeshRouteSpecHttp2RouteAction struct {
	// weighted_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/appmesh_route#weighted_target AppmeshRoute#weighted_target}
	WeightedTarget interface{} `field:"required" json:"weightedTarget" yaml:"weightedTarget"`
}

