package cmn

// 构建Map
// 如：OfMap("n1",1, "n2","v2", "n3",nil)
func OfMap(kvs ...any) map[string]any {
	m := make(map[string]any)
	for i := 0; i < len(kvs)-1; i += 2 {
		k := kvs[i].(string)
		m[k] = kvs[i+1]
	}
	return m
}
