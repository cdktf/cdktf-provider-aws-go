package ssm


type SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#name SsmMaintenanceWindowTask#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#values SsmMaintenanceWindowTask#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

