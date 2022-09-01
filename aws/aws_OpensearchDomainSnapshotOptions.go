// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type OpensearchDomainSnapshotOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#automated_snapshot_start_hour OpensearchDomain#automated_snapshot_start_hour}.
	AutomatedSnapshotStartHour *float64 `field:"required" json:"automatedSnapshotStartHour" yaml:"automatedSnapshotStartHour"`
}

