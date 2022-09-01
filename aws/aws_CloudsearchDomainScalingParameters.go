// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type CloudsearchDomainScalingParameters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudsearch_domain#desired_instance_type CloudsearchDomain#desired_instance_type}.
	DesiredInstanceType *string `field:"optional" json:"desiredInstanceType" yaml:"desiredInstanceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudsearch_domain#desired_partition_count CloudsearchDomain#desired_partition_count}.
	DesiredPartitionCount *float64 `field:"optional" json:"desiredPartitionCount" yaml:"desiredPartitionCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudsearch_domain#desired_replication_count CloudsearchDomain#desired_replication_count}.
	DesiredReplicationCount *float64 `field:"optional" json:"desiredReplicationCount" yaml:"desiredReplicationCount"`
}

