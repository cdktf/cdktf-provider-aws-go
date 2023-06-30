package elasticacheparametergroup


type ElasticacheParameterGroupParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/elasticache_parameter_group#name ElasticacheParameterGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/elasticache_parameter_group#value ElasticacheParameterGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

