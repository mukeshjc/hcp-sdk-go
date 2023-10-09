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

// TaxSettingsIDType tax settings Id type
//
// swagger:model TaxSettingsIdType
type TaxSettingsIDType string

func NewTaxSettingsIDType(value TaxSettingsIDType) *TaxSettingsIDType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated TaxSettingsIDType.
func (m TaxSettingsIDType) Pointer() *TaxSettingsIDType {
	return &m
}

const (

	// TaxSettingsIDTypeUNSET captures enum value "UNSET"
	TaxSettingsIDTypeUNSET TaxSettingsIDType = "UNSET"

	// TaxSettingsIDTypeAETRN captures enum value "AE_TRN"
	TaxSettingsIDTypeAETRN TaxSettingsIDType = "AE_TRN"

	// TaxSettingsIDTypeAUABN captures enum value "AU_ABN"
	TaxSettingsIDTypeAUABN TaxSettingsIDType = "AU_ABN"

	// TaxSettingsIDTypeBRCNPJ captures enum value "BR_CNPJ"
	TaxSettingsIDTypeBRCNPJ TaxSettingsIDType = "BR_CNPJ"

	// TaxSettingsIDTypeBRCPF captures enum value "BR_CPF"
	TaxSettingsIDTypeBRCPF TaxSettingsIDType = "BR_CPF"

	// TaxSettingsIDTypeCABN captures enum value "CA_BN"
	TaxSettingsIDTypeCABN TaxSettingsIDType = "CA_BN"

	// TaxSettingsIDTypeCAQST captures enum value "CA_QST"
	TaxSettingsIDTypeCAQST TaxSettingsIDType = "CA_QST"

	// TaxSettingsIDTypeCHVAT captures enum value "CH_VAT"
	TaxSettingsIDTypeCHVAT TaxSettingsIDType = "CH_VAT"

	// TaxSettingsIDTypeCLTIN captures enum value "CL_TIN"
	TaxSettingsIDTypeCLTIN TaxSettingsIDType = "CL_TIN"

	// TaxSettingsIDTypeESCIF captures enum value "ES_CIF"
	TaxSettingsIDTypeESCIF TaxSettingsIDType = "ES_CIF"

	// TaxSettingsIDTypeEUVAT captures enum value "EU_VAT"
	TaxSettingsIDTypeEUVAT TaxSettingsIDType = "EU_VAT"

	// TaxSettingsIDTypeHKBR captures enum value "HK_BR"
	TaxSettingsIDTypeHKBR TaxSettingsIDType = "HK_BR"

	// TaxSettingsIDTypeIDNPWP captures enum value "ID_NPWP"
	TaxSettingsIDTypeIDNPWP TaxSettingsIDType = "ID_NPWP"

	// TaxSettingsIDTypeINGST captures enum value "IN_GST"
	TaxSettingsIDTypeINGST TaxSettingsIDType = "IN_GST"

	// TaxSettingsIDTypeJPCN captures enum value "JP_CN"
	TaxSettingsIDTypeJPCN TaxSettingsIDType = "JP_CN"

	// TaxSettingsIDTypeKRBRN captures enum value "KR_BRN"
	TaxSettingsIDTypeKRBRN TaxSettingsIDType = "KR_BRN"

	// TaxSettingsIDTypeLIUID captures enum value "LI_UID"
	TaxSettingsIDTypeLIUID TaxSettingsIDType = "LI_UID"

	// TaxSettingsIDTypeMXRFC captures enum value "MX_RFC"
	TaxSettingsIDTypeMXRFC TaxSettingsIDType = "MX_RFC"

	// TaxSettingsIDTypeMYFRP captures enum value "MY_FRP"
	TaxSettingsIDTypeMYFRP TaxSettingsIDType = "MY_FRP"

	// TaxSettingsIDTypeMYITN captures enum value "MY_ITN"
	TaxSettingsIDTypeMYITN TaxSettingsIDType = "MY_ITN"

	// TaxSettingsIDTypeMYSST captures enum value "MY_SST"
	TaxSettingsIDTypeMYSST TaxSettingsIDType = "MY_SST"

	// TaxSettingsIDTypeNOVAT captures enum value "NO_VAT"
	TaxSettingsIDTypeNOVAT TaxSettingsIDType = "NO_VAT"

	// TaxSettingsIDTypeNZGST captures enum value "NZ_GST"
	TaxSettingsIDTypeNZGST TaxSettingsIDType = "NZ_GST"

	// TaxSettingsIDTypeRUINN captures enum value "RU_INN"
	TaxSettingsIDTypeRUINN TaxSettingsIDType = "RU_INN"

	// TaxSettingsIDTypeSAVAT captures enum value "SA_VAT"
	TaxSettingsIDTypeSAVAT TaxSettingsIDType = "SA_VAT"

	// TaxSettingsIDTypeSGGST captures enum value "SG_GST"
	TaxSettingsIDTypeSGGST TaxSettingsIDType = "SG_GST"

	// TaxSettingsIDTypeSGUEN captures enum value "SG_UEN"
	TaxSettingsIDTypeSGUEN TaxSettingsIDType = "SG_UEN"

	// TaxSettingsIDTypeTHVAT captures enum value "TH_VAT"
	TaxSettingsIDTypeTHVAT TaxSettingsIDType = "TH_VAT"

	// TaxSettingsIDTypeTWVAT captures enum value "TW_VAT"
	TaxSettingsIDTypeTWVAT TaxSettingsIDType = "TW_VAT"

	// TaxSettingsIDTypeUSEIN captures enum value "US_EIN"
	TaxSettingsIDTypeUSEIN TaxSettingsIDType = "US_EIN"

	// TaxSettingsIDTypeZAVAT captures enum value "ZA_VAT"
	TaxSettingsIDTypeZAVAT TaxSettingsIDType = "ZA_VAT"

	// TaxSettingsIDTypeGBVAT captures enum value "GB_VAT"
	TaxSettingsIDTypeGBVAT TaxSettingsIDType = "GB_VAT"
)

// for schema
var taxSettingsIdTypeEnum []interface{}

func init() {
	var res []TaxSettingsIDType
	if err := json.Unmarshal([]byte(`["UNSET","AE_TRN","AU_ABN","BR_CNPJ","BR_CPF","CA_BN","CA_QST","CH_VAT","CL_TIN","ES_CIF","EU_VAT","HK_BR","ID_NPWP","IN_GST","JP_CN","KR_BRN","LI_UID","MX_RFC","MY_FRP","MY_ITN","MY_SST","NO_VAT","NZ_GST","RU_INN","SA_VAT","SG_GST","SG_UEN","TH_VAT","TW_VAT","US_EIN","ZA_VAT","GB_VAT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taxSettingsIdTypeEnum = append(taxSettingsIdTypeEnum, v)
	}
}

func (m TaxSettingsIDType) validateTaxSettingsIDTypeEnum(path, location string, value TaxSettingsIDType) error {
	if err := validate.EnumCase(path, location, value, taxSettingsIdTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this tax settings Id type
func (m TaxSettingsIDType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTaxSettingsIDTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this tax settings Id type based on context it is used
func (m TaxSettingsIDType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
