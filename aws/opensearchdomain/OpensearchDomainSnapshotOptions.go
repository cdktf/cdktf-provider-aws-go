package opensearchdomain


type OpensearchDomainSnapshotOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/opensearch_domain#automated_snapshot_start_hour OpensearchDomain#automated_snapshot_start_hour}.
	AutomatedSnapshotStartHour *float64 `field:"required" json:"automatedSnapshotStartHour" yaml:"automatedSnapshotStartHour"`
}

