package quicksightdashboard


type QuicksightDashboardParametersIntegerParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/quicksight_dashboard#name QuicksightDashboard#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/quicksight_dashboard#values QuicksightDashboard#values}.
	Values *[]*float64 `field:"required" json:"values" yaml:"values"`
}

