// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferserver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransferServerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#certificate TransferServer#certificate}.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#directory_id TransferServer#directory_id}.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#domain TransferServer#domain}.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// endpoint_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#endpoint_details TransferServer#endpoint_details}
	EndpointDetails *TransferServerEndpointDetails `field:"optional" json:"endpointDetails" yaml:"endpointDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#endpoint_type TransferServer#endpoint_type}.
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#force_destroy TransferServer#force_destroy}.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#function TransferServer#function}.
	Function *string `field:"optional" json:"function" yaml:"function"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#host_key TransferServer#host_key}.
	HostKey *string `field:"optional" json:"hostKey" yaml:"hostKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#id TransferServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#identity_provider_type TransferServer#identity_provider_type}.
	IdentityProviderType *string `field:"optional" json:"identityProviderType" yaml:"identityProviderType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#invocation_role TransferServer#invocation_role}.
	InvocationRole *string `field:"optional" json:"invocationRole" yaml:"invocationRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#logging_role TransferServer#logging_role}.
	LoggingRole *string `field:"optional" json:"loggingRole" yaml:"loggingRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#post_authentication_login_banner TransferServer#post_authentication_login_banner}.
	PostAuthenticationLoginBanner *string `field:"optional" json:"postAuthenticationLoginBanner" yaml:"postAuthenticationLoginBanner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#pre_authentication_login_banner TransferServer#pre_authentication_login_banner}.
	PreAuthenticationLoginBanner *string `field:"optional" json:"preAuthenticationLoginBanner" yaml:"preAuthenticationLoginBanner"`
	// protocol_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#protocol_details TransferServer#protocol_details}
	ProtocolDetails *TransferServerProtocolDetails `field:"optional" json:"protocolDetails" yaml:"protocolDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#protocols TransferServer#protocols}.
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// s3_storage_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#s3_storage_options TransferServer#s3_storage_options}
	S3StorageOptions *TransferServerS3StorageOptions `field:"optional" json:"s3StorageOptions" yaml:"s3StorageOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#security_policy_name TransferServer#security_policy_name}.
	SecurityPolicyName *string `field:"optional" json:"securityPolicyName" yaml:"securityPolicyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#sftp_authentication_methods TransferServer#sftp_authentication_methods}.
	SftpAuthenticationMethods *string `field:"optional" json:"sftpAuthenticationMethods" yaml:"sftpAuthenticationMethods"`
	// This is a set of arns of destinations that will receive structured logs from the transfer server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#structured_log_destinations TransferServer#structured_log_destinations}
	StructuredLogDestinations *[]*string `field:"optional" json:"structuredLogDestinations" yaml:"structuredLogDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#tags TransferServer#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#tags_all TransferServer#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#url TransferServer#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// workflow_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transfer_server#workflow_details TransferServer#workflow_details}
	WorkflowDetails *TransferServerWorkflowDetails `field:"optional" json:"workflowDetails" yaml:"workflowDetails"`
}

