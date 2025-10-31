// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildproject

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodebuildProjectConfig struct {
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
	// artifacts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#artifacts CodebuildProject#artifacts}
	Artifacts *CodebuildProjectArtifacts `field:"required" json:"artifacts" yaml:"artifacts"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#environment CodebuildProject#environment}
	Environment *CodebuildProjectEnvironment `field:"required" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#name CodebuildProject#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#service_role CodebuildProject#service_role}.
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#source CodebuildProject#source}
	Source *CodebuildProjectSource `field:"required" json:"source" yaml:"source"`
	// Maximum number of additional automatic retries after a failed build. The default value is 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#auto_retry_limit CodebuildProject#auto_retry_limit}
	AutoRetryLimit *float64 `field:"optional" json:"autoRetryLimit" yaml:"autoRetryLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#badge_enabled CodebuildProject#badge_enabled}.
	BadgeEnabled interface{} `field:"optional" json:"badgeEnabled" yaml:"badgeEnabled"`
	// build_batch_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#build_batch_config CodebuildProject#build_batch_config}
	BuildBatchConfig *CodebuildProjectBuildBatchConfig `field:"optional" json:"buildBatchConfig" yaml:"buildBatchConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#build_timeout CodebuildProject#build_timeout}.
	BuildTimeout *float64 `field:"optional" json:"buildTimeout" yaml:"buildTimeout"`
	// cache block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#cache CodebuildProject#cache}
	Cache *CodebuildProjectCache `field:"optional" json:"cache" yaml:"cache"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#concurrent_build_limit CodebuildProject#concurrent_build_limit}.
	ConcurrentBuildLimit *float64 `field:"optional" json:"concurrentBuildLimit" yaml:"concurrentBuildLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#description CodebuildProject#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#encryption_key CodebuildProject#encryption_key}.
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// file_system_locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#file_system_locations CodebuildProject#file_system_locations}
	FileSystemLocations interface{} `field:"optional" json:"fileSystemLocations" yaml:"fileSystemLocations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#id CodebuildProject#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// logs_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#logs_config CodebuildProject#logs_config}
	LogsConfig *CodebuildProjectLogsConfig `field:"optional" json:"logsConfig" yaml:"logsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#project_visibility CodebuildProject#project_visibility}.
	ProjectVisibility *string `field:"optional" json:"projectVisibility" yaml:"projectVisibility"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#queued_timeout CodebuildProject#queued_timeout}.
	QueuedTimeout *float64 `field:"optional" json:"queuedTimeout" yaml:"queuedTimeout"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#region CodebuildProject#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#resource_access_role CodebuildProject#resource_access_role}.
	ResourceAccessRole *string `field:"optional" json:"resourceAccessRole" yaml:"resourceAccessRole"`
	// secondary_artifacts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#secondary_artifacts CodebuildProject#secondary_artifacts}
	SecondaryArtifacts interface{} `field:"optional" json:"secondaryArtifacts" yaml:"secondaryArtifacts"`
	// secondary_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#secondary_sources CodebuildProject#secondary_sources}
	SecondarySources interface{} `field:"optional" json:"secondarySources" yaml:"secondarySources"`
	// secondary_source_version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#secondary_source_version CodebuildProject#secondary_source_version}
	SecondarySourceVersion interface{} `field:"optional" json:"secondarySourceVersion" yaml:"secondarySourceVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#source_version CodebuildProject#source_version}.
	SourceVersion *string `field:"optional" json:"sourceVersion" yaml:"sourceVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#tags CodebuildProject#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#tags_all CodebuildProject#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/codebuild_project#vpc_config CodebuildProject#vpc_config}
	VpcConfig *CodebuildProjectVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

