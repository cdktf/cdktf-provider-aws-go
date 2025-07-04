// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkinterface

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.networkInterface.NetworkInterface",
		reflect.TypeOf((*NetworkInterface)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "attachment", GoGetter: "Attachment"},
			_jsii_.MemberProperty{JsiiProperty: "attachmentInput", GoGetter: "AttachmentInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "enablePrimaryIpv6", GoGetter: "EnablePrimaryIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "enablePrimaryIpv6Input", GoGetter: "EnablePrimaryIpv6Input"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "interfaceType", GoGetter: "InterfaceType"},
			_jsii_.MemberProperty{JsiiProperty: "interfaceTypeInput", GoGetter: "InterfaceTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4PrefixCount", GoGetter: "Ipv4PrefixCount"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4PrefixCountInput", GoGetter: "Ipv4PrefixCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4Prefixes", GoGetter: "Ipv4Prefixes"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4PrefixesInput", GoGetter: "Ipv4PrefixesInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6AddressCount", GoGetter: "Ipv6AddressCount"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6AddressCountInput", GoGetter: "Ipv6AddressCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Addresses", GoGetter: "Ipv6Addresses"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6AddressesInput", GoGetter: "Ipv6AddressesInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6AddressList", GoGetter: "Ipv6AddressList"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6AddressListEnabled", GoGetter: "Ipv6AddressListEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6AddressListEnabledInput", GoGetter: "Ipv6AddressListEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6AddressListInput", GoGetter: "Ipv6AddressListInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6PrefixCount", GoGetter: "Ipv6PrefixCount"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6PrefixCountInput", GoGetter: "Ipv6PrefixCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Prefixes", GoGetter: "Ipv6Prefixes"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6PrefixesInput", GoGetter: "Ipv6PrefixesInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "macAddress", GoGetter: "MacAddress"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outpostArn", GoGetter: "OutpostArn"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ownerId", GoGetter: "OwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsName", GoGetter: "PrivateDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "privateIp", GoGetter: "PrivateIp"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpInput", GoGetter: "PrivateIpInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpList", GoGetter: "PrivateIpList"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpListEnabled", GoGetter: "PrivateIpListEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpListEnabledInput", GoGetter: "PrivateIpListEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpListInput", GoGetter: "PrivateIpListInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateIps", GoGetter: "PrivateIps"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpsCount", GoGetter: "PrivateIpsCount"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpsCountInput", GoGetter: "PrivateIpsCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpsInput", GoGetter: "PrivateIpsInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putAttachment", GoMethod: "PutAttachment"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAttachment", GoMethod: "ResetAttachment"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnablePrimaryIpv6", GoMethod: "ResetEnablePrimaryIpv6"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetInterfaceType", GoMethod: "ResetInterfaceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv4PrefixCount", GoMethod: "ResetIpv4PrefixCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv4Prefixes", GoMethod: "ResetIpv4Prefixes"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6AddressCount", GoMethod: "ResetIpv6AddressCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6Addresses", GoMethod: "ResetIpv6Addresses"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6AddressList", GoMethod: "ResetIpv6AddressList"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6AddressListEnabled", GoMethod: "ResetIpv6AddressListEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6PrefixCount", GoMethod: "ResetIpv6PrefixCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6Prefixes", GoMethod: "ResetIpv6Prefixes"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateIp", GoMethod: "ResetPrivateIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateIpList", GoMethod: "ResetPrivateIpList"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateIpListEnabled", GoMethod: "ResetPrivateIpListEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateIps", GoMethod: "ResetPrivateIps"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateIpsCount", GoMethod: "ResetPrivateIpsCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityGroups", GoMethod: "ResetSecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceDestCheck", GoMethod: "ResetSourceDestCheck"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupsInput", GoGetter: "SecurityGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceDestCheck", GoGetter: "SourceDestCheck"},
			_jsii_.MemberProperty{JsiiProperty: "sourceDestCheckInput", GoGetter: "SourceDestCheckInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdInput", GoGetter: "SubnetIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkInterface{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.networkInterface.NetworkInterfaceAttachment",
		reflect.TypeOf((*NetworkInterfaceAttachment)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.networkInterface.NetworkInterfaceAttachmentList",
		reflect.TypeOf((*NetworkInterfaceAttachmentList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkInterfaceAttachmentList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.networkInterface.NetworkInterfaceAttachmentOutputReference",
		reflect.TypeOf((*NetworkInterfaceAttachmentOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "attachmentId", GoGetter: "AttachmentId"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deviceIndex", GoGetter: "DeviceIndex"},
			_jsii_.MemberProperty{JsiiProperty: "deviceIndexInput", GoGetter: "DeviceIndexInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instance", GoGetter: "Instance"},
			_jsii_.MemberProperty{JsiiProperty: "instanceInput", GoGetter: "InstanceInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkInterfaceAttachmentOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.networkInterface.NetworkInterfaceConfig",
		reflect.TypeOf((*NetworkInterfaceConfig)(nil)).Elem(),
	)
}
