// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecrawler

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlueCrawlerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#database_name GlueCrawler#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#name GlueCrawler#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#role GlueCrawler#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// catalog_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#catalog_target GlueCrawler#catalog_target}
	CatalogTarget interface{} `field:"optional" json:"catalogTarget" yaml:"catalogTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#classifiers GlueCrawler#classifiers}.
	Classifiers *[]*string `field:"optional" json:"classifiers" yaml:"classifiers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#configuration GlueCrawler#configuration}.
	Configuration *string `field:"optional" json:"configuration" yaml:"configuration"`
	// delta_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#delta_target GlueCrawler#delta_target}
	DeltaTarget interface{} `field:"optional" json:"deltaTarget" yaml:"deltaTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#description GlueCrawler#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dynamodb_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#dynamodb_target GlueCrawler#dynamodb_target}
	DynamodbTarget interface{} `field:"optional" json:"dynamodbTarget" yaml:"dynamodbTarget"`
	// hudi_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#hudi_target GlueCrawler#hudi_target}
	HudiTarget interface{} `field:"optional" json:"hudiTarget" yaml:"hudiTarget"`
	// iceberg_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#iceberg_target GlueCrawler#iceberg_target}
	IcebergTarget interface{} `field:"optional" json:"icebergTarget" yaml:"icebergTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#id GlueCrawler#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// jdbc_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#jdbc_target GlueCrawler#jdbc_target}
	JdbcTarget interface{} `field:"optional" json:"jdbcTarget" yaml:"jdbcTarget"`
	// lake_formation_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#lake_formation_configuration GlueCrawler#lake_formation_configuration}
	LakeFormationConfiguration *GlueCrawlerLakeFormationConfiguration `field:"optional" json:"lakeFormationConfiguration" yaml:"lakeFormationConfiguration"`
	// lineage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#lineage_configuration GlueCrawler#lineage_configuration}
	LineageConfiguration *GlueCrawlerLineageConfiguration `field:"optional" json:"lineageConfiguration" yaml:"lineageConfiguration"`
	// mongodb_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#mongodb_target GlueCrawler#mongodb_target}
	MongodbTarget interface{} `field:"optional" json:"mongodbTarget" yaml:"mongodbTarget"`
	// recrawl_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#recrawl_policy GlueCrawler#recrawl_policy}
	RecrawlPolicy *GlueCrawlerRecrawlPolicy `field:"optional" json:"recrawlPolicy" yaml:"recrawlPolicy"`
	// s3_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#s3_target GlueCrawler#s3_target}
	S3Target interface{} `field:"optional" json:"s3Target" yaml:"s3Target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#schedule GlueCrawler#schedule}.
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// schema_change_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#schema_change_policy GlueCrawler#schema_change_policy}
	SchemaChangePolicy *GlueCrawlerSchemaChangePolicy `field:"optional" json:"schemaChangePolicy" yaml:"schemaChangePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#security_configuration GlueCrawler#security_configuration}.
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#table_prefix GlueCrawler#table_prefix}.
	TablePrefix *string `field:"optional" json:"tablePrefix" yaml:"tablePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#tags GlueCrawler#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/glue_crawler#tags_all GlueCrawler#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

