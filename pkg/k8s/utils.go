package k8s

func Int32Ptr(i int) *int32 {
	scale := new(int32)
	*scale = int32(i)
	return scale
}
