package ssmmaintenancewindowtask


type SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ssm_maintenance_window_task#client_context SsmMaintenanceWindowTask#client_context}.
	ClientContext *string `field:"optional" json:"clientContext" yaml:"clientContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ssm_maintenance_window_task#payload SsmMaintenanceWindowTask#payload}.
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ssm_maintenance_window_task#qualifier SsmMaintenanceWindowTask#qualifier}.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

