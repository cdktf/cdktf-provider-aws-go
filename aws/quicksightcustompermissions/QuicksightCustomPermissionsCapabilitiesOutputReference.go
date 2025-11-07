// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightcustompermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/quicksightcustompermissions/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightCustomPermissionsCapabilitiesOutputReference interface {
	cdktf.ComplexObject
	AddOrRunAnomalyDetectionForAnalyses() *string
	SetAddOrRunAnomalyDetectionForAnalyses(val *string)
	AddOrRunAnomalyDetectionForAnalysesInput() *string
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
	CreateAndUpdateDashboardEmailReports() *string
	SetCreateAndUpdateDashboardEmailReports(val *string)
	CreateAndUpdateDashboardEmailReportsInput() *string
	CreateAndUpdateDatasets() *string
	SetCreateAndUpdateDatasets(val *string)
	CreateAndUpdateDatasetsInput() *string
	CreateAndUpdateDataSources() *string
	SetCreateAndUpdateDataSources(val *string)
	CreateAndUpdateDataSourcesInput() *string
	CreateAndUpdateThemes() *string
	SetCreateAndUpdateThemes(val *string)
	CreateAndUpdateThemesInput() *string
	CreateAndUpdateThresholdAlerts() *string
	SetCreateAndUpdateThresholdAlerts(val *string)
	CreateAndUpdateThresholdAlertsInput() *string
	CreateSharedFolders() *string
	SetCreateSharedFolders(val *string)
	CreateSharedFoldersInput() *string
	CreateSpiceDataset() *string
	SetCreateSpiceDataset(val *string)
	CreateSpiceDatasetInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExportToCsv() *string
	SetExportToCsv(val *string)
	ExportToCsvInput() *string
	ExportToCsvInScheduledReports() *string
	SetExportToCsvInScheduledReports(val *string)
	ExportToCsvInScheduledReportsInput() *string
	ExportToExcel() *string
	SetExportToExcel(val *string)
	ExportToExcelInput() *string
	ExportToExcelInScheduledReports() *string
	SetExportToExcelInScheduledReports(val *string)
	ExportToExcelInScheduledReportsInput() *string
	ExportToPdf() *string
	SetExportToPdf(val *string)
	ExportToPdfInput() *string
	ExportToPdfInScheduledReports() *string
	SetExportToPdfInScheduledReports(val *string)
	ExportToPdfInScheduledReportsInput() *string
	// Experimental.
	Fqn() *string
	IncludeContentInScheduledReportsEmail() *string
	SetIncludeContentInScheduledReportsEmail(val *string)
	IncludeContentInScheduledReportsEmailInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PrintReports() *string
	SetPrintReports(val *string)
	PrintReportsInput() *string
	RenameSharedFolders() *string
	SetRenameSharedFolders(val *string)
	RenameSharedFoldersInput() *string
	ShareAnalyses() *string
	SetShareAnalyses(val *string)
	ShareAnalysesInput() *string
	ShareDashboards() *string
	SetShareDashboards(val *string)
	ShareDashboardsInput() *string
	ShareDatasets() *string
	SetShareDatasets(val *string)
	ShareDatasetsInput() *string
	ShareDataSources() *string
	SetShareDataSources(val *string)
	ShareDataSourcesInput() *string
	SubscribeDashboardEmailReports() *string
	SetSubscribeDashboardEmailReports(val *string)
	SubscribeDashboardEmailReportsInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ViewAccountSpiceCapacity() *string
	SetViewAccountSpiceCapacity(val *string)
	ViewAccountSpiceCapacityInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetAddOrRunAnomalyDetectionForAnalyses()
	ResetCreateAndUpdateDashboardEmailReports()
	ResetCreateAndUpdateDatasets()
	ResetCreateAndUpdateDataSources()
	ResetCreateAndUpdateThemes()
	ResetCreateAndUpdateThresholdAlerts()
	ResetCreateSharedFolders()
	ResetCreateSpiceDataset()
	ResetExportToCsv()
	ResetExportToCsvInScheduledReports()
	ResetExportToExcel()
	ResetExportToExcelInScheduledReports()
	ResetExportToPdf()
	ResetExportToPdfInScheduledReports()
	ResetIncludeContentInScheduledReportsEmail()
	ResetPrintReports()
	ResetRenameSharedFolders()
	ResetShareAnalyses()
	ResetShareDashboards()
	ResetShareDatasets()
	ResetShareDataSources()
	ResetSubscribeDashboardEmailReports()
	ResetViewAccountSpiceCapacity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightCustomPermissionsCapabilitiesOutputReference
type jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AddOrRunAnomalyDetectionForAnalyses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addOrRunAnomalyDetectionForAnalyses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AddOrRunAnomalyDetectionForAnalysesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addOrRunAnomalyDetectionForAnalysesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDashboardEmailReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDashboardEmailReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDashboardEmailReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDashboardEmailReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDatasets() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDatasets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDatasetsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDatasetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDataSources() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDataSourcesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateThemes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateThemes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateThemesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateThemesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateThresholdAlerts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateThresholdAlerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateThresholdAlertsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateThresholdAlertsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSharedFolders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSharedFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSharedFoldersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSharedFoldersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSpiceDataset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSpiceDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSpiceDatasetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSpiceDatasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToCsv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToCsv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToCsvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToCsvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToCsvInScheduledReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToCsvInScheduledReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToCsvInScheduledReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToCsvInScheduledReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToExcel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToExcel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToExcelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToExcelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToExcelInScheduledReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToExcelInScheduledReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToExcelInScheduledReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToExcelInScheduledReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToPdf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToPdf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToPdfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToPdfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToPdfInScheduledReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToPdfInScheduledReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToPdfInScheduledReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToPdfInScheduledReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) IncludeContentInScheduledReportsEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeContentInScheduledReportsEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) IncludeContentInScheduledReportsEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeContentInScheduledReportsEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) PrintReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"printReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) PrintReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"printReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) RenameSharedFolders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renameSharedFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) RenameSharedFoldersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renameSharedFoldersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAnalyses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAnalyses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAnalysesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAnalysesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDashboards() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDashboards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDashboardsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDashboardsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDatasets() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDatasets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDatasetsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDatasetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDataSources() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDataSourcesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SubscribeDashboardEmailReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscribeDashboardEmailReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SubscribeDashboardEmailReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscribeDashboardEmailReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ViewAccountSpiceCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewAccountSpiceCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ViewAccountSpiceCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewAccountSpiceCapacityInput",
		&returns,
	)
	return returns
}


