# \CommonOperationsApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewEmailAddressUsingPOST**](CommonOperationsApi.md#CreateNewEmailAddressUsingPOST) | **Post** /newEmailAddress | Create new email address
[**FetchLatestEmailUsingGET**](CommonOperationsApi.md#FetchLatestEmailUsingGET) | **Get** /fetchLatestEmail | Fetch inbox&#39;s latest email or if empty wait for email to arrive
[**SendEmailSimpleUsingPOST**](CommonOperationsApi.md#SendEmailSimpleUsingPOST) | **Post** /sendEmail | Send an email from a random email address


# **CreateNewEmailAddressUsingPOST**
> Inbox CreateNewEmailAddressUsingPOST(ctx, )
Create new email address

Returns an Inbox with an `id` and an `emailAddress`

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Inbox**](Inbox.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchLatestEmailUsingGET**
> Email FetchLatestEmailUsingGET(ctx, optional)
Fetch inbox's latest email or if empty wait for email to arrive

Will return either the last received email or wait for an email to arrive and return that. If you need to wait for an email for a non-empty inbox see the other receive methods.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FetchLatestEmailUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FetchLatestEmailUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxEmailAddress** | **optional.String**| Email address of the inbox we are fetching emails from | 
 **inboxId** | [**optional.Interface of string**](.md)| Id of the inbox we are fetching emails from | 

### Return type

[**Email**](Email.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendEmailSimpleUsingPOST**
> SendEmailSimpleUsingPOST(ctx, sendEmailOptions)
Send an email from a random email address

To specify an email address first create an inbox and use that with the other send email methods

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sendEmailOptions** | [**SendEmailOptions**](SendEmailOptions.md)| sendEmailOptions | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
