package redshiftparametergroup


type RedshiftParameterGroupParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/redshift_parameter_group#name RedshiftParameterGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/redshift_parameter_group#value RedshiftParameterGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

