// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cognitoidentity_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/stowelly/aws-sdk-go/aws"
	"github.com/stowelly/aws-sdk-go/aws/session"
	"github.com/stowelly/aws-sdk-go/service/cognitoidentity"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCognitoIdentity_CreateIdentityPool() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.CreateIdentityPoolInput{
		AllowUnauthenticatedIdentities: aws.Bool(true),                 // Required
		IdentityPoolName:               aws.String("IdentityPoolName"), // Required
		CognitoIdentityProviders: []*cognitoidentity.Provider{
			{ // Required
				ClientId:     aws.String("ProviderClientId"),
				ProviderName: aws.String("ProviderName"),
			},
			// More values...
		},
		DeveloperProviderName: aws.String("DeveloperProviderName"),
		OpenIdConnectProviderARNs: []*string{
			aws.String("ARNString"), // Required
			// More values...
		},
		SamlProviderARNs: []*string{
			aws.String("ARNString"), // Required
			// More values...
		},
		SupportedLoginProviders: map[string]*string{
			"Key": aws.String("IdentityProviderId"), // Required
			// More values...
		},
	}
	resp, err := svc.CreateIdentityPool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_DeleteIdentities() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.DeleteIdentitiesInput{
		IdentityIdsToDelete: []*string{ // Required
			aws.String("IdentityId"), // Required
			// More values...
		},
	}
	resp, err := svc.DeleteIdentities(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_DeleteIdentityPool() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.DeleteIdentityPoolInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.DeleteIdentityPool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_DescribeIdentity() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.DescribeIdentityInput{
		IdentityId: aws.String("IdentityId"), // Required
	}
	resp, err := svc.DescribeIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_DescribeIdentityPool() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.DescribeIdentityPoolInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.DescribeIdentityPool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_GetCredentialsForIdentity() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.GetCredentialsForIdentityInput{
		IdentityId:    aws.String("IdentityId"), // Required
		CustomRoleArn: aws.String("ARNString"),
		Logins: map[string]*string{
			"Key": aws.String("IdentityProviderToken"), // Required
			// More values...
		},
	}
	resp, err := svc.GetCredentialsForIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_GetId() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.GetIdInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
		AccountId:      aws.String("AccountId"),
		Logins: map[string]*string{
			"Key": aws.String("IdentityProviderToken"), // Required
			// More values...
		},
	}
	resp, err := svc.GetId(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_GetIdentityPoolRoles() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.GetIdentityPoolRolesInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.GetIdentityPoolRoles(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_GetOpenIdToken() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.GetOpenIdTokenInput{
		IdentityId: aws.String("IdentityId"), // Required
		Logins: map[string]*string{
			"Key": aws.String("IdentityProviderToken"), // Required
			// More values...
		},
	}
	resp, err := svc.GetOpenIdToken(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_GetOpenIdTokenForDeveloperIdentity() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
		Logins: map[string]*string{ // Required
			"Key": aws.String("IdentityProviderToken"), // Required
			// More values...
		},
		IdentityId:    aws.String("IdentityId"),
		TokenDuration: aws.Int64(1),
	}
	resp, err := svc.GetOpenIdTokenForDeveloperIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_ListIdentities() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.ListIdentitiesInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
		MaxResults:     aws.Int64(1),                 // Required
		HideDisabled:   aws.Bool(true),
		NextToken:      aws.String("PaginationKey"),
	}
	resp, err := svc.ListIdentities(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_ListIdentityPools() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.ListIdentityPoolsInput{
		MaxResults: aws.Int64(1), // Required
		NextToken:  aws.String("PaginationKey"),
	}
	resp, err := svc.ListIdentityPools(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_LookupDeveloperIdentity() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.LookupDeveloperIdentityInput{
		IdentityPoolId:          aws.String("IdentityPoolId"), // Required
		DeveloperUserIdentifier: aws.String("DeveloperUserIdentifier"),
		IdentityId:              aws.String("IdentityId"),
		MaxResults:              aws.Int64(1),
		NextToken:               aws.String("PaginationKey"),
	}
	resp, err := svc.LookupDeveloperIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_MergeDeveloperIdentities() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.MergeDeveloperIdentitiesInput{
		DestinationUserIdentifier: aws.String("DeveloperUserIdentifier"), // Required
		DeveloperProviderName:     aws.String("DeveloperProviderName"),   // Required
		IdentityPoolId:            aws.String("IdentityPoolId"),          // Required
		SourceUserIdentifier:      aws.String("DeveloperUserIdentifier"), // Required
	}
	resp, err := svc.MergeDeveloperIdentities(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_SetIdentityPoolRoles() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.SetIdentityPoolRolesInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
		Roles: map[string]*string{ // Required
			"Key": aws.String("ARNString"), // Required
			// More values...
		},
		RoleMappings: map[string]*cognitoidentity.RoleMapping{
			"Key": { // Required
				Type: aws.String("RoleMappingType"), // Required
				AmbiguousRoleResolution: aws.String("AmbiguousRoleResolutionType"),
				RulesConfiguration: &cognitoidentity.RulesConfigurationType{
					Rules: []*cognitoidentity.MappingRule{ // Required
						{ // Required
							Claim:     aws.String("ClaimName"),            // Required
							MatchType: aws.String("MappingRuleMatchType"), // Required
							RoleARN:   aws.String("ARNString"),            // Required
							Value:     aws.String("ClaimValue"),           // Required
						},
						// More values...
					},
				},
			},
			// More values...
		},
	}
	resp, err := svc.SetIdentityPoolRoles(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_UnlinkDeveloperIdentity() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.UnlinkDeveloperIdentityInput{
		DeveloperProviderName:   aws.String("DeveloperProviderName"),   // Required
		DeveloperUserIdentifier: aws.String("DeveloperUserIdentifier"), // Required
		IdentityId:              aws.String("IdentityId"),              // Required
		IdentityPoolId:          aws.String("IdentityPoolId"),          // Required
	}
	resp, err := svc.UnlinkDeveloperIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_UnlinkIdentity() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.UnlinkIdentityInput{
		IdentityId: aws.String("IdentityId"), // Required
		Logins: map[string]*string{ // Required
			"Key": aws.String("IdentityProviderToken"), // Required
			// More values...
		},
		LoginsToRemove: []*string{ // Required
			aws.String("IdentityProviderName"), // Required
			// More values...
		},
	}
	resp, err := svc.UnlinkIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_UpdateIdentityPool() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cognitoidentity.New(sess)

	params := &cognitoidentity.IdentityPool{
		AllowUnauthenticatedIdentities: aws.Bool(true),                 // Required
		IdentityPoolId:                 aws.String("IdentityPoolId"),   // Required
		IdentityPoolName:               aws.String("IdentityPoolName"), // Required
		CognitoIdentityProviders: []*cognitoidentity.Provider{
			{ // Required
				ClientId:     aws.String("ProviderClientId"),
				ProviderName: aws.String("ProviderName"),
			},
			// More values...
		},
		DeveloperProviderName: aws.String("DeveloperProviderName"),
		OpenIdConnectProviderARNs: []*string{
			aws.String("ARNString"), // Required
			// More values...
		},
		SamlProviderARNs: []*string{
			aws.String("ARNString"), // Required
			// More values...
		},
		SupportedLoginProviders: map[string]*string{
			"Key": aws.String("IdentityProviderId"), // Required
			// More values...
		},
	}
	resp, err := svc.UpdateIdentityPool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
