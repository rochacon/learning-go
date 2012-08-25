package sqrt

import "testing"

func TestSqrt_1(t *testing.T) {
    in, out := 1.0, 1.0
    if x := sqrt(in); x != out {
        t.Errorf("sqrt(%v) was not %v: %v", in, out, x)
    }
}

func TestSqrt_4(t *testing.T) {
    in, out := 4.0, 2.0
    if x := sqrt(in); x != out {
        t.Errorf("sqrt(%v) was not %v: %v", in, out, x)
    }
}

func TestSqrt_9(t *testing.T) {
    in, out := 9.0, 3.0
    if x := sqrt(in); x != out {
        t.Errorf("sqrt(%v) was not %v: %v", in, out, x)
    }
}

func TestSqrt_16(t *testing.T) {
    in, out := 16.0, 4.0
    if x := sqrt(in); x != out {
        t.Errorf("sqrt(%v) was not %v: %v", in, out, x)
    }
}

func TestSqrt_25(t *testing.T) {
    in, out := 25.0, 5.0
    if x := sqrt(in); x != out {
        t.Errorf("sqrt(%v) was not %v: %v", in, out, x)
    }
}

func TestSqrt_36(t *testing.T) {
    in, out := 36.0, 6.0
    if x := sqrt(in); x != out {
        t.Errorf("sqrt(%v) was not %v: %v", in, out, x)
    }
}

func TestSqrt_64(t *testing.T) {
    in, out := 64.0, 8.0
    if x := sqrt(in); x != out {
        t.Errorf("sqrt(%v) was not %v: %v", in, out, x)
    }
}

func TestSqrt_81(t *testing.T) {
    in, out := 81.0, 9.0
    if x := sqrt(in); x != out {
        t.Errorf("sqrt(%v) was not %v: %v", in, out, x)
    }
}

func TestSqrt_100(t *testing.T) {
    in, out := 100.0, 10.0
    if x := sqrt(in); x != out {
        t.Errorf("sqrt(%v) was not %v: %v", in, out, x)
    }
}

func TestSqrt_144(t *testing.T) {
    in, out := 144.0, 12.0
    if x := sqrt(in); x != out {
        t.Errorf("sqrt(%v) was not %v: %v", in, out, x)
    }
}

func TestSqrt_10000(t *testing.T) {
    in, out := 10000.0, 100.0
    if x := sqrt(in); x != out {
        t.Errorf("sqrt(%v) was not %v: %v", in, out, x)
    }
}

func TestSqrt_1000000(t *testing.T) {
    in, out := 1000000.0, 1000.0
    if x := sqrt(in); x != out {
        t.Errorf("sqrt(%v) was not %v: %v", in, out, x)
    }
}

func TestSqrt_1000000000000(t *testing.T) {
    in, out := 1000000000000.0, 1000000.0
    if x := sqrt(in); x != out {
        t.Errorf("sqrt(%v) was not %v: %v", in, out, x)
    }
}
