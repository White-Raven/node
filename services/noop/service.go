/*
 * Copyright (C) 2018 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package noop

import (
	"errors"
	"sync"

	log "github.com/cihub/seelog"
	"github.com/mysteriumnetwork/node/core/ip"
	"github.com/mysteriumnetwork/node/core/location"
	"github.com/mysteriumnetwork/node/identity"
	"github.com/mysteriumnetwork/node/money"
	dto_discovery "github.com/mysteriumnetwork/node/service_discovery/dto"
	"github.com/mysteriumnetwork/node/session"
)

const logPrefix = "[service-noop] "

// ErrAlreadyStarted is the error we return when the start is called multiple times
var ErrAlreadyStarted = errors.New("Service already started")

// NewManager creates new instance of Noop service
func NewManager(locationResolver location.Resolver, ipResolver ip.Resolver) *Manager {
	return &Manager{
		locationResolver: locationResolver,
		ipResolver:       ipResolver,
	}
}

// Manager represents entrypoint for Noop service
type Manager struct {
	process          sync.WaitGroup
	locationResolver location.Resolver
	ipResolver       ip.Resolver
	isStarted        bool
}

// Start starts service - does not block
func (manager *Manager) Start(providerID identity.Identity) (dto_discovery.ServiceProposal, session.ConfigProvider, error) {
	sessionConfigProvider := func() (session.ServiceConfiguration, error) {
		return nil, nil
	}

	if manager.isStarted {
		return dto_discovery.ServiceProposal{}, sessionConfigProvider, ErrAlreadyStarted
	}

	manager.process.Add(1)
	manager.isStarted = true
	log.Info(logPrefix, "Noop service started successfully")

	publicIP, err := manager.ipResolver.GetPublicIP()
	if err != nil {
		return dto_discovery.ServiceProposal{}, sessionConfigProvider, err
	}

	country, err := manager.locationResolver.ResolveCountry(publicIP)
	if err != nil {
		return dto_discovery.ServiceProposal{}, sessionConfigProvider, err
	}

	proposal := dto_discovery.ServiceProposal{
		ServiceType: ServiceType,
		ServiceDefinition: ServiceDefinition{
			Location: dto_discovery.Location{Country: country},
		},
		PaymentMethodType: PaymentMethodNoop,
		PaymentMethod: PaymentNoop{
			Price: money.NewMoney(0, money.CURRENCY_MYST),
		},
	}

	return proposal, sessionConfigProvider, nil
}

// Wait blocks until service is stopped
func (manager *Manager) Wait() error {
	if !manager.isStarted {
		return nil
	}
	manager.process.Wait()
	return nil
}

// Stop stops service
func (manager *Manager) Stop() error {
	if !manager.isStarted {
		return nil
	}

	manager.process.Done()
	manager.isStarted = false
	log.Info(logPrefix, "Noop service stopped")
	return nil
}