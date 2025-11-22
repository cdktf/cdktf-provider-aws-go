// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleScripts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/dlm_lifecycle_policy#execution_handler DlmLifecyclePolicy#execution_handler}.
	ExecutionHandler *string `field:"required" json:"executionHandler" yaml:"executionHandler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/dlm_lifecycle_policy#execute_operation_on_script_failure DlmLifecyclePolicy#execute_operation_on_script_failure}.
	ExecuteOperationOnScriptFailure interface{} `field:"optional" json:"executeOperationOnScriptFailure" yaml:"executeOperationOnScriptFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/dlm_lifecycle_policy#execution_handler_service DlmLifecyclePolicy#execution_handler_service}.
	ExecutionHandlerService *string `field:"optional" json:"executionHandlerService" yaml:"executionHandlerService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/dlm_lifecycle_policy#execution_timeout DlmLifecyclePolicy#execution_timeout}.
	ExecutionTimeout *float64 `field:"optional" json:"executionTimeout" yaml:"executionTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/dlm_lifecycle_policy#maximum_retry_count DlmLifecyclePolicy#maximum_retry_count}.
	MaximumRetryCount *float64 `field:"optional" json:"maximumRetryCount" yaml:"maximumRetryCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/dlm_lifecycle_policy#stages DlmLifecyclePolicy#stages}.
	Stages *[]*string `field:"optional" json:"stages" yaml:"stages"`
}

