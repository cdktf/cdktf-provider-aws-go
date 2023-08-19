package redshiftscheduledaction


type RedshiftScheduledActionTargetActionPauseCluster struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/redshift_scheduled_action#cluster_identifier RedshiftScheduledAction#cluster_identifier}.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
}

