package appmeshroute


type AppmeshRouteSpecHttpRouteMatchQueryParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/appmesh_route#name AppmeshRoute#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/appmesh_route#match AppmeshRoute#match}
	Match *AppmeshRouteSpecHttpRouteMatchQueryParameterMatch `field:"optional" json:"match" yaml:"match"`
}

