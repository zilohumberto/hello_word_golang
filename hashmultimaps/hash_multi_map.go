package hashmultimaps

func New() *HashMultiMap{
	return &HashMultiMap{make(map[interface{}][]interface{})}
}

type HashMultiMap struct {
	data map[interface{}][]interface{}
}

func (h *HashMultiMap) Merge(maps ...*HashMultiMap){
	for _, multiMap := range maps{
		for _, key := range multiMap.GetKeys(){
			values := multiMap.GetValues(key)
			h.PutAll(key, values...)
		}
	}
}

func (h *HashMultiMap) Put(key interface{}, value interface{}){
	values, ok := h.data[key]
	if !ok {
		values = make([]interface{}, 0)
	}
	values = append(values, value)
	h.data[key] = values
}

func (h *HashMultiMap) PutAll(key interface{}, values ...interface{}){
	for _, value := range values{
		h.Put(key, value)
	}
}

func (h * HashMultiMap) GetKeys() []interface{}{
	keys := make([]interface{}, 0, h.Size())
	for k := range h.data{
		keys = append(keys, k)
	}
	return keys
}

func (h * HashMultiMap) Size() int{
	return len(h.data)
}

func (h * HashMultiMap) Contains(key interface{}) bool{
	_, ok := h.data[key]
	return ok
}

func (h * HashMultiMap) GetValues(key interface{}) []interface{}{
	values, ok := h.data[key]
	if !ok{
		return make([]interface{}, 0)
	}
	return values
}

func (h * HashMultiMap) Clear(){
	h.data = make(map[interface{}][]interface{})
}

func (h * HashMultiMap) RemoveKeys(keys ...interface{}){
	for _, key := range keys{
		delete(h.data, key)
	}
}

func (h * HashMultiMap) Remove(key interface{}, value interface{}){
	values, ok := h.data[key]
	if !ok {
		return
	}
	index := h.getIndex(value, values...)
	if index == -1 {
		return
	}
	newValues := make([]interface{}, 0)
	newValues = append(newValues, values[:index]...)
	h.data[key] = append(newValues, values[index+1:]...)
}

func (h * HashMultiMap) getIndex(value interface{}, values ...interface{}) int{
	for i, v := range values{
		if v==value{
			return i
		}
	}
	return -1
}
