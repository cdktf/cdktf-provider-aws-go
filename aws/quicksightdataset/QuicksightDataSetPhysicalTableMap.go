// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetPhysicalTableMap struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/quicksight_data_set#physical_table_map_id QuicksightDataSet#physical_table_map_id}.
	PhysicalTableMapId *string `field:"required" json:"physicalTableMapId" yaml:"physicalTableMapId"`
	// custom_sql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/quicksight_data_set#custom_sql QuicksightDataSet#custom_sql}
	CustomSql *QuicksightDataSetPhysicalTableMapCustomSql `field:"optional" json:"customSql" yaml:"customSql"`
	// relational_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/quicksight_data_set#relational_table QuicksightDataSet#relational_table}
	RelationalTable *QuicksightDataSetPhysicalTableMapRelationalTable `field:"optional" json:"relationalTable" yaml:"relationalTable"`
	// s3_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/quicksight_data_set#s3_source QuicksightDataSet#s3_source}
	S3Source *QuicksightDataSetPhysicalTableMapS3Source `field:"optional" json:"s3Source" yaml:"s3Source"`
}

