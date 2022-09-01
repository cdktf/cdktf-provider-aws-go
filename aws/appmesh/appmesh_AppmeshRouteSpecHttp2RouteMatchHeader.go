package appmesh


type AppmeshRouteSpecHttp2RouteMatchHeader struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#name AppmeshRoute#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#invert AppmeshRoute#invert}.
	Invert interface{} `field:"optional" json:"invert" yaml:"invert"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#match AppmeshRoute#match}
	Match *AppmeshRouteSpecHttp2RouteMatchHeaderMatch `field:"optional" json:"match" yaml:"match"`
}

