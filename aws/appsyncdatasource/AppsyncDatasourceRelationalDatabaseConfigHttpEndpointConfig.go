package appsyncdatasource


type AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appsync_datasource#aws_secret_store_arn AppsyncDatasource#aws_secret_store_arn}.
	AwsSecretStoreArn *string `field:"required" json:"awsSecretStoreArn" yaml:"awsSecretStoreArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appsync_datasource#db_cluster_identifier AppsyncDatasource#db_cluster_identifier}.
	DbClusterIdentifier *string `field:"required" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appsync_datasource#database_name AppsyncDatasource#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appsync_datasource#region AppsyncDatasource#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appsync_datasource#schema AppsyncDatasource#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

