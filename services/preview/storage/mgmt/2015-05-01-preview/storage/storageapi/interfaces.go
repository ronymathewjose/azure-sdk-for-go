package storageapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/storage/mgmt/2015-05-01-preview/storage"
	"github.com/Azure/go-autorest/autorest"
)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	CheckNameAvailability(ctx context.Context, accountName storage.AccountCheckNameAvailabilityParameters) (result storage.CheckNameAvailabilityResult, err error)
	Create(ctx context.Context, resourceGroupName string, accountName string, parameters storage.AccountCreateParameters) (result storage.AccountsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	GetProperties(ctx context.Context, resourceGroupName string, accountName string) (result storage.Account, err error)
	List(ctx context.Context) (result storage.AccountListResultPage, err error)
	ListComplete(ctx context.Context) (result storage.AccountListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result storage.AccountListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result storage.AccountListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result storage.AccountKeys, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, accountName string, regenerateKey storage.AccountRegenerateKeyParameters) (result storage.AccountKeys, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, parameters storage.AccountUpdateParameters) (result storage.Account, err error)
}

var _ AccountsClientAPI = (*storage.AccountsClient)(nil)

// UsageClientAPI contains the set of methods on the UsageClient type.
type UsageClientAPI interface {
	List(ctx context.Context) (result storage.UsageListResult, err error)
}

var _ UsageClientAPI = (*storage.UsageClient)(nil)
