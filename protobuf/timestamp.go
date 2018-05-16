package protobuf

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes"
	"time"
	"fmt"
)

type Timestamp timestamp.Timestamp

func (g *Timestamp) Scan(src interface{}) error {
	t, _ := ptypes.TimestampProto(src.(time.Time))
	g.Seconds = t.Seconds
	g.Nanos = t.Nanos
	return nil
}

func (g *Timestamp) ToTime() (time.Time,error) {
	return TimestampToTime(g)
}

func (g *Timestamp) MarshalJSON() ([]byte, error) {
	t,err := TimestampToTime(g)
	if err != nil {
		return nil,err
	}
	var stamp = fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}