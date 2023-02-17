package appmeshroute


type AppmeshRouteSpecTcpRoute struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#action AppmeshRoute#action}
	Action *AppmeshRouteSpecTcpRouteAction `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#match AppmeshRoute#match}
	Match *AppmeshRouteSpecTcpRouteMatch `field:"optional" json:"match" yaml:"match"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#timeout AppmeshRoute#timeout}
	Timeout *AppmeshRouteSpecTcpRouteTimeout `field:"optional" json:"timeout" yaml:"timeout"`
}

