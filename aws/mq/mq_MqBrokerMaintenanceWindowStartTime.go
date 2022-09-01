package mq


type MqBrokerMaintenanceWindowStartTime struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker#day_of_week MqBroker#day_of_week}.
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker#time_of_day MqBroker#time_of_day}.
	TimeOfDay *string `field:"required" json:"timeOfDay" yaml:"timeOfDay"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker#time_zone MqBroker#time_zone}.
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
}

