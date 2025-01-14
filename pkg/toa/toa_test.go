// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

package toa

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	assertions "github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func buildLoRaDownlinkFromParameters(payloadSize int, frequency uint64, dataRate *ttnpb.DataRate, codingRate string) (downlink *ttnpb.DownlinkMessage, err error) {
	payload := bytes.Repeat([]byte{0x0}, payloadSize)
	scheduled := &ttnpb.TxSettings{
		Frequency:  frequency,
		DataRate:   dataRate,
		CodingRate: codingRate,
	}
	downlink = &ttnpb.DownlinkMessage{
		RawPayload: payload,
		Settings: &ttnpb.DownlinkMessage_Scheduled{
			Scheduled: scheduled,
		},
	}
	return downlink, nil
}

func TestInvalidLoRa(t *testing.T) {
	a := assertions.New(t)

	// Invalid coding rate.
	{
		downlink, err := buildLoRaDownlinkFromParameters(10, 868100000, &ttnpb.DataRate{
			Modulation: &ttnpb.DataRate_Lora{
				Lora: &ttnpb.LoRaDataRate{
					SpreadingFactor: 10,
					Bandwidth:       125000,
				},
			},
		}, "1/9")
		a.So(err, should.BeNil)
		_, err = Compute(len(downlink.RawPayload), downlink.GetScheduled())
		a.So(err, should.NotBeNil)
	}

	// Invalid spreading factor.
	{
		downlink, err := buildLoRaDownlinkFromParameters(10, 868100000, &ttnpb.DataRate{
			Modulation: &ttnpb.DataRate_Lora{
				Lora: &ttnpb.LoRaDataRate{
					SpreadingFactor: 0,
					Bandwidth:       125000,
				},
			},
		}, "4/5")
		a.So(err, should.BeNil)
		_, err = Compute(len(downlink.RawPayload), downlink.GetScheduled())
		a.So(err, should.NotBeNil)
	}

	// Invalid bandwidth.
	{
		downlink, err := buildLoRaDownlinkFromParameters(10, 868100000, &ttnpb.DataRate{
			Modulation: &ttnpb.DataRate_Lora{
				Lora: &ttnpb.LoRaDataRate{
					SpreadingFactor: 7,
					Bandwidth:       0,
				},
			},
		}, "4/5")
		a.So(err, should.BeNil)
		_, err = Compute(len(downlink.RawPayload), downlink.GetScheduled())
		a.So(err, should.NotBeNil)
	}
}

func TestDifferentLoRaSFs(t *testing.T) {
	a := assertions.New(t)
	sfTests := []struct {
		dr       *ttnpb.DataRate
		expected uint
	}{
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 7, Bandwidth: 125000}}},
			expected: 41216,
		},
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 8, Bandwidth: 125000}}},
			expected: 72192,
		},
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 9, Bandwidth: 125000}}},
			expected: 144384,
		},
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 10, Bandwidth: 125000}}},
			expected: 288768,
		},
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 11, Bandwidth: 125000}}},
			expected: 577536,
		},
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 12, Bandwidth: 125000}}},
			expected: 991232,
		},
	}
	for _, tt := range sfTests {
		dl, err := buildLoRaDownlinkFromParameters(10, 868100000, tt.dr, "4/5")
		a.So(err, should.BeNil)
		toa, err := Compute(len(dl.RawPayload), dl.GetScheduled())
		a.So(err, should.BeNil)
		a.So(toa, should.AlmostEqual, time.Duration(tt.expected)*time.Microsecond)
	}
}

func TestDifferentLoRaBWs(t *testing.T) {
	a := assertions.New(t)
	bwTests := []struct {
		dr       *ttnpb.DataRate
		expected uint
	}{
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 7, Bandwidth: 125000}}},
			expected: 41216,
		},
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 7, Bandwidth: 250000}}},
			expected: 20608,
		},
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 7, Bandwidth: 500000}}},
			expected: 10304,
		},
	}
	for _, tt := range bwTests {
		dl, err := buildLoRaDownlinkFromParameters(10, 868100000, tt.dr, "4/5")
		a.So(err, should.BeNil)
		toa, err := Compute(len(dl.RawPayload), dl.GetScheduled())
		a.So(err, should.BeNil)
		a.So(toa, should.AlmostEqual, time.Duration(tt.expected)*time.Microsecond)
	}
}

