package glue

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// AWS Glue.
type GluePartitionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition#database_name GluePartition#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition#partition_values GluePartition#partition_values}.
	PartitionValues *[]*string `field:"required" json:"partitionValues" yaml:"partitionValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition#table_name GluePartition#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition#catalog_id GluePartition#catalog_id}.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition#id GluePartition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition#parameters GluePartition#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// storage_descriptor block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition#storage_descriptor GluePartition#storage_descriptor}
	StorageDescriptor *GluePartitionStorageDescriptor `field:"optional" json:"storageDescriptor" yaml:"storageDescriptor"`
}

