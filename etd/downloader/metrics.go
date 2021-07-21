// Copyright 2015 The go-etherdata Authors
// This file is part of the go-etherdata library.
//
// The go-etherdata library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-etherdata library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-etherdata library. If not, see <http://www.gnu.org/licenses/>.

// Contains the metrics collected by the downloader.

package downloader

import (
	"github.com/crypyto-panel/go-etherdata/metrics"
)

var (
	headerInMeter      = metrics.NewRegisteredMeter("etd/downloader/headers/in", nil)
	headerReqTimer     = metrics.NewRegisteredTimer("etd/downloader/headers/req", nil)
	headerDropMeter    = metrics.NewRegisteredMeter("etd/downloader/headers/drop", nil)
	headerTimeoutMeter = metrics.NewRegisteredMeter("etd/downloader/headers/timeout", nil)

	bodyInMeter      = metrics.NewRegisteredMeter("etd/downloader/bodies/in", nil)
	bodyReqTimer     = metrics.NewRegisteredTimer("etd/downloader/bodies/req", nil)
	bodyDropMeter    = metrics.NewRegisteredMeter("etd/downloader/bodies/drop", nil)
	bodyTimeoutMeter = metrics.NewRegisteredMeter("etd/downloader/bodies/timeout", nil)

	receiptInMeter      = metrics.NewRegisteredMeter("etd/downloader/receipts/in", nil)
	receiptReqTimer     = metrics.NewRegisteredTimer("etd/downloader/receipts/req", nil)
	receiptDropMeter    = metrics.NewRegisteredMeter("etd/downloader/receipts/drop", nil)
	receiptTimeoutMeter = metrics.NewRegisteredMeter("etd/downloader/receipts/timeout", nil)

	stateInMeter   = metrics.NewRegisteredMeter("etd/downloader/states/in", nil)
	stateDropMeter = metrics.NewRegisteredMeter("etd/downloader/states/drop", nil)

	throttleCounter = metrics.NewRegisteredCounter("etd/downloader/throttle", nil)
)
