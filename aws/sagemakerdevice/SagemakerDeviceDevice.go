// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdevice


type SagemakerDeviceDevice struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/sagemaker_device#device_name SagemakerDevice#device_name}.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/sagemaker_device#description SagemakerDevice#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/sagemaker_device#iot_thing_name SagemakerDevice#iot_thing_name}.
	IotThingName *string `field:"optional" json:"iotThingName" yaml:"iotThingName"`
}

