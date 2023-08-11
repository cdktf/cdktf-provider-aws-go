package appmeshroute


type AppmeshRouteSpecTcpRouteTimeout struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/appmesh_route#idle AppmeshRoute#idle}
	Idle *AppmeshRouteSpecTcpRouteTimeoutIdle `field:"optional" json:"idle" yaml:"idle"`
}

