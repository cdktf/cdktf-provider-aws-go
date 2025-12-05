// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lambda_function#capacity_provider_arn LambdaFunction#capacity_provider_arn}.
	CapacityProviderArn *string `field:"required" json:"capacityProviderArn" yaml:"capacityProviderArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lambda_function#execution_environment_memory_gib_per_vcpu LambdaFunction#execution_environment_memory_gib_per_vcpu}.
	ExecutionEnvironmentMemoryGibPerVcpu *float64 `field:"optional" json:"executionEnvironmentMemoryGibPerVcpu" yaml:"executionEnvironmentMemoryGibPerVcpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lambda_function#per_execution_environment_max_concurrency LambdaFunction#per_execution_environment_max_concurrency}.
	PerExecutionEnvironmentMaxConcurrency *float64 `field:"optional" json:"perExecutionEnvironmentMaxConcurrency" yaml:"perExecutionEnvironmentMaxConcurrency"`
}

