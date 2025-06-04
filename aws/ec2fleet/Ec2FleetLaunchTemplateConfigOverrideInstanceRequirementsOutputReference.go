// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2fleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/ec2fleet/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference interface {
	cdktf.ComplexObject
	AcceleratorCount() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorCountOutputReference
	AcceleratorCountInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorCount
	AcceleratorManufacturers() *[]*string
	SetAcceleratorManufacturers(val *[]*string)
	AcceleratorManufacturersInput() *[]*string
	AcceleratorNames() *[]*string
	SetAcceleratorNames(val *[]*string)
	AcceleratorNamesInput() *[]*string
	AcceleratorTotalMemoryMib() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorTotalMemoryMibOutputReference
	AcceleratorTotalMemoryMibInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorTotalMemoryMib
	AcceleratorTypes() *[]*string
	SetAcceleratorTypes(val *[]*string)
	AcceleratorTypesInput() *[]*string
	AllowedInstanceTypes() *[]*string
	SetAllowedInstanceTypes(val *[]*string)
	AllowedInstanceTypesInput() *[]*string
	BareMetal() *string
	SetBareMetal(val *string)
	BareMetalInput() *string
	BaselineEbsBandwidthMbps() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference
	BaselineEbsBandwidthMbpsInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsBaselineEbsBandwidthMbps
	BurstablePerformance() *string
	SetBurstablePerformance(val *string)
	BurstablePerformanceInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	CpuManufacturers() *[]*string
	SetCpuManufacturers(val *[]*string)
	CpuManufacturersInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExcludedInstanceTypes() *[]*string
	SetExcludedInstanceTypes(val *[]*string)
	ExcludedInstanceTypesInput() *[]*string
	// Experimental.
	Fqn() *string
	InstanceGenerations() *[]*string
	SetInstanceGenerations(val *[]*string)
	InstanceGenerationsInput() *[]*string
	InternalValue() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirements
	SetInternalValue(val *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirements)
	LocalStorage() *string
	SetLocalStorage(val *string)
	LocalStorageInput() *string
	LocalStorageTypes() *[]*string
	SetLocalStorageTypes(val *[]*string)
	LocalStorageTypesInput() *[]*string
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice() *float64
	SetMaxSpotPriceAsPercentageOfOptimalOnDemandPrice(val *float64)
	MaxSpotPriceAsPercentageOfOptimalOnDemandPriceInput() *float64
	MemoryGibPerVcpu() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryGibPerVcpuOutputReference
	MemoryGibPerVcpuInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryGibPerVcpu
	MemoryMib() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryMibOutputReference
	MemoryMibInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryMib
	NetworkBandwidthGbps() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkBandwidthGbpsOutputReference
	NetworkBandwidthGbpsInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkBandwidthGbps
	NetworkInterfaceCount() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkInterfaceCountOutputReference
	NetworkInterfaceCountInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkInterfaceCount
	OnDemandMaxPricePercentageOverLowestPrice() *float64
	SetOnDemandMaxPricePercentageOverLowestPrice(val *float64)
	OnDemandMaxPricePercentageOverLowestPriceInput() *float64
	RequireHibernateSupport() interface{}
	SetRequireHibernateSupport(val interface{})
	RequireHibernateSupportInput() interface{}
	SpotMaxPricePercentageOverLowestPrice() *float64
	SetSpotMaxPricePercentageOverLowestPrice(val *float64)
	SpotMaxPricePercentageOverLowestPriceInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TotalLocalStorageGb() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsTotalLocalStorageGbOutputReference
	TotalLocalStorageGbInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsTotalLocalStorageGb
	VcpuCount() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsVcpuCountOutputReference
	VcpuCountInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsVcpuCount
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAcceleratorCount(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorCount)
	PutAcceleratorTotalMemoryMib(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorTotalMemoryMib)
	PutBaselineEbsBandwidthMbps(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsBaselineEbsBandwidthMbps)
	PutMemoryGibPerVcpu(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryGibPerVcpu)
	PutMemoryMib(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryMib)
	PutNetworkBandwidthGbps(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkBandwidthGbps)
	PutNetworkInterfaceCount(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkInterfaceCount)
	PutTotalLocalStorageGb(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsTotalLocalStorageGb)
	PutVcpuCount(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsVcpuCount)
	ResetAcceleratorCount()
	ResetAcceleratorManufacturers()
	ResetAcceleratorNames()
	ResetAcceleratorTotalMemoryMib()
	ResetAcceleratorTypes()
	ResetAllowedInstanceTypes()
	ResetBareMetal()
	ResetBaselineEbsBandwidthMbps()
	ResetBurstablePerformance()
	ResetCpuManufacturers()
	ResetExcludedInstanceTypes()
	ResetInstanceGenerations()
	ResetLocalStorage()
	ResetLocalStorageTypes()
	ResetMaxSpotPriceAsPercentageOfOptimalOnDemandPrice()
	ResetMemoryGibPerVcpu()
	ResetNetworkBandwidthGbps()
	ResetNetworkInterfaceCount()
	ResetOnDemandMaxPricePercentageOverLowestPrice()
	ResetRequireHibernateSupport()
	ResetSpotMaxPricePercentageOverLowestPrice()
	ResetTotalLocalStorageGb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference
type jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) AcceleratorCount() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorCountOutputReference {
	var returns Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorCountOutputReference
	_jsii_.Get(
		j,
		"acceleratorCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) AcceleratorCountInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorCount {
	var returns *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorCount
	_jsii_.Get(
		j,
		"acceleratorCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) AcceleratorManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) AcceleratorManufacturersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorManufacturersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) AcceleratorNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) AcceleratorNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) AcceleratorTotalMemoryMib() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorTotalMemoryMibOutputReference {
	var returns Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorTotalMemoryMibOutputReference
	_jsii_.Get(
		j,
		"acceleratorTotalMemoryMib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) AcceleratorTotalMemoryMibInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorTotalMemoryMib {
	var returns *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorTotalMemoryMib
	_jsii_.Get(
		j,
		"acceleratorTotalMemoryMibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) AcceleratorTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) AcceleratorTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) AllowedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) AllowedInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) BareMetal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) BareMetalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) BaselineEbsBandwidthMbps() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference {
	var returns Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference
	_jsii_.Get(
		j,
		"baselineEbsBandwidthMbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) BaselineEbsBandwidthMbpsInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsBaselineEbsBandwidthMbps {
	var returns *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsBaselineEbsBandwidthMbps
	_jsii_.Get(
		j,
		"baselineEbsBandwidthMbpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) BurstablePerformance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"burstablePerformance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) BurstablePerformanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"burstablePerformanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) CpuManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cpuManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) CpuManufacturersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cpuManufacturersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ExcludedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ExcludedInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) InstanceGenerations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGenerations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) InstanceGenerationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGenerationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) InternalValue() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirements {
	var returns *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirements
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) LocalStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) LocalStorageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) LocalStorageTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localStorageTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) LocalStorageTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localStorageTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) MaxSpotPriceAsPercentageOfOptimalOnDemandPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSpotPriceAsPercentageOfOptimalOnDemandPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) MaxSpotPriceAsPercentageOfOptimalOnDemandPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSpotPriceAsPercentageOfOptimalOnDemandPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) MemoryGibPerVcpu() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryGibPerVcpuOutputReference {
	var returns Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryGibPerVcpuOutputReference
	_jsii_.Get(
		j,
		"memoryGibPerVcpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) MemoryGibPerVcpuInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryGibPerVcpu {
	var returns *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryGibPerVcpu
	_jsii_.Get(
		j,
		"memoryGibPerVcpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) MemoryMib() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryMibOutputReference {
	var returns Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryMibOutputReference
	_jsii_.Get(
		j,
		"memoryMib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) MemoryMibInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryMib {
	var returns *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryMib
	_jsii_.Get(
		j,
		"memoryMibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) NetworkBandwidthGbps() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkBandwidthGbpsOutputReference {
	var returns Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkBandwidthGbpsOutputReference
	_jsii_.Get(
		j,
		"networkBandwidthGbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) NetworkBandwidthGbpsInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkBandwidthGbps {
	var returns *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkBandwidthGbps
	_jsii_.Get(
		j,
		"networkBandwidthGbpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) NetworkInterfaceCount() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkInterfaceCountOutputReference {
	var returns Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkInterfaceCountOutputReference
	_jsii_.Get(
		j,
		"networkInterfaceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) NetworkInterfaceCountInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkInterfaceCount {
	var returns *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkInterfaceCount
	_jsii_.Get(
		j,
		"networkInterfaceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) OnDemandMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) OnDemandMaxPricePercentageOverLowestPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandMaxPricePercentageOverLowestPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) RequireHibernateSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHibernateSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) RequireHibernateSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHibernateSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) SpotMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) SpotMaxPricePercentageOverLowestPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotMaxPricePercentageOverLowestPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) TotalLocalStorageGb() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsTotalLocalStorageGbOutputReference {
	var returns Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsTotalLocalStorageGbOutputReference
	_jsii_.Get(
		j,
		"totalLocalStorageGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) TotalLocalStorageGbInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsTotalLocalStorageGb {
	var returns *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsTotalLocalStorageGb
	_jsii_.Get(
		j,
		"totalLocalStorageGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) VcpuCount() Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsVcpuCountOutputReference {
	var returns Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsVcpuCountOutputReference
	_jsii_.Get(
		j,
		"vcpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) VcpuCountInput() *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsVcpuCount {
	var returns *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsVcpuCount
	_jsii_.Get(
		j,
		"vcpuCountInput",
		&returns,
	)
	return returns
}


func NewEc2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2Fleet.Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference_Override(e Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2Fleet.Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetAcceleratorManufacturers(val *[]*string) {
	if err := j.validateSetAcceleratorManufacturersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorManufacturers",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetAcceleratorNames(val *[]*string) {
	if err := j.validateSetAcceleratorNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorNames",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetAcceleratorTypes(val *[]*string) {
	if err := j.validateSetAcceleratorTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorTypes",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetAllowedInstanceTypes(val *[]*string) {
	if err := j.validateSetAllowedInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetBareMetal(val *string) {
	if err := j.validateSetBareMetalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bareMetal",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetBurstablePerformance(val *string) {
	if err := j.validateSetBurstablePerformanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"burstablePerformance",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetCpuManufacturers(val *[]*string) {
	if err := j.validateSetCpuManufacturersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuManufacturers",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetExcludedInstanceTypes(val *[]*string) {
	if err := j.validateSetExcludedInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetInstanceGenerations(val *[]*string) {
	if err := j.validateSetInstanceGenerationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceGenerations",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetInternalValue(val *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirements) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetLocalStorage(val *string) {
	if err := j.validateSetLocalStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localStorage",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetLocalStorageTypes(val *[]*string) {
	if err := j.validateSetLocalStorageTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localStorageTypes",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetMaxSpotPriceAsPercentageOfOptimalOnDemandPrice(val *float64) {
	if err := j.validateSetMaxSpotPriceAsPercentageOfOptimalOnDemandPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSpotPriceAsPercentageOfOptimalOnDemandPrice",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetOnDemandMaxPricePercentageOverLowestPrice(val *float64) {
	if err := j.validateSetOnDemandMaxPricePercentageOverLowestPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandMaxPricePercentageOverLowestPrice",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetRequireHibernateSupport(val interface{}) {
	if err := j.validateSetRequireHibernateSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireHibernateSupport",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetSpotMaxPricePercentageOverLowestPrice(val *float64) {
	if err := j.validateSetSpotMaxPricePercentageOverLowestPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotMaxPricePercentageOverLowestPrice",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) PutAcceleratorCount(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorCount) {
	if err := e.validatePutAcceleratorCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAcceleratorCount",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) PutAcceleratorTotalMemoryMib(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorTotalMemoryMib) {
	if err := e.validatePutAcceleratorTotalMemoryMibParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAcceleratorTotalMemoryMib",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) PutBaselineEbsBandwidthMbps(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsBaselineEbsBandwidthMbps) {
	if err := e.validatePutBaselineEbsBandwidthMbpsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putBaselineEbsBandwidthMbps",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) PutMemoryGibPerVcpu(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryGibPerVcpu) {
	if err := e.validatePutMemoryGibPerVcpuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMemoryGibPerVcpu",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) PutMemoryMib(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryMib) {
	if err := e.validatePutMemoryMibParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMemoryMib",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) PutNetworkBandwidthGbps(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkBandwidthGbps) {
	if err := e.validatePutNetworkBandwidthGbpsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkBandwidthGbps",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) PutNetworkInterfaceCount(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkInterfaceCount) {
	if err := e.validatePutNetworkInterfaceCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkInterfaceCount",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) PutTotalLocalStorageGb(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsTotalLocalStorageGb) {
	if err := e.validatePutTotalLocalStorageGbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTotalLocalStorageGb",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) PutVcpuCount(value *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsVcpuCount) {
	if err := e.validatePutVcpuCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVcpuCount",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetAcceleratorCount() {
	_jsii_.InvokeVoid(
		e,
		"resetAcceleratorCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetAcceleratorManufacturers() {
	_jsii_.InvokeVoid(
		e,
		"resetAcceleratorManufacturers",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetAcceleratorNames() {
	_jsii_.InvokeVoid(
		e,
		"resetAcceleratorNames",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetAcceleratorTotalMemoryMib() {
	_jsii_.InvokeVoid(
		e,
		"resetAcceleratorTotalMemoryMib",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetAcceleratorTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetAcceleratorTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetAllowedInstanceTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetAllowedInstanceTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetBareMetal() {
	_jsii_.InvokeVoid(
		e,
		"resetBareMetal",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetBaselineEbsBandwidthMbps() {
	_jsii_.InvokeVoid(
		e,
		"resetBaselineEbsBandwidthMbps",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetBurstablePerformance() {
	_jsii_.InvokeVoid(
		e,
		"resetBurstablePerformance",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetCpuManufacturers() {
	_jsii_.InvokeVoid(
		e,
		"resetCpuManufacturers",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetExcludedInstanceTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetExcludedInstanceTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetInstanceGenerations() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceGenerations",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetLocalStorage() {
	_jsii_.InvokeVoid(
		e,
		"resetLocalStorage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetLocalStorageTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetLocalStorageTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetMaxSpotPriceAsPercentageOfOptimalOnDemandPrice() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxSpotPriceAsPercentageOfOptimalOnDemandPrice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetMemoryGibPerVcpu() {
	_jsii_.InvokeVoid(
		e,
		"resetMemoryGibPerVcpu",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetNetworkBandwidthGbps() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkBandwidthGbps",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetNetworkInterfaceCount() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkInterfaceCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetOnDemandMaxPricePercentageOverLowestPrice() {
	_jsii_.InvokeVoid(
		e,
		"resetOnDemandMaxPricePercentageOverLowestPrice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetRequireHibernateSupport() {
	_jsii_.InvokeVoid(
		e,
		"resetRequireHibernateSupport",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetSpotMaxPricePercentageOverLowestPrice() {
	_jsii_.InvokeVoid(
		e,
		"resetSpotMaxPricePercentageOverLowestPrice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ResetTotalLocalStorageGb() {
	_jsii_.InvokeVoid(
		e,
		"resetTotalLocalStorageGb",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

