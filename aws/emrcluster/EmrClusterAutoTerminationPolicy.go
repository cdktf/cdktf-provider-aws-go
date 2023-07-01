package emrcluster


type EmrClusterAutoTerminationPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/emr_cluster#idle_timeout EmrCluster#idle_timeout}.
	IdleTimeout *float64 `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
}

