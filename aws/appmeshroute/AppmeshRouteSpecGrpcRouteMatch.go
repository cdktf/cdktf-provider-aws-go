package appmeshroute


type AppmeshRouteSpecGrpcRouteMatch struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/appmesh_route#metadata AppmeshRoute#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/appmesh_route#method_name AppmeshRoute#method_name}.
	MethodName *string `field:"optional" json:"methodName" yaml:"methodName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/appmesh_route#port AppmeshRoute#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/appmesh_route#prefix AppmeshRoute#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/appmesh_route#service_name AppmeshRoute#service_name}.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

