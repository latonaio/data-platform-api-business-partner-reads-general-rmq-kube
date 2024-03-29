package dpfm_api_output_formatter

import (
	"data-platform-api-business-partner-reads-general-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToGeneral(rows *sql.Rows) (*General, error) {
	defer rows.Close()
	pm := &requests.General{}

	i := 0
	for rows.Next() {
		i++
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
			&pm.BPTag1,
			&pm.BPTag2,
			&pm.BPTag3,
			&pm.BPTag4,
			&pm.OrganizationFoundationDate,
			&pm.OrganizationLiquidationDate,
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
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
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
		BPTag1:                        data.BPTag1,
		BPTag2:                        data.BPTag2,
		BPTag3:                        data.BPTag3,
		BPTag4:                        data.BPTag4,
		OrganizationFoundationDate:    data.OrganizationFoundationDate,
		OrganizationLiquidationDate:   data.OrganizationLiquidationDate,
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

func ConvertToGenerals(rows *sql.Rows) (*[]General, error) {
	defer rows.Close()
	general := make([]General, 0)

	i := 0
	for rows.Next() {
		i++
		pm := General{}

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
			&pm.BPTag1,
			&pm.BPTag2,
			&pm.BPTag3,
			&pm.BPTag4,
			&pm.OrganizationFoundationDate,
			&pm.OrganizationLiquidationDate,
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

		data := pm
		general = append(general, General{
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
			BPTag1:                        data.BPTag1,
			BPTag2:                        data.BPTag2,
			BPTag3:                        data.BPTag3,
			BPTag4:                        data.BPTag4,
			OrganizationFoundationDate:    data.OrganizationFoundationDate,
			OrganizationLiquidationDate:   data.OrganizationLiquidationDate,
			BusinessPartnerBirthplaceName: data.BusinessPartnerBirthplaceName,
			BusinessPartnerDeathDate:      data.BusinessPartnerDeathDate,
			BusinessPartnerIsBlocked:      data.BusinessPartnerIsBlocked,
			GroupBusinessPartnerName1:     data.GroupBusinessPartnerName1,
			GroupBusinessPartnerName2:     data.GroupBusinessPartnerName2,
			AddressID:                     data.AddressID,
			BusinessPartnerIDByExtSystem:  data.BusinessPartnerIDByExtSystem,
			IsMarkedForDeletion:           data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &general, nil
}

func ConvertToGeneralPDF(rows *sql.Rows) (*GeneralPDF, error) {
	defer rows.Close()
	pm := &requests.GeneralPDF{}

	i := 0
	for rows.Next() {
		i++
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
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
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

func ConvertToRole(rows *sql.Rows) (*Role, error) {
	defer rows.Close()
	pm := &requests.Role{}

	i := 0
	for rows.Next() {
		i++
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
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
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

func ConvertToFinInst(rows *sql.Rows) (*FinInst, error) {
	defer rows.Close()
	pm := &requests.FinInst{}

	i := 0
	for rows.Next() {
		i++
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.FinInstIdentification,
			&pm.ValidityEndDate,
			&pm.ValidityStartDate,
			&pm.FinInstCountry,
			&pm.FinInstCode,
			&pm.FinInstBranchCode,
			&pm.FinInstFullCode,
			&pm.FinInstName,
			&pm.FinInstBranchName,
			&pm.SWIFTCode,
			&pm.InternalFinInstCustomerID,
			&pm.InternalFinInstAccountID,
			&pm.FinInstControlKey,
			&pm.FinInstAccountName,
			&pm.FinInstAccount,
			&pm.HouseBank,
			&pm.HouseBankAccount,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	data := pm
	finInst := &FinInst{
		BusinessPartner:           data.BusinessPartner,
		FinInstIdentification:     data.FinInstIdentification,
		ValidityEndDate:           data.ValidityEndDate,
		ValidityStartDate:         data.ValidityStartDate,
		FinInstCountry:            data.FinInstCountry,
		FinInstCode:               data.FinInstCode,
		FinInstBranchCode:         data.FinInstBranchCode,
		FinInstFullCode:           data.FinInstFullCode,
		FinInstName:               data.FinInstName,
		FinInstBranchName:         data.FinInstBranchName,
		SWIFTCode:                 data.SWIFTCode,
		InternalFinInstCustomerID: data.InternalFinInstCustomerID,
		InternalFinInstAccountID:  data.InternalFinInstAccountID,
		FinInstControlKey:         data.FinInstControlKey,
		FinInstAccountName:        data.FinInstAccountName,
		FinInstAccount:            data.FinInstAccount,
		HouseBank:                 data.HouseBank,
		HouseBankAccount:          data.HouseBankAccount,
		IsMarkedForDeletion:       data.IsMarkedForDeletion,
	}

	return finInst, nil
}

func ConvertToRelationship(rows *sql.Rows) (*Relationship, error) {
	defer rows.Close()
	pm := &requests.Relationship{}

	i := 0
	for rows.Next() {
		i++
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
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
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

func ConvertToAccounting(rows *sql.Rows) (*Accounting, error) {
	defer rows.Close()
	pm := &requests.Accounting{}

	i := 0
	for rows.Next() {
		i++
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
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
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
