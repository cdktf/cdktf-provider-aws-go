package appmeshroute


type AppmeshRouteSpecHttpRouteMatchHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appmesh_route#name AppmeshRoute#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appmesh_route#invert AppmeshRoute#invert}.
	Invert interface{} `field:"optional" json:"invert" yaml:"invert"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appmesh_route#match AppmeshRoute#match}
	Match *AppmeshRouteSpecHttpRouteMatchHeaderMatch `field:"optional" json:"match" yaml:"match"`
}

