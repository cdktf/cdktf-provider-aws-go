package rds

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// AWS Relational Database Service.
type DataAwsRdsEngineVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/rds_engine_version#engine DataAwsRdsEngineVersion#engine}.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/rds_engine_version#id DataAwsRdsEngineVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/rds_engine_version#parameter_group_family DataAwsRdsEngineVersion#parameter_group_family}.
	ParameterGroupFamily *string `field:"optional" json:"parameterGroupFamily" yaml:"parameterGroupFamily"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/rds_engine_version#preferred_versions DataAwsRdsEngineVersion#preferred_versions}.
	PreferredVersions *[]*string `field:"optional" json:"preferredVersions" yaml:"preferredVersions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/rds_engine_version#version DataAwsRdsEngineVersion#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

