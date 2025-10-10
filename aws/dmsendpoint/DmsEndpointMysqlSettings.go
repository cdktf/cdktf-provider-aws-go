// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint


type DmsEndpointMysqlSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/dms_endpoint#after_connect_script DmsEndpoint#after_connect_script}.
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/dms_endpoint#authentication_method DmsEndpoint#authentication_method}.
	AuthenticationMethod *string `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/dms_endpoint#clean_source_metadata_on_mismatch DmsEndpoint#clean_source_metadata_on_mismatch}.
	CleanSourceMetadataOnMismatch interface{} `field:"optional" json:"cleanSourceMetadataOnMismatch" yaml:"cleanSourceMetadataOnMismatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/dms_endpoint#events_poll_interval DmsEndpoint#events_poll_interval}.
	EventsPollInterval *float64 `field:"optional" json:"eventsPollInterval" yaml:"eventsPollInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/dms_endpoint#execute_timeout DmsEndpoint#execute_timeout}.
	ExecuteTimeout *float64 `field:"optional" json:"executeTimeout" yaml:"executeTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/dms_endpoint#max_file_size DmsEndpoint#max_file_size}.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/dms_endpoint#parallel_load_threads DmsEndpoint#parallel_load_threads}.
	ParallelLoadThreads *float64 `field:"optional" json:"parallelLoadThreads" yaml:"parallelLoadThreads"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/dms_endpoint#server_timezone DmsEndpoint#server_timezone}.
	ServerTimezone *string `field:"optional" json:"serverTimezone" yaml:"serverTimezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/dms_endpoint#service_access_role_arn DmsEndpoint#service_access_role_arn}.
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/dms_endpoint#target_db_type DmsEndpoint#target_db_type}.
	TargetDbType *string `field:"optional" json:"targetDbType" yaml:"targetDbType"`
}

