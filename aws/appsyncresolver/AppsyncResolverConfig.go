// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncresolver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsyncResolverConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/appsync_resolver#api_id AppsyncResolver#api_id}.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/appsync_resolver#field AppsyncResolver#field}.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/appsync_resolver#type AppsyncResolver#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// caching_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/appsync_resolver#caching_config AppsyncResolver#caching_config}
	CachingConfig *AppsyncResolverCachingConfig `field:"optional" json:"cachingConfig" yaml:"cachingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/appsync_resolver#code AppsyncResolver#code}.
	Code *string `field:"optional" json:"code" yaml:"code"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/appsync_resolver#data_source AppsyncResolver#data_source}.
	DataSource *string `field:"optional" json:"dataSource" yaml:"dataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/appsync_resolver#id AppsyncResolver#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/appsync_resolver#kind AppsyncResolver#kind}.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/appsync_resolver#max_batch_size AppsyncResolver#max_batch_size}.
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// pipeline_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/appsync_resolver#pipeline_config AppsyncResolver#pipeline_config}
	PipelineConfig *AppsyncResolverPipelineConfig `field:"optional" json:"pipelineConfig" yaml:"pipelineConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/appsync_resolver#request_template AppsyncResolver#request_template}.
	RequestTemplate *string `field:"optional" json:"requestTemplate" yaml:"requestTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/appsync_resolver#response_template AppsyncResolver#response_template}.
	ResponseTemplate *string `field:"optional" json:"responseTemplate" yaml:"responseTemplate"`
	// runtime block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/appsync_resolver#runtime AppsyncResolver#runtime}
	Runtime *AppsyncResolverRuntime `field:"optional" json:"runtime" yaml:"runtime"`
	// sync_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/appsync_resolver#sync_config AppsyncResolver#sync_config}
	SyncConfig *AppsyncResolverSyncConfig `field:"optional" json:"syncConfig" yaml:"syncConfig"`
}

