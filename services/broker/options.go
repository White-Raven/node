/*
 * Copyright (C) 2020 The "MysteriumNetwork/node" Authors.
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

package broker

import (
	"encoding/json"

	"github.com/mysteriumnetwork/node/core/service"
)

// GetOptions returns effective broker service options from application configuration.
func GetOptions() service.Options {
	return nil
}

// ParseJSONOptions function fills in Broker options from JSON request
func ParseJSONOptions(_ *json.RawMessage) (service.Options, error) {
	return nil, nil
}