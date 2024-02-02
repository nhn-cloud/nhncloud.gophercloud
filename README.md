# Extensions for Gophercloud: an NHN Cloud SDK for Go

This is the Go SDK extension for [Gophercloud](https://github.com/gophercloud/gophercloud) to use NHN Cloud services.

The following NHN Cloud services' SDKs are included.

* VPC
* VPC subnet

## How to install

Reference a `nhncloud.gophercloud` package in your code.

```go
import "github.com/nhn-cloud/nhncloud.gophercloud"
```

Then update your `go.mod`.

```shell
go mod tidy
```

---

Or use the `go get` command to import packages.

```shell
go get github.com/nhn-cloud/nhncloud.gophercloud
```
