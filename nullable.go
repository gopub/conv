package conv

func ToNullInt(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

func ToNullInt64(i int64) *int64 {
	if i == 0 {
		return nil
	}
	return &i
}

func ToNullBool(b bool) *bool {
	if b == false {
		return nil
	}
	return &b
}

func ToNullString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func ToNullFloat64(f float64) *float64 {
	if f == 0 {
		return nil
	}
	return &f
}
