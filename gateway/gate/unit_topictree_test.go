package gateway

import (
	"testing"
)

func eok(e error, t *testing.T) {
	if e != nil {
		t.Fatalf("ERROR %s\n", e)
	}
}

func enok(e error, t *testing.T) {
	if e == nil {
		t.Fatalf("ERROR (NO ERROR)")
	}
}

func chkb(b bool, e bool, t *testing.T) {
	if b != e {
		t.Fatalf("ERROR bool\n")
	}
}

func Test_NewTopicTree(t *testing.T) {
	tt := NewTopicTree()
	if tt == nil {
		t.Fatalf("NewTopicTree was nil")
	}
	if tt.root == nil {
		t.Fatalf("NewTopicTree.Root was nil")
	}
}

func Test_AddSubscription_alpha(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("c1", conn, addr)
	tt := NewTopicTree()
	b, e := tt.AddSubscription(c, "alpha")
	eok(e, t)
	chkb(b, true, t)
	if len(tt.root.clients) != 0 {
		t.Fatalf("AddSub root had client")
	}
	if len(tt.root.children["alpha"].clients) != 1 {
		t.Fatalf("AddSub alpha != 1")
	}
}

func Test_AddSubscription_Salpha(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("c2", conn, addr)
	tt := NewTopicTree()
	b, e := tt.AddSubscription(c, "/alpha")
	eok(e, t)
	chkb(b, true, t)
}

func Test_AddSubscription_alphaS(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("c2.5", conn, addr)
	tt := NewTopicTree()
	b, e := tt.AddSubscription(c, "alpha/")
	enok(e, t)
	chkb(b, false, t)
}

func Test_AddSubscription_aSbScSd(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("c3", conn, addr)
	tt := NewTopicTree()
	b, e := tt.AddSubscription(c, "a/b/c/d")
	eok(e, t)
	chkb(b, true, t)
}

func Test_AddSubscription_aSSb(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("c4", conn, addr)
	tt := NewTopicTree()
	b, e := tt.AddSubscription(c, "a//b")
	enok(e, t)
	chkb(b, false, t)
}

func Test_AddSubscription_H(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("c5", conn, addr)
	tt := NewTopicTree()
	b, e := tt.AddSubscription(c, "#")
	eok(e, t)
	chkb(b, true, t)
}

func Test_AddSubscription_SH(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("c6", conn, addr)
	tt := NewTopicTree()
	b, e := tt.AddSubscription(c, "/#")
	eok(e, t)
	chkb(b, true, t)
}

func Test_AddSubscription_aSHSb(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("c7", conn, addr)
	tt := NewTopicTree()
	b, e := tt.AddSubscription(c, "a/#/b")
	enok(e, t)
	chkb(b, false, t)
}

func Test_AddSubscription_aSHS(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("c8", conn, addr)
	tt := NewTopicTree()
	b, e := tt.AddSubscription(c, "a/#/")
	enok(e, t)
	chkb(b, false, t)
}

func Test_AddSubscription_P(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("c9", conn, addr)
	tt := NewTopicTree()
	b, e := tt.AddSubscription(c, "+")
	eok(e, t)
	chkb(b, true, t)
}

func Test_AddSubscription_SPSbSP(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("c10", conn, addr)
	tt := NewTopicTree()
	b, e := tt.AddSubscription(c, "/+/b/+")
	eok(e, t)
	chkb(b, true, t)
}

func Test_AddSubscription(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("c", conn, addr)
	tt := NewTopicTree()

	topics := []string{
		"a",
		"a/b",
		"a/b/c",
		"a/b/d",
		"a/b/e",
		"a/b/f",
		"a/b/c/d/e/f/g/h/i/j/k",
		"a/b/c/d/e/f/g/h/i/j/+",
		"a/b/c/d/e/f/g/h/i/j/k/+",
		"a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t/u/v/w/x/y/z",
		"a/b/c/d/e/f/g/h/i/j/k",
		"+/b/c",
		"a/+/b",
		"a/b/+",
		"a/b/#",
		"#",
	}

	for _, topic := range topics {
		_, e := tt.AddSubscription(c, topic)
		eok(e, t)
	}
}

