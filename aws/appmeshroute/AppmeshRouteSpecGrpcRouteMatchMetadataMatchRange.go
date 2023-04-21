package appmeshroute


type AppmeshRouteSpecGrpcRouteMatchMetadataMatchRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/appmesh_route#end AppmeshRoute#end}.
	End *float64 `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/appmesh_route#start AppmeshRoute#start}.
	Start *float64 `field:"required" json:"start" yaml:"start"`
}

