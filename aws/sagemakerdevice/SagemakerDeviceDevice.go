package sagemakerdevice


type SagemakerDeviceDevice struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/sagemaker_device#device_name SagemakerDevice#device_name}.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/sagemaker_device#description SagemakerDevice#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/sagemaker_device#iot_thing_name SagemakerDevice#iot_thing_name}.
	IotThingName *string `field:"optional" json:"iotThingName" yaml:"iotThingName"`
}

