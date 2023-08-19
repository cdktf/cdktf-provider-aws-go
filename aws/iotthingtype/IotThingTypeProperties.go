package iotthingtype


type IotThingTypeProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/iot_thing_type#description IotThingType#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/iot_thing_type#searchable_attributes IotThingType#searchable_attributes}.
	SearchableAttributes *[]*string `field:"optional" json:"searchableAttributes" yaml:"searchableAttributes"`
}

