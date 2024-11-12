// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsredshiftserverlesscredentials

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsRedshiftserverlessCredentialsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/redshiftserverless_credentials#workgroup_name DataAwsRedshiftserverlessCredentials#workgroup_name}.
	WorkgroupName *string `field:"required" json:"workgroupName" yaml:"workgroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/redshiftserverless_credentials#db_name DataAwsRedshiftserverlessCredentials#db_name}.
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/redshiftserverless_credentials#duration_seconds DataAwsRedshiftserverlessCredentials#duration_seconds}.
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/redshiftserverless_credentials#id DataAwsRedshiftserverlessCredentials#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

