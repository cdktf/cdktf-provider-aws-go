package elasticachereplicationgroup


type ElasticacheReplicationGroupClusterMode struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticache_replication_group#num_node_groups ElasticacheReplicationGroup#num_node_groups}.
	NumNodeGroups *float64 `field:"optional" json:"numNodeGroups" yaml:"numNodeGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticache_replication_group#replicas_per_node_group ElasticacheReplicationGroup#replicas_per_node_group}.
	ReplicasPerNodeGroup *float64 `field:"optional" json:"replicasPerNodeGroup" yaml:"replicasPerNodeGroup"`
}

