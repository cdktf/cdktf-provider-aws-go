package spotfleetrequest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/spotfleetrequest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpotFleetRequestLaunchSpecificationOutputReference interface {
	cdktf.ComplexObject
	Ami() *string
	SetAmi(val *string)
	AmiInput() *string
	AssociatePublicIpAddress() interface{}
	SetAssociatePublicIpAddress(val interface{})
	AssociatePublicIpAddressInput() interface{}
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EbsBlockDevice() SpotFleetRequestLaunchSpecificationEbsBlockDeviceList
	EbsBlockDeviceInput() interface{}
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	EbsOptimizedInput() interface{}
	EphemeralBlockDevice() SpotFleetRequestLaunchSpecificationEphemeralBlockDeviceList
	EphemeralBlockDeviceInput() interface{}
	// Experimental.
	Fqn() *string
	IamInstanceProfile() *string
	SetIamInstanceProfile(val *string)
	IamInstanceProfileArn() *string
	SetIamInstanceProfileArn(val *string)
	IamInstanceProfileArnInput() *string
	IamInstanceProfileInput() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	Monitoring() interface{}
	SetMonitoring(val interface{})
	MonitoringInput() interface{}
	PlacementGroup() *string
	SetPlacementGroup(val *string)
	PlacementGroupInput() *string
	PlacementTenancy() *string
	SetPlacementTenancy(val *string)
	PlacementTenancyInput() *string
	RootBlockDevice() SpotFleetRequestLaunchSpecificationRootBlockDeviceList
	RootBlockDeviceInput() interface{}
	SpotPrice() *string
	SetSpotPrice(val *string)
	SpotPriceInput() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	VpcSecurityGroupIdsInput() *[]*string
	WeightedCapacity() *string
	SetWeightedCapacity(val *string)
	WeightedCapacityInput() *string
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
	PutEbsBlockDevice(value interface{})
	PutEphemeralBlockDevice(value interface{})
	PutRootBlockDevice(value interface{})
	ResetAssociatePublicIpAddress()
	ResetAvailabilityZone()
	ResetEbsBlockDevice()
	ResetEbsOptimized()
	ResetEphemeralBlockDevice()
	ResetIamInstanceProfile()
	ResetIamInstanceProfileArn()
	ResetKeyName()
	ResetMonitoring()
	ResetPlacementGroup()
	ResetPlacementTenancy()
	ResetRootBlockDevice()
	ResetSpotPrice()
	ResetSubnetId()
	ResetTags()
	ResetUserData()
	ResetVpcSecurityGroupIds()
	ResetWeightedCapacity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpotFleetRequestLaunchSpecificationOutputReference
type jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) Ami() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ami",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) AmiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) AssociatePublicIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) AssociatePublicIpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) EbsBlockDevice() SpotFleetRequestLaunchSpecificationEbsBlockDeviceList {
	var returns SpotFleetRequestLaunchSpecificationEbsBlockDeviceList
	_jsii_.Get(
		j,
		"ebsBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) EbsBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) EphemeralBlockDevice() SpotFleetRequestLaunchSpecificationEphemeralBlockDeviceList {
	var returns SpotFleetRequestLaunchSpecificationEphemeralBlockDeviceList
	_jsii_.Get(
		j,
		"ephemeralBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) EphemeralBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) IamInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) IamInstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) Monitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) MonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) PlacementGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) PlacementGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) PlacementTenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementTenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) PlacementTenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementTenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) RootBlockDevice() SpotFleetRequestLaunchSpecificationRootBlockDeviceList {
	var returns SpotFleetRequestLaunchSpecificationRootBlockDeviceList
	_jsii_.Get(
		j,
		"rootBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) RootBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) SpotPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) SpotPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) WeightedCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weightedCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) WeightedCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weightedCapacityInput",
		&returns,
	)
	return returns
}


func NewSpotFleetRequestLaunchSpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SpotFleetRequestLaunchSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewSpotFleetRequestLaunchSpecificationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.spotFleetRequest.SpotFleetRequestLaunchSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSpotFleetRequestLaunchSpecificationOutputReference_Override(s SpotFleetRequestLaunchSpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.spotFleetRequest.SpotFleetRequestLaunchSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetAmi(val *string) {
	if err := j.validateSetAmiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ami",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetAssociatePublicIpAddress(val interface{}) {
	if err := j.validateSetAssociatePublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetIamInstanceProfileArn(val *string) {
	if err := j.validateSetIamInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetMonitoring(val interface{}) {
	if err := j.validateSetMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoring",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetPlacementGroup(val *string) {
	if err := j.validateSetPlacementGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementGroup",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetPlacementTenancy(val *string) {
	if err := j.validateSetPlacementTenancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementTenancy",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetSpotPrice(val *string) {
	if err := j.validateSetSpotPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotPrice",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference)SetWeightedCapacity(val *string) {
	if err := j.validateSetWeightedCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weightedCapacity",
		val,
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) PutEbsBlockDevice(value interface{}) {
	if err := s.validatePutEbsBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEbsBlockDevice",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) PutEphemeralBlockDevice(value interface{}) {
	if err := s.validatePutEphemeralBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEphemeralBlockDevice",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) PutRootBlockDevice(value interface{}) {
	if err := s.validatePutRootBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRootBlockDevice",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetAssociatePublicIpAddress() {
	_jsii_.InvokeVoid(
		s,
		"resetAssociatePublicIpAddress",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		s,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetEbsBlockDevice() {
	_jsii_.InvokeVoid(
		s,
		"resetEbsBlockDevice",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		s,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetEphemeralBlockDevice() {
	_jsii_.InvokeVoid(
		s,
		"resetEphemeralBlockDevice",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		s,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetIamInstanceProfileArn() {
	_jsii_.InvokeVoid(
		s,
		"resetIamInstanceProfileArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetKeyName() {
	_jsii_.InvokeVoid(
		s,
		"resetKeyName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetMonitoring() {
	_jsii_.InvokeVoid(
		s,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetPlacementGroup() {
	_jsii_.InvokeVoid(
		s,
		"resetPlacementGroup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetPlacementTenancy() {
	_jsii_.InvokeVoid(
		s,
		"resetPlacementTenancy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetRootBlockDevice() {
	_jsii_.InvokeVoid(
		s,
		"resetRootBlockDevice",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetSpotPrice() {
	_jsii_.InvokeVoid(
		s,
		"resetSpotPrice",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		s,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetUserData() {
	_jsii_.InvokeVoid(
		s,
		"resetUserData",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		s,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ResetWeightedCapacity() {
	_jsii_.InvokeVoid(
		s,
		"resetWeightedCapacity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SpotFleetRequestLaunchSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

