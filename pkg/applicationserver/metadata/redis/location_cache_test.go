// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
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

package redis_test

import (
	"testing"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/metadata/redis"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

var (
	registeredEndDeviceIDs = ttnpb.EndDeviceIdentifiers{
		ApplicationIds: &ttnpb.ApplicationIdentifiers{
			ApplicationId: "foo",
		},
		DeviceId: "bar",
	}

	locationA = map[string]*ttnpb.Location{
		"foo": {
			Latitude:  123,
			Longitude: 234,
		},
		"bar": {
			Latitude:  345,
			Longitude: 456,
		},
	}
	locationB = map[string]*ttnpb.Location{
		"baz": {
			Latitude: 567,
		},
	}
	errUnavailable = errors.DefineUnavailable("unavailable", "unavailable")
)

func TestLocationCache(t *testing.T) {
	a, ctx := test.New(t)
	cl, flush := test.NewRedis(ctx, "metadata_redis_test")
	defer flush()
	cache := &redis.EndDeviceLocationCache{
		Redis: cl,
	}

	_, _, err := cache.Get(ctx, registeredEndDeviceIDs)
	a.So(err, should.NotBeNil)
	a.So(errors.IsNotFound(err), should.BeTrue)

	err = cache.SetLocations(ctx, registeredEndDeviceIDs, locationA, 10*time.Second)
	a.So(err, should.BeNil)

	locations, ttl, err := cache.Get(ctx, registeredEndDeviceIDs)
	if a.So(err, should.BeNil) {
		a.So(ttl, should.BeGreaterThan, 0)
		a.So(ttl, should.BeLessThanOrEqualTo, 10*time.Second)
		a.So(len(locations), should.Equal, len(locationA))
		for k, v := range locations {
			a.So(locationA[k], should.Resemble, v)
		}
	}

	err = cache.SetErrorDetails(ctx, registeredEndDeviceIDs, ttnpb.ErrorDetailsToProto(errUnavailable.New()), 5*time.Second)
	a.So(err, should.BeNil)

	_, ttl, err = cache.Get(ctx, registeredEndDeviceIDs)
	a.So(err, should.NotBeNil)
	a.So(err, should.HaveSameErrorDefinitionAs, errUnavailable)
	a.So(ttl, should.BeGreaterThan, 0)
	a.So(ttl, should.BeLessThanOrEqualTo, 5*time.Second)

	err = cache.SetLocations(ctx, registeredEndDeviceIDs, locationB, 10*time.Second)
	a.So(err, should.BeNil)

	locations, ttl, err = cache.Get(ctx, registeredEndDeviceIDs)
	if a.So(err, should.BeNil) {
		a.So(ttl, should.BeGreaterThan, 0)
		a.So(ttl, should.BeLessThanOrEqualTo, 10*time.Second)
		a.So(len(locations), should.Equal, len(locationB))
		for k, v := range locations {
			a.So(locationB[k], should.Resemble, v)
		}
	}

	err = cache.Delete(ctx, registeredEndDeviceIDs)
	a.So(err, should.BeNil)

	_, _, err = cache.Get(ctx, registeredEndDeviceIDs)
	a.So(err, should.NotBeNil)
	a.So(errors.IsNotFound(err), should.BeTrue)
}
