package hash

// HashMap 字符串键哈希表,链地址法,负载因子超过 1 时倍增扩容
type HashMap struct {
	buckets [][]entry
	size    int
}

type entry struct {
	key   string
	value any
}

const (
	initialBuckets = 16
	maxLoadFactor  = 1.0
)

func NewHashMap() *HashMap {
	return &HashMap{
		buckets: make([][]entry, initialBuckets),
	}
}

func (this *HashMap) Size() int {
	return this.size
}

func (this *HashMap) IsEmpty() bool {
	return this.size == 0
}

// Put 写入键值,已存在则覆盖并返回旧值
func (this *HashMap) Put(key string, value any) (old any, existed bool) {
	if float64(this.size+1)/float64(len(this.buckets)) > maxLoadFactor {
		this.resize()
	}
	idx := this.indexOf(key)
	for i := range this.buckets[idx] {
		if this.buckets[idx][i].key == key {
			old = this.buckets[idx][i].value
			this.buckets[idx][i].value = value
			return old, true
		}
	}
	this.buckets[idx] = append(this.buckets[idx], entry{key: key, value: value})
	this.size++
	return nil, false
}

// Get 读取键值,不存在返回 false
func (this *HashMap) Get(key string) (any, bool) {
	idx := this.indexOf(key)
	for i := range this.buckets[idx] {
		if this.buckets[idx][i].key == key {
			return this.buckets[idx][i].value, true
		}
	}
	return nil, false
}

// Contains 判断键是否存在
func (this *HashMap) Contains(key string) bool {
	_, ok := this.Get(key)
	return ok
}

// Delete 删除键,存在返回旧值
func (this *HashMap) Delete(key string) (old any, existed bool) {
	idx := this.indexOf(key)
	for i := range this.buckets[idx] {
		if this.buckets[idx][i].key == key {
			old = this.buckets[idx][i].value
			this.buckets[idx] = append(this.buckets[idx][:i], this.buckets[idx][i+1:]...)
			this.size--
			return old, true
		}
	}
	return nil, false
}

// Keys 返回全部键(顺序不定)
func (this *HashMap) Keys() []string {
	keys := make([]string, 0, this.size)
	for _, bucket := range this.buckets {
		for _, e := range bucket {
			keys = append(keys, e.key)
		}
	}
	return keys
}

// fnv1a FNV-1a 32 位哈希
func fnv1a(key string) uint32 {
	var h uint32 = 2166136261
	for i := 0; i < len(key); i++ {
		h ^= uint32(key[i])
		h *= 16777619
	}
	return h
}

func (this *HashMap) indexOf(key string) int {
	return int(fnv1a(key) & uint32(len(this.buckets)-1))
}

func (this *HashMap) resize() {
	old := this.buckets
	this.buckets = make([][]entry, len(old)*2)
	this.size = 0
	for _, bucket := range old {
		for _, e := range bucket {
			idx := this.indexOf(e.key)
			this.buckets[idx] = append(this.buckets[idx], e)
			this.size++
		}
	}
}
