package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/ec2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference interface {
	cdktf.ComplexObject
	AcceleratorCount() SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsAcceleratorCountOutputReference
	AcceleratorCountInput() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsAcceleratorCount
	AcceleratorManufacturers() *[]*string
	SetAcceleratorManufacturers(val *[]*string)
	AcceleratorManufacturersInput() *[]*string
	AcceleratorNames() *[]*string
	SetAcceleratorNames(val *[]*string)
	AcceleratorNamesInput() *[]*string
	AcceleratorTotalMemoryMib() SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsAcceleratorTotalMemoryMibOutputReference
	AcceleratorTotalMemoryMibInput() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsAcceleratorTotalMemoryMib
	AcceleratorTypes() *[]*string
	SetAcceleratorTypes(val *[]*string)
	AcceleratorTypesInput() *[]*string
	BareMetal() *string
	SetBareMetal(val *string)
	BareMetalInput() *string
	BaselineEbsBandwidthMbps() SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference
	BaselineEbsBandwidthMbpsInput() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsBaselineEbsBandwidthMbps
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
	InternalValue() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirements
	SetInternalValue(val *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirements)
	LocalStorage() *string
	SetLocalStorage(val *string)
	LocalStorageInput() *string
	LocalStorageTypes() *[]*string
	SetLocalStorageTypes(val *[]*string)
	LocalStorageTypesInput() *[]*string
	MemoryGibPerVcpu() SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsMemoryGibPerVcpuOutputReference
	MemoryGibPerVcpuInput() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsMemoryGibPerVcpu
	MemoryMib() SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsMemoryMibOutputReference
	MemoryMibInput() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsMemoryMib
	NetworkInterfaceCount() SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsNetworkInterfaceCountOutputReference
	NetworkInterfaceCountInput() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsNetworkInterfaceCount
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
	TotalLocalStorageGb() SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsTotalLocalStorageGbOutputReference
	TotalLocalStorageGbInput() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsTotalLocalStorageGb
	VcpuCount() SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsVcpuCountOutputReference
	VcpuCountInput() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsVcpuCount
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
	PutAcceleratorCount(value *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsAcceleratorCount)
	PutAcceleratorTotalMemoryMib(value *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsAcceleratorTotalMemoryMib)
	PutBaselineEbsBandwidthMbps(value *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsBaselineEbsBandwidthMbps)
	PutMemoryGibPerVcpu(value *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsMemoryGibPerVcpu)
	PutMemoryMib(value *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsMemoryMib)
	PutNetworkInterfaceCount(value *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsNetworkInterfaceCount)
	PutTotalLocalStorageGb(value *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsTotalLocalStorageGb)
	PutVcpuCount(value *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsVcpuCount)
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

// The jsii proxy struct for SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference
type jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) AcceleratorCount() SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsAcceleratorCountOutputReference {
	var returns SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsAcceleratorCountOutputReference
	_jsii_.Get(
		j,
		"acceleratorCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) AcceleratorCountInput() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsAcceleratorCount {
	var returns *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsAcceleratorCount
	_jsii_.Get(
		j,
		"acceleratorCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) AcceleratorManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) AcceleratorManufacturersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorManufacturersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) AcceleratorNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) AcceleratorNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) AcceleratorTotalMemoryMib() SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsAcceleratorTotalMemoryMibOutputReference {
	var returns SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsAcceleratorTotalMemoryMibOutputReference
	_jsii_.Get(
		j,
		"acceleratorTotalMemoryMib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) AcceleratorTotalMemoryMibInput() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsAcceleratorTotalMemoryMib {
	var returns *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsAcceleratorTotalMemoryMib
	_jsii_.Get(
		j,
		"acceleratorTotalMemoryMibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) AcceleratorTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) AcceleratorTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) BareMetal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) BareMetalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) BaselineEbsBandwidthMbps() SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference {
	var returns SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference
	_jsii_.Get(
		j,
		"baselineEbsBandwidthMbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) BaselineEbsBandwidthMbpsInput() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsBaselineEbsBandwidthMbps {
	var returns *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsBaselineEbsBandwidthMbps
	_jsii_.Get(
		j,
		"baselineEbsBandwidthMbpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) BurstablePerformance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"burstablePerformance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) BurstablePerformanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"burstablePerformanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) CpuManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cpuManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) CpuManufacturersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cpuManufacturersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ExcludedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ExcludedInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) InstanceGenerations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGenerations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) InstanceGenerationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGenerationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) InternalValue() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirements {
	var returns *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirements
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) LocalStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) LocalStorageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) LocalStorageTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localStorageTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) LocalStorageTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localStorageTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) MemoryGibPerVcpu() SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsMemoryGibPerVcpuOutputReference {
	var returns SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsMemoryGibPerVcpuOutputReference
	_jsii_.Get(
		j,
		"memoryGibPerVcpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) MemoryGibPerVcpuInput() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsMemoryGibPerVcpu {
	var returns *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsMemoryGibPerVcpu
	_jsii_.Get(
		j,
		"memoryGibPerVcpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) MemoryMib() SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsMemoryMibOutputReference {
	var returns SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsMemoryMibOutputReference
	_jsii_.Get(
		j,
		"memoryMib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) MemoryMibInput() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsMemoryMib {
	var returns *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsMemoryMib
	_jsii_.Get(
		j,
		"memoryMibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) NetworkInterfaceCount() SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsNetworkInterfaceCountOutputReference {
	var returns SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsNetworkInterfaceCountOutputReference
	_jsii_.Get(
		j,
		"networkInterfaceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) NetworkInterfaceCountInput() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsNetworkInterfaceCount {
	var returns *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsNetworkInterfaceCount
	_jsii_.Get(
		j,
		"networkInterfaceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) OnDemandMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) OnDemandMaxPricePercentageOverLowestPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandMaxPricePercentageOverLowestPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) RequireHibernateSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHibernateSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) RequireHibernateSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHibernateSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) SpotMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) SpotMaxPricePercentageOverLowestPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotMaxPricePercentageOverLowestPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) TotalLocalStorageGb() SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsTotalLocalStorageGbOutputReference {
	var returns SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsTotalLocalStorageGbOutputReference
	_jsii_.Get(
		j,
		"totalLocalStorageGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) TotalLocalStorageGbInput() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsTotalLocalStorageGb {
	var returns *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsTotalLocalStorageGb
	_jsii_.Get(
		j,
		"totalLocalStorageGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) VcpuCount() SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsVcpuCountOutputReference {
	var returns SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsVcpuCountOutputReference
	_jsii_.Get(
		j,
		"vcpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) VcpuCountInput() *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsVcpuCount {
	var returns *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsVcpuCount
	_jsii_.Get(
		j,
		"vcpuCountInput",
		&returns,
	)
	return returns
}


func NewSpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference {
	_init_.Initialize()

	if err := validateNewSpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2.SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference_Override(s SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2.SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetAcceleratorManufacturers(val *[]*string) {
	if err := j.validateSetAcceleratorManufacturersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorManufacturers",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetAcceleratorNames(val *[]*string) {
	if err := j.validateSetAcceleratorNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorNames",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetAcceleratorTypes(val *[]*string) {
	if err := j.validateSetAcceleratorTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorTypes",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetBareMetal(val *string) {
	if err := j.validateSetBareMetalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bareMetal",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetBurstablePerformance(val *string) {
	if err := j.validateSetBurstablePerformanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"burstablePerformance",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetCpuManufacturers(val *[]*string) {
	if err := j.validateSetCpuManufacturersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuManufacturers",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetExcludedInstanceTypes(val *[]*string) {
	if err := j.validateSetExcludedInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetInstanceGenerations(val *[]*string) {
	if err := j.validateSetInstanceGenerationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceGenerations",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetInternalValue(val *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirements) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetLocalStorage(val *string) {
	if err := j.validateSetLocalStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localStorage",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetLocalStorageTypes(val *[]*string) {
	if err := j.validateSetLocalStorageTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localStorageTypes",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetOnDemandMaxPricePercentageOverLowestPrice(val *float64) {
	if err := j.validateSetOnDemandMaxPricePercentageOverLowestPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandMaxPricePercentageOverLowestPrice",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetRequireHibernateSupport(val interface{}) {
	if err := j.validateSetRequireHibernateSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireHibernateSupport",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetSpotMaxPricePercentageOverLowestPrice(val *float64) {
	if err := j.validateSetSpotMaxPricePercentageOverLowestPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotMaxPricePercentageOverLowestPrice",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) PutAcceleratorCount(value *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsAcceleratorCount) {
	if err := s.validatePutAcceleratorCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAcceleratorCount",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) PutAcceleratorTotalMemoryMib(value *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsAcceleratorTotalMemoryMib) {
	if err := s.validatePutAcceleratorTotalMemoryMibParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAcceleratorTotalMemoryMib",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) PutBaselineEbsBandwidthMbps(value *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsBaselineEbsBandwidthMbps) {
	if err := s.validatePutBaselineEbsBandwidthMbpsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBaselineEbsBandwidthMbps",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) PutMemoryGibPerVcpu(value *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsMemoryGibPerVcpu) {
	if err := s.validatePutMemoryGibPerVcpuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMemoryGibPerVcpu",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) PutMemoryMib(value *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsMemoryMib) {
	if err := s.validatePutMemoryMibParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMemoryMib",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) PutNetworkInterfaceCount(value *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsNetworkInterfaceCount) {
	if err := s.validatePutNetworkInterfaceCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkInterfaceCount",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) PutTotalLocalStorageGb(value *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsTotalLocalStorageGb) {
	if err := s.validatePutTotalLocalStorageGbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTotalLocalStorageGb",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) PutVcpuCount(value *SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsVcpuCount) {
	if err := s.validatePutVcpuCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVcpuCount",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetAcceleratorCount() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceleratorCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetAcceleratorManufacturers() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceleratorManufacturers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetAcceleratorNames() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceleratorNames",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetAcceleratorTotalMemoryMib() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceleratorTotalMemoryMib",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetAcceleratorTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceleratorTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetBareMetal() {
	_jsii_.InvokeVoid(
		s,
		"resetBareMetal",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetBaselineEbsBandwidthMbps() {
	_jsii_.InvokeVoid(
		s,
		"resetBaselineEbsBandwidthMbps",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetBurstablePerformance() {
	_jsii_.InvokeVoid(
		s,
		"resetBurstablePerformance",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetCpuManufacturers() {
	_jsii_.InvokeVoid(
		s,
		"resetCpuManufacturers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetExcludedInstanceTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludedInstanceTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetInstanceGenerations() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceGenerations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetLocalStorage() {
	_jsii_.InvokeVoid(
		s,
		"resetLocalStorage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetLocalStorageTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetLocalStorageTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetMemoryGibPerVcpu() {
	_jsii_.InvokeVoid(
		s,
		"resetMemoryGibPerVcpu",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetMemoryMib() {
	_jsii_.InvokeVoid(
		s,
		"resetMemoryMib",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetNetworkInterfaceCount() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkInterfaceCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetOnDemandMaxPricePercentageOverLowestPrice() {
	_jsii_.InvokeVoid(
		s,
		"resetOnDemandMaxPricePercentageOverLowestPrice",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetRequireHibernateSupport() {
	_jsii_.InvokeVoid(
		s,
		"resetRequireHibernateSupport",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetSpotMaxPricePercentageOverLowestPrice() {
	_jsii_.InvokeVoid(
		s,
		"resetSpotMaxPricePercentageOverLowestPrice",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetTotalLocalStorageGb() {
	_jsii_.InvokeVoid(
		s,
		"resetTotalLocalStorageGb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ResetVcpuCount() {
	_jsii_.InvokeVoid(
		s,
		"resetVcpuCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