func Benchmark_AddSubscription(b *testing.B) {
	var conn uConn
	var addr uAddr
	c := NewClient("b", conn, addr)
	tt := NewTopicTree()

	topics := []string{
		"a",
		"a/b",
		"a/b/c",
		"a/b/d",
		"a/b/e",
		"a/b/f",
		"a/b/c/d/e/f/g/h/i/j/k",
		"a/b/c/d/e/f/g/h/i/j/+",
		"a/b/c/d/e/f/g/h/i/j/k/+",
		"a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t/u/v/w/x/y/z",
		"a/b/c/d/e/f/g/h/i/j/k",
		"+/b/c",
		"a/+/b",
		"a/b/+",
		"a/b/#",
		"#",
	}

	for i := 0; i < b.N; i++ {
		for _, topic := range topics {
			tt.AddSubscription(c, topic)
		}
	}
}

func Test_SubscribersOf_none(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("so_A", conn, addr)
	tt := NewTopicTree()

	sum := 0
	sum += len(tt.SubscribersOf("+"))
	sum += len(tt.SubscribersOf("#"))
	sum += len(tt.SubscribersOf("a"))
	sum += len(tt.SubscribersOf("/+"))
	sum += len(tt.SubscribersOf("/#"))
	sum += len(tt.SubscribersOf("/a"))

	tt.AddSubscription(c, "kappa")
	sum += len(tt.SubscribersOf("+"))
	sum += len(tt.SubscribersOf("#"))
	sum += len(tt.SubscribersOf("a"))
	sum += len(tt.SubscribersOf("/+"))
	sum += len(tt.SubscribersOf("/#"))
	sum += len(tt.SubscribersOf("/a"))

	tt.AddSubscription(c, "/kappa")
	sum += len(tt.SubscribersOf("+"))
	sum += len(tt.SubscribersOf("#"))
	sum += len(tt.SubscribersOf("a"))
	sum += len(tt.SubscribersOf("/+"))
	sum += len(tt.SubscribersOf("/#"))
	sum += len(tt.SubscribersOf("/a"))

	tt.AddSubscription(c, "kappa/#")
	sum += len(tt.SubscribersOf("+"))
	sum += len(tt.SubscribersOf("#"))
	sum += len(tt.SubscribersOf("a"))
	sum += len(tt.SubscribersOf("/+"))
	sum += len(tt.SubscribersOf("/#"))
	sum += len(tt.SubscribersOf("/a"))

	tt.AddSubscription(c, "kappa/+")
	sum += len(tt.SubscribersOf("+"))
	sum += len(tt.SubscribersOf("#"))
	sum += len(tt.SubscribersOf("a"))
	sum += len(tt.SubscribersOf("/+"))
	sum += len(tt.SubscribersOf("/#"))
	sum += len(tt.SubscribersOf("/a"))

	tt.AddSubscription(c, "a/b")
	sum += len(tt.SubscribersOf("+"))
	sum += len(tt.SubscribersOf("#"))
	sum += len(tt.SubscribersOf("a"))
	sum += len(tt.SubscribersOf("/+"))
	sum += len(tt.SubscribersOf("/#"))
	sum += len(tt.SubscribersOf("/a"))

	tt.AddSubscription(c, "b/a")
	sum += len(tt.SubscribersOf("+"))
	sum += len(tt.SubscribersOf("#"))
	sum += len(tt.SubscribersOf("a"))
	sum += len(tt.SubscribersOf("/+"))
	sum += len(tt.SubscribersOf("/#"))
	sum += len(tt.SubscribersOf("/a"))

	tt.AddSubscription(c, "/a/#")
	sum += len(tt.SubscribersOf("+"))
	sum += len(tt.SubscribersOf("#"))
	sum += len(tt.SubscribersOf("a"))
	sum += len(tt.SubscribersOf("/+"))
	sum += len(tt.SubscribersOf("/#"))
	sum += len(tt.SubscribersOf("/a"))

	if sum != 0 {
		t.Fatalf("SubscribersOf_none had subscriber")
	}
}

