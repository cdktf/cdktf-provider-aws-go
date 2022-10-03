package iotprovisioningtemplate


type IotProvisioningTemplatePreProvisioningHook struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_provisioning_template#target_arn IotProvisioningTemplate#target_arn}.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_provisioning_template#payload_version IotProvisioningTemplate#payload_version}.
	PayloadVersion *string `field:"optional" json:"payloadVersion" yaml:"payloadVersion"`
}

