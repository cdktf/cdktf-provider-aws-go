// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devicefarmnetworkprofile

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.devicefarmNetworkProfile.DevicefarmNetworkProfile",
		reflect.TypeOf((*DevicefarmNetworkProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "downlinkBandwidthBits", GoGetter: "DownlinkBandwidthBits"},
			_jsii_.MemberProperty{JsiiProperty: "downlinkBandwidthBitsInput", GoGetter: "DownlinkBandwidthBitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "downlinkDelayMs", GoGetter: "DownlinkDelayMs"},
			_jsii_.MemberProperty{JsiiProperty: "downlinkDelayMsInput", GoGetter: "DownlinkDelayMsInput"},
			_jsii_.MemberProperty{JsiiProperty: "downlinkJitterMs", GoGetter: "DownlinkJitterMs"},
			_jsii_.MemberProperty{JsiiProperty: "downlinkJitterMsInput", GoGetter: "DownlinkJitterMsInput"},
			_jsii_.MemberProperty{JsiiProperty: "downlinkLossPercent", GoGetter: "DownlinkLossPercent"},
			_jsii_.MemberProperty{JsiiProperty: "downlinkLossPercentInput", GoGetter: "DownlinkLossPercentInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "projectArn", GoGetter: "ProjectArn"},
			_jsii_.MemberProperty{JsiiProperty: "projectArnInput", GoGetter: "ProjectArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetDownlinkBandwidthBits", GoMethod: "ResetDownlinkBandwidthBits"},
			_jsii_.MemberMethod{JsiiMethod: "resetDownlinkDelayMs", GoMethod: "ResetDownlinkDelayMs"},
			_jsii_.MemberMethod{JsiiMethod: "resetDownlinkJitterMs", GoMethod: "ResetDownlinkJitterMs"},
			_jsii_.MemberMethod{JsiiMethod: "resetDownlinkLossPercent", GoMethod: "ResetDownlinkLossPercent"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resetUplinkBandwidthBits", GoMethod: "ResetUplinkBandwidthBits"},
			_jsii_.MemberMethod{JsiiMethod: "resetUplinkDelayMs", GoMethod: "ResetUplinkDelayMs"},
			_jsii_.MemberMethod{JsiiMethod: "resetUplinkJitterMs", GoMethod: "ResetUplinkJitterMs"},
			_jsii_.MemberMethod{JsiiMethod: "resetUplinkLossPercent", GoMethod: "ResetUplinkLossPercent"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "uplinkBandwidthBits", GoGetter: "UplinkBandwidthBits"},
			_jsii_.MemberProperty{JsiiProperty: "uplinkBandwidthBitsInput", GoGetter: "UplinkBandwidthBitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "uplinkDelayMs", GoGetter: "UplinkDelayMs"},
			_jsii_.MemberProperty{JsiiProperty: "uplinkDelayMsInput", GoGetter: "UplinkDelayMsInput"},
			_jsii_.MemberProperty{JsiiProperty: "uplinkJitterMs", GoGetter: "UplinkJitterMs"},
			_jsii_.MemberProperty{JsiiProperty: "uplinkJitterMsInput", GoGetter: "UplinkJitterMsInput"},
			_jsii_.MemberProperty{JsiiProperty: "uplinkLossPercent", GoGetter: "UplinkLossPercent"},
			_jsii_.MemberProperty{JsiiProperty: "uplinkLossPercentInput", GoGetter: "UplinkLossPercentInput"},
		},
		func() interface{} {
			j := jsiiProxy_DevicefarmNetworkProfile{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.devicefarmNetworkProfile.DevicefarmNetworkProfileConfig",
		reflect.TypeOf((*DevicefarmNetworkProfileConfig)(nil)).Elem(),
	)
}
