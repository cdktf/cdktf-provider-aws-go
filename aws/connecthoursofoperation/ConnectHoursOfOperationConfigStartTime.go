package connecthoursofoperation


type ConnectHoursOfOperationConfigStartTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/connect_hours_of_operation#hours ConnectHoursOfOperation#hours}.
	Hours *float64 `field:"required" json:"hours" yaml:"hours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/connect_hours_of_operation#minutes ConnectHoursOfOperation#minutes}.
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
}

