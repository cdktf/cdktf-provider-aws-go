package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerOutlierDetection struct {
	// base_ejection_duration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#base_ejection_duration AppmeshVirtualNode#base_ejection_duration}
	BaseEjectionDuration *AppmeshVirtualNodeSpecListenerOutlierDetectionBaseEjectionDuration `field:"required" json:"baseEjectionDuration" yaml:"baseEjectionDuration"`
	// interval block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#interval AppmeshVirtualNode#interval}
	Interval *AppmeshVirtualNodeSpecListenerOutlierDetectionInterval `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#max_ejection_percent AppmeshVirtualNode#max_ejection_percent}.
	MaxEjectionPercent *float64 `field:"required" json:"maxEjectionPercent" yaml:"maxEjectionPercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#max_server_errors AppmeshVirtualNode#max_server_errors}.
	MaxServerErrors *float64 `field:"required" json:"maxServerErrors" yaml:"maxServerErrors"`
}

