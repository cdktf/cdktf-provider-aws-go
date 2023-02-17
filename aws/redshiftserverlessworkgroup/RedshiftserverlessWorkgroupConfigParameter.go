package redshiftserverlessworkgroup


type RedshiftserverlessWorkgroupConfigParameter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/redshiftserverless_workgroup#parameter_key RedshiftserverlessWorkgroup#parameter_key}.
	ParameterKey *string `field:"required" json:"parameterKey" yaml:"parameterKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/redshiftserverless_workgroup#parameter_value RedshiftserverlessWorkgroup#parameter_value}.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

