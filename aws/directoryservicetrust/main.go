// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryservicetrust

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.directoryServiceTrust.DirectoryServiceTrust",
		reflect.TypeOf((*DirectoryServiceTrust)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "conditionalForwarderIpAddrs", GoGetter: "ConditionalForwarderIpAddrs"},
			_jsii_.MemberProperty{JsiiProperty: "conditionalForwarderIpAddrsInput", GoGetter: "ConditionalForwarderIpAddrsInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdDateTime", GoGetter: "CreatedDateTime"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAssociatedConditionalForwarder", GoGetter: "DeleteAssociatedConditionalForwarder"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAssociatedConditionalForwarderInput", GoGetter: "DeleteAssociatedConditionalForwarderInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "directoryId", GoGetter: "DirectoryId"},
			_jsii_.MemberProperty{JsiiProperty: "directoryIdInput", GoGetter: "DirectoryIdInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedDateTime", GoGetter: "LastUpdatedDateTime"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "remoteDomainName", GoGetter: "RemoteDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "remoteDomainNameInput", GoGetter: "RemoteDomainNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetConditionalForwarderIpAddrs", GoMethod: "ResetConditionalForwarderIpAddrs"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeleteAssociatedConditionalForwarder", GoMethod: "ResetDeleteAssociatedConditionalForwarder"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSelectiveAuth", GoMethod: "ResetSelectiveAuth"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrustType", GoMethod: "ResetTrustType"},
			_jsii_.MemberProperty{JsiiProperty: "selectiveAuth", GoGetter: "SelectiveAuth"},
			_jsii_.MemberProperty{JsiiProperty: "selectiveAuthInput", GoGetter: "SelectiveAuthInput"},
			_jsii_.MemberProperty{JsiiProperty: "stateLastUpdatedDateTime", GoGetter: "StateLastUpdatedDateTime"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "trustDirection", GoGetter: "TrustDirection"},
			_jsii_.MemberProperty{JsiiProperty: "trustDirectionInput", GoGetter: "TrustDirectionInput"},
			_jsii_.MemberProperty{JsiiProperty: "trustPassword", GoGetter: "TrustPassword"},
			_jsii_.MemberProperty{JsiiProperty: "trustPasswordInput", GoGetter: "TrustPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "trustState", GoGetter: "TrustState"},
			_jsii_.MemberProperty{JsiiProperty: "trustStateReason", GoGetter: "TrustStateReason"},
			_jsii_.MemberProperty{JsiiProperty: "trustType", GoGetter: "TrustType"},
			_jsii_.MemberProperty{JsiiProperty: "trustTypeInput", GoGetter: "TrustTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_DirectoryServiceTrust{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.directoryServiceTrust.DirectoryServiceTrustConfig",
		reflect.TypeOf((*DirectoryServiceTrustConfig)(nil)).Elem(),
	)
}
