package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#property_group_id Kinesisanalyticsv2Application#property_group_id}.
	PropertyGroupId *string `field:"required" json:"propertyGroupId" yaml:"propertyGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#property_map Kinesisanalyticsv2Application#property_map}.
	PropertyMap *map[string]*string `field:"required" json:"propertyMap" yaml:"propertyMap"`
}

