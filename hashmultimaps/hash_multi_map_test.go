package hashmultimaps

import (
	"testing"
)


func TestPut(t *testing.T){
	h := New()
	size := h.Size()
	if size != 0{
		t.Errorf("Excepted 0, got %d", size)
	}
	h.Put("#1", 1)
	size = h.Size()
	if size != 1{
		t.Errorf("Excepted 1, got %d", size)
	}

	h.Put("#2", 2)
	size = h.Size()
	if size != 2{
		t.Errorf("Excepted 2, got %d", size)
	}
	h.Put("#2", 23)
	size = h.Size()
	if size != 2{
		t.Errorf("Excepted 2, got %d", size)
	}
}

func TestPutAll(t *testing.T){
	h := New()
	size := h.Size()
	if size != 0{
		t.Errorf("Excepted 0, got %d", size)
	}
	h.PutAll("#1", "1", "111", "11")
	size = h.Size()
	if size != 1{
		t.Errorf("Excepted 1, got %d", size)
	}

	h.PutAll("#2", "2", "222", "22")
	size = h.Size()
	if size != 2{
		t.Errorf("Excepted 2, got %d", size)
	}
	h.PutAll("#2", "2222")
	size = h.Size()
	if size != 2{
		t.Errorf("Excepted 2, got %d", size)
	}
}

func TestGetKey(t *testing.T){
	h:= New()
	emptyKeys := h.GetKeys()
	if len(emptyKeys) != 0{
		t.Errorf("Excepted 0 keys, got %d", len(emptyKeys))
	}
	h.Put("Holamundo", "Hello")
	oneElementKeys := h.GetKeys()
	if len(oneElementKeys) != 1{
		t.Errorf("Excepted 1 keys, got %d", len(oneElementKeys))
	}
}

func TestContains(t *testing.T){
	h:= New()
	NoContains := h.Contains("Hello")
	if NoContains{
		t.Errorf("Excepted false contains, got %v", NoContains)
	}
	h.Put("Hello", "HelloWorld")
	YesContains := h.Contains("Hello")
	if YesContains != true{
		t.Errorf("Excepted true contains, got %v", YesContains)
	}
}

func TestGetValues(t *testing.T){
	h := New()
	values := h.GetValues("hello")

	if len(values) != 0{
		t.Errorf("Excepted 0 values, got %d", len(values))
	}

	h.Put("hello", "sayHello")
	valuesOne := h.GetValues("hello")
	if len(valuesOne) != 1{
		t.Errorf("Excepted 1 value, got %d", len(valuesOne))
	}
}

func TestClear(t *testing.T){
	h := New()

	h.Put("hello", "sayHello")
	h.PutAll("gophers", "zilo", "huber", "humberto")
	h.Put("language", "Spanish")
	if h.Size() != 3{
		t.Errorf("Excepted size 3, got %d", h.Size())
	}
	h.Clear()
	if h.Size() != 0{
		t.Errorf("Excepted size 0, got %d", h.Size())
	}
}

func TestRemoveKeys(t *testing.T){
	h := New()

	h.Put("hello", "sayHello")
	h.PutAll("gophers", "zilo", "huber", "humberto")
	h.Put("language", "Spanish")

	if h.Size() != 3{
		t.Errorf("Excepted size 3, got %d", h.Size())
	}
	h.RemoveKeys("hello")
	if h.Size() != 2{
		t.Errorf("Excepted size 2, got %d", h.Size())
	}
	h.RemoveKeys("language", "gophers")
	if h.Size() != 0{
		t.Errorf("Excepted size 0, got %d", h.Size())
	}
	h.Put("booleans", "true")
	h.RemoveKeys("bool")
	if h.Size() != 1{
		t.Errorf("Excepted size 1, got %d", h.Size())
	}
}

func TestRemove(t *testing.T) {
	h := New()

	h.Put("hello", "sayHello")
	h.PutAll("gophers", "zilo", "huber", "humberto")
	h.Put("language", "Spanish")
	h.Remove("motors", "Spanish")
	h.Remove("gophers", "Luis")
	h.Remove("gophers", "huber")
	if h.Size() != 3{
		t.Errorf("Excepted size 3, got %d", h.Size())
	}
	valuesGophers := h.GetValues("gophers")
	if len(valuesGophers) != 2{
		t.Errorf("Excepted size 2, got %d", valuesGophers)
	}

	h.Remove("hello", "sayHello")
	valuesHello := h.GetValues("hello")
	if len(valuesHello) != 0{
		t.Errorf("Excepted size 0, got %d", len(valuesHello))
	}
}

func TestMerge(t *testing.T){
	mapOne := New()
	mapTwo := New()

	mapOne.Put("languages", "Spanish")
	mapOne.Put("languages", "Russian")
	mapOne.Put("languages", "Italian")
	mapOne.Put("aws", "s3")

	mapTwo.Put("languages", "English")
	mapTwo.Put("databases", "mongo")
	mapTwo.Put("fruits", "Orange")
	keysOne := mapOne.GetKeys()

	if len(keysOne) != 2{
		t.Errorf("Expected keyOne 2 values, got %d", len(keysOne))
	}

	keysTwo := mapTwo.GetKeys()
	if len(keysTwo) != 3{
		t.Errorf("Expected keyTwo 3 values, got %d", len(keysTwo))
	}

	valuesOne := mapOne.GetValues("languages")
	if len(valuesOne) != 3{
		t.Errorf("Expected 3 values, got %d", valuesOne)
	}

	valuesTwo := mapTwo.GetValues("languages")

	if len(valuesTwo) != 1{
		t.Errorf("Excepted 1 value, got %d", len(valuesTwo))
	}

	mapOne.Merge(mapTwo)

	keysOne = mapOne.GetKeys()
	if len(keysOne) != 4{
		t.Errorf("Expected keyOne 4 values, got %d", len(keysOne))
	}

	valuesOne = mapOne.GetValues("languages")
	if len(valuesOne) != 4 {
		t.Errorf("Expected 4 values, got %d", len(valuesOne))
	}
}