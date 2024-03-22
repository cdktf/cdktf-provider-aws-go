// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mqbroker


type MqBrokerMaintenanceWindowStartTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/mq_broker#day_of_week MqBroker#day_of_week}.
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/mq_broker#time_of_day MqBroker#time_of_day}.
	TimeOfDay *string `field:"required" json:"timeOfDay" yaml:"timeOfDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/mq_broker#time_zone MqBroker#time_zone}.
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
}

