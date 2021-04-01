package main

import (
	"fmt"
	fc "github.com/fullcontact/fullcontact-go/fc"
	"log"
)

type CustomRetryHandler struct{}

func (crh CustomRetryHandler) ShouldRetry(responseCode int) bool {
	if responseCode == 429 {
		return true
	}
	return false
}

func (crh CustomRetryHandler) RetryAttempts() int {
	return 2
}

func (crh CustomRetryHandler) RetryDelayMillis() int {
	return 2000
}

func main() {
	//Creating a Credentials Provider
	cp, err := fc.NewStaticCredentialsProvider("your-api-key")
	if err != nil {
		log.Fatalln(err)
		return
	}

	//Creating a FullContact Client
	fcClient, err := fc.NewFullContactClient(
		fc.WithCredentialsProvider(cp),
		fc.WithHeaders(map[string]string{"Reporting-Key": "FC_GoClient_1.0.0"}),
		fc.WithTimeout(3000),
		fc.WithRetryHandler(&CustomRetryHandler{}))

	if err != nil {
		log.Fatalln(err)
		return
	}

	//Person Enrich
	profile, err := fc.NewProfile(
		fc.WithUsername("bartlorang"),
		fc.WithService("twitter"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	personRequest, err := fc.NewPersonRequest(
		fc.WithEmail("bart@fullcontact.com"),
		fc.WithProfile(profile))
	if err != nil {
		log.Fatalln(err)
		return
	}
	//Sending Person Enrich request which returns a channel of type `APIResponse`
	ch := fcClient.PersonEnrich(personRequest)
	resp := <-ch
	fmt.Printf("Person Enrich API Response: %v", resp)
	if resp.IsSuccessful == true {
		fmt.Printf("Person Response: %v", *resp.PersonResponse)
		fmt.Println(resp.PersonResponse.FullName)
		fmt.Println(resp.PersonResponse.Details.Identifiers.LiNonId)
	}

	//Company Enrich
	companyEnrichRequest, err := fc.NewCompanyRequest(fc.WithDomain("fullcontact.com"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	resp = <-fcClient.CompanyEnrich(companyEnrichRequest)
	fmt.Printf("\n\nCompany Enrich API Response: %v", resp)
	if resp.IsSuccessful {
		fmt.Printf("Company Name: %v", resp.CompanyResponse.Name)
	}

	//Company Search
	companySearchRequest, err := fc.NewCompanyRequest(fc.WithCompanyName("FullContact"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	resp = <-fcClient.CompanySearch(companySearchRequest)
	fmt.Printf("\n\nCompany Search API Response: %v", resp)
	if resp.IsSuccessful {
		fmt.Printf("Company Lookup Domain: %v", resp.CompanySearchResponse[0].LookupDomain)
	}

	//Resolve
	//Identity Map
	resolveRequest, err := fc.NewResolveRequest(
		fc.WithRecordIdForResolve("r1"),
		fc.WithEmailForResolve("bart@fullcontact.com"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	resp = <-fcClient.IdentityMap(resolveRequest)
	fmt.Printf("\n\nIdentity Map API Response: %v", resp)
	if resp.IsSuccessful {
		fmt.Printf("RecordIds Mapped: %v", resp.ResolveResponse.RecordIds)
	}

	//Identity Resolve
	resolveRequest, err = fc.NewResolveRequest(fc.WithRecordIdForResolve("r1"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	resp = <-fcClient.IdentityResolve(resolveRequest)
	fmt.Printf("\n\nIdentity Resolve API Response: %v", resp)
	if resp.IsSuccessful {
		fmt.Printf("PersonIds Mapped: %v", resp.ResolveResponse.PersonIds)
	}

	//Identity Resolve With Tags
	resolveRequest, err = fc.NewResolveRequest(fc.WithRecordIdForResolve("r1"))
	if err != nil {
		log.Fatalln(err)
		return
	}

	resp = <-fcClient.IdentityResolveWithTags(resolveRequest)
	fmt.Printf("\n\nIdentity Resolve with Tags API Response: %v", resp)
	if resp.IsSuccessful {
		fmt.Printf("Tags: %v", resp.ResolveResponseWithTags.Tags)
	}

	//Identity Delete
	resolveRequest, err = fc.NewResolveRequest(fc.WithRecordIdForResolve("r1"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	resp = <-fcClient.IdentityDelete(resolveRequest)
	fmt.Printf("\n\nIdentity Delete API Response: %v", resp)
	if resp.IsSuccessful {
		fmt.Println("Record Deleted Successfully!")
	}

	//Tags Create
	tagsRequest, err := fc.NewTagsRequest(fc.WithRecordIdForTags("k1"),
		fc.WithTag(
			fc.NewTag(fc.WithTagKey("gender"), fc.WithTagValue("male"))))
	if err != nil {
		log.Fatalln(err)
		return
	}
	resp = <-fcClient.TagsCreate(tagsRequest)
	fmt.Printf("\n\nTags Create API Response: %v", resp.TagsResponse)

	//Tags Get
	resp = <-fcClient.TagsGet("k1")
	fmt.Printf("\n\nTags Get API Response: %v", resp.TagsResponse)

	//Tags Delete
	resp = <-fcClient.TagsDelete(tagsRequest)
	fmt.Printf("\n\nTags Delete API Response: %v", resp.Status)

	//Audience Create
	audienceRequest, err := fc.NewAudienceRequest(fc.WithWebhookUrlForAudience("your-webhookUrl"),
		fc.WithTagForAudience(fc.NewTag(fc.WithTagKey("gender"), fc.WithTagValue("male"))))
	if err != nil {
		log.Fatalln(err)
		return
	}

	resp = <-fcClient.AudienceCreate(audienceRequest)
	fmt.Printf("\n\nAudience Create API Response: %v", resp.AudienceResponse)
	if resp.IsSuccessful {
		fmt.Println(resp.AudienceResponse.RequestId)
	}

	//Audience Download
	requestId := "730000fd-009a-00fc-8008-100e000085f0"
	resp = <-fcClient.AudienceDownload(requestId)
	fmt.Printf("\n\nAudience Download API Response: %v", resp.AudienceResponse)
	if resp.IsSuccessful {
		fmt.Println(resp.AudienceResponse.WriteAudienceBytesToFile(requestId + "_audienceFile.json.gz"))
	}

	//Email Verification
	resp = <-fcClient.EmailVerification("bart@fullcontact.com")
	fmt.Printf("\n\nEmail Verification API Response: %v", resp)
	if resp.IsSuccessful {
		fmt.Println(resp.EmailVerificationResponse)
	}

	//Permission
	//Permission Create

	consentPurposes, err := fc.NewConsentPurposes(
		fc.WithConsentPurposeId(1),
		fc.WithConsentPurposeChannel("web"),
		fc.WithConsentPurposeTtl(365),
		fc.WithConsentPurposeEnabled(true))

	permissionCreateRequest, err := fc.NewPermissionRequest(
		fc.WithEmailForPermission("bart@fullcontact.com"),
		fc.WithConsentPurposeForPermission(consentPurposes),
		fc.WithCollectionMethodForPermission("cookiePopUp"),
		fc.WithCollectionLocationForPermission("Can we get a snapshot of where someone is opting in/out here?"),
		fc.WithPolicyUrlForPermission("https://www.fullcontact.com/privacy/privacy-policy"),
		fc.WithTermsServiceForPermission("https://www.fullcontact.com/privacy/terms-of-use"))
	if err != nil {
		log.Fatalln(err)
		return
	}

	//Sending Permission create request
	permissionCreateCh := fcClient.PermissionCreate(permissionCreateRequest)
	permissionCreateResp := <-permissionCreateCh
	fmt.Printf("Permission Create API Response: %v", permissionCreateResp)
	if permissionCreateResp.IsSuccessful == true {
		fmt.Printf("Permission Create Response: %v", permissionCreateResp)
	}

	//Permission Find
	permissionFindRequest, err := fc.NewPermissionRequest(
		fc.WithEmailForPermission("bart@fullcontact.com"))

	//Sending Permission Find request which returns a channel of type `APIResponse`
	permissionFindCh := fcClient.PermissionFind(permissionFindRequest)
	permissionFindResp := <-permissionFindCh
	fmt.Printf("Permission Find API Response: %v", permissionFindResp)
	if permissionFindResp.IsSuccessful == true {
		fmt.Printf("Permission Find Response: %v", permissionFindResp.PermissionFindResponse)
	}

	//Permission Current
	permissionCurrentRequest, err := fc.NewPermissionRequest(
		fc.WithEmailForPermission("bart@fullcontact.com"))

	//Sending Permission Find request which returns a channel of type `APIResponse`
	permissionCurrentCh := fcClient.PermissionCurrent(permissionCurrentRequest)
	permissionCurrentResp := <-permissionCurrentCh
	fmt.Printf("Permission Current API Response: %v", permissionCurrentResp)
	if permissionCurrentResp.IsSuccessful == true {
		fmt.Printf("Permission Current Response: %v", permissionCurrentResp.PermissionCurrentResponse)
	}

	//Permission Verify
	permissionVerifyRequest, err := fc.NewPermissionRequest(
		fc.WithEmailForPermission("bart@fullcontact.com"),
		fc.WithPurposeIdForPermission(1),
		fc.WithChannelForPermission("web"))

	//Sending Permission Find request which returns a channel of type `APIResponse`
	permissionVerifyCh := fcClient.PermissionVerify(permissionVerifyRequest)
	permissionVerifyResp := <-permissionVerifyCh
	fmt.Printf("Permission Verify API Response: %v", permissionVerifyResp)
	if permissionVerifyResp.IsSuccessful == true {
		fmt.Printf("Permission Verify Response: %v", permissionVerifyResp.PermissionVerifyResponse)
	}

	//Permission Delete
	permissionDeleteRequest, err := fc.NewPermissionRequest(
		fc.WithEmailForPermission("bart@fullcontact.com"))

	//Sending Permission Find request which returns a channel of type `APIResponse`
	permissionDeleteCh := fcClient.PermissionDelete(permissionDeleteRequest)
	permissionDeleteResp := <-permissionDeleteCh
	if permissionDeleteResp.IsSuccessful == true {
		fmt.Printf("Permission Delete API Response: %v", permissionDeleteResp)
	}
}
