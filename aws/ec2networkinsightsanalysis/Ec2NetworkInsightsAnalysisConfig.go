// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2networkinsightsanalysis

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2NetworkInsightsAnalysisConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/ec2_network_insights_analysis#network_insights_path_id Ec2NetworkInsightsAnalysis#network_insights_path_id}.
	NetworkInsightsPathId *string `field:"required" json:"networkInsightsPathId" yaml:"networkInsightsPathId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/ec2_network_insights_analysis#filter_in_arns Ec2NetworkInsightsAnalysis#filter_in_arns}.
	FilterInArns *[]*string `field:"optional" json:"filterInArns" yaml:"filterInArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/ec2_network_insights_analysis#id Ec2NetworkInsightsAnalysis#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/ec2_network_insights_analysis#tags Ec2NetworkInsightsAnalysis#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/ec2_network_insights_analysis#tags_all Ec2NetworkInsightsAnalysis#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/ec2_network_insights_analysis#wait_for_completion Ec2NetworkInsightsAnalysis#wait_for_completion}.
	WaitForCompletion interface{} `field:"optional" json:"waitForCompletion" yaml:"waitForCompletion"`
}

