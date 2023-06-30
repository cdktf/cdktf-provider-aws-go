package launchtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/launchtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LaunchTemplateInstanceRequirementsOutputReference interface {
	cdktf.ComplexObject
	AcceleratorCount() LaunchTemplateInstanceRequirementsAcceleratorCountOutputReference
	AcceleratorCountInput() *LaunchTemplateInstanceRequirementsAcceleratorCount
	AcceleratorManufacturers() *[]*string
	SetAcceleratorManufacturers(val *[]*string)
	AcceleratorManufacturersInput() *[]*string
	AcceleratorNames() *[]*string
	SetAcceleratorNames(val *[]*string)
	AcceleratorNamesInput() *[]*string
	AcceleratorTotalMemoryMib() LaunchTemplateInstanceRequirementsAcceleratorTotalMemoryMibOutputReference
	AcceleratorTotalMemoryMibInput() *LaunchTemplateInstanceRequirementsAcceleratorTotalMemoryMib
	AcceleratorTypes() *[]*string
	SetAcceleratorTypes(val *[]*string)
	AcceleratorTypesInput() *[]*string
	AllowedInstanceTypes() *[]*string
	SetAllowedInstanceTypes(val *[]*string)
	AllowedInstanceTypesInput() *[]*string
	BareMetal() *string
	SetBareMetal(val *string)
	BareMetalInput() *string
	BaselineEbsBandwidthMbps() LaunchTemplateInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference
	BaselineEbsBandwidthMbpsInput() *LaunchTemplateInstanceRequirementsBaselineEbsBandwidthMbps
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
	InternalValue() *LaunchTemplateInstanceRequirements
	SetInternalValue(val *LaunchTemplateInstanceRequirements)
	LocalStorage() *string
	SetLocalStorage(val *string)
	LocalStorageInput() *string
	LocalStorageTypes() *[]*string
	SetLocalStorageTypes(val *[]*string)
	LocalStorageTypesInput() *[]*string
	MemoryGibPerVcpu() LaunchTemplateInstanceRequirementsMemoryGibPerVcpuOutputReference
	MemoryGibPerVcpuInput() *LaunchTemplateInstanceRequirementsMemoryGibPerVcpu
	MemoryMib() LaunchTemplateInstanceRequirementsMemoryMibOutputReference
	MemoryMibInput() *LaunchTemplateInstanceRequirementsMemoryMib
	NetworkBandwidthGbps() LaunchTemplateInstanceRequirementsNetworkBandwidthGbpsOutputReference
	NetworkBandwidthGbpsInput() *LaunchTemplateInstanceRequirementsNetworkBandwidthGbps
	NetworkInterfaceCount() LaunchTemplateInstanceRequirementsNetworkInterfaceCountOutputReference
	NetworkInterfaceCountInput() *LaunchTemplateInstanceRequirementsNetworkInterfaceCount
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
	TotalLocalStorageGb() LaunchTemplateInstanceRequirementsTotalLocalStorageGbOutputReference
	TotalLocalStorageGbInput() *LaunchTemplateInstanceRequirementsTotalLocalStorageGb
	VcpuCount() LaunchTemplateInstanceRequirementsVcpuCountOutputReference
	VcpuCountInput() *LaunchTemplateInstanceRequirementsVcpuCount
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
	PutAcceleratorCount(value *LaunchTemplateInstanceRequirementsAcceleratorCount)
	PutAcceleratorTotalMemoryMib(value *LaunchTemplateInstanceRequirementsAcceleratorTotalMemoryMib)
	PutBaselineEbsBandwidthMbps(value *LaunchTemplateInstanceRequirementsBaselineEbsBandwidthMbps)
	PutMemoryGibPerVcpu(value *LaunchTemplateInstanceRequirementsMemoryGibPerVcpu)
	PutMemoryMib(value *LaunchTemplateInstanceRequirementsMemoryMib)
	PutNetworkBandwidthGbps(value *LaunchTemplateInstanceRequirementsNetworkBandwidthGbps)
	PutNetworkInterfaceCount(value *LaunchTemplateInstanceRequirementsNetworkInterfaceCount)
	PutTotalLocalStorageGb(value *LaunchTemplateInstanceRequirementsTotalLocalStorageGb)
	PutVcpuCount(value *LaunchTemplateInstanceRequirementsVcpuCount)
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

// The jsii proxy struct for LaunchTemplateInstanceRequirementsOutputReference
type jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) AcceleratorCount() LaunchTemplateInstanceRequirementsAcceleratorCountOutputReference {
	var returns LaunchTemplateInstanceRequirementsAcceleratorCountOutputReference
	_jsii_.Get(
		j,
		"acceleratorCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) AcceleratorCountInput() *LaunchTemplateInstanceRequirementsAcceleratorCount {
	var returns *LaunchTemplateInstanceRequirementsAcceleratorCount
	_jsii_.Get(
		j,
		"acceleratorCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) AcceleratorManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) AcceleratorManufacturersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorManufacturersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) AcceleratorNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) AcceleratorNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) AcceleratorTotalMemoryMib() LaunchTemplateInstanceRequirementsAcceleratorTotalMemoryMibOutputReference {
	var returns LaunchTemplateInstanceRequirementsAcceleratorTotalMemoryMibOutputReference
	_jsii_.Get(
		j,
		"acceleratorTotalMemoryMib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) AcceleratorTotalMemoryMibInput() *LaunchTemplateInstanceRequirementsAcceleratorTotalMemoryMib {
	var returns *LaunchTemplateInstanceRequirementsAcceleratorTotalMemoryMib
	_jsii_.Get(
		j,
		"acceleratorTotalMemoryMibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) AcceleratorTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) AcceleratorTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) AllowedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) AllowedInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) BareMetal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) BareMetalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) BaselineEbsBandwidthMbps() LaunchTemplateInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference {
	var returns LaunchTemplateInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference
	_jsii_.Get(
		j,
		"baselineEbsBandwidthMbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) BaselineEbsBandwidthMbpsInput() *LaunchTemplateInstanceRequirementsBaselineEbsBandwidthMbps {
	var returns *LaunchTemplateInstanceRequirementsBaselineEbsBandwidthMbps
	_jsii_.Get(
		j,
		"baselineEbsBandwidthMbpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) BurstablePerformance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"burstablePerformance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) BurstablePerformanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"burstablePerformanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) CpuManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cpuManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) CpuManufacturersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cpuManufacturersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ExcludedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ExcludedInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) InstanceGenerations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGenerations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) InstanceGenerationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGenerationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) InternalValue() *LaunchTemplateInstanceRequirements {
	var returns *LaunchTemplateInstanceRequirements
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) LocalStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) LocalStorageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) LocalStorageTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localStorageTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) LocalStorageTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localStorageTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) MemoryGibPerVcpu() LaunchTemplateInstanceRequirementsMemoryGibPerVcpuOutputReference {
	var returns LaunchTemplateInstanceRequirementsMemoryGibPerVcpuOutputReference
	_jsii_.Get(
		j,
		"memoryGibPerVcpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) MemoryGibPerVcpuInput() *LaunchTemplateInstanceRequirementsMemoryGibPerVcpu {
	var returns *LaunchTemplateInstanceRequirementsMemoryGibPerVcpu
	_jsii_.Get(
		j,
		"memoryGibPerVcpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) MemoryMib() LaunchTemplateInstanceRequirementsMemoryMibOutputReference {
	var returns LaunchTemplateInstanceRequirementsMemoryMibOutputReference
	_jsii_.Get(
		j,
		"memoryMib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) MemoryMibInput() *LaunchTemplateInstanceRequirementsMemoryMib {
	var returns *LaunchTemplateInstanceRequirementsMemoryMib
	_jsii_.Get(
		j,
		"memoryMibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) NetworkBandwidthGbps() LaunchTemplateInstanceRequirementsNetworkBandwidthGbpsOutputReference {
	var returns LaunchTemplateInstanceRequirementsNetworkBandwidthGbpsOutputReference
	_jsii_.Get(
		j,
		"networkBandwidthGbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) NetworkBandwidthGbpsInput() *LaunchTemplateInstanceRequirementsNetworkBandwidthGbps {
	var returns *LaunchTemplateInstanceRequirementsNetworkBandwidthGbps
	_jsii_.Get(
		j,
		"networkBandwidthGbpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) NetworkInterfaceCount() LaunchTemplateInstanceRequirementsNetworkInterfaceCountOutputReference {
	var returns LaunchTemplateInstanceRequirementsNetworkInterfaceCountOutputReference
	_jsii_.Get(
		j,
		"networkInterfaceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) NetworkInterfaceCountInput() *LaunchTemplateInstanceRequirementsNetworkInterfaceCount {
	var returns *LaunchTemplateInstanceRequirementsNetworkInterfaceCount
	_jsii_.Get(
		j,
		"networkInterfaceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) OnDemandMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) OnDemandMaxPricePercentageOverLowestPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandMaxPricePercentageOverLowestPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) RequireHibernateSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHibernateSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) RequireHibernateSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHibernateSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) SpotMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) SpotMaxPricePercentageOverLowestPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotMaxPricePercentageOverLowestPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) TotalLocalStorageGb() LaunchTemplateInstanceRequirementsTotalLocalStorageGbOutputReference {
	var returns LaunchTemplateInstanceRequirementsTotalLocalStorageGbOutputReference
	_jsii_.Get(
		j,
		"totalLocalStorageGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) TotalLocalStorageGbInput() *LaunchTemplateInstanceRequirementsTotalLocalStorageGb {
	var returns *LaunchTemplateInstanceRequirementsTotalLocalStorageGb
	_jsii_.Get(
		j,
		"totalLocalStorageGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) VcpuCount() LaunchTemplateInstanceRequirementsVcpuCountOutputReference {
	var returns LaunchTemplateInstanceRequirementsVcpuCountOutputReference
	_jsii_.Get(
		j,
		"vcpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) VcpuCountInput() *LaunchTemplateInstanceRequirementsVcpuCount {
	var returns *LaunchTemplateInstanceRequirementsVcpuCount
	_jsii_.Get(
		j,
		"vcpuCountInput",
		&returns,
	)
	return returns
}


func NewLaunchTemplateInstanceRequirementsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LaunchTemplateInstanceRequirementsOutputReference {
	_init_.Initialize()

	if err := validateNewLaunchTemplateInstanceRequirementsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.launchTemplate.LaunchTemplateInstanceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLaunchTemplateInstanceRequirementsOutputReference_Override(l LaunchTemplateInstanceRequirementsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.launchTemplate.LaunchTemplateInstanceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetAcceleratorManufacturers(val *[]*string) {
	if err := j.validateSetAcceleratorManufacturersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorManufacturers",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetAcceleratorNames(val *[]*string) {
	if err := j.validateSetAcceleratorNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorNames",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetAcceleratorTypes(val *[]*string) {
	if err := j.validateSetAcceleratorTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorTypes",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetAllowedInstanceTypes(val *[]*string) {
	if err := j.validateSetAllowedInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetBareMetal(val *string) {
	if err := j.validateSetBareMetalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bareMetal",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetBurstablePerformance(val *string) {
	if err := j.validateSetBurstablePerformanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"burstablePerformance",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetCpuManufacturers(val *[]*string) {
	if err := j.validateSetCpuManufacturersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuManufacturers",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetExcludedInstanceTypes(val *[]*string) {
	if err := j.validateSetExcludedInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetInstanceGenerations(val *[]*string) {
	if err := j.validateSetInstanceGenerationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceGenerations",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetInternalValue(val *LaunchTemplateInstanceRequirements) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetLocalStorage(val *string) {
	if err := j.validateSetLocalStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localStorage",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetLocalStorageTypes(val *[]*string) {
	if err := j.validateSetLocalStorageTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localStorageTypes",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetOnDemandMaxPricePercentageOverLowestPrice(val *float64) {
	if err := j.validateSetOnDemandMaxPricePercentageOverLowestPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandMaxPricePercentageOverLowestPrice",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetRequireHibernateSupport(val interface{}) {
	if err := j.validateSetRequireHibernateSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireHibernateSupport",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetSpotMaxPricePercentageOverLowestPrice(val *float64) {
	if err := j.validateSetSpotMaxPricePercentageOverLowestPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotMaxPricePercentageOverLowestPrice",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) PutAcceleratorCount(value *LaunchTemplateInstanceRequirementsAcceleratorCount) {
	if err := l.validatePutAcceleratorCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAcceleratorCount",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) PutAcceleratorTotalMemoryMib(value *LaunchTemplateInstanceRequirementsAcceleratorTotalMemoryMib) {
	if err := l.validatePutAcceleratorTotalMemoryMibParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAcceleratorTotalMemoryMib",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) PutBaselineEbsBandwidthMbps(value *LaunchTemplateInstanceRequirementsBaselineEbsBandwidthMbps) {
	if err := l.validatePutBaselineEbsBandwidthMbpsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putBaselineEbsBandwidthMbps",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) PutMemoryGibPerVcpu(value *LaunchTemplateInstanceRequirementsMemoryGibPerVcpu) {
	if err := l.validatePutMemoryGibPerVcpuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMemoryGibPerVcpu",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) PutMemoryMib(value *LaunchTemplateInstanceRequirementsMemoryMib) {
	if err := l.validatePutMemoryMibParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMemoryMib",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) PutNetworkBandwidthGbps(value *LaunchTemplateInstanceRequirementsNetworkBandwidthGbps) {
	if err := l.validatePutNetworkBandwidthGbpsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putNetworkBandwidthGbps",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) PutNetworkInterfaceCount(value *LaunchTemplateInstanceRequirementsNetworkInterfaceCount) {
	if err := l.validatePutNetworkInterfaceCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putNetworkInterfaceCount",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) PutTotalLocalStorageGb(value *LaunchTemplateInstanceRequirementsTotalLocalStorageGb) {
	if err := l.validatePutTotalLocalStorageGbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTotalLocalStorageGb",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) PutVcpuCount(value *LaunchTemplateInstanceRequirementsVcpuCount) {
	if err := l.validatePutVcpuCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putVcpuCount",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetAcceleratorCount() {
	_jsii_.InvokeVoid(
		l,
		"resetAcceleratorCount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetAcceleratorManufacturers() {
	_jsii_.InvokeVoid(
		l,
		"resetAcceleratorManufacturers",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetAcceleratorNames() {
	_jsii_.InvokeVoid(
		l,
		"resetAcceleratorNames",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetAcceleratorTotalMemoryMib() {
	_jsii_.InvokeVoid(
		l,
		"resetAcceleratorTotalMemoryMib",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetAcceleratorTypes() {
	_jsii_.InvokeVoid(
		l,
		"resetAcceleratorTypes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetAllowedInstanceTypes() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowedInstanceTypes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetBareMetal() {
	_jsii_.InvokeVoid(
		l,
		"resetBareMetal",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetBaselineEbsBandwidthMbps() {
	_jsii_.InvokeVoid(
		l,
		"resetBaselineEbsBandwidthMbps",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetBurstablePerformance() {
	_jsii_.InvokeVoid(
		l,
		"resetBurstablePerformance",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetCpuManufacturers() {
	_jsii_.InvokeVoid(
		l,
		"resetCpuManufacturers",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetExcludedInstanceTypes() {
	_jsii_.InvokeVoid(
		l,
		"resetExcludedInstanceTypes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetInstanceGenerations() {
	_jsii_.InvokeVoid(
		l,
		"resetInstanceGenerations",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetLocalStorage() {
	_jsii_.InvokeVoid(
		l,
		"resetLocalStorage",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetLocalStorageTypes() {
	_jsii_.InvokeVoid(
		l,
		"resetLocalStorageTypes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetMemoryGibPerVcpu() {
	_jsii_.InvokeVoid(
		l,
		"resetMemoryGibPerVcpu",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetNetworkBandwidthGbps() {
	_jsii_.InvokeVoid(
		l,
		"resetNetworkBandwidthGbps",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetNetworkInterfaceCount() {
	_jsii_.InvokeVoid(
		l,
		"resetNetworkInterfaceCount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetOnDemandMaxPricePercentageOverLowestPrice() {
	_jsii_.InvokeVoid(
		l,
		"resetOnDemandMaxPricePercentageOverLowestPrice",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetRequireHibernateSupport() {
	_jsii_.InvokeVoid(
		l,
		"resetRequireHibernateSupport",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetSpotMaxPricePercentageOverLowestPrice() {
	_jsii_.InvokeVoid(
		l,
		"resetSpotMaxPricePercentageOverLowestPrice",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ResetTotalLocalStorageGb() {
	_jsii_.InvokeVoid(
		l,
		"resetTotalLocalStorageGb",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateInstanceRequirementsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

