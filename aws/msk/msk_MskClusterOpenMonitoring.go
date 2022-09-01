package msk


type MskClusterOpenMonitoring struct {
	// prometheus block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/msk_cluster#prometheus MskCluster#prometheus}
	Prometheus *MskClusterOpenMonitoringPrometheus `field:"required" json:"prometheus" yaml:"prometheus"`
}

