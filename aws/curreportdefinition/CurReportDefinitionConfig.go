// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package curreportdefinition

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CurReportDefinitionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/cur_report_definition#additional_schema_elements CurReportDefinition#additional_schema_elements}.
	AdditionalSchemaElements *[]*string `field:"required" json:"additionalSchemaElements" yaml:"additionalSchemaElements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/cur_report_definition#compression CurReportDefinition#compression}.
	Compression *string `field:"required" json:"compression" yaml:"compression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/cur_report_definition#format CurReportDefinition#format}.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/cur_report_definition#report_name CurReportDefinition#report_name}.
	ReportName *string `field:"required" json:"reportName" yaml:"reportName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/cur_report_definition#s3_bucket CurReportDefinition#s3_bucket}.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/cur_report_definition#s3_region CurReportDefinition#s3_region}.
	S3Region *string `field:"required" json:"s3Region" yaml:"s3Region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/cur_report_definition#time_unit CurReportDefinition#time_unit}.
	TimeUnit *string `field:"required" json:"timeUnit" yaml:"timeUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/cur_report_definition#additional_artifacts CurReportDefinition#additional_artifacts}.
	AdditionalArtifacts *[]*string `field:"optional" json:"additionalArtifacts" yaml:"additionalArtifacts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/cur_report_definition#id CurReportDefinition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/cur_report_definition#refresh_closed_reports CurReportDefinition#refresh_closed_reports}.
	RefreshClosedReports interface{} `field:"optional" json:"refreshClosedReports" yaml:"refreshClosedReports"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/cur_report_definition#report_versioning CurReportDefinition#report_versioning}.
	ReportVersioning *string `field:"optional" json:"reportVersioning" yaml:"reportVersioning"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/cur_report_definition#s3_prefix CurReportDefinition#s3_prefix}.
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
}

