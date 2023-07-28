package appmeshroute


type AppmeshRouteSpecTcpRouteAction struct {
	// weighted_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/appmesh_route#weighted_target AppmeshRoute#weighted_target}
	WeightedTarget interface{} `field:"required" json:"weightedTarget" yaml:"weightedTarget"`
}

