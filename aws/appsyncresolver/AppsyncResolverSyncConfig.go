// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncresolver


type AppsyncResolverSyncConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/appsync_resolver#conflict_detection AppsyncResolver#conflict_detection}.
	ConflictDetection *string `field:"optional" json:"conflictDetection" yaml:"conflictDetection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/appsync_resolver#conflict_handler AppsyncResolver#conflict_handler}.
	ConflictHandler *string `field:"optional" json:"conflictHandler" yaml:"conflictHandler"`
	// lambda_conflict_handler_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/appsync_resolver#lambda_conflict_handler_config AppsyncResolver#lambda_conflict_handler_config}
	LambdaConflictHandlerConfig *AppsyncResolverSyncConfigLambdaConflictHandlerConfig `field:"optional" json:"lambdaConflictHandlerConfig" yaml:"lambdaConflictHandlerConfig"`
}

