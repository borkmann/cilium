// Copyright 2019 Authors of Cilium
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

// +build !privileged_tests

package mock

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type MockSuite struct{}

var _ = check.Suite(&MockSuite{})

func (e *MockSuite) TestMock(c *check.C) {
	api := NewMockMetrics()
	api.IncENIAllocationAttempt("foo", "s-1")
	c.Assert(api.ENIAllocationAttempts("foo", "s-1"), check.Equals, int64(1))
	api.AddIPAllocation("s-1", 10)
	api.AddIPAllocation("s-1", 20)
	c.Assert(api.IPAllocations("s-1"), check.Equals, int64(30))
	api.SetAllocatedIPs("used", 200)
	c.Assert(api.AllocatedIPs("used"), check.Equals, 200)
	api.SetAvailableENIs(10)
	c.Assert(api.AvailableENIs(), check.Equals, 10)
	api.SetNodesAtCapacity(5)
	c.Assert(api.NodesAtCapacity(), check.Equals, 5)
	api.ObserveEC2APICall("DescribeNetworkInterfaces", "success", 2.0)
	c.Assert(api.EC2APICall("DescribeNetworkInterfaces", "success"), check.Equals, 2.0)
	api.IncEC2RateLimit("DescribeNetworkInterfaces")
	c.Assert(api.EC2RateLimit("DescribeNetworkInterfaces"), check.Equals, int64(1))
	api.IncResyncCount()
	c.Assert(api.ResyncCount(), check.Equals, int64(1))
}