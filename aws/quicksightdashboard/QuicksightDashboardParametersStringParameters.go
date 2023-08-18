package quicksightdashboard


type QuicksightDashboardParametersStringParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/quicksight_dashboard#name QuicksightDashboard#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/quicksight_dashboard#values QuicksightDashboard#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

