// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package amplifybackendenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AmplifyBackendEnvironmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/amplify_backend_environment#app_id AmplifyBackendEnvironment#app_id}.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/amplify_backend_environment#environment_name AmplifyBackendEnvironment#environment_name}.
	EnvironmentName *string `field:"required" json:"environmentName" yaml:"environmentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/amplify_backend_environment#deployment_artifacts AmplifyBackendEnvironment#deployment_artifacts}.
	DeploymentArtifacts *string `field:"optional" json:"deploymentArtifacts" yaml:"deploymentArtifacts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/amplify_backend_environment#id AmplifyBackendEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/amplify_backend_environment#stack_name AmplifyBackendEnvironment#stack_name}.
	StackName *string `field:"optional" json:"stackName" yaml:"stackName"`
}

