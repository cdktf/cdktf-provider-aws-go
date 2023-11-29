// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2fleet


type Ec2FleetLaunchTemplateConfigOverrideInstanceRequirements struct {
	// memory_mib block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#memory_mib Ec2Fleet#memory_mib}
	MemoryMib *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryMib `field:"required" json:"memoryMib" yaml:"memoryMib"`
	// vcpu_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#vcpu_count Ec2Fleet#vcpu_count}
	VcpuCount *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsVcpuCount `field:"required" json:"vcpuCount" yaml:"vcpuCount"`
	// accelerator_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#accelerator_count Ec2Fleet#accelerator_count}
	AcceleratorCount *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorCount `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#accelerator_manufacturers Ec2Fleet#accelerator_manufacturers}.
	AcceleratorManufacturers *[]*string `field:"optional" json:"acceleratorManufacturers" yaml:"acceleratorManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#accelerator_names Ec2Fleet#accelerator_names}.
	AcceleratorNames *[]*string `field:"optional" json:"acceleratorNames" yaml:"acceleratorNames"`
	// accelerator_total_memory_mib block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#accelerator_total_memory_mib Ec2Fleet#accelerator_total_memory_mib}
	AcceleratorTotalMemoryMib *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorTotalMemoryMib `field:"optional" json:"acceleratorTotalMemoryMib" yaml:"acceleratorTotalMemoryMib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#accelerator_types Ec2Fleet#accelerator_types}.
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#allowed_instance_types Ec2Fleet#allowed_instance_types}.
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#bare_metal Ec2Fleet#bare_metal}.
	BareMetal *string `field:"optional" json:"bareMetal" yaml:"bareMetal"`
	// baseline_ebs_bandwidth_mbps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#baseline_ebs_bandwidth_mbps Ec2Fleet#baseline_ebs_bandwidth_mbps}
	BaselineEbsBandwidthMbps *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsBaselineEbsBandwidthMbps `field:"optional" json:"baselineEbsBandwidthMbps" yaml:"baselineEbsBandwidthMbps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#burstable_performance Ec2Fleet#burstable_performance}.
	BurstablePerformance *string `field:"optional" json:"burstablePerformance" yaml:"burstablePerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#cpu_manufacturers Ec2Fleet#cpu_manufacturers}.
	CpuManufacturers *[]*string `field:"optional" json:"cpuManufacturers" yaml:"cpuManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#excluded_instance_types Ec2Fleet#excluded_instance_types}.
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#instance_generations Ec2Fleet#instance_generations}.
	InstanceGenerations *[]*string `field:"optional" json:"instanceGenerations" yaml:"instanceGenerations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#local_storage Ec2Fleet#local_storage}.
	LocalStorage *string `field:"optional" json:"localStorage" yaml:"localStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#local_storage_types Ec2Fleet#local_storage_types}.
	LocalStorageTypes *[]*string `field:"optional" json:"localStorageTypes" yaml:"localStorageTypes"`
	// memory_gib_per_vcpu block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#memory_gib_per_vcpu Ec2Fleet#memory_gib_per_vcpu}
	MemoryGibPerVcpu *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryGibPerVcpu `field:"optional" json:"memoryGibPerVcpu" yaml:"memoryGibPerVcpu"`
	// network_bandwidth_gbps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#network_bandwidth_gbps Ec2Fleet#network_bandwidth_gbps}
	NetworkBandwidthGbps *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkBandwidthGbps `field:"optional" json:"networkBandwidthGbps" yaml:"networkBandwidthGbps"`
	// network_interface_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#network_interface_count Ec2Fleet#network_interface_count}
	NetworkInterfaceCount *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkInterfaceCount `field:"optional" json:"networkInterfaceCount" yaml:"networkInterfaceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#on_demand_max_price_percentage_over_lowest_price Ec2Fleet#on_demand_max_price_percentage_over_lowest_price}.
	OnDemandMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"onDemandMaxPricePercentageOverLowestPrice" yaml:"onDemandMaxPricePercentageOverLowestPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#require_hibernate_support Ec2Fleet#require_hibernate_support}.
	RequireHibernateSupport interface{} `field:"optional" json:"requireHibernateSupport" yaml:"requireHibernateSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#spot_max_price_percentage_over_lowest_price Ec2Fleet#spot_max_price_percentage_over_lowest_price}.
	SpotMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"spotMaxPricePercentageOverLowestPrice" yaml:"spotMaxPricePercentageOverLowestPrice"`
	// total_local_storage_gb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ec2_fleet#total_local_storage_gb Ec2Fleet#total_local_storage_gb}
	TotalLocalStorageGb *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsTotalLocalStorageGb `field:"optional" json:"totalLocalStorageGb" yaml:"totalLocalStorageGb"`
}

