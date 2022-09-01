package msk


type MskconnectConnectorCapacityProvisionedCapacity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mskconnect_connector#worker_count MskconnectConnector#worker_count}.
	WorkerCount *float64 `field:"required" json:"workerCount" yaml:"workerCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mskconnect_connector#mcu_count MskconnectConnector#mcu_count}.
	McuCount *float64 `field:"optional" json:"mcuCount" yaml:"mcuCount"`
}

