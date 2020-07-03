package conv

import "time"

func AbsInt64(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}

func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func AbsInt32(i int32) int32 {
	if i < 0 {
		return -i
	}
	return i
}

func AbsFloat32(i float32) float32 {
	if i < 0 {
		return -i
	}
	return i
}

func AbsDuration(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}
