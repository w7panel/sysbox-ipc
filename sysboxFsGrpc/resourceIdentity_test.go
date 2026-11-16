package sysboxFsGrpc

import (
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	pb "github.com/nestybox/sysbox-ipc/sysboxFsGrpc/sysboxFsProtobuf"
)

func TestResourceIdentityRoundTrip(t *testing.T) {
	for _, input := range []ContainerData{
		{Id: "legacy", Ctime: time.Unix(123, 0)},
		{Id: "seeded", Ctime: time.Unix(123, 0), CgroupData: "0::/containers/test\n", Pidns: "pid:[123]", InitStartTime: 456, ResourceCacheTtlNanos: int64(25 * time.Millisecond)},
	} {
		wire, err := containerDataToPbData(&input)
		if err != nil {
			t.Fatal(err)
		}
		raw, err := proto.Marshal(wire)
		if err != nil {
			t.Fatal(err)
		}
		var decoded pb.ContainerData
		if err := proto.Unmarshal(raw, &decoded); err != nil {
			t.Fatal(err)
		}
		got, err := pbDatatoContainerData(&decoded)
		if err != nil {
			t.Fatal(err)
		}
		if got.Id != input.Id || got.CgroupData != input.CgroupData || got.Pidns != input.Pidns || got.InitStartTime != input.InitStartTime || got.ResourceCacheTtlNanos != input.ResourceCacheTtlNanos {
			t.Fatalf("round trip: %#v, want %#v", got, input)
		}
	}
}
