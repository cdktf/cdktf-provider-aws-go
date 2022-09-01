package storagegateway


type StoragegatewayGatewayMaintenanceStartTime struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway#hour_of_day StoragegatewayGateway#hour_of_day}.
	HourOfDay *float64 `field:"required" json:"hourOfDay" yaml:"hourOfDay"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway#day_of_month StoragegatewayGateway#day_of_month}.
	DayOfMonth *string `field:"optional" json:"dayOfMonth" yaml:"dayOfMonth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway#day_of_week StoragegatewayGateway#day_of_week}.
	DayOfWeek *string `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway#minute_of_hour StoragegatewayGateway#minute_of_hour}.
	MinuteOfHour *float64 `field:"optional" json:"minuteOfHour" yaml:"minuteOfHour"`
}

