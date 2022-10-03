package quicksightdatasource


type QuicksightDataSourceParametersTwitter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source#max_rows QuicksightDataSource#max_rows}.
	MaxRows *float64 `field:"required" json:"maxRows" yaml:"maxRows"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source#query QuicksightDataSource#query}.
	Query *string `field:"required" json:"query" yaml:"query"`
}

