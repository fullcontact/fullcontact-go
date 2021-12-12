## FullContact Go Client
[![GoDoc](https://godoc.org/github.com/fullcontact/fullcontact-go/fc?status.svg)](https://godoc.org/github.com/fullcontact/fullcontact-go/fc)


The official [FullContact](https://www.fullcontact.com/) Golang Client Library for the FullContact V3 APIs.

## Installation

To install FullContact Go client, use `go get`:
```sh
go get github.com/fullcontact/fullcontact-go/fc
```

Then, reference fullcontact-go in a Go program with `import`:

``` go
import fc "github.com/fullcontact/fullcontact-go/fc"
```

To update FullContact Go client to the latest version use:
```sh
go get -u github.com/fullcontact/fullcontact-go/fc
```

## Working with FullContact Client
FullContact client supports v3 Enrich and Resolve APIs which are super simplified to easily 
enrich Person and Company data and Resolve fragmented customer data. All the API requests are over HTTPS 
using POST method with content sent as JSON. This library supports Multi Field Request, 
Person Enrichment & Data Packs, Company Enrichment & Data Packs and Webhooks. Just build a FullContact 
Client with your API Key, make a enrich request and get a response object back.

### Quick Overview
If you are not familiar with the Enrich API, complete details can be found 
@[API documentation](https://platform.fullcontact.com/docs/apis/enrich/introduction) 


FullContact Client provides an object layer to FullContact API communication, 
but understanding Enrich API, webhooks, request and response parameters, 
and common snags is still important.

Once you’re on board with the API behavior, FullContact Client library should simplify 
your integration.

### Supported APIs
- _[Enrich](https://platform.fullcontact.com/docs/apis/enrich/introduction)_
    - `person.enrich`
    - `company.enrich`
    - `company.search`

- Private Identity Cloud
    - _[Resolve](https://platform.fullcontact.com/docs/apis/resolve/introduction)_
        - `identity.map`
        - `identity.resolve`
        - `identity.mapResolve`
        - `identity.delete`
    - [Tags](https://platform.fullcontact.com/docs/apis/resolve/customer-tags)
        - `tags.create`
        - `tags.get`
        - `tags.delete`
    - [Audience](https://platform.fullcontact.com/docs/apis/resolve/customer-tags)
        - `audience.create`
        - `audience.download`
- _[Verification](https://platform.fullcontact.com/docs/apis/verification/introduction)_
    - `v2/verification/email`

- _[Permission](https://platform.fullcontact.com/docs/apis/permission/introduction)_
    - `permission.create`
    - `permission.delete`
    - `permission.find`
    - `permission.current`
    - `permission.verify`

## Providing Authentication to FullContact Client
FullContact client uses ```CredentialsProvider``` interface for Authentication. Different ways 
to provide authentication:

- __Static API Key provider__: 
```go
cp, err := fc.NewStaticCredentialsProvider("your-api-key")
```
- __Through System Environment Variable__:
```go
//API Key is stored as Environment variable FC_API_KEY
cp, err := fc.NewDefaultCredentialsProvider("FC_API_KEY")
```
- If __no__ ```CredentialsProvider``` is specified while making FullContact Client,
it automatically looks for API key from Environment variable `"FC_API_KEY"`

(Don't have an API key? You can pick one up for free [right here.](https://www.fullcontact.com/developer-portal/))

## Making a FullContact Client
Make your fcClient with:

| Parameters | Description | Default value | isOptional |
| ---------------- | ----------- | ------------- | ---------- |
| `WithCredentialsProvider`| Used for Authentication | API Key through Environment variable```"FC_API_KEY"``` | No | 
| `WithHeaders` | Any Custom Headers you want to add with every request, can include `Reporting-Key` as well. | No additional header | Yes |
| `WithTimeout` | Connection timeout in millis for request | 3000ms | Yes |
| `WithRetryHandler` | type RetryHandler  | `DefaultRetryHandler` | Yes |

 
__Please note that you don't have to provide `Authorization` and `Content-Type` in the 
custom Headers map as these will be automatically added.__ 
Custom headers provided will remain same and will be sent with every request made with this client. 
If you wish to change the headers, make a new client with new custom headers.

#### RetryHandler
```go
type RetryHandler interface {
	ShouldRetry(responseCode int) bool
	RetryAttempts() int
	RetryDelayMillis() int
}
```
In case of failure, FullContact Client will auto-retry for same request based on certain conditions set in RetryHandler
- Although optional, a custom Retry handler can be created by implementing `RetryHandler` interface and then be 
used to make FC client. 
By default, client will use `DefaultRetryHandler` to schedule a retry for same request, with `retryAttempts = 1`, 
`retryDelayMillis = 1000`, and in case of `429`(rate limit error) or `503`(capacity limit error).

- This Client will auto-retry for a maximum of 5 times, even if higher value 
is set in the custom Retry Handler.

```go
fcClient, err := fc.NewFullContactClient(
		fc.WithCredentialsProvider(cp),
		fc.WithHeader(map[string]string{"Reporting-Key": "FC_GoClient_1.0.0"}),
		fc.WithTimeout(3000))
```
## Enrich
[Enrich API Reference](https://platform.fullcontact.com/docs/apis/enrich/introduction)
- `person.enrich`
- `company.enrich`
- `company.search`
#### Making a Person Enrich Request
Our V3 Person Enrich supports __Multi Field Request:__ ability to match on __one or many__ input fields

You can build a Person Request using `NewPersonRequest`
and setting different input parameters that you have. If you want to use Webhook, you can specify
it in `webhookUrl` field.
API can lookup and enrich individuals by sending any identifiers you may already have, 
such as: 

- `Emails`: _[]string_
- `Phones`: _[]string_
- `Location`: _*Location_
    - `AddressLine1`: _string_
    - `AddressLine2`: _string_
    - `City`: _string_
    - `Region`: _string_
    - `RegionCode`: _string_
    - `PostalCode`: _string_
- `Name`: _*PersonName_
    - `Full`: _string_
    - `Given`: _string_
    - `Family`: _string_
- `Profiles`: _[]*Profile_
    - `Service`: _string_
    - `Username`: _string_
    - `Userid`: _string_
    - `Url`: _string_
- `DataFilters`: _[]string_
- `Maids`: _[]string_
- `Confidence`: _string_
- `Infer`: _bool_
- `WebhookUrl`: _string_
- `RecordId`: _string_
- `PersonId`: _string_
- `LiNonId`: _string_
- `PartnerId`: _string_
- `Placekey`: _string_
- `VerifiedPhysical`: _bool_
- `MaxMaids`: _int_
- `PanoramaId`:_string_


```go
profile, err := fc.NewProfile(
		fc.WithUsername("bartlorang"),
		fc.WithService("twitter"))
if err != nil {
    log.Fatalln(err)
    return
}
personRequest, err := fc.NewPersonRequest(
    fc.WithEmail("bart@fullcontact.com"),
    fc.WithEmail("bart.lorang@fullcontact.com"),
    fc.WithName(&fc.PersonName{Full:"Bart Lorang",}),
    fc.WithLocation(fc.NewLocation(
        fc.WithAddressLine1("123 Main street"),
        fc.WithAddressLine1("Unit2"),
        fc.WithCity("Denver"),
        fc.WithRegionForLocation("Colorado"))),
    fc.WithProfile(profile),
    fc.WithWebhookUrl(""),
    fc.WithRecordId("customer123"),
    fc.WithPersonId("eYxWc0B-dKRxerTw_uQpxCssM_GyPaLErj0Eu3y2FrU6py1J"))
```
#### Person Enrich Request and Response
You can send a request by calling `PersonEnrich` on fcClient and passing `personRequest` 
as a argument. It sends a Asynchronous request and a `channel` of type `APIResponse` 
is returned as response. You can then `receive` response from this channel.
There is a flag ```isSuccessful``` on `APIResponse` to check 
if the request was successful or not. If the request was unsuccessful, you can check the status code 
and message to determine the cause. 
```go
//Sending Person Enrich request which returns a channel of type `APIResponse`
ch := fcClient.PersonEnrich(personRequest)
resp := <-ch
fmt.Printf("Person Enrich API Response: %v", resp)
if resp.IsSuccessful == true {
    fmt.Printf("Person Response: %v", *resp.PersonResponse)
    fmt.Println(resp.PersonResponse.FullName)
}
```

#### Company Enrich Request and Response
To Enrich Company data FullContact library provides two methods __Lookup by Company Domain__ or
__Search by Company Name__. More data is available through the Lookup by Company Domain, 
but if the domain is unknown, use our Search by Company Name API to find the list of domains 
that could be related to the Company you are looking for and then call the Lookup by 
Company Domain with that domain to get the full information about the company.

##### Lookup by Company Domain
- Request:
```go
companyEnrichRequest, err := fc.NewCompanyRequest(fc.WithDomain("fullcontact.com"))
if err != nil {
    log.Fatalln(err)
    return
}
```
- Response:
```go
resp := <-fcClient.CompanyEnrich(companyEnrichRequest)
fmt.Printf("Company Enrich API Response: %v", resp)
if resp.IsSuccessful {
    fmt.Printf("Company Name: %v", resp.CompanyResponse.Name)
}
```

##### Search by Company Name
- Request:
    - Parameters:
        - `companyName`
        - `webhookUrl` 
        - `location`
        - `locality` 
        - `region`
        - `country`
        - `sort`
```go
companySearchRequest, err := fc.NewCompanyRequest(fc.WithCompanyName("FullContact"))
if err != nil {
    log.Fatalln(err)
    return
}
```
- Response: It returns an array of `CompanySearchResponse`.
```go
resp := <-fcClient.CompanySearch(companySearchRequest)
fmt.Printf("Company Search API Response: %v", resp)
if resp.IsSuccessful {
    fmt.Printf("Company Lookup Domain: %v", resp.CompanySearchResponse[0].LookupDomain)
}
```

## Resolve
[Resolve API Reference](https://platform.fullcontact.com/docs/apis/resolve/introduction)
- `identity.map`
- `identity.resolve`
- `identity.delete`
- `identity.mapResolve`
#### Resolve Request
Resolve uses `ResolveRequest` type for its request which supports
 __Multi Field Request:__ ability to match on __one or many__ input fields

You can build a Resolve Request by using `NewResolveRequest`
and setting different input parameters that you have.

Note: For `identity.map` and `identity.mapResolve` any of `email`, `phone`, `profile`, `name & location` 
must be present.
 
API can lookup and resolve individuals by sending any identifiers you may already have, 
such as: 

- `Emails`: _[]string_
- `Phones`: _[]string_
- `Location`: _*Location_
    - `AddressLine1`: _string_
    - `AddressLine2`: _string_
    - `City`: _string_
    - `Region`: _string_
    - `RegionCode`: _string_
    - `PostalCode`: _string_
- `Name`: _*PersonName_
    - `Full`: _string_
    - `Given`: _string_
    - `Family`: _string_
- `Profiles`: _[]*Profile_
    - `Service`: _string_
    - `Username`: _string_
    - `Userid`: _string_
    - `Url`: _string_
- `Maids`: _[]string_
- `Tags`: _[]*Tag_
- `RecordId`: _string_
- `PersonId`: _string_
- `LiNonId`: _string_
- `PartnerId`: _string_
- `Placekey`: _string_
- `PanoramaId`:_string_
- `GeneratePid`:_bool_

```go
resolveRequest, err := fc.NewResolveRequest(
		fc.WithRecordIdForResolve("r1"),
		fc.WithEmailForResolve("bart@fullcontact.com"))
```

#### Resolve Response
All resolve methods returns a `channel` of type `APIResponse` from which you can get `ResolveResponse`

```go
resp := <-fcClient.IdentityMap(resolveRequest)
fmt.Printf("Identity Map API Response: %v", resp)
if resp.IsSuccessful {
    fmt.Printf("RecordIds Mapped: %v", resp.ResolveResponse.RecordIds)
}
resp = <-fcClient.IdentityResolve(resolveRequest)
fmt.Printf("Identity Resolve API Response: %v", resp)
if resp.IsSuccessful {
    fmt.Printf("PersonIds Mapped: %v", resp.ResolveResponse.PersonIds)
}

resp = <-fcClient.IdentityMapResolve(resolveRequest)
fmt.Printf("Identity Map Resolve API Response: %v", resp)
if resp.IsSuccessful {
    fmt.Printf("PersonIds Mapped: %v", resp.ResolveResponse.PersonIds)
}

resp = <-fcClient.IdentityDelete(resolveRequest)
fmt.Printf("Identity Delete API Response: %v", resp)
if resp.IsSuccessful {
    fmt.Println("Record Deleted Successfully!")
}
```

### Tags/Metadata

[Tags API Reference](https://platform.fullcontact.com/docs/apis/resolve/customer-tags)
- `tags.create`
- `tags.get`
- `tags.delete`

FullContact provides the ability to store customer tags/metadata to each record within a customer's Private Identity 
Cloud for continuous updates, retrievals and deletes across both 1st party as well as 2nd party data partnerships.

#### Creating Tags
Tags can be added while mapping records using `identity.map` API or later using `tags.create` API. 
Once a Customer Record ID has been mapped, customer tags can continue to be added to the originally provided Record ID

##### Tags Request
- Request Parameters:
    - `RecordId`: _string_
    - `Tags`: _[]*Tag_
        - `Key`: _string_
        - `Value`: _string_
       
```go
tagsRequest, err := fc.NewTagsRequest(fc.WithRecordIdForTags("k1"),
		fc.WithTag(fc.NewTag(fc.WithTagKey("gender"), fc.WithTagValue("male"))))
if err != nil {
		log.Fatalln(err)
		return
	}

// Sending Request
resp := <-fcClient.TagsCreate(tagsRequest)
fmt.Printf("\n\nTags Create API Response: %v", resp.TagsResponse)
```

#### Get Tags
This will return all customer tags that are associated to a mapped record using `recordId`.

```java
resp := <-fcClient.TagsGet("recordId")
	fmt.Printf("\n\nTags Get API Response: %v", resp.TagsResponse)
```

#### Delete Tags
This will remove specific or all customer tags that are attached to a mapped record.

```go
tagsRequest, err := fc.NewTagsRequest(fc.WithRecordIdForTags("k1"),
		fc.WithTag(fc.NewTag(fc.WithTagKey("gender"), fc.WithTagValue("male"))))
if err != nil {
		log.Fatalln(err)
		return
	}

// Sending Request
resp := <-fcClient.TagsDelete(tagsRequest)
	fmt.Printf("\n\nTags Delete API Response: %v", resp.Status)
```

### Audience
- `audience.create`
- `audience.download`

This endpoint can be used in order to obtain multiple individuals based upon the key, value 
tag inputs (both are required as input) in order to suppress or take action upon certain audiences 
for data onboarding or audience analysis.

#### Audience Create
The Audience Creation endpoint requires a at least one `Tag` and valid `webhookURL` to be present in order to 
send a message when the audience creation is complete and ready to be downloaded.

```go
audienceRequest, err := fc.NewAudienceRequest(fc.WithWebhookUrlForAudience("your-webhookUrl"),
		fc.WithTagForAudience(fc.NewTag(fc.WithTagKey("gender"), fc.WithTagValue("male"))))
if err != nil {
    log.Fatalln(err)
    return
}

resp := <-fcClient.AudienceCreate(audienceRequest)
fmt.Printf("\n\nAudience Create API Response: %v", resp.AudienceResponse)
if resp.IsSuccessful {
    fmt.Println(resp.AudienceResponse.RequestId)
}
```

#### Audience Download
When `audience.create` result is ready, `requestId` from its response can be used to download the audience data.
A utility method is provided `WriteAudienceBytesToFile(fileName string)` which generates a file in `json.gz` format
with audience data bytes.
```go
requestId := "730000fd-009a-00fc-8008-100e000085f0"  //From the response of 'AudienceCreate'
resp := <-fcClient.AudienceDownload(requestId)
fmt.Printf("\n\nAudience Download API Response: %v", resp.AudienceResponse)
if resp.IsSuccessful {
    resp.AudienceResponse.WriteAudienceBytesToFile(requestId + "_audienceFile.json.gz")
}
```

## Verification
[EmailVerification API Reference](https://platform.fullcontact.com/docs/apis/verification/introduction)
- `v2/verification/email`

FullContact Email Verification API accepts single `email` request, as a `string`. Requests are sent 
using HTTP GET and request email is set as a query parameter.
```go
resp := <-fcClient.EmailVerification("bart@fullcontact.com")
fmt.Printf("\n\nEmail Verification API Response: %v", resp)
if resp.IsSuccessful {
    fmt.Println(resp.EmailVerificationResponse)
}
```

## Permission
[Permission API Reference](https://platform.fullcontact.com/docs/apis/permission/introduction)
- `permission.create`
- `permission.delete`
- `permission.find`
- `permission.current`
- `permission.verify`

#### Permission Request
Permission uses the following type of parameters for it's requests

`PermissionRequest` for
- `permission.create`
- `permission.verify`

and `MultifieldRequest` for
- `permission.delete`
- `permission.find`
- `permission.current`

You can build a Permission Request by using `NewPermissionRequest`
and setting different input parameters that you have.
 
All Permission Api requires `MultifieldRequest` requests, which can be constructed by using `NewMultifieldRequest` and following are it's parameters.

- `Emails`: _[]string_
- `Phones`: _[]string_
- `Location`: _*Location_
    - `AddressLine1`: _string_
    - `AddressLine2`: _string_
    - `City`: _string_
    - `Region`: _string_
    - `RegionCode`: _string_
    - `PostalCode`: _string_
- `Name`: _*PersonName_
    - `Full`: _string_
    - `Given`: _string_
    - `Family`: _string_
- `Profiles`: _[]*Profile_
    - `Service`: _string_
    - `Username`: _string_
    - `Userid`: _string_
    - `Url`: _string_
- `Maids`: _[]string_
- `RecordId`: _string_
- `PersonId`: _string_
- `LiNonId`: _string_
- `PartnerId`: _string_
- `Placekey`: _string_
- `PanoramaId`:_string_

```go
multifieldRequest, err := fc.NewMultifieldRequest(
		fc.WithEmailForMultifieldRequest("bart@fullcontact.com"))
permissionRequest, err := fc.NewPermissionRequest(
		fc.WithMultifieldRequestForPermission(multifieldRequest))
```

### Permission Request
All permission methods returns a `channel` of type `APIResponse` from which you can get corresponding response classes.

The following are the corresponding response classes
- `PermissionCreatedResponse` - permission.create
- `PermissionDeleteResponse` - permission.delete
- `PermissionVerifyResponse` - permission.verify
- `PermissionFindResponse` - permission.find
- `PermissionCurrentResponse` - permission.current

`PermissionCreate` and `PermissionVerify` requires a `ResolveRequest` as parameter while the rest requires `MultifieldRequest` as parameter

### Permission Create

#### Parameters:
Supported fields in query:
- `query`: MultifieldRequest - [required]
- `consentPurposes`: List[ConsentPurposes] - [required]
- `locale`: string
- `ipAddress`: string
- `language`: string
- `collectionMethod`: string - [required]
- `collectionLocation`: string - [required]
- `policyUrl`: string - [required]
- `termsService`: string - [required]
- `tcf`: string
- `timestamp`: int

#### Returns:
class: `PermissionCreateResponse`. A basic API response with response code as 202 if successful.

### Permission Verify
#### Parameters:
Supported fields in query:
- `query`: MultifieldRequest - [required]
- `purposeId`: int - [required]
- `channel`: string - [required]

#### Returns:
class: `PermissionVerifyResponse` with following fields.

- `ttl`: string
- `enabled`: bool
- `channel`: string
- `purposeId`: int
- `purposeName`: string
- `timestamp`: int

### Permission Delete
#### Parameters:
Query takes a `MultiFieldReq`

#### Returns:
class: `PermissionDeleteResponse`. A basic API response with response code as 202 if successful.

### Permission Find
#### Parameters:
Query takes a `MultiFieldReq`

#### Returns:
class: `PermissionFindResponse` with list of Permissions.

### Permission Current
#### Parameters:
Query takes a `MultiFieldReq`

#### Returns:
class: `PermissionCurrentResponse` with set of current permissions

```go
resp := <-fcClient.PermissionCreate(permissionRequest)
fmt.Printf("Permission Create API Response: %v", resp)
if resp.IsSuccessful {
    fmt.Println("Permission Created Successfully!")
}
resp = <-fcClient.PermissionVerify(permissionRequest)
fmt.Printf("Permission Verify API Response: %v", resp)
if resp.IsSuccessful {
    fmt.Printf("Permissions List: %v", resp.PermissionVerifyResponse)
}

resp = <-fcClient.PermissionDelete(multifieldRequest)
fmt.Printf("Permission Delete API Response: %v", resp)
if resp.IsSuccessful {
    fmt.Println("Permission Deleted Successfully!")
}

resp = <-fcClient.PermissionFind(multifieldRequest)
fmt.Printf("Permission Find API Response: %v", resp)
if resp.IsSuccessful {
    fmt.Printf("Permission Find: %v", resp.PermissionFindResponse)
}

resp = <-fcClient.PermissionCurrent(multifieldRequest)
fmt.Printf("Permission Current API Response: %v", resp)
if resp.IsSuccessful {
    fmt.Printf("Permission Current: %v", resp.PermissionCurrentResponse)
}
```
