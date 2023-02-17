package appmeshroute


type AppmeshRouteSpecHttpRouteTimeoutIdle struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#unit AppmeshRoute#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#value AppmeshRoute#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

