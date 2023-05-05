package mskcluster


type MskClusterOpenMonitoring struct {
	// prometheus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/msk_cluster#prometheus MskCluster#prometheus}
	Prometheus *MskClusterOpenMonitoringPrometheus `field:"required" json:"prometheus" yaml:"prometheus"`
}

