// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasynctask

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatasyncTaskConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/datasync_task#destination_location_arn DatasyncTask#destination_location_arn}.
	DestinationLocationArn *string `field:"required" json:"destinationLocationArn" yaml:"destinationLocationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/datasync_task#source_location_arn DatasyncTask#source_location_arn}.
	SourceLocationArn *string `field:"required" json:"sourceLocationArn" yaml:"sourceLocationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/datasync_task#cloudwatch_log_group_arn DatasyncTask#cloudwatch_log_group_arn}.
	CloudwatchLogGroupArn *string `field:"optional" json:"cloudwatchLogGroupArn" yaml:"cloudwatchLogGroupArn"`
	// excludes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/datasync_task#excludes DatasyncTask#excludes}
	Excludes *DatasyncTaskExcludes `field:"optional" json:"excludes" yaml:"excludes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/datasync_task#id DatasyncTask#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// includes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/datasync_task#includes DatasyncTask#includes}
	Includes *DatasyncTaskIncludes `field:"optional" json:"includes" yaml:"includes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/datasync_task#name DatasyncTask#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/datasync_task#options DatasyncTask#options}
	Options *DatasyncTaskOptions `field:"optional" json:"options" yaml:"options"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/datasync_task#region DatasyncTask#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/datasync_task#schedule DatasyncTask#schedule}
	Schedule *DatasyncTaskSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/datasync_task#tags DatasyncTask#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/datasync_task#tags_all DatasyncTask#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/datasync_task#task_mode DatasyncTask#task_mode}.
	TaskMode *string `field:"optional" json:"taskMode" yaml:"taskMode"`
	// task_report_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/datasync_task#task_report_config DatasyncTask#task_report_config}
	TaskReportConfig *DatasyncTaskTaskReportConfig `field:"optional" json:"taskReportConfig" yaml:"taskReportConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/datasync_task#timeouts DatasyncTask#timeouts}
	Timeouts *DatasyncTaskTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

