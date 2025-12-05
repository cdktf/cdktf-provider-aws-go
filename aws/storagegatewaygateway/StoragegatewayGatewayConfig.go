// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagegatewaygateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StoragegatewayGatewayConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#gateway_name StoragegatewayGateway#gateway_name}.
	GatewayName *string `field:"required" json:"gatewayName" yaml:"gatewayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#gateway_timezone StoragegatewayGateway#gateway_timezone}.
	GatewayTimezone *string `field:"required" json:"gatewayTimezone" yaml:"gatewayTimezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#activation_key StoragegatewayGateway#activation_key}.
	ActivationKey *string `field:"optional" json:"activationKey" yaml:"activationKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#average_download_rate_limit_in_bits_per_sec StoragegatewayGateway#average_download_rate_limit_in_bits_per_sec}.
	AverageDownloadRateLimitInBitsPerSec *float64 `field:"optional" json:"averageDownloadRateLimitInBitsPerSec" yaml:"averageDownloadRateLimitInBitsPerSec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#average_upload_rate_limit_in_bits_per_sec StoragegatewayGateway#average_upload_rate_limit_in_bits_per_sec}.
	AverageUploadRateLimitInBitsPerSec *float64 `field:"optional" json:"averageUploadRateLimitInBitsPerSec" yaml:"averageUploadRateLimitInBitsPerSec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#cloudwatch_log_group_arn StoragegatewayGateway#cloudwatch_log_group_arn}.
	CloudwatchLogGroupArn *string `field:"optional" json:"cloudwatchLogGroupArn" yaml:"cloudwatchLogGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#gateway_ip_address StoragegatewayGateway#gateway_ip_address}.
	GatewayIpAddress *string `field:"optional" json:"gatewayIpAddress" yaml:"gatewayIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#gateway_type StoragegatewayGateway#gateway_type}.
	GatewayType *string `field:"optional" json:"gatewayType" yaml:"gatewayType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#gateway_vpc_endpoint StoragegatewayGateway#gateway_vpc_endpoint}.
	GatewayVpcEndpoint *string `field:"optional" json:"gatewayVpcEndpoint" yaml:"gatewayVpcEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#id StoragegatewayGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// maintenance_start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#maintenance_start_time StoragegatewayGateway#maintenance_start_time}
	MaintenanceStartTime *StoragegatewayGatewayMaintenanceStartTime `field:"optional" json:"maintenanceStartTime" yaml:"maintenanceStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#medium_changer_type StoragegatewayGateway#medium_changer_type}.
	MediumChangerType *string `field:"optional" json:"mediumChangerType" yaml:"mediumChangerType"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#region StoragegatewayGateway#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// smb_active_directory_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#smb_active_directory_settings StoragegatewayGateway#smb_active_directory_settings}
	SmbActiveDirectorySettings *StoragegatewayGatewaySmbActiveDirectorySettings `field:"optional" json:"smbActiveDirectorySettings" yaml:"smbActiveDirectorySettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#smb_file_share_visibility StoragegatewayGateway#smb_file_share_visibility}.
	SmbFileShareVisibility interface{} `field:"optional" json:"smbFileShareVisibility" yaml:"smbFileShareVisibility"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#smb_guest_password StoragegatewayGateway#smb_guest_password}.
	SmbGuestPassword *string `field:"optional" json:"smbGuestPassword" yaml:"smbGuestPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#smb_security_strategy StoragegatewayGateway#smb_security_strategy}.
	SmbSecurityStrategy *string `field:"optional" json:"smbSecurityStrategy" yaml:"smbSecurityStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#tags StoragegatewayGateway#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#tags_all StoragegatewayGateway#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#tape_drive_type StoragegatewayGateway#tape_drive_type}.
	TapeDriveType *string `field:"optional" json:"tapeDriveType" yaml:"tapeDriveType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/storagegateway_gateway#timeouts StoragegatewayGateway#timeouts}
	Timeouts *StoragegatewayGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

