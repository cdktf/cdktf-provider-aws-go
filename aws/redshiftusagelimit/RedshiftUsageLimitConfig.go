// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshiftusagelimit

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedshiftUsageLimitConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/redshift_usage_limit#amount RedshiftUsageLimit#amount}.
	Amount *float64 `field:"required" json:"amount" yaml:"amount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/redshift_usage_limit#cluster_identifier RedshiftUsageLimit#cluster_identifier}.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/redshift_usage_limit#feature_type RedshiftUsageLimit#feature_type}.
	FeatureType *string `field:"required" json:"featureType" yaml:"featureType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/redshift_usage_limit#limit_type RedshiftUsageLimit#limit_type}.
	LimitType *string `field:"required" json:"limitType" yaml:"limitType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/redshift_usage_limit#breach_action RedshiftUsageLimit#breach_action}.
	BreachAction *string `field:"optional" json:"breachAction" yaml:"breachAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/redshift_usage_limit#id RedshiftUsageLimit#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/redshift_usage_limit#period RedshiftUsageLimit#period}.
	Period *string `field:"optional" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/redshift_usage_limit#tags RedshiftUsageLimit#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/redshift_usage_limit#tags_all RedshiftUsageLimit#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

