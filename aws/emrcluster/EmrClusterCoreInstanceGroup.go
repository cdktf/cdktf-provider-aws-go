package emrcluster


type EmrClusterCoreInstanceGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/emr_cluster#instance_type EmrCluster#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/emr_cluster#autoscaling_policy EmrCluster#autoscaling_policy}.
	AutoscalingPolicy *string `field:"optional" json:"autoscalingPolicy" yaml:"autoscalingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/emr_cluster#bid_price EmrCluster#bid_price}.
	BidPrice *string `field:"optional" json:"bidPrice" yaml:"bidPrice"`
	// ebs_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/emr_cluster#ebs_config EmrCluster#ebs_config}
	EbsConfig interface{} `field:"optional" json:"ebsConfig" yaml:"ebsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/emr_cluster#instance_count EmrCluster#instance_count}.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/emr_cluster#name EmrCluster#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

