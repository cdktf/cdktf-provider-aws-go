package mskconnectconnector


type MskconnectConnectorCapacityAutoscaling struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/mskconnect_connector#max_worker_count MskconnectConnector#max_worker_count}.
	MaxWorkerCount *float64 `field:"required" json:"maxWorkerCount" yaml:"maxWorkerCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/mskconnect_connector#min_worker_count MskconnectConnector#min_worker_count}.
	MinWorkerCount *float64 `field:"required" json:"minWorkerCount" yaml:"minWorkerCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/mskconnect_connector#mcu_count MskconnectConnector#mcu_count}.
	McuCount *float64 `field:"optional" json:"mcuCount" yaml:"mcuCount"`
	// scale_in_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/mskconnect_connector#scale_in_policy MskconnectConnector#scale_in_policy}
	ScaleInPolicy *MskconnectConnectorCapacityAutoscalingScaleInPolicy `field:"optional" json:"scaleInPolicy" yaml:"scaleInPolicy"`
	// scale_out_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/mskconnect_connector#scale_out_policy MskconnectConnector#scale_out_policy}
	ScaleOutPolicy *MskconnectConnectorCapacityAutoscalingScaleOutPolicy `field:"optional" json:"scaleOutPolicy" yaml:"scaleOutPolicy"`
}

