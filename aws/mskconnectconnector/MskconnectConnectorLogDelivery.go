package mskconnectconnector


type MskconnectConnectorLogDelivery struct {
	// worker_log_delivery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/mskconnect_connector#worker_log_delivery MskconnectConnector#worker_log_delivery}
	WorkerLogDelivery *MskconnectConnectorLogDeliveryWorkerLogDelivery `field:"required" json:"workerLogDelivery" yaml:"workerLogDelivery"`
}

