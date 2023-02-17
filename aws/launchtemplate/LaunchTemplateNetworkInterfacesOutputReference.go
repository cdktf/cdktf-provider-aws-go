package launchtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/launchtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LaunchTemplateNetworkInterfacesOutputReference interface {
	cdktf.ComplexObject
	AssociateCarrierIpAddress() *string
	SetAssociateCarrierIpAddress(val *string)
	AssociateCarrierIpAddressInput() *string
	AssociatePublicIpAddress() *string
	SetAssociatePublicIpAddress(val *string)
	AssociatePublicIpAddressInput() *string
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
	DeleteOnTermination() *string
	SetDeleteOnTermination(val *string)
	DeleteOnTerminationInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DeviceIndex() *float64
	SetDeviceIndex(val *float64)
	DeviceIndexInput() *float64
	// Experimental.
	Fqn() *string
	InterfaceType() *string
	SetInterfaceType(val *string)
	InterfaceTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ipv4AddressCount() *float64
	SetIpv4AddressCount(val *float64)
	Ipv4AddressCountInput() *float64
	Ipv4Addresses() *[]*string
	SetIpv4Addresses(val *[]*string)
	Ipv4AddressesInput() *[]*string
	Ipv4PrefixCount() *float64
	SetIpv4PrefixCount(val *float64)
	Ipv4PrefixCountInput() *float64
	Ipv4Prefixes() *[]*string
	SetIpv4Prefixes(val *[]*string)
	Ipv4PrefixesInput() *[]*string
	Ipv6AddressCount() *float64
	SetIpv6AddressCount(val *float64)
	Ipv6AddressCountInput() *float64
	Ipv6Addresses() *[]*string
	SetIpv6Addresses(val *[]*string)
	Ipv6AddressesInput() *[]*string
	Ipv6PrefixCount() *float64
	SetIpv6PrefixCount(val *float64)
	Ipv6PrefixCountInput() *float64
	Ipv6Prefixes() *[]*string
	SetIpv6Prefixes(val *[]*string)
	Ipv6PrefixesInput() *[]*string
	NetworkCardIndex() *float64
	SetNetworkCardIndex(val *float64)
	NetworkCardIndexInput() *float64
	NetworkInterfaceId() *string
	SetNetworkInterfaceId(val *string)
	NetworkInterfaceIdInput() *string
	PrivateIpAddress() *string
	SetPrivateIpAddress(val *string)
	PrivateIpAddressInput() *string
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetAssociateCarrierIpAddress()
	ResetAssociatePublicIpAddress()
	ResetDeleteOnTermination()
	ResetDescription()
	ResetDeviceIndex()
	ResetInterfaceType()
	ResetIpv4AddressCount()
	ResetIpv4Addresses()
	ResetIpv4PrefixCount()
	ResetIpv4Prefixes()
	ResetIpv6AddressCount()
	ResetIpv6Addresses()
	ResetIpv6PrefixCount()
	ResetIpv6Prefixes()
	ResetNetworkCardIndex()
	ResetNetworkInterfaceId()
	ResetPrivateIpAddress()
	ResetSecurityGroups()
	ResetSubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LaunchTemplateNetworkInterfacesOutputReference
type jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) AssociateCarrierIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associateCarrierIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) AssociateCarrierIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associateCarrierIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) AssociatePublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) AssociatePublicIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatePublicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) DeleteOnTermination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteOnTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) DeleteOnTerminationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteOnTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) DeviceIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) DeviceIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) InterfaceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) InterfaceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Ipv4AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Ipv4AddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4AddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Ipv4Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Ipv4AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Ipv4PrefixCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4PrefixCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Ipv4PrefixCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4PrefixCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Ipv4Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Ipv4PrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4PrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Ipv6AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Ipv6AddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Ipv6Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Ipv6AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Ipv6PrefixCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6PrefixCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Ipv6PrefixCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6PrefixCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Ipv6Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Ipv6PrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6PrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) NetworkCardIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkCardIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) NetworkCardIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkCardIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) NetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) NetworkInterfaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) PrivateIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLaunchTemplateNetworkInterfacesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LaunchTemplateNetworkInterfacesOutputReference {
	_init_.Initialize()

	if err := validateNewLaunchTemplateNetworkInterfacesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.launchTemplate.LaunchTemplateNetworkInterfacesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLaunchTemplateNetworkInterfacesOutputReference_Override(l LaunchTemplateNetworkInterfacesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.launchTemplate.LaunchTemplateNetworkInterfacesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetAssociateCarrierIpAddress(val *string) {
	if err := j.validateSetAssociateCarrierIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associateCarrierIpAddress",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetAssociatePublicIpAddress(val *string) {
	if err := j.validateSetAssociatePublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetDeleteOnTermination(val *string) {
	if err := j.validateSetDeleteOnTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOnTermination",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetDeviceIndex(val *float64) {
	if err := j.validateSetDeviceIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceIndex",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetInterfaceType(val *string) {
	if err := j.validateSetInterfaceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interfaceType",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetIpv4AddressCount(val *float64) {
	if err := j.validateSetIpv4AddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4AddressCount",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetIpv4Addresses(val *[]*string) {
	if err := j.validateSetIpv4AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Addresses",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetIpv4PrefixCount(val *float64) {
	if err := j.validateSetIpv4PrefixCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4PrefixCount",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetIpv4Prefixes(val *[]*string) {
	if err := j.validateSetIpv4PrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Prefixes",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetIpv6AddressCount(val *float64) {
	if err := j.validateSetIpv6AddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6AddressCount",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetIpv6Addresses(val *[]*string) {
	if err := j.validateSetIpv6AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Addresses",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetIpv6PrefixCount(val *float64) {
	if err := j.validateSetIpv6PrefixCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6PrefixCount",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetIpv6Prefixes(val *[]*string) {
	if err := j.validateSetIpv6PrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Prefixes",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetNetworkCardIndex(val *float64) {
	if err := j.validateSetNetworkCardIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkCardIndex",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetNetworkInterfaceId(val *string) {
	if err := j.validateSetNetworkInterfaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaceId",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetPrivateIpAddress(val *string) {
	if err := j.validateSetPrivateIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddress",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetAssociateCarrierIpAddress() {
	_jsii_.InvokeVoid(
		l,
		"resetAssociateCarrierIpAddress",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetAssociatePublicIpAddress() {
	_jsii_.InvokeVoid(
		l,
		"resetAssociatePublicIpAddress",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetDeleteOnTermination() {
	_jsii_.InvokeVoid(
		l,
		"resetDeleteOnTermination",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetDeviceIndex() {
	_jsii_.InvokeVoid(
		l,
		"resetDeviceIndex",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetInterfaceType() {
	_jsii_.InvokeVoid(
		l,
		"resetInterfaceType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetIpv4AddressCount() {
	_jsii_.InvokeVoid(
		l,
		"resetIpv4AddressCount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetIpv4Addresses() {
	_jsii_.InvokeVoid(
		l,
		"resetIpv4Addresses",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetIpv4PrefixCount() {
	_jsii_.InvokeVoid(
		l,
		"resetIpv4PrefixCount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetIpv4Prefixes() {
	_jsii_.InvokeVoid(
		l,
		"resetIpv4Prefixes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetIpv6AddressCount() {
	_jsii_.InvokeVoid(
		l,
		"resetIpv6AddressCount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetIpv6Addresses() {
	_jsii_.InvokeVoid(
		l,
		"resetIpv6Addresses",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetIpv6PrefixCount() {
	_jsii_.InvokeVoid(
		l,
		"resetIpv6PrefixCount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetIpv6Prefixes() {
	_jsii_.InvokeVoid(
		l,
		"resetIpv6Prefixes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetNetworkCardIndex() {
	_jsii_.InvokeVoid(
		l,
		"resetNetworkCardIndex",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetNetworkInterfaceId() {
	_jsii_.InvokeVoid(
		l,
		"resetNetworkInterfaceId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetPrivateIpAddress() {
	_jsii_.InvokeVoid(
		l,
		"resetPrivateIpAddress",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		l,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		l,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

