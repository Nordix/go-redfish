package client

import (
	"context"
	"net/http"

	client "github.com/Nordix/go-redfish/client"
)

//go:generate mockery -name=RedfishAPI -output ./mocks
type RedfishAPI interface {
	EjectVirtualMedia(context.Context,
		string,
		string,
		map[string]interface{},
	) (client.RedfishError,
		*http.Response,
		error,
	)

	FirmwareInventory(context.Context,
	) (client.Collection,
		*http.Response,
		error,
	)

	FirmwareInventoryDownloadImage(context.Context,
		*client.FirmwareInventoryDownloadImageOpts,
	) (client.RedfishError,
		*http.Response,
		error,
	)

	GetManager(context.Context,
		string,
	) (client.Manager,
		*http.Response,
		error,
	)

	GetManagerVirtualMedia(context.Context,
		string,
		string,
	) (client.VirtualMedia,
		*http.Response,
		error,
	)

	GetRoot(context.Context,
	) (client.Root,
		*http.Response,
		error,
	)

	GetSystem(context.Context,
		string,
	) (client.ComputerSystem,
		*http.Response,
		error,
	)

	GetTask(context.Context,
		string,
	) (client.Task,
		*http.Response,
		error,
	)

	InsertVirtualMedia(context.Context,
		string,
		string,
		client.InsertMediaRequestBody,
	) (client.RedfishError,
		*http.Response,
		error,
	)

	ListManagerVirtualMedia(context.Context,
		string,
	) (client.Collection,
		*http.Response,
		error,
	)

	ListManagers(context.Context,
	) (client.Collection,
		*http.Response,
		error,
	)

	ListSystems(context.Context,
	) (client.Collection,
		*http.Response,
		error,
	)

	ResetSystem(context.Context,
		string,
		client.ResetRequestBody,
	) (client.RedfishError,
		*http.Response,
		error,
	)

	SetSystem(context.Context,
		string,
		client.ComputerSystem,
	) (client.ComputerSystem,
		*http.Response,
		error,
	)

	UpdateService(context.Context,
	) (client.UpdateService,
		*http.Response,
		error,
	)

	UpdateServiceSimpleUpdate(context.Context,
		client.SimpleUpdateRequestBody,
	) (client.RedfishError,
		*http.Response,
		error,
	)
}
