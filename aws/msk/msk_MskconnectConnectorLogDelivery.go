package msk


type MskconnectConnectorLogDelivery struct {
	// worker_log_delivery block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mskconnect_connector#worker_log_delivery MskconnectConnector#worker_log_delivery}
	WorkerLogDelivery *MskconnectConnectorLogDeliveryWorkerLogDelivery `field:"required" json:"workerLogDelivery" yaml:"workerLogDelivery"`
}

