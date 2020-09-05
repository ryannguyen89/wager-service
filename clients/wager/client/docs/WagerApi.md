# \WagerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuyWagersPost**](WagerApi.md#BuyWagersPost) | **Post** /buy/{wager_id} | 
[**CreateWagerPost**](WagerApi.md#CreateWagerPost) | **Post** /wagers | 
[**ListWagersGet**](WagerApi.md#ListWagersGet) | **Get** /wagers | 



## BuyWagersPost

> []WagerPurchase BuyWagersPost(ctx, body, wagerId)



Buy wager

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**BuyWagerRequest**](BuyWagerRequest.md)| body for sending message | 
**wagerId** | **int32**|  | 

### Return type

[**[]WagerPurchase**](WagerPurchase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWagerPost

> Wager CreateWagerPost(ctx, body)



Create a wager

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**CreateWagerRequest**](CreateWagerRequest.md)| body for sending message | 

### Return type

[**Wager**](Wager.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWagersGet

> []Wager ListWagersGet(ctx, page, limit)



List wagers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**page** | **int32**|  | 
**limit** | **string**|  | 

### Return type

[**[]Wager**](Wager.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