func alen(exp, act, i int, t *testing.T) {
	if exp != act {
		t.Errorf("assert #%d expected %d, got %d", i, exp, act)
	}
}

func Test_SubscribersOf_1(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("so_1", conn, addr)
	tt := NewTopicTree()

	tt.AddSubscription(c, "a")

	alen(1, len(tt.SubscribersOf("a")), 1, t)
	alen(0, len(tt.SubscribersOf("b")), 2, t)
	alen(0, len(tt.SubscribersOf("/a")), 3, t)
	alen(0, len(tt.SubscribersOf("/b")), 4, t)
}

func Test_SubscribersOf_1h(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("so_1h", conn, addr)
	tt := NewTopicTree()

	tt.AddSubscription(c, "#")

	alen(1, len(tt.SubscribersOf("a")), 1, t)
	alen(1, len(tt.SubscribersOf("b")), 2, t)
	alen(1, len(tt.SubscribersOf("/a")), 3, t)
	alen(1, len(tt.SubscribersOf("/b")), 4, t)
}

func Test_SubscribersOf_1sh(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("so_1h", conn, addr)
	tt := NewTopicTree()

	tt.AddSubscription(c, "/#")

	alen(0, len(tt.SubscribersOf("a")), 1, t)
	alen(0, len(tt.SubscribersOf("b")), 2, t)
	alen(1, len(tt.SubscribersOf("/a")), 3, t)
	alen(1, len(tt.SubscribersOf("/b")), 4, t)
}

func Test_SubscribersOf_1p(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("so_1h", conn, addr)
	tt := NewTopicTree()

	tt.AddSubscription(c, "+")

	alen(1, len(tt.SubscribersOf("a")), 1, t)
	alen(1, len(tt.SubscribersOf("b")), 2, t)
	alen(0, len(tt.SubscribersOf("/a")), 3, t)
	alen(0, len(tt.SubscribersOf("/b")), 4, t)
}

func Test_SubscribersOf_1sp(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("so_1h", conn, addr)
	tt := NewTopicTree()

	tt.AddSubscription(c, "/+")

	alen(0, len(tt.SubscribersOf("a")), 1, t)
	alen(0, len(tt.SubscribersOf("b")), 2, t)
	alen(1, len(tt.SubscribersOf("/a")), 3, t)
	alen(1, len(tt.SubscribersOf("/b")), 4, t)
}

func Test_SubscribersOf_2(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("so_1", conn, addr)
	tt := NewTopicTree()

	tt.AddSubscription(c, "a/b")

	alen(0, len(tt.SubscribersOf("a")), 1, t)
	alen(0, len(tt.SubscribersOf("b")), 2, t)
	alen(0, len(tt.SubscribersOf("/a")), 3, t)
	alen(0, len(tt.SubscribersOf("/b")), 4, t)
	alen(1, len(tt.SubscribersOf("a/b")), 5, t)
	alen(0, len(tt.SubscribersOf("b/a")), 6, t)
	alen(0, len(tt.SubscribersOf("a/a")), 7, t)
	alen(0, len(tt.SubscribersOf("a/b/c")), 8, t)
}

func Test_SubscribersOf_2ash(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("so_1", conn, addr)
	tt := NewTopicTree()

	tt.AddSubscription(c, "a/#")

	alen(0, len(tt.SubscribersOf("a")), 1, t)
	alen(0, len(tt.SubscribersOf("b")), 2, t)
	alen(0, len(tt.SubscribersOf("/a")), 3, t)
	alen(0, len(tt.SubscribersOf("/b")), 4, t)
	alen(1, len(tt.SubscribersOf("a/b")), 5, t)
	alen(0, len(tt.SubscribersOf("b/a")), 6, t)
	alen(1, len(tt.SubscribersOf("a/a")), 7, t)
	alen(1, len(tt.SubscribersOf("a/b/c")), 8, t)
}

