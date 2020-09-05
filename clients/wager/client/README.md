# Go API client for wager

APIs for a wager system

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./wager"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*WagerApi* | [**BuyWagersPost**](docs/WagerApi.md#buywagerspost) | **Post** /buy/{wager_id} | 
*WagerApi* | [**CreateWagerPost**](docs/WagerApi.md#createwagerpost) | **Post** /wagers | 
*WagerApi* | [**ListWagersGet**](docs/WagerApi.md#listwagersget) | **Get** /wagers | 


## Documentation For Models

 - [BuyWagerRequest](docs/BuyWagerRequest.md)
 - [CreateWagerRequest](docs/CreateWagerRequest.md)
 - [Error](docs/Error.md)
 - [Wager](docs/Wager.md)
 - [WagerPurchase](docs/WagerPurchase.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author


