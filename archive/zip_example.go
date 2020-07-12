package archive

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
)
type Files struct{
	Name string
	Body string
}
var files = []Files{
	{"readme.txt", "This archive contains some text files."},
	{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
	{"todo.txt", "Get animal handling licence.\nWrite more examples."},
}

func ZipIt(filename string) error{
	buf, err := os.Create(filename)
	if err != nil{
		return err
	}
	w := zip.NewWriter(buf)

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			return err
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			return err
		}
	}
	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		return err
	}
	return nil
}

func ReadZip(filename string) ([]Files, error){
	r, err := zip.OpenReader(filename)
	if err != nil{
		return nil, nil
	}
	defer r.Close()
	filesContent := make([]Files, 0)
	for _, f := range r.File{
		rc, err := f.Open()
		if err != nil{
			return nil, err
		}
		var content []byte
		buf := bytes.NewBuffer(content)
		_, err = io.Copy(buf, rc)
		if err != nil {
			return nil, err
		}
		filesContent = append(filesContent, Files{f.Name, string(buf.Bytes())})
		err = rc.Close()
		if err != nil{
			return nil, err
		}
	}
	return filesContent, nil
}
