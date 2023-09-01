// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingGroupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#max_size AutoscalingGroup#max_size}.
	MaxSize *float64 `field:"required" json:"maxSize" yaml:"maxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#min_size AutoscalingGroup#min_size}.
	MinSize *float64 `field:"required" json:"minSize" yaml:"minSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#availability_zones AutoscalingGroup#availability_zones}.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#capacity_rebalance AutoscalingGroup#capacity_rebalance}.
	CapacityRebalance interface{} `field:"optional" json:"capacityRebalance" yaml:"capacityRebalance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#context AutoscalingGroup#context}.
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#default_cooldown AutoscalingGroup#default_cooldown}.
	DefaultCooldown *float64 `field:"optional" json:"defaultCooldown" yaml:"defaultCooldown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#default_instance_warmup AutoscalingGroup#default_instance_warmup}.
	DefaultInstanceWarmup *float64 `field:"optional" json:"defaultInstanceWarmup" yaml:"defaultInstanceWarmup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#desired_capacity AutoscalingGroup#desired_capacity}.
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#desired_capacity_type AutoscalingGroup#desired_capacity_type}.
	DesiredCapacityType *string `field:"optional" json:"desiredCapacityType" yaml:"desiredCapacityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#enabled_metrics AutoscalingGroup#enabled_metrics}.
	EnabledMetrics *[]*string `field:"optional" json:"enabledMetrics" yaml:"enabledMetrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#force_delete AutoscalingGroup#force_delete}.
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#force_delete_warm_pool AutoscalingGroup#force_delete_warm_pool}.
	ForceDeleteWarmPool interface{} `field:"optional" json:"forceDeleteWarmPool" yaml:"forceDeleteWarmPool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#health_check_grace_period AutoscalingGroup#health_check_grace_period}.
	HealthCheckGracePeriod *float64 `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#health_check_type AutoscalingGroup#health_check_type}.
	HealthCheckType *string `field:"optional" json:"healthCheckType" yaml:"healthCheckType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#id AutoscalingGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#ignore_failed_scaling_activities AutoscalingGroup#ignore_failed_scaling_activities}.
	IgnoreFailedScalingActivities interface{} `field:"optional" json:"ignoreFailedScalingActivities" yaml:"ignoreFailedScalingActivities"`
	// initial_lifecycle_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#initial_lifecycle_hook AutoscalingGroup#initial_lifecycle_hook}
	InitialLifecycleHook interface{} `field:"optional" json:"initialLifecycleHook" yaml:"initialLifecycleHook"`
	// instance_refresh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#instance_refresh AutoscalingGroup#instance_refresh}
	InstanceRefresh *AutoscalingGroupInstanceRefresh `field:"optional" json:"instanceRefresh" yaml:"instanceRefresh"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#launch_configuration AutoscalingGroup#launch_configuration}.
	LaunchConfiguration *string `field:"optional" json:"launchConfiguration" yaml:"launchConfiguration"`
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#launch_template AutoscalingGroup#launch_template}
	LaunchTemplate *AutoscalingGroupLaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#load_balancers AutoscalingGroup#load_balancers}.
	LoadBalancers *[]*string `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#max_instance_lifetime AutoscalingGroup#max_instance_lifetime}.
	MaxInstanceLifetime *float64 `field:"optional" json:"maxInstanceLifetime" yaml:"maxInstanceLifetime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#metrics_granularity AutoscalingGroup#metrics_granularity}.
	MetricsGranularity *string `field:"optional" json:"metricsGranularity" yaml:"metricsGranularity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#min_elb_capacity AutoscalingGroup#min_elb_capacity}.
	MinElbCapacity *float64 `field:"optional" json:"minElbCapacity" yaml:"minElbCapacity"`
	// mixed_instances_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#mixed_instances_policy AutoscalingGroup#mixed_instances_policy}
	MixedInstancesPolicy *AutoscalingGroupMixedInstancesPolicy `field:"optional" json:"mixedInstancesPolicy" yaml:"mixedInstancesPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#name AutoscalingGroup#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#name_prefix AutoscalingGroup#name_prefix}.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#placement_group AutoscalingGroup#placement_group}.
	PlacementGroup *string `field:"optional" json:"placementGroup" yaml:"placementGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#protect_from_scale_in AutoscalingGroup#protect_from_scale_in}.
	ProtectFromScaleIn interface{} `field:"optional" json:"protectFromScaleIn" yaml:"protectFromScaleIn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#service_linked_role_arn AutoscalingGroup#service_linked_role_arn}.
	ServiceLinkedRoleArn *string `field:"optional" json:"serviceLinkedRoleArn" yaml:"serviceLinkedRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#suspended_processes AutoscalingGroup#suspended_processes}.
	SuspendedProcesses *[]*string `field:"optional" json:"suspendedProcesses" yaml:"suspendedProcesses"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#tag AutoscalingGroup#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#target_group_arns AutoscalingGroup#target_group_arns}.
	TargetGroupArns *[]*string `field:"optional" json:"targetGroupArns" yaml:"targetGroupArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#termination_policies AutoscalingGroup#termination_policies}.
	TerminationPolicies *[]*string `field:"optional" json:"terminationPolicies" yaml:"terminationPolicies"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#timeouts AutoscalingGroup#timeouts}
	Timeouts *AutoscalingGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// traffic_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#traffic_source AutoscalingGroup#traffic_source}
	TrafficSource interface{} `field:"optional" json:"trafficSource" yaml:"trafficSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#vpc_zone_identifier AutoscalingGroup#vpc_zone_identifier}.
	VpcZoneIdentifier *[]*string `field:"optional" json:"vpcZoneIdentifier" yaml:"vpcZoneIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#wait_for_capacity_timeout AutoscalingGroup#wait_for_capacity_timeout}.
	WaitForCapacityTimeout *string `field:"optional" json:"waitForCapacityTimeout" yaml:"waitForCapacityTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#wait_for_elb_capacity AutoscalingGroup#wait_for_elb_capacity}.
	WaitForElbCapacity *float64 `field:"optional" json:"waitForElbCapacity" yaml:"waitForElbCapacity"`
	// warm_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/autoscaling_group#warm_pool AutoscalingGroup#warm_pool}
	WarmPool *AutoscalingGroupWarmPool `field:"optional" json:"warmPool" yaml:"warmPool"`
}

