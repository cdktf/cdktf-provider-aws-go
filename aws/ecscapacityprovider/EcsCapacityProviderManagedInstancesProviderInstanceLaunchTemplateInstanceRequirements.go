// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecscapacityprovider


type EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirements struct {
	// memory_mib block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#memory_mib EcsCapacityProvider#memory_mib}
	MemoryMib *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirementsMemoryMib `field:"required" json:"memoryMib" yaml:"memoryMib"`
	// vcpu_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#vcpu_count EcsCapacityProvider#vcpu_count}
	VcpuCount *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirementsVcpuCount `field:"required" json:"vcpuCount" yaml:"vcpuCount"`
	// accelerator_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#accelerator_count EcsCapacityProvider#accelerator_count}
	AcceleratorCount *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirementsAcceleratorCount `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#accelerator_manufacturers EcsCapacityProvider#accelerator_manufacturers}.
	AcceleratorManufacturers *[]*string `field:"optional" json:"acceleratorManufacturers" yaml:"acceleratorManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#accelerator_names EcsCapacityProvider#accelerator_names}.
	AcceleratorNames *[]*string `field:"optional" json:"acceleratorNames" yaml:"acceleratorNames"`
	// accelerator_total_memory_mib block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#accelerator_total_memory_mib EcsCapacityProvider#accelerator_total_memory_mib}
	AcceleratorTotalMemoryMib *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirementsAcceleratorTotalMemoryMib `field:"optional" json:"acceleratorTotalMemoryMib" yaml:"acceleratorTotalMemoryMib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#accelerator_types EcsCapacityProvider#accelerator_types}.
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#allowed_instance_types EcsCapacityProvider#allowed_instance_types}.
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#bare_metal EcsCapacityProvider#bare_metal}.
	BareMetal *string `field:"optional" json:"bareMetal" yaml:"bareMetal"`
	// baseline_ebs_bandwidth_mbps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#baseline_ebs_bandwidth_mbps EcsCapacityProvider#baseline_ebs_bandwidth_mbps}
	BaselineEbsBandwidthMbps *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirementsBaselineEbsBandwidthMbps `field:"optional" json:"baselineEbsBandwidthMbps" yaml:"baselineEbsBandwidthMbps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#burstable_performance EcsCapacityProvider#burstable_performance}.
	BurstablePerformance *string `field:"optional" json:"burstablePerformance" yaml:"burstablePerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#cpu_manufacturers EcsCapacityProvider#cpu_manufacturers}.
	CpuManufacturers *[]*string `field:"optional" json:"cpuManufacturers" yaml:"cpuManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#excluded_instance_types EcsCapacityProvider#excluded_instance_types}.
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#instance_generations EcsCapacityProvider#instance_generations}.
	InstanceGenerations *[]*string `field:"optional" json:"instanceGenerations" yaml:"instanceGenerations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#local_storage EcsCapacityProvider#local_storage}.
	LocalStorage *string `field:"optional" json:"localStorage" yaml:"localStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#local_storage_types EcsCapacityProvider#local_storage_types}.
	LocalStorageTypes *[]*string `field:"optional" json:"localStorageTypes" yaml:"localStorageTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#max_spot_price_as_percentage_of_optimal_on_demand_price EcsCapacityProvider#max_spot_price_as_percentage_of_optimal_on_demand_price}.
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice *float64 `field:"optional" json:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice" yaml:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice"`
	// memory_gib_per_vcpu block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#memory_gib_per_vcpu EcsCapacityProvider#memory_gib_per_vcpu}
	MemoryGibPerVcpu *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirementsMemoryGibPerVcpu `field:"optional" json:"memoryGibPerVcpu" yaml:"memoryGibPerVcpu"`
	// network_bandwidth_gbps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#network_bandwidth_gbps EcsCapacityProvider#network_bandwidth_gbps}
	NetworkBandwidthGbps *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirementsNetworkBandwidthGbps `field:"optional" json:"networkBandwidthGbps" yaml:"networkBandwidthGbps"`
	// network_interface_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#network_interface_count EcsCapacityProvider#network_interface_count}
	NetworkInterfaceCount *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirementsNetworkInterfaceCount `field:"optional" json:"networkInterfaceCount" yaml:"networkInterfaceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#on_demand_max_price_percentage_over_lowest_price EcsCapacityProvider#on_demand_max_price_percentage_over_lowest_price}.
	OnDemandMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"onDemandMaxPricePercentageOverLowestPrice" yaml:"onDemandMaxPricePercentageOverLowestPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#require_hibernate_support EcsCapacityProvider#require_hibernate_support}.
	RequireHibernateSupport interface{} `field:"optional" json:"requireHibernateSupport" yaml:"requireHibernateSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#spot_max_price_percentage_over_lowest_price EcsCapacityProvider#spot_max_price_percentage_over_lowest_price}.
	SpotMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"spotMaxPricePercentageOverLowestPrice" yaml:"spotMaxPricePercentageOverLowestPrice"`
	// total_local_storage_gb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#total_local_storage_gb EcsCapacityProvider#total_local_storage_gb}
	TotalLocalStorageGb *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirementsTotalLocalStorageGb `field:"optional" json:"totalLocalStorageGb" yaml:"totalLocalStorageGb"`
}

