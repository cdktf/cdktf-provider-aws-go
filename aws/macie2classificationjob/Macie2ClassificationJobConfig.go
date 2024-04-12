// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package macie2classificationjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Macie2ClassificationJobConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/macie2_classification_job#job_type Macie2ClassificationJob#job_type}.
	JobType *string `field:"required" json:"jobType" yaml:"jobType"`
	// s3_job_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/macie2_classification_job#s3_job_definition Macie2ClassificationJob#s3_job_definition}
	S3JobDefinition *Macie2ClassificationJobS3JobDefinition `field:"required" json:"s3JobDefinition" yaml:"s3JobDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/macie2_classification_job#custom_data_identifier_ids Macie2ClassificationJob#custom_data_identifier_ids}.
	CustomDataIdentifierIds *[]*string `field:"optional" json:"customDataIdentifierIds" yaml:"customDataIdentifierIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/macie2_classification_job#description Macie2ClassificationJob#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/macie2_classification_job#id Macie2ClassificationJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/macie2_classification_job#initial_run Macie2ClassificationJob#initial_run}.
	InitialRun interface{} `field:"optional" json:"initialRun" yaml:"initialRun"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/macie2_classification_job#job_status Macie2ClassificationJob#job_status}.
	JobStatus *string `field:"optional" json:"jobStatus" yaml:"jobStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/macie2_classification_job#name Macie2ClassificationJob#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/macie2_classification_job#name_prefix Macie2ClassificationJob#name_prefix}.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/macie2_classification_job#sampling_percentage Macie2ClassificationJob#sampling_percentage}.
	SamplingPercentage *float64 `field:"optional" json:"samplingPercentage" yaml:"samplingPercentage"`
	// schedule_frequency block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/macie2_classification_job#schedule_frequency Macie2ClassificationJob#schedule_frequency}
	ScheduleFrequency *Macie2ClassificationJobScheduleFrequency `field:"optional" json:"scheduleFrequency" yaml:"scheduleFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/macie2_classification_job#tags Macie2ClassificationJob#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/macie2_classification_job#tags_all Macie2ClassificationJob#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

