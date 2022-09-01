package iot


type IotThingGroupPropertiesAttributePayload struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iot_thing_group#attributes IotThingGroup#attributes}.
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
}

