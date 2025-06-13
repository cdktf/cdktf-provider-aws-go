// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceParameters struct {
	// amazon_elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#amazon_elasticsearch QuicksightDataSource#amazon_elasticsearch}
	AmazonElasticsearch *QuicksightDataSourceParametersAmazonElasticsearch `field:"optional" json:"amazonElasticsearch" yaml:"amazonElasticsearch"`
	// athena block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#athena QuicksightDataSource#athena}
	Athena *QuicksightDataSourceParametersAthena `field:"optional" json:"athena" yaml:"athena"`
	// aurora block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#aurora QuicksightDataSource#aurora}
	Aurora *QuicksightDataSourceParametersAurora `field:"optional" json:"aurora" yaml:"aurora"`
	// aurora_postgresql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#aurora_postgresql QuicksightDataSource#aurora_postgresql}
	AuroraPostgresql *QuicksightDataSourceParametersAuroraPostgresql `field:"optional" json:"auroraPostgresql" yaml:"auroraPostgresql"`
	// aws_iot_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#aws_iot_analytics QuicksightDataSource#aws_iot_analytics}
	AwsIotAnalytics *QuicksightDataSourceParametersAwsIotAnalytics `field:"optional" json:"awsIotAnalytics" yaml:"awsIotAnalytics"`
	// databricks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#databricks QuicksightDataSource#databricks}
	Databricks *QuicksightDataSourceParametersDatabricks `field:"optional" json:"databricks" yaml:"databricks"`
	// jira block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#jira QuicksightDataSource#jira}
	Jira *QuicksightDataSourceParametersJira `field:"optional" json:"jira" yaml:"jira"`
	// maria_db block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#maria_db QuicksightDataSource#maria_db}
	MariaDb *QuicksightDataSourceParametersMariaDb `field:"optional" json:"mariaDb" yaml:"mariaDb"`
	// mysql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#mysql QuicksightDataSource#mysql}
	Mysql *QuicksightDataSourceParametersMysql `field:"optional" json:"mysql" yaml:"mysql"`
	// oracle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#oracle QuicksightDataSource#oracle}
	Oracle *QuicksightDataSourceParametersOracle `field:"optional" json:"oracle" yaml:"oracle"`
	// postgresql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#postgresql QuicksightDataSource#postgresql}
	Postgresql *QuicksightDataSourceParametersPostgresql `field:"optional" json:"postgresql" yaml:"postgresql"`
	// presto block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#presto QuicksightDataSource#presto}
	Presto *QuicksightDataSourceParametersPresto `field:"optional" json:"presto" yaml:"presto"`
	// rds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#rds QuicksightDataSource#rds}
	Rds *QuicksightDataSourceParametersRds `field:"optional" json:"rds" yaml:"rds"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#redshift QuicksightDataSource#redshift}
	Redshift *QuicksightDataSourceParametersRedshift `field:"optional" json:"redshift" yaml:"redshift"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#s3 QuicksightDataSource#s3}
	S3 *QuicksightDataSourceParametersS3 `field:"optional" json:"s3" yaml:"s3"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#service_now QuicksightDataSource#service_now}
	ServiceNow *QuicksightDataSourceParametersServiceNow `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// snowflake block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#snowflake QuicksightDataSource#snowflake}
	Snowflake *QuicksightDataSourceParametersSnowflake `field:"optional" json:"snowflake" yaml:"snowflake"`
	// spark block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#spark QuicksightDataSource#spark}
	Spark *QuicksightDataSourceParametersSpark `field:"optional" json:"spark" yaml:"spark"`
	// sql_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#sql_server QuicksightDataSource#sql_server}
	SqlServer *QuicksightDataSourceParametersSqlServer `field:"optional" json:"sqlServer" yaml:"sqlServer"`
	// teradata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#teradata QuicksightDataSource#teradata}
	Teradata *QuicksightDataSourceParametersTeradata `field:"optional" json:"teradata" yaml:"teradata"`
	// twitter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/quicksight_data_source#twitter QuicksightDataSource#twitter}
	Twitter *QuicksightDataSourceParametersTwitter `field:"optional" json:"twitter" yaml:"twitter"`
}

