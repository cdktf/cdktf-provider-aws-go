package iotthinggroup


type IotThingGroupPropertiesAttributePayload struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/iot_thing_group#attributes IotThingGroup#attributes}.
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
}

