// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrinstancegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrInstanceGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/emr_instance_group#cluster_id EmrInstanceGroup#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/emr_instance_group#instance_type EmrInstanceGroup#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/emr_instance_group#autoscaling_policy EmrInstanceGroup#autoscaling_policy}.
	AutoscalingPolicy *string `field:"optional" json:"autoscalingPolicy" yaml:"autoscalingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/emr_instance_group#bid_price EmrInstanceGroup#bid_price}.
	BidPrice *string `field:"optional" json:"bidPrice" yaml:"bidPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/emr_instance_group#configurations_json EmrInstanceGroup#configurations_json}.
	ConfigurationsJson *string `field:"optional" json:"configurationsJson" yaml:"configurationsJson"`
	// ebs_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/emr_instance_group#ebs_config EmrInstanceGroup#ebs_config}
	EbsConfig interface{} `field:"optional" json:"ebsConfig" yaml:"ebsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/emr_instance_group#ebs_optimized EmrInstanceGroup#ebs_optimized}.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/emr_instance_group#id EmrInstanceGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/emr_instance_group#instance_count EmrInstanceGroup#instance_count}.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/emr_instance_group#name EmrInstanceGroup#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

