## FullContact Go Client
The official [FullContact](https://www.fullcontact.com/) Golang Client Library for the FullContact V3 APIs.

## Installation

Make sure your project is using Go Modules (it will have a `go.mod` file in its
root if it already is):

``` sh
go mod init
```

Then, reference fullcontact-go in a Go program with `import`:

``` go
import fc "github.com/fullcontact/fullcontact-go/fc"
```

## Working with FullContact Client
FullContact client supports v3 Enrich APIs which are super simplified to easily 
enrich Person and Company data. All the API requests are over HTTPS using POST method 
with content sent as JSON. This library supports Multi Field Request, Person Enrichment 
& Data Packs, Company Enrichment & Data Packs and Webhooks. Just build a FullContact 
Client with your API Key, make a enrich request and get a response object back.

### Quick Overview
If you are not familiar with the Enrich API, complete details can be found 
@[API documentation](https://www.fullcontact.com/developer/docs/) 


FullContact Client provides an object layer to FullContact API communication, 
but understanding Enrich API, webhooks, request and response parameters, 
and common snags is still important.

Once you’re on board with the API behavior, FullContact Client library should simplify 
your integration.

### Supported APIs
- _[Enrich](https://dashboard.fullcontact.com/api-ref#enrich)_
    - `person.enrich`
    - `company.enrich`
    - `company.search`
- _[Resolve](https://dashboard.fullcontact.com/api-ref#resolve-2)_
    - `identity.map`
    - `identity.resolve`
    - `identity.delete`

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
[Enrich API Reference](https://dashboard.fullcontact.com/api-ref#enrich)
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
- `EebhookUrl`: _string_
- `RecordId`: _string_
- `PersonId`: _string_


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
[Resolve API Reference](https://dashboard.fullcontact.com/api-ref#resolve-2)
- `identity.map`
- `identity.resolve`
- `identity.delete`
#### Resolve Request
Resolve uses `ResolveRequest` type for its request which supports
 __Multi Field Request:__ ability to match on __one or many__ input fields

You can build a Resolve Request by using `NewResolveRequest`
and setting different input parameters that you have.

Note: For `identity.map` any of `email`, `phone`, `profile`, `name & location` 
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
- `RecordId`: _string_
- `PersonId`: _string_

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

resp = <-fcClient.IdentityDelete(resolveRequest)
fmt.Printf("Identity Delete API Response: %v", resp)
if resp.IsSuccessful {
    fmt.Println("Record Deleted Successfully!")
}
```