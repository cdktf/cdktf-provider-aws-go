package redshift


type RedshiftdataStatementParameters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/redshiftdata_statement#name RedshiftdataStatement#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/redshiftdata_statement#value RedshiftdataStatement#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

