// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightanalysis

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightAnalysisConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/quicksight_analysis#analysis_id QuicksightAnalysis#analysis_id}.
	AnalysisId *string `field:"required" json:"analysisId" yaml:"analysisId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/quicksight_analysis#name QuicksightAnalysis#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/quicksight_analysis#aws_account_id QuicksightAnalysis#aws_account_id}.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/quicksight_analysis#definition QuicksightAnalysis#definition}
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/quicksight_analysis#id QuicksightAnalysis#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/quicksight_analysis#parameters QuicksightAnalysis#parameters}
	Parameters *QuicksightAnalysisParameters `field:"optional" json:"parameters" yaml:"parameters"`
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/quicksight_analysis#permissions QuicksightAnalysis#permissions}
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/quicksight_analysis#recovery_window_in_days QuicksightAnalysis#recovery_window_in_days}.
	RecoveryWindowInDays *float64 `field:"optional" json:"recoveryWindowInDays" yaml:"recoveryWindowInDays"`
	// source_entity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/quicksight_analysis#source_entity QuicksightAnalysis#source_entity}
	SourceEntity *QuicksightAnalysisSourceEntity `field:"optional" json:"sourceEntity" yaml:"sourceEntity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/quicksight_analysis#tags QuicksightAnalysis#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/quicksight_analysis#tags_all QuicksightAnalysis#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/quicksight_analysis#theme_arn QuicksightAnalysis#theme_arn}.
	ThemeArn *string `field:"optional" json:"themeArn" yaml:"themeArn"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/quicksight_analysis#timeouts QuicksightAnalysis#timeouts}
	Timeouts *QuicksightAnalysisTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

