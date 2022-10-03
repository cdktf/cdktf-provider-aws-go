package appmeshroute


type AppmeshRouteSpecHttp2RouteMatchHeaderMatchRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#end AppmeshRoute#end}.
	End *float64 `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#start AppmeshRoute#start}.
	Start *float64 `field:"required" json:"start" yaml:"start"`
}

