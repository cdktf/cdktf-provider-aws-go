package appmeshroute


type AppmeshRouteSpecTcpRouteActionWeightedTarget struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#virtual_node AppmeshRoute#virtual_node}.
	VirtualNode *string `field:"required" json:"virtualNode" yaml:"virtualNode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#weight AppmeshRoute#weight}.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#port AppmeshRoute#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