func TestDifferentLoRaCRs(t *testing.T) {
	a := assertions.New(t)
	crTests := map[string]uint{
		"4/5": 41216,
		"4/6": 45312,
		"4/7": 49408,
		"4/8": 53504,
	}
	for cr, us := range crTests {
		dl, err := buildLoRaDownlinkFromParameters(10, 868100000, &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 7, Bandwidth: 125000}}}, cr)
		a.So(err, should.BeNil)
		toa, err := Compute(len(dl.RawPayload), dl.GetScheduled())
		a.So(err, should.BeNil)
		a.So(toa, should.AlmostEqual, time.Duration(us)*time.Microsecond)
	}
}

func TestDifferentLoRaPayloadSizes(t *testing.T) {
	a := assertions.New(t)
	plTests := map[int]uint{
		13: 46336,
		14: 46336,
		15: 46336,
		16: 51456,
		17: 51456,
		18: 51456,
		19: 51456,
	}
	for size, us := range plTests {
		dl, err := buildLoRaDownlinkFromParameters(size, 868100000, &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 7, Bandwidth: 125000}}}, "4/5")
		a.So(err, should.BeNil)
		toa, err := Compute(len(dl.RawPayload), dl.GetScheduled())
		a.So(err, should.BeNil)
		a.So(toa, should.AlmostEqual, time.Duration(us)*time.Microsecond)
	}
}

func TestFSK(t *testing.T) {
	a := assertions.New(t)
	payloadSize := 200
	scheduled := &ttnpb.TxSettings{
		Frequency: 868300000,
		DataRate: &ttnpb.DataRate{
			Modulation: &ttnpb.DataRate_Fsk{
				Fsk: &ttnpb.FSKDataRate{
					BitRate: 50000,
				},
			},
		},
	}
	toa, err := Compute(payloadSize, scheduled)
	a.So(err, should.BeNil)
	a.So(toa, should.AlmostEqual, 33760*time.Microsecond)
}

func getDownlink() ttnpb.DownlinkMessage { return ttnpb.DownlinkMessage{} }

func ExampleCompute() {
	downlink := getDownlink()
	toa, err := Compute(len(downlink.RawPayload), downlink.GetScheduled())
	if err != nil {
		panic(err)
	}

	fmt.Println("Time on air:", toa)
}

func TestInvalidLoRa2400(t *testing.T) {
	a := assertions.New(t)

	// Invalid coding rate.
	{
		downlink, err := buildLoRaDownlinkFromParameters(10, 2422000000, &ttnpb.DataRate{
			Modulation: &ttnpb.DataRate_Lora{
				Lora: &ttnpb.LoRaDataRate{
					SpreadingFactor: 10,
					Bandwidth:       812000,
				},
			},
		}, "1/9LI")
		a.So(err, should.BeNil)
		_, err = Compute(len(downlink.RawPayload), downlink.GetScheduled())
		a.So(err, should.NotBeNil)
	}

	// Invalid spreading factor.
	{
		downlink, err := buildLoRaDownlinkFromParameters(10, 2422000000, &ttnpb.DataRate{
			Modulation: &ttnpb.DataRate_Lora{
				Lora: &ttnpb.LoRaDataRate{
					SpreadingFactor: 0,
					Bandwidth:       812000,
				},
			},
		}, "4/5LI")
		a.So(err, should.BeNil)
		_, err = Compute(len(downlink.RawPayload), downlink.GetScheduled())
		a.So(err, should.NotBeNil)
	}

	// Invalid bandwidth.
	{
		downlink, err := buildLoRaDownlinkFromParameters(10, 2422000000, &ttnpb.DataRate{
			Modulation: &ttnpb.DataRate_Lora{
				Lora: &ttnpb.LoRaDataRate{
					SpreadingFactor: 7,
					Bandwidth:       0,
				},
			},
		}, "4/5LI")
		a.So(err, should.BeNil)
		_, err = Compute(len(downlink.RawPayload), downlink.GetScheduled())
		a.So(err, should.NotBeNil)
	}
}

