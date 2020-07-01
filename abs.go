package conv

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
