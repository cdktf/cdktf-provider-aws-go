package appmesh


type AppmeshRouteSpecTcpRouteTimeout struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#idle AppmeshRoute#idle}
	Idle *AppmeshRouteSpecTcpRouteTimeoutIdle `field:"optional" json:"idle" yaml:"idle"`
}

