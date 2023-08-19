package opsworksecsclusterlayer


type OpsworksEcsClusterLayerLoadBasedAutoScalingDownscaling struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/opsworks_ecs_cluster_layer#alarms OpsworksEcsClusterLayer#alarms}.
	Alarms *[]*string `field:"optional" json:"alarms" yaml:"alarms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/opsworks_ecs_cluster_layer#cpu_threshold OpsworksEcsClusterLayer#cpu_threshold}.
	CpuThreshold *float64 `field:"optional" json:"cpuThreshold" yaml:"cpuThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/opsworks_ecs_cluster_layer#ignore_metrics_time OpsworksEcsClusterLayer#ignore_metrics_time}.
	IgnoreMetricsTime *float64 `field:"optional" json:"ignoreMetricsTime" yaml:"ignoreMetricsTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/opsworks_ecs_cluster_layer#instance_count OpsworksEcsClusterLayer#instance_count}.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/opsworks_ecs_cluster_layer#load_threshold OpsworksEcsClusterLayer#load_threshold}.
	LoadThreshold *float64 `field:"optional" json:"loadThreshold" yaml:"loadThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/opsworks_ecs_cluster_layer#memory_threshold OpsworksEcsClusterLayer#memory_threshold}.
	MemoryThreshold *float64 `field:"optional" json:"memoryThreshold" yaml:"memoryThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/opsworks_ecs_cluster_layer#thresholds_wait_time OpsworksEcsClusterLayer#thresholds_wait_time}.
	ThresholdsWaitTime *float64 `field:"optional" json:"thresholdsWaitTime" yaml:"thresholdsWaitTime"`
}

