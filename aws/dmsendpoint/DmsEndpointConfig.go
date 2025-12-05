// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#endpoint_id DmsEndpoint#endpoint_id}.
	EndpointId *string `field:"required" json:"endpointId" yaml:"endpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#endpoint_type DmsEndpoint#endpoint_type}.
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#engine_name DmsEndpoint#engine_name}.
	EngineName *string `field:"required" json:"engineName" yaml:"engineName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#certificate_arn DmsEndpoint#certificate_arn}.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#database_name DmsEndpoint#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// elasticsearch_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#elasticsearch_settings DmsEndpoint#elasticsearch_settings}
	ElasticsearchSettings *DmsEndpointElasticsearchSettings `field:"optional" json:"elasticsearchSettings" yaml:"elasticsearchSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#extra_connection_attributes DmsEndpoint#extra_connection_attributes}.
	ExtraConnectionAttributes *string `field:"optional" json:"extraConnectionAttributes" yaml:"extraConnectionAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#id DmsEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kafka_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#kafka_settings DmsEndpoint#kafka_settings}
	KafkaSettings *DmsEndpointKafkaSettings `field:"optional" json:"kafkaSettings" yaml:"kafkaSettings"`
	// kinesis_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#kinesis_settings DmsEndpoint#kinesis_settings}
	KinesisSettings *DmsEndpointKinesisSettings `field:"optional" json:"kinesisSettings" yaml:"kinesisSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#kms_key_arn DmsEndpoint#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// mongodb_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#mongodb_settings DmsEndpoint#mongodb_settings}
	MongodbSettings *DmsEndpointMongodbSettings `field:"optional" json:"mongodbSettings" yaml:"mongodbSettings"`
	// mysql_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#mysql_settings DmsEndpoint#mysql_settings}
	MysqlSettings *DmsEndpointMysqlSettings `field:"optional" json:"mysqlSettings" yaml:"mysqlSettings"`
	// oracle_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#oracle_settings DmsEndpoint#oracle_settings}
	OracleSettings *DmsEndpointOracleSettings `field:"optional" json:"oracleSettings" yaml:"oracleSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#password DmsEndpoint#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#pause_replication_tasks DmsEndpoint#pause_replication_tasks}.
	PauseReplicationTasks interface{} `field:"optional" json:"pauseReplicationTasks" yaml:"pauseReplicationTasks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#port DmsEndpoint#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// postgres_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#postgres_settings DmsEndpoint#postgres_settings}
	PostgresSettings *DmsEndpointPostgresSettings `field:"optional" json:"postgresSettings" yaml:"postgresSettings"`
	// redis_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#redis_settings DmsEndpoint#redis_settings}
	RedisSettings *DmsEndpointRedisSettings `field:"optional" json:"redisSettings" yaml:"redisSettings"`
	// redshift_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#redshift_settings DmsEndpoint#redshift_settings}
	RedshiftSettings *DmsEndpointRedshiftSettings `field:"optional" json:"redshiftSettings" yaml:"redshiftSettings"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#region DmsEndpoint#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#secrets_manager_access_role_arn DmsEndpoint#secrets_manager_access_role_arn}.
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#secrets_manager_arn DmsEndpoint#secrets_manager_arn}.
	SecretsManagerArn *string `field:"optional" json:"secretsManagerArn" yaml:"secretsManagerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#server_name DmsEndpoint#server_name}.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#service_access_role DmsEndpoint#service_access_role}.
	ServiceAccessRole *string `field:"optional" json:"serviceAccessRole" yaml:"serviceAccessRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#ssl_mode DmsEndpoint#ssl_mode}.
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#tags DmsEndpoint#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#tags_all DmsEndpoint#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#timeouts DmsEndpoint#timeouts}
	Timeouts *DmsEndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dms_endpoint#username DmsEndpoint#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

