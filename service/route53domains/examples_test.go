// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package route53domains_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/stowelly/aws-sdk-go/aws"
	"github.com/stowelly/aws-sdk-go/aws/session"
	"github.com/stowelly/aws-sdk-go/service/route53domains"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleRoute53Domains_CheckDomainAvailability() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.CheckDomainAvailabilityInput{
		DomainName:  aws.String("DomainName"), // Required
		IdnLangCode: aws.String("LangCode"),
	}
	resp, err := svc.CheckDomainAvailability(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_DeleteTagsForDomain() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.DeleteTagsForDomainInput{
		DomainName: aws.String("DomainName"), // Required
		TagsToDelete: []*string{ // Required
			aws.String("TagKey"), // Required
			// More values...
		},
	}
	resp, err := svc.DeleteTagsForDomain(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_DisableDomainAutoRenew() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.DisableDomainAutoRenewInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.DisableDomainAutoRenew(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_DisableDomainTransferLock() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.DisableDomainTransferLockInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.DisableDomainTransferLock(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_EnableDomainAutoRenew() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.EnableDomainAutoRenewInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.EnableDomainAutoRenew(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_EnableDomainTransferLock() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.EnableDomainTransferLockInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.EnableDomainTransferLock(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_GetContactReachabilityStatus() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.GetContactReachabilityStatusInput{
		DomainName: aws.String("DomainName"),
	}
	resp, err := svc.GetContactReachabilityStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_GetDomainDetail() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.GetDomainDetailInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.GetDomainDetail(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_GetDomainSuggestions() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.GetDomainSuggestionsInput{
		DomainName:      aws.String("DomainName"), // Required
		OnlyAvailable:   aws.Bool(true),           // Required
		SuggestionCount: aws.Int64(1),             // Required
	}
	resp, err := svc.GetDomainSuggestions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_GetOperationDetail() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.GetOperationDetailInput{
		OperationId: aws.String("OperationId"), // Required
	}
	resp, err := svc.GetOperationDetail(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_ListDomains() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.ListDomainsInput{
		Marker:   aws.String("PageMarker"),
		MaxItems: aws.Int64(1),
	}
	resp, err := svc.ListDomains(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_ListOperations() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.ListOperationsInput{
		Marker:   aws.String("PageMarker"),
		MaxItems: aws.Int64(1),
	}
	resp, err := svc.ListOperations(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_ListTagsForDomain() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.ListTagsForDomainInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.ListTagsForDomain(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_RegisterDomain() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.RegisterDomainInput{
		AdminContact: &route53domains.ContactDetail{ // Required
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
		DomainName:      aws.String("DomainName"), // Required
		DurationInYears: aws.Int64(1),             // Required
		RegistrantContact: &route53domains.ContactDetail{ // Required
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
		TechContact: &route53domains.ContactDetail{ // Required
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
		AutoRenew:                       aws.Bool(true),
		IdnLangCode:                     aws.String("LangCode"),
		PrivacyProtectAdminContact:      aws.Bool(true),
		PrivacyProtectRegistrantContact: aws.Bool(true),
		PrivacyProtectTechContact:       aws.Bool(true),
	}
	resp, err := svc.RegisterDomain(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_RenewDomain() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.RenewDomainInput{
		CurrentExpiryYear: aws.Int64(1),             // Required
		DomainName:        aws.String("DomainName"), // Required
		DurationInYears:   aws.Int64(1),
	}
	resp, err := svc.RenewDomain(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_ResendContactReachabilityEmail() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.ResendContactReachabilityEmailInput{
		DomainName: aws.String("DomainName"),
	}
	resp, err := svc.ResendContactReachabilityEmail(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_RetrieveDomainAuthCode() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.RetrieveDomainAuthCodeInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.RetrieveDomainAuthCode(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_TransferDomain() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.TransferDomainInput{
		AdminContact: &route53domains.ContactDetail{ // Required
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
		DomainName:      aws.String("DomainName"), // Required
		DurationInYears: aws.Int64(1),             // Required
		RegistrantContact: &route53domains.ContactDetail{ // Required
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
		TechContact: &route53domains.ContactDetail{ // Required
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
		AuthCode:    aws.String("DomainAuthCode"),
		AutoRenew:   aws.Bool(true),
		IdnLangCode: aws.String("LangCode"),
		Nameservers: []*route53domains.Nameserver{
			{ // Required
				Name: aws.String("HostName"), // Required
				GlueIps: []*string{
					aws.String("GlueIp"), // Required
					// More values...
				},
			},
			// More values...
		},
		PrivacyProtectAdminContact:      aws.Bool(true),
		PrivacyProtectRegistrantContact: aws.Bool(true),
		PrivacyProtectTechContact:       aws.Bool(true),
	}
	resp, err := svc.TransferDomain(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_UpdateDomainContact() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.UpdateDomainContactInput{
		DomainName: aws.String("DomainName"), // Required
		AdminContact: &route53domains.ContactDetail{
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
		RegistrantContact: &route53domains.ContactDetail{
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
		TechContact: &route53domains.ContactDetail{
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
	}
	resp, err := svc.UpdateDomainContact(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_UpdateDomainContactPrivacy() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.UpdateDomainContactPrivacyInput{
		DomainName:        aws.String("DomainName"), // Required
		AdminPrivacy:      aws.Bool(true),
		RegistrantPrivacy: aws.Bool(true),
		TechPrivacy:       aws.Bool(true),
	}
	resp, err := svc.UpdateDomainContactPrivacy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_UpdateDomainNameservers() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.UpdateDomainNameserversInput{
		DomainName: aws.String("DomainName"), // Required
		Nameservers: []*route53domains.Nameserver{ // Required
			{ // Required
				Name: aws.String("HostName"), // Required
				GlueIps: []*string{
					aws.String("GlueIp"), // Required
					// More values...
				},
			},
			// More values...
		},
		FIAuthKey: aws.String("FIAuthKey"),
	}
	resp, err := svc.UpdateDomainNameservers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_UpdateTagsForDomain() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.UpdateTagsForDomainInput{
		DomainName: aws.String("DomainName"), // Required
		TagsToUpdate: []*route53domains.Tag{
			{ // Required
				Key:   aws.String("TagKey"),
				Value: aws.String("TagValue"),
			},
			// More values...
		},
	}
	resp, err := svc.UpdateTagsForDomain(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53Domains_ViewBilling() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53domains.New(sess)

	params := &route53domains.ViewBillingInput{
		End:      aws.Time(time.Now()),
		Marker:   aws.String("PageMarker"),
		MaxItems: aws.Int64(1),
		Start:    aws.Time(time.Now()),
	}
	resp, err := svc.ViewBilling(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
