// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ClusterValidationID cluster validation id
//
// swagger:model cluster-validation-id
type ClusterValidationID string

func NewClusterValidationID(value ClusterValidationID) *ClusterValidationID {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ClusterValidationID.
func (m ClusterValidationID) Pointer() *ClusterValidationID {
	return &m
}

const (

	// ClusterValidationIDMachineCidrDefined captures enum value "machine-cidr-defined"
	ClusterValidationIDMachineCidrDefined ClusterValidationID = "machine-cidr-defined"

	// ClusterValidationIDClusterCidrDefined captures enum value "cluster-cidr-defined"
	ClusterValidationIDClusterCidrDefined ClusterValidationID = "cluster-cidr-defined"

	// ClusterValidationIDServiceCidrDefined captures enum value "service-cidr-defined"
	ClusterValidationIDServiceCidrDefined ClusterValidationID = "service-cidr-defined"

	// ClusterValidationIDNoCidrsOverlapping captures enum value "no-cidrs-overlapping"
	ClusterValidationIDNoCidrsOverlapping ClusterValidationID = "no-cidrs-overlapping"

	// ClusterValidationIDNetworksSameAddressFamilies captures enum value "networks-same-address-families"
	ClusterValidationIDNetworksSameAddressFamilies ClusterValidationID = "networks-same-address-families"

	// ClusterValidationIDNetworkPrefixValid captures enum value "network-prefix-valid"
	ClusterValidationIDNetworkPrefixValid ClusterValidationID = "network-prefix-valid"

	// ClusterValidationIDMachineCidrEqualsToCalculatedCidr captures enum value "machine-cidr-equals-to-calculated-cidr"
	ClusterValidationIDMachineCidrEqualsToCalculatedCidr ClusterValidationID = "machine-cidr-equals-to-calculated-cidr"

	// ClusterValidationIDAPIVipsDefined captures enum value "api-vips-defined"
	ClusterValidationIDAPIVipsDefined ClusterValidationID = "api-vips-defined"

	// ClusterValidationIDAPIVipsValid captures enum value "api-vips-valid"
	ClusterValidationIDAPIVipsValid ClusterValidationID = "api-vips-valid"

	// ClusterValidationIDIngressVipsDefined captures enum value "ingress-vips-defined"
	ClusterValidationIDIngressVipsDefined ClusterValidationID = "ingress-vips-defined"

	// ClusterValidationIDIngressVipsValid captures enum value "ingress-vips-valid"
	ClusterValidationIDIngressVipsValid ClusterValidationID = "ingress-vips-valid"

	// ClusterValidationIDAllHostsAreReadyToInstall captures enum value "all-hosts-are-ready-to-install"
	ClusterValidationIDAllHostsAreReadyToInstall ClusterValidationID = "all-hosts-are-ready-to-install"

	// ClusterValidationIDSufficientMastersCount captures enum value "sufficient-masters-count"
	ClusterValidationIDSufficientMastersCount ClusterValidationID = "sufficient-masters-count"

	// ClusterValidationIDDNSDomainDefined captures enum value "dns-domain-defined"
	ClusterValidationIDDNSDomainDefined ClusterValidationID = "dns-domain-defined"

	// ClusterValidationIDPullSecretSet captures enum value "pull-secret-set"
	ClusterValidationIDPullSecretSet ClusterValidationID = "pull-secret-set"

	// ClusterValidationIDNtpServerConfigured captures enum value "ntp-server-configured"
	ClusterValidationIDNtpServerConfigured ClusterValidationID = "ntp-server-configured"

	// ClusterValidationIDLsoRequirementsSatisfied captures enum value "lso-requirements-satisfied"
	ClusterValidationIDLsoRequirementsSatisfied ClusterValidationID = "lso-requirements-satisfied"

	// ClusterValidationIDOcsRequirementsSatisfied captures enum value "ocs-requirements-satisfied"
	ClusterValidationIDOcsRequirementsSatisfied ClusterValidationID = "ocs-requirements-satisfied"

	// ClusterValidationIDOdfRequirementsSatisfied captures enum value "odf-requirements-satisfied"
	ClusterValidationIDOdfRequirementsSatisfied ClusterValidationID = "odf-requirements-satisfied"

	// ClusterValidationIDCnvRequirementsSatisfied captures enum value "cnv-requirements-satisfied"
	ClusterValidationIDCnvRequirementsSatisfied ClusterValidationID = "cnv-requirements-satisfied"

	// ClusterValidationIDLvmRequirementsSatisfied captures enum value "lvm-requirements-satisfied"
	ClusterValidationIDLvmRequirementsSatisfied ClusterValidationID = "lvm-requirements-satisfied"

	// ClusterValidationIDMceRequirementsSatisfied captures enum value "mce-requirements-satisfied"
	ClusterValidationIDMceRequirementsSatisfied ClusterValidationID = "mce-requirements-satisfied"

	// ClusterValidationIDNetworkTypeValid captures enum value "network-type-valid"
	ClusterValidationIDNetworkTypeValid ClusterValidationID = "network-type-valid"
)

// for schema
var clusterValidationIdEnum []interface{}

func init() {
	var res []ClusterValidationID
	if err := json.Unmarshal([]byte(`["machine-cidr-defined","cluster-cidr-defined","service-cidr-defined","no-cidrs-overlapping","networks-same-address-families","network-prefix-valid","machine-cidr-equals-to-calculated-cidr","api-vips-defined","api-vips-valid","ingress-vips-defined","ingress-vips-valid","all-hosts-are-ready-to-install","sufficient-masters-count","dns-domain-defined","pull-secret-set","ntp-server-configured","lso-requirements-satisfied","ocs-requirements-satisfied","odf-requirements-satisfied","cnv-requirements-satisfied","lvm-requirements-satisfied","mce-requirements-satisfied","network-type-valid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterValidationIdEnum = append(clusterValidationIdEnum, v)
	}
}

func (m ClusterValidationID) validateClusterValidationIDEnum(path, location string, value ClusterValidationID) error {
	if err := validate.EnumCase(path, location, value, clusterValidationIdEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cluster validation id
func (m ClusterValidationID) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClusterValidationIDEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cluster validation id based on context it is used
func (m ClusterValidationID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
