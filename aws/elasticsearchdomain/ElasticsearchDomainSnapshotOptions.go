package elasticsearchdomain


type ElasticsearchDomainSnapshotOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/elasticsearch_domain#automated_snapshot_start_hour ElasticsearchDomain#automated_snapshot_start_hour}.
	AutomatedSnapshotStartHour *float64 `field:"required" json:"automatedSnapshotStartHour" yaml:"automatedSnapshotStartHour"`
}

