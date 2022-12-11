package dpfm_api_output_formatter

import (
	"data-platform-api-business-partner-reads-general-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-business-partner-reads-general-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToGeneral(sdc *api_input_reader.SDC, rows *sql.Rows) (*General, error) {
	pm := &requests.General{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.Industry,
			&pm.LegalEntityRegistration,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.OrganizationBPName1,
			&pm.OrganizationBPName2,
			&pm.OrganizationBPName3,
			&pm.OrganizationBPName4,
			&pm.BPGroup1,
			&pm.BPGroup2,
			&pm.BPGroup3,
			&pm.BPGroup4,
			&pm.BPGroup5,
			&pm.OrganizationFoundationDate,
			&pm.OrganizationLiquidationDate,
			&pm.SearchTerm1,
			&pm.SearchTerm2,
			&pm.BusinessPartnerBirthplaceName,
			&pm.BusinessPartnerDeathDate,
			&pm.BusinessPartnerIsBlocked,
			&pm.GroupBusinessPartnerName1,
			&pm.GroupBusinessPartnerName2,
			&pm.AddressID,
			&pm.BusinessPartnerIDByExtSystem,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	general := &General{
		BusinessPartner:               data.BusinessPartner,
		BusinessPartnerFullName:       data.BusinessPartnerFullName,
		BusinessPartnerName:           data.BusinessPartnerName,
		CreationDate:                  data.CreationDate,
		CreationTime:                  data.CreationTime,
		Industry:                      data.Industry,
		LegalEntityRegistration:       data.LegalEntityRegistration,
		Country:                       data.Country,
		Language:                      data.Language,
		Currency:                      data.Currency,
		LastChangeDate:                data.LastChangeDate,
		LastChangeTime:                data.LastChangeTime,
		OrganizationBPName1:           data.OrganizationBPName1,
		OrganizationBPName2:           data.OrganizationBPName2,
		OrganizationBPName3:           data.OrganizationBPName3,
		OrganizationBPName4:           data.OrganizationBPName4,
		BPGroup1:                      data.BPGroup1,
		BPGroup2:                      data.BPGroup2,
		BPGroup3:                      data.BPGroup3,
		BPGroup4:                      data.BPGroup4,
		BPGroup5:                      data.BPGroup5,
		OrganizationFoundationDate:    data.OrganizationFoundationDate,
		OrganizationLiquidationDate:   data.OrganizationLiquidationDate,
		SearchTerm1:                   data.SearchTerm1,
		SearchTerm2:                   data.SearchTerm2,
		BusinessPartnerBirthplaceName: data.BusinessPartnerBirthplaceName,
		BusinessPartnerDeathDate:      data.BusinessPartnerDeathDate,
		BusinessPartnerIsBlocked:      data.BusinessPartnerIsBlocked,
		GroupBusinessPartnerName1:     data.GroupBusinessPartnerName1,
		GroupBusinessPartnerName2:     data.GroupBusinessPartnerName2,
		AddressID:                     data.AddressID,
		BusinessPartnerIDByExtSystem:  data.BusinessPartnerIDByExtSystem,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
	return general, nil
}

func ConvertToGeneralPDF(sdc *api_input_reader.SDC, rows *sql.Rows) (*GeneralPDF, error) {
	pm := &requests.GeneralPDF{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.DocType,
			&pm.DocVersionID,
			&pm.DocID,
			&pm.DocIssuerBusinessPartner,
			&pm.FileName,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	generalPDF := &GeneralPDF{
		BusinessPartner:          data.BusinessPartner,
		DocType:                  data.DocType,
		DocVersionID:             data.DocVersionID,
		DocID:                    data.DocID,
		DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
		FileName:                 data.FileName,
	}
	return generalPDF, nil
}

func ConvertToRole(sdc *api_input_reader.SDC, rows *sql.Rows) (*Role, error) {
	pm := &requests.Role{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.BusinessPartnerRole,
			&pm.ValidityEndDate,
			&pm.ValidityStartDate,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	role := &Role{
		BusinessPartner:     data.BusinessPartner,
		BusinessPartnerRole: data.BusinessPartnerRole,
		ValidityEndDate:     data.ValidityEndDate,
		ValidityStartDate:   data.ValidityStartDate,
	}
	return role, nil
}

func ConvertToFinInst(sdc *api_input_reader.SDC, rows *sql.Rows) (*FinInst, error) {
	pm := &requests.FinInst{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.FinInstIdentification,
			&pm.ValidityEndDate,
			&pm.ValidityStartDate,
			&pm.FinInstCountry,
			&pm.FinInstNumber,
			&pm.FinInstName,
			&pm.FinInstBranchName,
			&pm.SWIFTCode,
			&pm.InternalFinInstCustomerID,
			&pm.InternalFinInstAccountID,
			&pm.FinInstControlKey,
			&pm.FinInstAccountName,
			&pm.FinInstAccount,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	finInst := &FinInst{
		BusinessPartner:           data.BusinessPartner,
		FinInstIdentification:     data.FinInstIdentification,
		ValidityEndDate:           data.ValidityEndDate,
		ValidityStartDate:         data.ValidityStartDate,
		FinInstCountry:            data.FinInstCountry,
		FinInstNumber:             data.FinInstNumber,
		FinInstName:               data.FinInstName,
		FinInstBranchName:         data.FinInstBranchName,
		SWIFTCode:                 data.SWIFTCode,
		InternalFinInstCustomerID: data.InternalFinInstCustomerID,
		InternalFinInstAccountID:  data.InternalFinInstAccountID,
		FinInstControlKey:         data.FinInstControlKey,
		FinInstAccountName:        data.FinInstAccountName,
		FinInstAccount:            data.FinInstAccount,
		IsMarkedForDeletion:       data.IsMarkedForDeletion,
	}
	return finInst, nil
}

func ConvertToRelationship(sdc *api_input_reader.SDC, rows *sql.Rows) (*Relationship, error) {
	pm := &requests.Relationship{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.RelationshipNumber,
			&pm.ValidityEndDate,
			&pm.ValidityStartDate,
			&pm.RelationshipCategory,
			&pm.RelationshipBusinessPartner,
			&pm.BusinessPartnerPerson,
			&pm.IsStandardRelationship,
			&pm.IsMarkedForDeletion)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	relationship := &Relationship{
		BusinessPartner:             data.BusinessPartner,
		RelationshipNumber:          data.RelationshipNumber,
		ValidityEndDate:             data.ValidityEndDate,
		ValidityStartDate:           data.ValidityStartDate,
		RelationshipCategory:        data.RelationshipCategory,
		RelationshipBusinessPartner: data.RelationshipBusinessPartner,
		BusinessPartnerPerson:       data.BusinessPartnerPerson,
		IsStandardRelationship:      data.IsStandardRelationship,
		IsMarkedForDeletion:         data.IsMarkedForDeletion,
	}
	return relationship, nil
}

func ConvertToAccounting(sdc *api_input_reader.SDC, rows *sql.Rows) (*Accounting, error) {
	pm := &requests.Accounting{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.ChartOfAccounts,
			&pm.FiscalYearVariant,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	accounting := &Accounting{
		BusinessPartner:     data.BusinessPartner,
		ChartOfAccounts:     data.ChartOfAccounts,
		FiscalYearVariant:   data.FiscalYearVariant,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
	return accounting, nil
}
