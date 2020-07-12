package archive

import (
	"os"
	"testing"
)

var filename = "test_zip.zip"
func TestMain(m *testing.M) {
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func shutdown(){
	_ = os.Remove(filename)
}

func TestZipIt(t *testing.T){
	if err := ZipIt(filename); err!=nil{
		t.Errorf("exceptec zip without error. got %v", err.Error())
	}
	_, err := os.Stat(filename)
	if err != nil{
		t.Errorf("Excepted filename %s create. got %v", filename, err.Error())
	}
}

func TestReadZip(t *testing.T) {
	filesTest, err := ReadZip(filename)
	if err != nil{
		t.Errorf("Excepted reader without error. got %v", err.Error())
	}
	for index:=0; index<len(filesTest); index++{
		if filesTest[index].Name != files[index].Name{
			t.Errorf("Excepted name %s, got %s in index %d", filesTest[index].Name, files[index].Name, index)
		}
		if filesTest[index].Body != files[index].Body{
			t.Errorf("Excepted body %s, got %s in index %d", filesTest[index].Body, files[index].Body, index)
		}
	}
}
