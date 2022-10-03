package appmeshroute


type AppmeshRouteSpecHttpRouteTimeout struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#idle AppmeshRoute#idle}
	Idle *AppmeshRouteSpecHttpRouteTimeoutIdle `field:"optional" json:"idle" yaml:"idle"`
	// per_request block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#per_request AppmeshRoute#per_request}
	PerRequest *AppmeshRouteSpecHttpRouteTimeoutPerRequest `field:"optional" json:"perRequest" yaml:"perRequest"`
}

