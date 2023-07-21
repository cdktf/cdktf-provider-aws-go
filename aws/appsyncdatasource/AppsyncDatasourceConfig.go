package appsyncdatasource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsyncDatasourceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/appsync_datasource#api_id AppsyncDatasource#api_id}.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/appsync_datasource#name AppsyncDatasource#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/appsync_datasource#type AppsyncDatasource#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/appsync_datasource#description AppsyncDatasource#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dynamodb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/appsync_datasource#dynamodb_config AppsyncDatasource#dynamodb_config}
	DynamodbConfig *AppsyncDatasourceDynamodbConfig `field:"optional" json:"dynamodbConfig" yaml:"dynamodbConfig"`
	// elasticsearch_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/appsync_datasource#elasticsearch_config AppsyncDatasource#elasticsearch_config}
	ElasticsearchConfig *AppsyncDatasourceElasticsearchConfig `field:"optional" json:"elasticsearchConfig" yaml:"elasticsearchConfig"`
	// event_bridge_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/appsync_datasource#event_bridge_config AppsyncDatasource#event_bridge_config}
	EventBridgeConfig *AppsyncDatasourceEventBridgeConfig `field:"optional" json:"eventBridgeConfig" yaml:"eventBridgeConfig"`
	// http_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/appsync_datasource#http_config AppsyncDatasource#http_config}
	HttpConfig *AppsyncDatasourceHttpConfig `field:"optional" json:"httpConfig" yaml:"httpConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/appsync_datasource#id AppsyncDatasource#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// lambda_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/appsync_datasource#lambda_config AppsyncDatasource#lambda_config}
	LambdaConfig *AppsyncDatasourceLambdaConfig `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
	// opensearchservice_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/appsync_datasource#opensearchservice_config AppsyncDatasource#opensearchservice_config}
	OpensearchserviceConfig *AppsyncDatasourceOpensearchserviceConfig `field:"optional" json:"opensearchserviceConfig" yaml:"opensearchserviceConfig"`
	// relational_database_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/appsync_datasource#relational_database_config AppsyncDatasource#relational_database_config}
	RelationalDatabaseConfig *AppsyncDatasourceRelationalDatabaseConfig `field:"optional" json:"relationalDatabaseConfig" yaml:"relationalDatabaseConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/appsync_datasource#service_role_arn AppsyncDatasource#service_role_arn}.
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
}

