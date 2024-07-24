# \SlurmAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlurmV0038Diag**](SlurmAPI.md#SlurmV0038Diag) | **Get** /slurm/v0.0.38/diag | get diagnostics
[**SlurmV0038GetJob**](SlurmAPI.md#SlurmV0038GetJob) | **Get** /slurm/v0.0.38/job/{job_id} | get job info
[**SlurmV0038GetJobs**](SlurmAPI.md#SlurmV0038GetJobs) | **Get** /slurm/v0.0.38/jobs | get list of jobs
[**SlurmV0038GetNode**](SlurmAPI.md#SlurmV0038GetNode) | **Get** /slurm/v0.0.38/node/{node_name} | get node info
[**SlurmV0038GetNodes**](SlurmAPI.md#SlurmV0038GetNodes) | **Get** /slurm/v0.0.38/nodes | get all node info
[**SlurmV0038GetPartition**](SlurmAPI.md#SlurmV0038GetPartition) | **Get** /slurm/v0.0.38/partition/{partition_name} | get partition info
[**SlurmV0038GetPartitions**](SlurmAPI.md#SlurmV0038GetPartitions) | **Get** /slurm/v0.0.38/partitions | get all partition info
[**SlurmV0038GetReservation**](SlurmAPI.md#SlurmV0038GetReservation) | **Get** /slurm/v0.0.38/reservation/{reservation_name} | get reservation info
[**SlurmV0038GetReservations**](SlurmAPI.md#SlurmV0038GetReservations) | **Get** /slurm/v0.0.38/reservations | get all reservation info
[**SlurmV0038Ping**](SlurmAPI.md#SlurmV0038Ping) | **Get** /slurm/v0.0.38/ping | ping test
[**SlurmV0038SlurmctldGetLicenses**](SlurmAPI.md#SlurmV0038SlurmctldGetLicenses) | **Get** /slurm/v0.0.38/licenses | get all Slurm tracked license info



## SlurmV0038Diag

> V0038Diag SlurmV0038Diag(ctx).Execute()

get diagnostics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pcolladosoto/goslurm/oapigenro"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0038Diag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038Diag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0038Diag`: V0038Diag
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038Diag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038DiagRequest struct via the builder pattern


### Return type

[**V0038Diag**](V0038Diag.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038GetJob

> V0038JobsResponse SlurmV0038GetJob(ctx, jobId).Execute()

get job info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pcolladosoto/goslurm/oapigenro"
)

func main() {
	jobId := "jobId_example" // string | Slurm JobID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0038GetJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0038GetJob`: V0038JobsResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Slurm JobID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0038JobsResponse**](V0038JobsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038GetJobs

> V0038JobsResponse SlurmV0038GetJobs(ctx).UpdateTime(updateTime).Execute()

get list of jobs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pcolladosoto/goslurm/oapigenro"
)

func main() {
	updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0038GetJobs(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0038GetJobs`: V0038JobsResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0038JobsResponse**](V0038JobsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038GetNode

> V0038NodesResponse SlurmV0038GetNode(ctx, nodeName).Execute()

get node info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pcolladosoto/goslurm/oapigenro"
)

func main() {
	nodeName := "nodeName_example" // string | Slurm Node Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0038GetNode(context.Background(), nodeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038GetNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0038GetNode`: V0038NodesResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Slurm Node Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038GetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0038NodesResponse**](V0038NodesResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038GetNodes

> V0038NodesResponse SlurmV0038GetNodes(ctx).UpdateTime(updateTime).Execute()

get all node info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pcolladosoto/goslurm/oapigenro"
)

func main() {
	updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0038GetNodes(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038GetNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0038GetNodes`: V0038NodesResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038GetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038GetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0038NodesResponse**](V0038NodesResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038GetPartition

> V0038PartitionsResponse SlurmV0038GetPartition(ctx, partitionName).UpdateTime(updateTime).Execute()

get partition info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pcolladosoto/goslurm/oapigenro"
)

func main() {
	partitionName := "partitionName_example" // string | Slurm Partition Name
	updateTime := int64(789) // int64 | Filter if there were no partition changes (not limited to partition in URL endpoint) since update_time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0038GetPartition(context.Background(), partitionName).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038GetPartition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0038GetPartition`: V0038PartitionsResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038GetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionName** | **string** | Slurm Partition Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038GetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **int64** | Filter if there were no partition changes (not limited to partition in URL endpoint) since update_time. | 

### Return type

[**V0038PartitionsResponse**](V0038PartitionsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038GetPartitions

> V0038PartitionsResponse SlurmV0038GetPartitions(ctx).UpdateTime(updateTime).Execute()

get all partition info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pcolladosoto/goslurm/oapigenro"
)

func main() {
	updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0038GetPartitions(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038GetPartitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0038GetPartitions`: V0038PartitionsResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038GetPartitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038GetPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0038PartitionsResponse**](V0038PartitionsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038GetReservation

> V0038ReservationsResponse SlurmV0038GetReservation(ctx, reservationName).UpdateTime(updateTime).Execute()

get reservation info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pcolladosoto/goslurm/oapigenro"
)

func main() {
	reservationName := "reservationName_example" // string | Slurm Reservation Name
	updateTime := int64(789) // int64 | Filter if no reservation (not limited to reservation in URL) changed since update_time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0038GetReservation(context.Background(), reservationName).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038GetReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0038GetReservation`: V0038ReservationsResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038GetReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Slurm Reservation Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038GetReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **int64** | Filter if no reservation (not limited to reservation in URL) changed since update_time. | 

### Return type

[**V0038ReservationsResponse**](V0038ReservationsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038GetReservations

> V0038ReservationsResponse SlurmV0038GetReservations(ctx).UpdateTime(updateTime).Execute()

get all reservation info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pcolladosoto/goslurm/oapigenro"
)

func main() {
	updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0038GetReservations(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038GetReservations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0038GetReservations`: V0038ReservationsResponse
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038GetReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038GetReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0038ReservationsResponse**](V0038ReservationsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038Ping

> V0038Pings SlurmV0038Ping(ctx).Execute()

ping test

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pcolladosoto/goslurm/oapigenro"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0038Ping(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038Ping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0038Ping`: V0038Pings
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038Ping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038PingRequest struct via the builder pattern


### Return type

[**V0038Pings**](V0038Pings.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0038SlurmctldGetLicenses

> V0038Licenses SlurmV0038SlurmctldGetLicenses(ctx).Execute()

get all Slurm tracked license info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pcolladosoto/goslurm/oapigenro"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0038SlurmctldGetLicenses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0038SlurmctldGetLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0038SlurmctldGetLicenses`: V0038Licenses
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0038SlurmctldGetLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0038SlurmctldGetLicensesRequest struct via the builder pattern


### Return type

[**V0038Licenses**](V0038Licenses.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

