// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armdataboxedge.ClientFactory type.
type ServerFactory struct {
	AddonsServer                    AddonsServer
	AlertsServer                    AlertsServer
	AvailableSKUsServer             AvailableSKUsServer
	BandwidthSchedulesServer        BandwidthSchedulesServer
	ContainersServer                ContainersServer
	DevicesServer                   DevicesServer
	DiagnosticSettingsServer        DiagnosticSettingsServer
	JobsServer                      JobsServer
	MonitoringConfigServer          MonitoringConfigServer
	NodesServer                     NodesServer
	OperationsServer                OperationsServer
	OperationsStatusServer          OperationsStatusServer
	OrdersServer                    OrdersServer
	RolesServer                     RolesServer
	SharesServer                    SharesServer
	StorageAccountCredentialsServer StorageAccountCredentialsServer
	StorageAccountsServer           StorageAccountsServer
	SupportPackagesServer           SupportPackagesServer
	TriggersServer                  TriggersServer
	UsersServer                     UsersServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armdataboxedge.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armdataboxedge.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                               *ServerFactory
	trMu                              sync.Mutex
	trAddonsServer                    *AddonsServerTransport
	trAlertsServer                    *AlertsServerTransport
	trAvailableSKUsServer             *AvailableSKUsServerTransport
	trBandwidthSchedulesServer        *BandwidthSchedulesServerTransport
	trContainersServer                *ContainersServerTransport
	trDevicesServer                   *DevicesServerTransport
	trDiagnosticSettingsServer        *DiagnosticSettingsServerTransport
	trJobsServer                      *JobsServerTransport
	trMonitoringConfigServer          *MonitoringConfigServerTransport
	trNodesServer                     *NodesServerTransport
	trOperationsServer                *OperationsServerTransport
	trOperationsStatusServer          *OperationsStatusServerTransport
	trOrdersServer                    *OrdersServerTransport
	trRolesServer                     *RolesServerTransport
	trSharesServer                    *SharesServerTransport
	trStorageAccountCredentialsServer *StorageAccountCredentialsServerTransport
	trStorageAccountsServer           *StorageAccountsServerTransport
	trSupportPackagesServer           *SupportPackagesServerTransport
	trTriggersServer                  *TriggersServerTransport
	trUsersServer                     *UsersServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "AddonsClient":
		initServer(s, &s.trAddonsServer, func() *AddonsServerTransport { return NewAddonsServerTransport(&s.srv.AddonsServer) })
		resp, err = s.trAddonsServer.Do(req)
	case "AlertsClient":
		initServer(s, &s.trAlertsServer, func() *AlertsServerTransport { return NewAlertsServerTransport(&s.srv.AlertsServer) })
		resp, err = s.trAlertsServer.Do(req)
	case "AvailableSKUsClient":
		initServer(s, &s.trAvailableSKUsServer, func() *AvailableSKUsServerTransport {
			return NewAvailableSKUsServerTransport(&s.srv.AvailableSKUsServer)
		})
		resp, err = s.trAvailableSKUsServer.Do(req)
	case "BandwidthSchedulesClient":
		initServer(s, &s.trBandwidthSchedulesServer, func() *BandwidthSchedulesServerTransport {
			return NewBandwidthSchedulesServerTransport(&s.srv.BandwidthSchedulesServer)
		})
		resp, err = s.trBandwidthSchedulesServer.Do(req)
	case "ContainersClient":
		initServer(s, &s.trContainersServer, func() *ContainersServerTransport { return NewContainersServerTransport(&s.srv.ContainersServer) })
		resp, err = s.trContainersServer.Do(req)
	case "DevicesClient":
		initServer(s, &s.trDevicesServer, func() *DevicesServerTransport { return NewDevicesServerTransport(&s.srv.DevicesServer) })
		resp, err = s.trDevicesServer.Do(req)
	case "DiagnosticSettingsClient":
		initServer(s, &s.trDiagnosticSettingsServer, func() *DiagnosticSettingsServerTransport {
			return NewDiagnosticSettingsServerTransport(&s.srv.DiagnosticSettingsServer)
		})
		resp, err = s.trDiagnosticSettingsServer.Do(req)
	case "JobsClient":
		initServer(s, &s.trJobsServer, func() *JobsServerTransport { return NewJobsServerTransport(&s.srv.JobsServer) })
		resp, err = s.trJobsServer.Do(req)
	case "MonitoringConfigClient":
		initServer(s, &s.trMonitoringConfigServer, func() *MonitoringConfigServerTransport {
			return NewMonitoringConfigServerTransport(&s.srv.MonitoringConfigServer)
		})
		resp, err = s.trMonitoringConfigServer.Do(req)
	case "NodesClient":
		initServer(s, &s.trNodesServer, func() *NodesServerTransport { return NewNodesServerTransport(&s.srv.NodesServer) })
		resp, err = s.trNodesServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "OperationsStatusClient":
		initServer(s, &s.trOperationsStatusServer, func() *OperationsStatusServerTransport {
			return NewOperationsStatusServerTransport(&s.srv.OperationsStatusServer)
		})
		resp, err = s.trOperationsStatusServer.Do(req)
	case "OrdersClient":
		initServer(s, &s.trOrdersServer, func() *OrdersServerTransport { return NewOrdersServerTransport(&s.srv.OrdersServer) })
		resp, err = s.trOrdersServer.Do(req)
	case "RolesClient":
		initServer(s, &s.trRolesServer, func() *RolesServerTransport { return NewRolesServerTransport(&s.srv.RolesServer) })
		resp, err = s.trRolesServer.Do(req)
	case "SharesClient":
		initServer(s, &s.trSharesServer, func() *SharesServerTransport { return NewSharesServerTransport(&s.srv.SharesServer) })
		resp, err = s.trSharesServer.Do(req)
	case "StorageAccountCredentialsClient":
		initServer(s, &s.trStorageAccountCredentialsServer, func() *StorageAccountCredentialsServerTransport {
			return NewStorageAccountCredentialsServerTransport(&s.srv.StorageAccountCredentialsServer)
		})
		resp, err = s.trStorageAccountCredentialsServer.Do(req)
	case "StorageAccountsClient":
		initServer(s, &s.trStorageAccountsServer, func() *StorageAccountsServerTransport {
			return NewStorageAccountsServerTransport(&s.srv.StorageAccountsServer)
		})
		resp, err = s.trStorageAccountsServer.Do(req)
	case "SupportPackagesClient":
		initServer(s, &s.trSupportPackagesServer, func() *SupportPackagesServerTransport {
			return NewSupportPackagesServerTransport(&s.srv.SupportPackagesServer)
		})
		resp, err = s.trSupportPackagesServer.Do(req)
	case "TriggersClient":
		initServer(s, &s.trTriggersServer, func() *TriggersServerTransport { return NewTriggersServerTransport(&s.srv.TriggersServer) })
		resp, err = s.trTriggersServer.Do(req)
	case "UsersClient":
		initServer(s, &s.trUsersServer, func() *UsersServerTransport { return NewUsersServerTransport(&s.srv.UsersServer) })
		resp, err = s.trUsersServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
