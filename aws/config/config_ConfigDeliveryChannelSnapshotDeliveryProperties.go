package config


type ConfigDeliveryChannelSnapshotDeliveryProperties struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_delivery_channel#delivery_frequency ConfigDeliveryChannel#delivery_frequency}.
	DeliveryFrequency *string `field:"optional" json:"deliveryFrequency" yaml:"deliveryFrequency"`
}

