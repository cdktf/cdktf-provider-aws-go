// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package amplifybranch

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AmplifyBranchConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#app_id AmplifyBranch#app_id}.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#branch_name AmplifyBranch#branch_name}.
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#backend_environment_arn AmplifyBranch#backend_environment_arn}.
	BackendEnvironmentArn *string `field:"optional" json:"backendEnvironmentArn" yaml:"backendEnvironmentArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#basic_auth_credentials AmplifyBranch#basic_auth_credentials}.
	BasicAuthCredentials *string `field:"optional" json:"basicAuthCredentials" yaml:"basicAuthCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#description AmplifyBranch#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#display_name AmplifyBranch#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#enable_auto_build AmplifyBranch#enable_auto_build}.
	EnableAutoBuild interface{} `field:"optional" json:"enableAutoBuild" yaml:"enableAutoBuild"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#enable_basic_auth AmplifyBranch#enable_basic_auth}.
	EnableBasicAuth interface{} `field:"optional" json:"enableBasicAuth" yaml:"enableBasicAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#enable_notification AmplifyBranch#enable_notification}.
	EnableNotification interface{} `field:"optional" json:"enableNotification" yaml:"enableNotification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#enable_performance_mode AmplifyBranch#enable_performance_mode}.
	EnablePerformanceMode interface{} `field:"optional" json:"enablePerformanceMode" yaml:"enablePerformanceMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#enable_pull_request_preview AmplifyBranch#enable_pull_request_preview}.
	EnablePullRequestPreview interface{} `field:"optional" json:"enablePullRequestPreview" yaml:"enablePullRequestPreview"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#enable_skew_protection AmplifyBranch#enable_skew_protection}.
	EnableSkewProtection interface{} `field:"optional" json:"enableSkewProtection" yaml:"enableSkewProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#environment_variables AmplifyBranch#environment_variables}.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#framework AmplifyBranch#framework}.
	Framework *string `field:"optional" json:"framework" yaml:"framework"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#id AmplifyBranch#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#pull_request_environment_name AmplifyBranch#pull_request_environment_name}.
	PullRequestEnvironmentName *string `field:"optional" json:"pullRequestEnvironmentName" yaml:"pullRequestEnvironmentName"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#region AmplifyBranch#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#stage AmplifyBranch#stage}.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#tags AmplifyBranch#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#tags_all AmplifyBranch#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/amplify_branch#ttl AmplifyBranch#ttl}.
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

