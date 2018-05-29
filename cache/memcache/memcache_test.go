package memcache_test

import (
	"testing"
	"time"
	"github.com/qeelyn/go-common/cache/memcache"
	"fmt"
	"github.com/qeelyn/go-common/cache"
)

var (
	ins = memcache.NewMemCache()
)

type Foo struct {
	F1 string
	F2 int
}

type Bar struct {
	B1 string
	F1 Foo
}

func init() {
	ins.StartAndGC(map[string]interface{}{
		"addr": "127.0.0.1:11211",
		"prefix": "redis:",
	})
}

func initTestData(t *testing.T)  {
	var a, b, c = "abc", 123456, 12345.123456
	var sl, mp = []int{22, 33, 44, 55}, map[string]string{"a": "abc", "b": "efg"}
	var obj = Bar{B1: "Bar", F1: Foo{F1: "abc", F2: 1111}}
	var prt = string("eeeee")
	var err error
	var du time.Duration = 1 * time.Hour
	if err = ins.Set("a", a, du); err != nil {
		t.Fatal(err)
	}
	if err = ins.Set("b", b, du); err != nil {
		t.Fatal(err)
	}
	if err = ins.Set("c", c, du); err != nil {
		t.Fatal(err)
	}
	if err = ins.Set("sl", sl, du); err != nil {
		t.Fatal(err)
	}
	if err = ins.Set("mp", mp, du); err != nil {
		t.Fatal(err)
	}
	if err = ins.Set("obj", obj, du); err != nil {
		t.Fatal(err)
	}
	if err = ins.Set("ptr", &prt, du); err != nil {
		t.Fatal(err)
	}
}

func TestCacheWrapper_Set(t *testing.T) {
	initTestData(t)
}

func TestCacheWrapper_Get(t *testing.T) {
	initTestData(t)
	var (
		a   string
		b   int = 0
		c 	float64
		sl  []int
		err error
		obj Bar
	)
	if err = ins.Get("a", &a); err != nil {
		t.Fatal(err)
	}
	if a != "abc" {
		t.Fatal("a no equeal")
	}
	if err = ins.Get("b", &b); err != nil {
		t.Fatal(err)
	}
	if b != 123456 {
		t.Fatal("b no equeal")
	}
	if err = ins.Get("c", &c); err != nil {
		t.Fatal(err)
	}
	if c != 12345.123456 {
		t.Fatal("c no equeal")
	}
	if err = ins.Get("sl", &sl); err != nil {
		t.Fatal(err)
	}
	if sl[0] != 22 {
		t.Fatal("sl no equeal")
	}
	if err = ins.Get("obj", &obj); err != nil {
		t.Fatal(err)
	}
	if obj.B1 != "Bar" {
		t.Fatal("obj no equeal")
	}
}

func TestCache_GetMulti(t *testing.T) {
	initTestData(t)
	//var (
	//	a   string
	//	b   int = 0
	//	c 	float64
	//	sl  []int
	//	err error
	//	obj Bar
	//)
	vals := ins.GetMulti([]string{"a","b","c","sl","obj"})
	fmt.Print(vals...)
}

func TestCacheWrapper_Incr(t *testing.T) {
	initTestData(t)
	var b int
	if err := ins.Incr("b"); err != nil {
		t.Fatal(err)
	}
	if err := ins.Get("b",&b);err != nil {
		t.Fatal(err)
	}
	if b != 123457 {
		t.Fatal("incr error")
	}
}

func TestCache_GetMiss(t *testing.T) {
	var a string
	if err := ins.Get("noexist",&a);err != nil {
		if err != cache.ErrCacheMiss {
			t.Error()
		}
	} else {
		t.Error()
	}
}

