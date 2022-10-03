package quicksightdatasource


type QuicksightDataSourceVpcConnectionProperties struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source#vpc_connection_arn QuicksightDataSource#vpc_connection_arn}.
	VpcConnectionArn *string `field:"required" json:"vpcConnectionArn" yaml:"vpcConnectionArn"`
}

