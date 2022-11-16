package autoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/autoscalinggroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference interface {
	cdktf.ComplexObject
	AcceleratorCount() AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorCountOutputReference
	AcceleratorCountInput() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorCount
	AcceleratorManufacturers() *[]*string
	SetAcceleratorManufacturers(val *[]*string)
	AcceleratorManufacturersInput() *[]*string
	AcceleratorNames() *[]*string
	SetAcceleratorNames(val *[]*string)
	AcceleratorNamesInput() *[]*string
	AcceleratorTotalMemoryMib() AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorTotalMemoryMibOutputReference
	AcceleratorTotalMemoryMibInput() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorTotalMemoryMib
	AcceleratorTypes() *[]*string
	SetAcceleratorTypes(val *[]*string)
	AcceleratorTypesInput() *[]*string
	BareMetal() *string
	SetBareMetal(val *string)
	BareMetalInput() *string
	BaselineEbsBandwidthMbps() AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference
	BaselineEbsBandwidthMbpsInput() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsBaselineEbsBandwidthMbps
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
	InternalValue() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirements
	SetInternalValue(val *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirements)
	LocalStorage() *string
	SetLocalStorage(val *string)
	LocalStorageInput() *string
	LocalStorageTypes() *[]*string
	SetLocalStorageTypes(val *[]*string)
	LocalStorageTypesInput() *[]*string
	MemoryGibPerVcpu() AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryGibPerVcpuOutputReference
	MemoryGibPerVcpuInput() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryGibPerVcpu
	MemoryMib() AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference
	MemoryMibInput() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMib
	NetworkInterfaceCount() AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsNetworkInterfaceCountOutputReference
	NetworkInterfaceCountInput() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsNetworkInterfaceCount
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
	TotalLocalStorageGb() AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsTotalLocalStorageGbOutputReference
	TotalLocalStorageGbInput() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsTotalLocalStorageGb
	VcpuCount() AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsVcpuCountOutputReference
	VcpuCountInput() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsVcpuCount
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
	PutAcceleratorCount(value *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorCount)
	PutAcceleratorTotalMemoryMib(value *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorTotalMemoryMib)
	PutBaselineEbsBandwidthMbps(value *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsBaselineEbsBandwidthMbps)
	PutMemoryGibPerVcpu(value *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryGibPerVcpu)
	PutMemoryMib(value *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMib)
	PutNetworkInterfaceCount(value *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsNetworkInterfaceCount)
	PutTotalLocalStorageGb(value *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsTotalLocalStorageGb)
	PutVcpuCount(value *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsVcpuCount)
	ResetAcceleratorCount()
	ResetAcceleratorManufacturers()
	ResetAcceleratorNames()
	ResetAcceleratorTotalMemoryMib()
	ResetAcceleratorTypes()
	ResetBareMetal()
	ResetBaselineEbsBandwidthMbps()
	ResetBurstablePerformance()
	ResetCpuManufacturers()
	ResetExcludedInstanceTypes()
	ResetInstanceGenerations()
	ResetLocalStorage()
	ResetLocalStorageTypes()
	ResetMemoryGibPerVcpu()
	ResetMemoryMib()
	ResetNetworkInterfaceCount()
	ResetOnDemandMaxPricePercentageOverLowestPrice()
	ResetRequireHibernateSupport()
	ResetSpotMaxPricePercentageOverLowestPrice()
	ResetTotalLocalStorageGb()
	ResetVcpuCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference
type jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) AcceleratorCount() AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorCountOutputReference {
	var returns AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorCountOutputReference
	_jsii_.Get(
		j,
		"acceleratorCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) AcceleratorCountInput() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorCount {
	var returns *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorCount
	_jsii_.Get(
		j,
		"acceleratorCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) AcceleratorManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) AcceleratorManufacturersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorManufacturersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) AcceleratorNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) AcceleratorNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) AcceleratorTotalMemoryMib() AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorTotalMemoryMibOutputReference {
	var returns AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorTotalMemoryMibOutputReference
	_jsii_.Get(
		j,
		"acceleratorTotalMemoryMib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) AcceleratorTotalMemoryMibInput() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorTotalMemoryMib {
	var returns *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorTotalMemoryMib
	_jsii_.Get(
		j,
		"acceleratorTotalMemoryMibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) AcceleratorTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) AcceleratorTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) BareMetal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) BareMetalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) BaselineEbsBandwidthMbps() AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference {
	var returns AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference
	_jsii_.Get(
		j,
		"baselineEbsBandwidthMbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) BaselineEbsBandwidthMbpsInput() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsBaselineEbsBandwidthMbps {
	var returns *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsBaselineEbsBandwidthMbps
	_jsii_.Get(
		j,
		"baselineEbsBandwidthMbpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) BurstablePerformance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"burstablePerformance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) BurstablePerformanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"burstablePerformanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) CpuManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cpuManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) CpuManufacturersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cpuManufacturersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ExcludedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ExcludedInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) InstanceGenerations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGenerations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) InstanceGenerationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGenerationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) InternalValue() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirements {
	var returns *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirements
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) LocalStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) LocalStorageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) LocalStorageTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localStorageTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) LocalStorageTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localStorageTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) MemoryGibPerVcpu() AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryGibPerVcpuOutputReference {
	var returns AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryGibPerVcpuOutputReference
	_jsii_.Get(
		j,
		"memoryGibPerVcpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) MemoryGibPerVcpuInput() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryGibPerVcpu {
	var returns *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryGibPerVcpu
	_jsii_.Get(
		j,
		"memoryGibPerVcpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) MemoryMib() AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference {
	var returns AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference
	_jsii_.Get(
		j,
		"memoryMib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) MemoryMibInput() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMib {
	var returns *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMib
	_jsii_.Get(
		j,
		"memoryMibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) NetworkInterfaceCount() AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsNetworkInterfaceCountOutputReference {
	var returns AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsNetworkInterfaceCountOutputReference
	_jsii_.Get(
		j,
		"networkInterfaceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) NetworkInterfaceCountInput() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsNetworkInterfaceCount {
	var returns *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsNetworkInterfaceCount
	_jsii_.Get(
		j,
		"networkInterfaceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) OnDemandMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) OnDemandMaxPricePercentageOverLowestPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandMaxPricePercentageOverLowestPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) RequireHibernateSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHibernateSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) RequireHibernateSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHibernateSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) SpotMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) SpotMaxPricePercentageOverLowestPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotMaxPricePercentageOverLowestPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) TotalLocalStorageGb() AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsTotalLocalStorageGbOutputReference {
	var returns AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsTotalLocalStorageGbOutputReference
	_jsii_.Get(
		j,
		"totalLocalStorageGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) TotalLocalStorageGbInput() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsTotalLocalStorageGb {
	var returns *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsTotalLocalStorageGb
	_jsii_.Get(
		j,
		"totalLocalStorageGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) VcpuCount() AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsVcpuCountOutputReference {
	var returns AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsVcpuCountOutputReference
	_jsii_.Get(
		j,
		"vcpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) VcpuCountInput() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsVcpuCount {
	var returns *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsVcpuCount
	_jsii_.Get(
		j,
		"vcpuCountInput",
		&returns,
	)
	return returns
}


func NewAutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference_Override(a AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetAcceleratorManufacturers(val *[]*string) {
	if err := j.validateSetAcceleratorManufacturersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorManufacturers",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetAcceleratorNames(val *[]*string) {
	if err := j.validateSetAcceleratorNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorNames",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetAcceleratorTypes(val *[]*string) {
	if err := j.validateSetAcceleratorTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorTypes",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetBareMetal(val *string) {
	if err := j.validateSetBareMetalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bareMetal",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetBurstablePerformance(val *string) {
	if err := j.validateSetBurstablePerformanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"burstablePerformance",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetCpuManufacturers(val *[]*string) {
	if err := j.validateSetCpuManufacturersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuManufacturers",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetExcludedInstanceTypes(val *[]*string) {
	if err := j.validateSetExcludedInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetInstanceGenerations(val *[]*string) {
	if err := j.validateSetInstanceGenerationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceGenerations",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetInternalValue(val *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirements) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetLocalStorage(val *string) {
	if err := j.validateSetLocalStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localStorage",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetLocalStorageTypes(val *[]*string) {
	if err := j.validateSetLocalStorageTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localStorageTypes",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetOnDemandMaxPricePercentageOverLowestPrice(val *float64) {
	if err := j.validateSetOnDemandMaxPricePercentageOverLowestPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandMaxPricePercentageOverLowestPrice",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetRequireHibernateSupport(val interface{}) {
	if err := j.validateSetRequireHibernateSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireHibernateSupport",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetSpotMaxPricePercentageOverLowestPrice(val *float64) {
	if err := j.validateSetSpotMaxPricePercentageOverLowestPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotMaxPricePercentageOverLowestPrice",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) PutAcceleratorCount(value *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorCount) {
	if err := a.validatePutAcceleratorCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAcceleratorCount",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) PutAcceleratorTotalMemoryMib(value *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorTotalMemoryMib) {
	if err := a.validatePutAcceleratorTotalMemoryMibParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAcceleratorTotalMemoryMib",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) PutBaselineEbsBandwidthMbps(value *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsBaselineEbsBandwidthMbps) {
	if err := a.validatePutBaselineEbsBandwidthMbpsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBaselineEbsBandwidthMbps",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) PutMemoryGibPerVcpu(value *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryGibPerVcpu) {
	if err := a.validatePutMemoryGibPerVcpuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMemoryGibPerVcpu",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) PutMemoryMib(value *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMib) {
	if err := a.validatePutMemoryMibParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMemoryMib",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) PutNetworkInterfaceCount(value *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsNetworkInterfaceCount) {
	if err := a.validatePutNetworkInterfaceCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkInterfaceCount",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) PutTotalLocalStorageGb(value *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsTotalLocalStorageGb) {
	if err := a.validatePutTotalLocalStorageGbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTotalLocalStorageGb",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) PutVcpuCount(value *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsVcpuCount) {
	if err := a.validatePutVcpuCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVcpuCount",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetAcceleratorCount() {
	_jsii_.InvokeVoid(
		a,
		"resetAcceleratorCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetAcceleratorManufacturers() {
	_jsii_.InvokeVoid(
		a,
		"resetAcceleratorManufacturers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetAcceleratorNames() {
	_jsii_.InvokeVoid(
		a,
		"resetAcceleratorNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetAcceleratorTotalMemoryMib() {
	_jsii_.InvokeVoid(
		a,
		"resetAcceleratorTotalMemoryMib",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetAcceleratorTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetAcceleratorTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetBareMetal() {
	_jsii_.InvokeVoid(
		a,
		"resetBareMetal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetBaselineEbsBandwidthMbps() {
	_jsii_.InvokeVoid(
		a,
		"resetBaselineEbsBandwidthMbps",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetBurstablePerformance() {
	_jsii_.InvokeVoid(
		a,
		"resetBurstablePerformance",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetCpuManufacturers() {
	_jsii_.InvokeVoid(
		a,
		"resetCpuManufacturers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetExcludedInstanceTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludedInstanceTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetInstanceGenerations() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceGenerations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetLocalStorage() {
	_jsii_.InvokeVoid(
		a,
		"resetLocalStorage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetLocalStorageTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetLocalStorageTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetMemoryGibPerVcpu() {
	_jsii_.InvokeVoid(
		a,
		"resetMemoryGibPerVcpu",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetMemoryMib() {
	_jsii_.InvokeVoid(
		a,
		"resetMemoryMib",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetNetworkInterfaceCount() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkInterfaceCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetOnDemandMaxPricePercentageOverLowestPrice() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandMaxPricePercentageOverLowestPrice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetRequireHibernateSupport() {
	_jsii_.InvokeVoid(
		a,
		"resetRequireHibernateSupport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetSpotMaxPricePercentageOverLowestPrice() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotMaxPricePercentageOverLowestPrice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetTotalLocalStorageGb() {
	_jsii_.InvokeVoid(
		a,
		"resetTotalLocalStorageGb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ResetVcpuCount() {
	_jsii_.InvokeVoid(
		a,
		"resetVcpuCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