func NewQuicksightCustomPermissionsCapabilitiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) QuicksightCustomPermissionsCapabilitiesOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightCustomPermissionsCapabilitiesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightCustomPermissions.QuicksightCustomPermissionsCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewQuicksightCustomPermissionsCapabilitiesOutputReference_Override(q QuicksightCustomPermissionsCapabilitiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightCustomPermissions.QuicksightCustomPermissionsCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		q,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetAddOrRunAnomalyDetectionForAnalyses(val *string) {
	if err := j.validateSetAddOrRunAnomalyDetectionForAnalysesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addOrRunAnomalyDetectionForAnalyses",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateDashboardEmailReports(val *string) {
	if err := j.validateSetCreateAndUpdateDashboardEmailReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateDashboardEmailReports",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateDatasets(val *string) {
	if err := j.validateSetCreateAndUpdateDatasetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateDatasets",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateDataSources(val *string) {
	if err := j.validateSetCreateAndUpdateDataSourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateDataSources",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateThemes(val *string) {
	if err := j.validateSetCreateAndUpdateThemesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateThemes",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateThresholdAlerts(val *string) {
	if err := j.validateSetCreateAndUpdateThresholdAlertsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateThresholdAlerts",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateSharedFolders(val *string) {
	if err := j.validateSetCreateSharedFoldersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createSharedFolders",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateSpiceDataset(val *string) {
	if err := j.validateSetCreateSpiceDatasetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createSpiceDataset",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetExportToCsv(val *string) {
	if err := j.validateSetExportToCsvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToCsv",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetExportToCsvInScheduledReports(val *string) {
	if err := j.validateSetExportToCsvInScheduledReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToCsvInScheduledReports",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetExportToExcel(val *string) {
	if err := j.validateSetExportToExcelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToExcel",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetExportToExcelInScheduledReports(val *string) {
	if err := j.validateSetExportToExcelInScheduledReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToExcelInScheduledReports",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetExportToPdf(val *string) {
	if err := j.validateSetExportToPdfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToPdf",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetExportToPdfInScheduledReports(val *string) {
	if err := j.validateSetExportToPdfInScheduledReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToPdfInScheduledReports",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetIncludeContentInScheduledReportsEmail(val *string) {
	if err := j.validateSetIncludeContentInScheduledReportsEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeContentInScheduledReportsEmail",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetPrintReports(val *string) {
	if err := j.validateSetPrintReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"printReports",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetRenameSharedFolders(val *string) {
	if err := j.validateSetRenameSharedFoldersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renameSharedFolders",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareAnalyses(val *string) {
	if err := j.validateSetShareAnalysesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareAnalyses",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareDashboards(val *string) {
	if err := j.validateSetShareDashboardsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDashboards",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareDatasets(val *string) {
	if err := j.validateSetShareDatasetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDatasets",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareDataSources(val *string) {
	if err := j.validateSetShareDataSourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDataSources",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSubscribeDashboardEmailReports(val *string) {
	if err := j.validateSetSubscribeDashboardEmailReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscribeDashboardEmailReports",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetViewAccountSpiceCapacity(val *string) {
	if err := j.validateSetViewAccountSpiceCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewAccountSpiceCapacity",
		val,
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetAddOrRunAnomalyDetectionForAnalyses() {
	_jsii_.InvokeVoid(
		q,
		"resetAddOrRunAnomalyDetectionForAnalyses",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateDashboardEmailReports() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateDashboardEmailReports",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateDatasets() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateDatasets",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateDataSources() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateDataSources",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateThemes() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateThemes",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateThresholdAlerts() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateThresholdAlerts",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateSharedFolders() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateSharedFolders",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateSpiceDataset() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateSpiceDataset",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetExportToCsv() {
	_jsii_.InvokeVoid(
		q,
		"resetExportToCsv",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetExportToCsvInScheduledReports() {
	_jsii_.InvokeVoid(
		q,
		"resetExportToCsvInScheduledReports",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetExportToExcel() {
	_jsii_.InvokeVoid(
		q,
		"resetExportToExcel",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetExportToExcelInScheduledReports() {
	_jsii_.InvokeVoid(
		q,
		"resetExportToExcelInScheduledReports",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetExportToPdf() {
	_jsii_.InvokeVoid(
		q,
		"resetExportToPdf",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetExportToPdfInScheduledReports() {
	_jsii_.InvokeVoid(
		q,
		"resetExportToPdfInScheduledReports",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetIncludeContentInScheduledReportsEmail() {
	_jsii_.InvokeVoid(
		q,
		"resetIncludeContentInScheduledReportsEmail",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetPrintReports() {
	_jsii_.InvokeVoid(
		q,
		"resetPrintReports",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetRenameSharedFolders() {
	_jsii_.InvokeVoid(
		q,
		"resetRenameSharedFolders",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareAnalyses() {
	_jsii_.InvokeVoid(
		q,
		"resetShareAnalyses",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareDashboards() {
	_jsii_.InvokeVoid(
		q,
		"resetShareDashboards",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareDatasets() {
	_jsii_.InvokeVoid(
		q,
		"resetShareDatasets",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareDataSources() {
	_jsii_.InvokeVoid(
		q,
		"resetShareDataSources",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSubscribeDashboardEmailReports() {
	_jsii_.InvokeVoid(
		q,
		"resetSubscribeDashboardEmailReports",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetViewAccountSpiceCapacity() {
	_jsii_.InvokeVoid(
		q,
		"resetViewAccountSpiceCapacity",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := q.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

