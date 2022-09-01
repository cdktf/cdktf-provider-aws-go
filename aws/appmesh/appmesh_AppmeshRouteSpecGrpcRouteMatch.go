package appmesh


type AppmeshRouteSpecGrpcRouteMatch struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#metadata AppmeshRoute#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#method_name AppmeshRoute#method_name}.
	MethodName *string `field:"optional" json:"methodName" yaml:"methodName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#prefix AppmeshRoute#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#service_name AppmeshRoute#service_name}.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