func Test_SubscribersOf_2h(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("so_1", conn, addr)
	tt := NewTopicTree()

	tt.AddSubscription(c, "#")

	alen(1, len(tt.SubscribersOf("a")), 1, t)
	alen(1, len(tt.SubscribersOf("b")), 2, t)
	alen(1, len(tt.SubscribersOf("/a")), 3, t)
	alen(1, len(tt.SubscribersOf("/b")), 4, t)
	alen(1, len(tt.SubscribersOf("a/b")), 5, t)
	alen(1, len(tt.SubscribersOf("b/a")), 6, t)
	alen(1, len(tt.SubscribersOf("a/a")), 7, t)
	alen(1, len(tt.SubscribersOf("a/b/c")), 8, t)
}

func Test_SubscribersOf_2p(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("so_1", conn, addr)
	tt := NewTopicTree()

	tt.AddSubscription(c, "+")

	alen(1, len(tt.SubscribersOf("a")), 1, t)
	alen(1, len(tt.SubscribersOf("b")), 2, t)
	alen(0, len(tt.SubscribersOf("/a")), 3, t)
	alen(0, len(tt.SubscribersOf("/b")), 4, t)
	alen(0, len(tt.SubscribersOf("a/b")), 5, t)
	alen(0, len(tt.SubscribersOf("b/a")), 6, t)
	alen(0, len(tt.SubscribersOf("a/a")), 7, t)
	alen(0, len(tt.SubscribersOf("a/b/c")), 8, t)
}

func Test_SubscribersOf_2asp(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("so_1", conn, addr)
	tt := NewTopicTree()

	tt.AddSubscription(c, "a/+")

	alen(0, len(tt.SubscribersOf("a")), 1, t)
	alen(0, len(tt.SubscribersOf("b")), 2, t)
	alen(0, len(tt.SubscribersOf("/a")), 3, t)
	alen(0, len(tt.SubscribersOf("/b")), 4, t)
	alen(1, len(tt.SubscribersOf("a/b")), 5, t)
	alen(0, len(tt.SubscribersOf("b/a")), 6, t)
	alen(1, len(tt.SubscribersOf("a/a")), 7, t)
	alen(0, len(tt.SubscribersOf("a/b/c")), 8, t)
}

func Test_SubscribersOf_3aspsc(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("so_1", conn, addr)
	tt := NewTopicTree()

	tt.AddSubscription(c, "a/+/c")

	alen(0, len(tt.SubscribersOf("a")), 1, t)
	alen(0, len(tt.SubscribersOf("b")), 2, t)
	alen(0, len(tt.SubscribersOf("/a")), 3, t)
	alen(0, len(tt.SubscribersOf("/b")), 4, t)
	alen(0, len(tt.SubscribersOf("a/b")), 5, t)
	alen(0, len(tt.SubscribersOf("b/a")), 6, t)
	alen(0, len(tt.SubscribersOf("a/a")), 7, t)
	alen(1, len(tt.SubscribersOf("a/b/c")), 8, t)
	alen(0, len(tt.SubscribersOf("a/b/z")), 8, t)
}

func Test_SubscribersOf_3saspsc(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("so_1", conn, addr)
	tt := NewTopicTree()

	tt.AddSubscription(c, "/a/+/c")

	alen(0, len(tt.SubscribersOf("a")), 1, t)
	alen(0, len(tt.SubscribersOf("b")), 2, t)
	alen(0, len(tt.SubscribersOf("/a")), 3, t)
	alen(0, len(tt.SubscribersOf("/b")), 4, t)
	alen(0, len(tt.SubscribersOf("a/b")), 5, t)
	alen(0, len(tt.SubscribersOf("b/a")), 6, t)
	alen(0, len(tt.SubscribersOf("a/a")), 7, t)
	alen(0, len(tt.SubscribersOf("a/b/c")), 8, t)
	alen(0, len(tt.SubscribersOf("a/b/z")), 9, t)
}

