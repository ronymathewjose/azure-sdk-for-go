// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package resourcehealth

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/resourcehealth/mgmt/2017-07-01/resourcehealth"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AvailabilityStateValues = original.AvailabilityStateValues

const (
	Available   AvailabilityStateValues = original.Available
	Unavailable AvailabilityStateValues = original.Unavailable
	Unknown     AvailabilityStateValues = original.Unknown
)

type ReasonChronicityTypes = original.ReasonChronicityTypes

const (
	Persistent ReasonChronicityTypes = original.Persistent
	Transient  ReasonChronicityTypes = original.Transient
)

type SeverityValues = original.SeverityValues

const (
	Error       SeverityValues = original.Error
	Information SeverityValues = original.Information
	Warning     SeverityValues = original.Warning
)

type StageValues = original.StageValues

const (
	Active   StageValues = original.Active
	Archived StageValues = original.Archived
	Resolve  StageValues = original.Resolve
)

type AvailabilityStatus = original.AvailabilityStatus
type AvailabilityStatusListResult = original.AvailabilityStatusListResult
type AvailabilityStatusListResultIterator = original.AvailabilityStatusListResultIterator
type AvailabilityStatusListResultPage = original.AvailabilityStatusListResultPage
type AvailabilityStatusProperties = original.AvailabilityStatusProperties
type AvailabilityStatusPropertiesRecentlyResolvedState = original.AvailabilityStatusPropertiesRecentlyResolvedState
type AvailabilityStatusesClient = original.AvailabilityStatusesClient
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type ChildAvailabilityStatusesClient = original.ChildAvailabilityStatusesClient
type ChildResourcesClient = original.ChildResourcesClient
type EmergingIssue = original.EmergingIssue
type EmergingIssueImpact = original.EmergingIssueImpact
type EmergingIssueListResult = original.EmergingIssueListResult
type EmergingIssueListResultIterator = original.EmergingIssueListResultIterator
type EmergingIssueListResultPage = original.EmergingIssueListResultPage
type EmergingIssuesClient = original.EmergingIssuesClient
type EmergingIssuesGetResult = original.EmergingIssuesGetResult
type ErrorResponse = original.ErrorResponse
type ImpactedRegion = original.ImpactedRegion
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type RecommendedAction = original.RecommendedAction
type Resource = original.Resource
type ServiceImpactingEvent = original.ServiceImpactingEvent
type ServiceImpactingEventIncidentProperties = original.ServiceImpactingEventIncidentProperties
type ServiceImpactingEventStatus = original.ServiceImpactingEventStatus
type StatusActiveEvent = original.StatusActiveEvent
type StatusBanner = original.StatusBanner
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAvailabilityStatusListResultIterator(page AvailabilityStatusListResultPage) AvailabilityStatusListResultIterator {
	return original.NewAvailabilityStatusListResultIterator(page)
}
func NewAvailabilityStatusListResultPage(getNextPage func(context.Context, AvailabilityStatusListResult) (AvailabilityStatusListResult, error)) AvailabilityStatusListResultPage {
	return original.NewAvailabilityStatusListResultPage(getNextPage)
}
func NewAvailabilityStatusesClient(subscriptionID string) AvailabilityStatusesClient {
	return original.NewAvailabilityStatusesClient(subscriptionID)
}
func NewAvailabilityStatusesClientWithBaseURI(baseURI string, subscriptionID string) AvailabilityStatusesClient {
	return original.NewAvailabilityStatusesClientWithBaseURI(baseURI, subscriptionID)
}
func NewChildAvailabilityStatusesClient(subscriptionID string) ChildAvailabilityStatusesClient {
	return original.NewChildAvailabilityStatusesClient(subscriptionID)
}
func NewChildAvailabilityStatusesClientWithBaseURI(baseURI string, subscriptionID string) ChildAvailabilityStatusesClient {
	return original.NewChildAvailabilityStatusesClientWithBaseURI(baseURI, subscriptionID)
}
func NewChildResourcesClient(subscriptionID string) ChildResourcesClient {
	return original.NewChildResourcesClient(subscriptionID)
}
func NewChildResourcesClientWithBaseURI(baseURI string, subscriptionID string) ChildResourcesClient {
	return original.NewChildResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewEmergingIssueListResultIterator(page EmergingIssueListResultPage) EmergingIssueListResultIterator {
	return original.NewEmergingIssueListResultIterator(page)
}
func NewEmergingIssueListResultPage(getNextPage func(context.Context, EmergingIssueListResult) (EmergingIssueListResult, error)) EmergingIssueListResultPage {
	return original.NewEmergingIssueListResultPage(getNextPage)
}
func NewEmergingIssuesClient(subscriptionID string) EmergingIssuesClient {
	return original.NewEmergingIssuesClient(subscriptionID)
}
func NewEmergingIssuesClientWithBaseURI(baseURI string, subscriptionID string) EmergingIssuesClient {
	return original.NewEmergingIssuesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAvailabilityStateValuesValues() []AvailabilityStateValues {
	return original.PossibleAvailabilityStateValuesValues()
}
func PossibleReasonChronicityTypesValues() []ReasonChronicityTypes {
	return original.PossibleReasonChronicityTypesValues()
}
func PossibleSeverityValuesValues() []SeverityValues {
	return original.PossibleSeverityValuesValues()
}
func PossibleStageValuesValues() []StageValues {
	return original.PossibleStageValuesValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
