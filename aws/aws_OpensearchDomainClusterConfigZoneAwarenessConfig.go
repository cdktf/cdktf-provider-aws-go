// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type OpensearchDomainClusterConfigZoneAwarenessConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#availability_zone_count OpensearchDomain#availability_zone_count}.
	AvailabilityZoneCount *float64 `field:"optional" json:"availabilityZoneCount" yaml:"availabilityZoneCount"`
}

