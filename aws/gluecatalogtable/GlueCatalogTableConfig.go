package gluecatalogtable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlueCatalogTableConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/glue_catalog_table#database_name GlueCatalogTable#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/glue_catalog_table#name GlueCatalogTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/glue_catalog_table#catalog_id GlueCatalogTable#catalog_id}.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/glue_catalog_table#description GlueCatalogTable#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/glue_catalog_table#id GlueCatalogTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/glue_catalog_table#owner GlueCatalogTable#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/glue_catalog_table#parameters GlueCatalogTable#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// partition_index block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/glue_catalog_table#partition_index GlueCatalogTable#partition_index}
	PartitionIndex interface{} `field:"optional" json:"partitionIndex" yaml:"partitionIndex"`
	// partition_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/glue_catalog_table#partition_keys GlueCatalogTable#partition_keys}
	PartitionKeys interface{} `field:"optional" json:"partitionKeys" yaml:"partitionKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/glue_catalog_table#retention GlueCatalogTable#retention}.
	Retention *float64 `field:"optional" json:"retention" yaml:"retention"`
	// storage_descriptor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/glue_catalog_table#storage_descriptor GlueCatalogTable#storage_descriptor}
	StorageDescriptor *GlueCatalogTableStorageDescriptor `field:"optional" json:"storageDescriptor" yaml:"storageDescriptor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/glue_catalog_table#table_type GlueCatalogTable#table_type}.
	TableType *string `field:"optional" json:"tableType" yaml:"tableType"`
	// target_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/glue_catalog_table#target_table GlueCatalogTable#target_table}
	TargetTable *GlueCatalogTableTargetTable `field:"optional" json:"targetTable" yaml:"targetTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/glue_catalog_table#view_expanded_text GlueCatalogTable#view_expanded_text}.
	ViewExpandedText *string `field:"optional" json:"viewExpandedText" yaml:"viewExpandedText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/glue_catalog_table#view_original_text GlueCatalogTable#view_original_text}.
	ViewOriginalText *string `field:"optional" json:"viewOriginalText" yaml:"viewOriginalText"`
}