func Test_SubscribersOf_mix(t *testing.T) {
	var conn uConn
	var addr uAddr
	c := NewClient("so_1", conn, addr)
	tt := NewTopicTree()

	tt.AddSubscription(c, "/a/+/c/#")

	alen(0, len(tt.SubscribersOf("a")), 1, t)
	alen(0, len(tt.SubscribersOf("b")), 2, t)
	alen(0, len(tt.SubscribersOf("/a")), 3, t)
	alen(0, len(tt.SubscribersOf("/b")), 4, t)
	alen(0, len(tt.SubscribersOf("/a/b")), 5, t)
	alen(0, len(tt.SubscribersOf("/b/a")), 6, t)
	alen(0, len(tt.SubscribersOf("/a/a")), 7, t)
	alen(0, len(tt.SubscribersOf("/a/b/c")), 8, t)
	alen(0, len(tt.SubscribersOf("/a/b/z")), 9, t)
	alen(1, len(tt.SubscribersOf("/a/b/c/d")), 10, t)
	alen(1, len(tt.SubscribersOf("/a/b/c/d/e")), 11, t)
	alen(1, len(tt.SubscribersOf("/a/b/c/d/e/f")), 12, t)
	alen(0, len(tt.SubscribersOf("/a/b/z/d/e/f")), 13, t)
	alen(0, len(tt.SubscribersOf("/a/b/b/c/d/e/f")), 14, t)
	alen(0, len(tt.SubscribersOf("a/b/c/d")), 15, t)
	alen(0, len(tt.SubscribersOf("a/b/c/d/e")), 16, t)
	alen(0, len(tt.SubscribersOf("a/b/c/d/e/f")), 17, t)
}

func Test_SubscribersOf_multi(t *testing.T) {
	var conn uConn
	var addr uAddr
	c1 := NewClient("c1", conn, addr)
	c2 := NewClient("c2", conn, addr)
	c3 := NewClient("c3", conn, addr)
	c4 := NewClient("c4", conn, addr)
	tt := NewTopicTree()

	tt.AddSubscription(c1, "a")
	tt.AddSubscription(c2, "/a/+/c/d")
	tt.AddSubscription(c3, "/a/+/c/#")
	tt.AddSubscription(c4, "/a/+/c/d/+")

	alen(1, len(tt.SubscribersOf("a")), 1, t)
	alen(0, len(tt.SubscribersOf("b")), 2, t)
	alen(0, len(tt.SubscribersOf("/a")), 3, t)
	alen(0, len(tt.SubscribersOf("/b")), 4, t)
	alen(0, len(tt.SubscribersOf("/a/b")), 5, t)
	alen(0, len(tt.SubscribersOf("/b/a")), 6, t)
	alen(0, len(tt.SubscribersOf("/a/a")), 7, t)
	alen(0, len(tt.SubscribersOf("/a/b/c")), 8, t)
	alen(0, len(tt.SubscribersOf("/a/b/z")), 9, t)
	alen(2, len(tt.SubscribersOf("/a/b/c/d")), 10, t)
	alen(2, len(tt.SubscribersOf("/a/b/c/d/e")), 11, t)
	alen(1, len(tt.SubscribersOf("/a/b/c/d/e/f")), 12, t)
	alen(0, len(tt.SubscribersOf("/a/b/z/d/e/f")), 13, t)
	alen(0, len(tt.SubscribersOf("/a/b/b/c/d/e/f")), 14, t)
	alen(0, len(tt.SubscribersOf("a/b/c/d")), 15, t)
	alen(0, len(tt.SubscribersOf("a/b/c/d/e")), 16, t)
	alen(0, len(tt.SubscribersOf("a/b/c/d/e/f")), 17, t)
}