func TestDifferentLoRa2400SFs(t *testing.T) {
	a := assertions.New(t)
	sfTests := []struct {
		dr       *ttnpb.DataRate
		expected time.Duration
	}{
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 5, Bandwidth: 812000}}},
			expected: 1665000,
		},
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 6, Bandwidth: 812000}}},
			expected: 3093600,
		},
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 7, Bandwidth: 812000}}},
			expected: 5556700,
		},
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 8, Bandwidth: 812000}}},
			expected: 10482800,
		},
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 9, Bandwidth: 812000}}},
			expected: 19073900,
		},
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 10, Bandwidth: 812000}}},
			expected: 36886700,
		},
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 11, Bandwidth: 812000}}},
			expected: 73773400,
		},
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 12, Bandwidth: 812000}}},
			expected: 142502500,
		},
	}
	for _, tt := range sfTests {
		dl, err := buildLoRaDownlinkFromParameters(10, 2422000000, tt.dr, "4/5LI")
		a.So(err, should.BeNil)
		toa, err := Compute(len(dl.RawPayload), dl.GetScheduled())
		a.So(err, should.BeNil)
		a.So(toa, should.AlmostEqual, tt.expected, 50)
	}
}

func TestDifferentLoRa2400BWs(t *testing.T) {
	a := assertions.New(t)
	bwTests := []struct {
		dr       *ttnpb.DataRate
		expected time.Duration
	}{
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 7, Bandwidth: 203000}}},
			expected: 22226600,
		},
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 7, Bandwidth: 406000}}},
			expected: 11113300,
		},
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 7, Bandwidth: 812000}}},
			expected: 5556700,
		},
		{
			dr:       &ttnpb.DataRate{Modulation: &ttnpb.DataRate_Lora{Lora: &ttnpb.LoRaDataRate{SpreadingFactor: 7, Bandwidth: 1625000}}},
			expected: 2776600,
		},
	}
	for _, tt := range bwTests {
		dl, err := buildLoRaDownlinkFromParameters(10, 2422000000, tt.dr, "4/5LI")
		a.So(err, should.BeNil)
		toa, err := Compute(len(dl.RawPayload), dl.GetScheduled())
		a.So(err, should.BeNil)
		a.So(toa, should.AlmostEqual, tt.expected, 50)
	}
}

func TestDifferentLoRa2400CRs(t *testing.T) {
	a := assertions.New(t)
	crTests := map[string]time.Duration{
		"4/5LI": 5556700,
		"4/6LI": 6029600,
		"4/7LI": 6817700,
		"4/8LI": 6817700,
	}
	for cr, ns := range crTests {
		dl, err := buildLoRaDownlinkFromParameters(10, 2422000000, &ttnpb.DataRate{
			Modulation: &ttnpb.DataRate_Lora{
				Lora: &ttnpb.LoRaDataRate{
					SpreadingFactor: 7,
					Bandwidth:       812000,
				},
			},
		}, cr)
		a.So(err, should.BeNil)
		toa, err := Compute(len(dl.RawPayload), dl.GetScheduled())
		a.So(err, should.BeNil)
		a.So(toa, should.AlmostEqual, ns, 50)
	}
}

func TestDifferentLoRa2400PayloadSizes(t *testing.T) {
	a := assertions.New(t)
	plTests := map[int]time.Duration{
		1:   102147800,
		10:  142502500,
		20:  192945800,
		50:  344275900,
		100: 596492600,
		230: 1252256200,
	}
	for size, ns := range plTests {
		dl, err := buildLoRaDownlinkFromParameters(size, 2422000000, &ttnpb.DataRate{
			Modulation: &ttnpb.DataRate_Lora{
				Lora: &ttnpb.LoRaDataRate{
					SpreadingFactor: 12,
					Bandwidth:       812000,
				},
			},
		}, "4/5LI")
		a.So(err, should.BeNil)
		toa, err := Compute(len(dl.RawPayload), dl.GetScheduled())
		a.So(err, should.BeNil)
		a.So(toa, should.AlmostEqual, ns, 50)
	}
}

func TestDifferentLoRa2400CRCs(t *testing.T) {
	a := assertions.New(t)
	crcTests := map[bool]time.Duration{
		true:  6029600,
		false: 5556700,
	}
	for crc, ns := range crcTests {
		dl, err := buildLoRaDownlinkFromParameters(10, 2422000000, &ttnpb.DataRate{
			Modulation: &ttnpb.DataRate_Lora{
				Lora: &ttnpb.LoRaDataRate{
					SpreadingFactor: 7,
					Bandwidth:       812000,
				},
			},
		}, "4/5LI")
		dl.GetScheduled().EnableCrc = crc
		a.So(err, should.BeNil)
		toa, err := Compute(len(dl.RawPayload), dl.GetScheduled())
		a.So(err, should.BeNil)
		a.So(toa, should.AlmostEqual, ns, 50)
	}
}
