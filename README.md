# Go-Redfish - Golang client library for DMTF Redfish

## Introduction

go-redfish is the golang client library to interact with Redfish Server. This library do not implement complete redfish standard but it is meant to provide parallel functionality as of [sushy](https://opendev.org/openstack/sushy). `go-redfish` can be tested against [sushy-emulator](https://opendev.org/openstack/sushy-tools).

## Usage ##

To use `go-redfish` library in go code, you just need to import the library as below.

```go
import (
    "fmt"

    redfish "github.com/nordix-airship/go-redfish/client"
)
```

Then, get the client api object for the desired redfish server.

```go
cfg := &redfish.Configuration{
    BasePath:      "http://localhost:8000",
    DefaultHeader: make(map[string]string),
    UserAgent:     "go-redfish/client",
  }

redfishApi := rf_client.NewAPIClient(cfg).DefaultApi
```

Use `redfishApi` to interact with redfish server. This object contains get, set and list functions for different redfish resources. There are loose validations on the client side for the required paramaters of a certain API. In case of any error e.g. missing a required parameter; then server will generate the error. To see the available actions for API endpoints see [Endpoint Reference](https://github.com/nordix-airship/go-redfish/tree/master/client#documentation-for-api-endpoints) and [Model Definitions](https://github.com/nordix-airship/go-redfish/tree/master/client#documentation-for-models). See following examples for different operations.

## Examples ##

- *List available ComputerSystems*

```go
sl, _, _ := redfishApi.ListSystems(context.Background())
```
- *List available Managers*

```go
sl, _, _ := redfishApi.ListManagers(context.Background())
```
- *Get Desired ComputerSystem Resource*

```go
system_id := "dd9fd064-263b-469c-91d4-d45f341fe2c5"
system, _, _ := redfishApi.GetSystem(context.Background(), system_id)
```
- *Reset ComputerSystem*

```go
system_id := "dd9fd064-263b-469c-91d4-d45f341fe2c5"
systemReq := rf_client.ComputerSystem{ResetType: "ForceRestart"}
reset_resp, _, _ = redfishApi.ResetSystem(context.Background(), system_id, reset_type)
```
- *Insert VirtualMedia for manager*

```go
manager_id := "58893887-8974-2487-2389-841168418919"
insertReq := rf_client.InsertMediaRequestBody{}
insertReq.Image = "http://releases.ubuntu.com/19.04/ubuntu-19.04-live-server-amd64.iso"
insertReq.Inserted = true
redfishApi.InsertVirtualMedia(context.Background(), manager_id, "Cd", insertReq)
```
- *Set Boot Device For Next Boot*

```go
system_id := "dd9fd064-263b-469c-91d4-d45f341fe2c5"
systemReq := rf_client.ComputerSystem{}
systemReq.Boot.BootSourceOverrideTarget = "Cd"
redfishApi.SetSystem(context.Background(), system_id, systemReq)
```

- *Get available system's boot options*

```go
system_id := "dd9fd064-263b-469c-91d4-d45f341fe2c5"
system, _, _ := redfishApi.GetSystem(context.Background(), system_id)
fmt.Printf("%#v", system.Boot.BootSourceOverrideTargetRedfishAllowableValues)
```
Sample Output:
```go
[Pxe Cd Hdd]
```

## Contributing ##

This library is created using openapi specification and code is generated using openapi go client generator. So it is very easy to extend the client using openapi spec. If you need any extra api resource or see any issues then create a pull request or raise an issue.