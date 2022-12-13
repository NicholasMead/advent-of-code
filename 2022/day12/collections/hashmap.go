package collections

type Hashable interface {
	Hash() uint64
	comparable
}

type Map[K comparable, I any] interface {
	Store(K, I)
	Get(K) (I, bool)
}

func CreateHashmap[K Hashable, I any]() Map[K, I] {
	return &hashMap[K, I]{
		intMap: make(map[uint64]I),
	}
}

func CreateMap[K comparable, I any]() Map[K, I] {
	return &sndMap[K, I]{
		intMap: make(map[K]I),
	}
}

type sndMap[K comparable, I any] struct {
	intMap map[K]I
}

// Get implements Map
func (m *sndMap[K, I]) Get(key K) (I, bool) {
	i, f := m.intMap[key]
	return i, f
}

// Store implements Map
func (m *sndMap[K, I]) Store(key K, item I) {
	m.intMap[key] = item
}

type hashMap[K Hashable, I any] struct {
	intMap map[uint64]I
}

// Get implements HashMap
func (h *hashMap[K, I]) Get(key K) (I, bool) {
	item, found := h.intMap[key.Hash()]
	return item, found
}

// Store implements HashMap
func (h *hashMap[K, I]) Store(key K, item I) {
	h.intMap[key.Hash()] = item
}
