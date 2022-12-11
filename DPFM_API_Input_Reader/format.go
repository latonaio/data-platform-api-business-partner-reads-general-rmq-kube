package dpfm_api_input_reader

import (
	"data-platform-api-business-partner-reads-general-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToGeneral() *requests.General {
	data := sdc.General
	return &requests.General{
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
}

func (sdc *SDC) ConvertToGeneralPDF() *requests.GeneralPDF {
	dataGeneral := sdc.General
	data := sdc.General.GeneralPDF
	return &requests.GeneralPDF{
		BusinessPartner:          dataGeneral.BusinessPartner,
		DocType:                  data.DocType,
		DocVersionID:             data.DocVersionID,
		DocID:                    data.DocID,
		DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
		FileName:                 data.FileName,
	}
}

func (sdc *SDC) ConvertToRole() *requests.Role {
	dataGeneral := sdc.General
	data := sdc.General.Role
	return &requests.Role{
		BusinessPartner:     dataGeneral.BusinessPartner,
		BusinessPartnerRole: data.BusinessPartnerRole,
		ValidityEndDate:     data.ValidityEndDate,
		ValidityStartDate:   data.ValidityStartDate,
	}
}

func (sdc *SDC) ConvertToFinInst() *requests.FinInst {
	dataGeneral := sdc.General
	data := sdc.General.FinInst
	return &requests.FinInst{
		BusinessPartner:           dataGeneral.BusinessPartner,
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
}

func (sdc *SDC) ConvertToRelationship() *requests.Relationship {
	dataGeneral := sdc.General
	data := sdc.General.Relationship
	return &requests.Relationship{
		BusinessPartner:             dataGeneral.BusinessPartner,
		RelationshipNumber:          data.RelationshipNumber,
		ValidityEndDate:             data.ValidityEndDate,
		ValidityStartDate:           data.ValidityStartDate,
		RelationshipCategory:        data.RelationshipCategory,
		RelationshipBusinessPartner: data.RelationshipBusinessPartner,
		BusinessPartnerPerson:       data.BusinessPartnerPerson,
		IsStandardRelationship:      data.IsStandardRelationship,
		IsMarkedForDeletion:         data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToAccounting() *requests.Accounting {
	dataGeneral := sdc.General
	data := sdc.General.Accounting
	return &requests.Accounting{
		BusinessPartner:     dataGeneral.BusinessPartner,
		ChartOfAccounts:     data.ChartOfAccounts,
		FiscalYearVariant:   data.FiscalYearVariant,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}
