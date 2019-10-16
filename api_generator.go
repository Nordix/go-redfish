package main

import (
	"fmt"
	"reflect"

	. "github.com/Nordix/go-redfish/client"
)

func main() {

	header := `
package client

import (
  "context"
  "net/http"

  client "github.com/Nordix/go-redfish/client"
)

//go:generate mockery -name=RedfishAPI -output ./mocks
type RedfishAPI interface {
`

	apiServiceType := reflect.TypeOf(&DefaultApiService{})

	fmt.Println(header)
	for i := 0; i < apiServiceType.NumMethod(); i++ {
		method := apiServiceType.Method(i)

		// Get args of method
		args := ""
		for n_in := 1; n_in < method.Type.NumIn(); n_in++ {
			args += method.Type.In(n_in).String()
			args += ",\n"
		}

		// Get return type of method
		rets := ""
		for n_out := 0; n_out < method.Type.NumOut(); n_out++ {
			rets += method.Type.Out(n_out).String()
			rets += ",\n"
		}

		func_sig := fmt.Sprintf("%s(%s) (%s)", method.Name, args, rets)
		fmt.Println(func_sig)
		fmt.Println("")
	}

	fmt.Println("}")

}
