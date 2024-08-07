// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package acmcertificate

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AcmCertificateRenewalSummaryList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AcmCertificateRenewalSummaryList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AcmCertificateRenewalSummaryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AcmCertificateRenewalSummaryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AcmCertificateRenewalSummaryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AcmCertificateRenewalSummaryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAcmCertificateRenewalSummaryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

