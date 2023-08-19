package dataawsssmmaintenancewindows


type DataAwsSsmMaintenanceWindowsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/data-sources/ssm_maintenance_windows#name DataAwsSsmMaintenanceWindows#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/data-sources/ssm_maintenance_windows#values DataAwsSsmMaintenanceWindows#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

