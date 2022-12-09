package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"

	"golang.org/x/mod/modfile"
)

func GetCountInDir(path string) (int, error) {
	files, err := ioutil.ReadDir(path)
	return len(files), err
}

func GetGoModName(workdir string) (string, error) {

	isExist, err := CheckPathIfExist(workdir + "/go.mod")
	if !isExist || err != nil {
		return "", errors.New("go.mod files not found in current directory")
	}
	res, err := ReadFile(workdir + "/go.mod")
	if err != nil {
		return "", err
	}
	gomod := modfile.ModulePath([]byte(res))
	return gomod, nil
}

func RunCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ReadFile(path string) (string, error) {

	isExist, _ := CheckPathIfExist(path)
	if !isExist {
		return "", errors.New("path does'nt exist")
	}

	input, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.New("failed to read file")
	}

	return string(input), nil
}

func ReplaceTextInFile(path string, oldString string, newString string) {

	input, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte(oldString), []byte(newString), -1)

	if err = ioutil.WriteFile(path, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func WriteFile(path string, content string) error {
	if err := ioutil.WriteFile(path, []byte(content), 0666); err != nil {
		return err
	}
	return nil
}

func Mkdir(path string) error {
	isExist, err := CheckPathIfExist(path)
	if isExist {
		return errors.New("directory already exist")
	}
	err = os.Mkdir(path, 0777)
	if err != nil {
		return err
	}
	return nil
}

func GetOS() string {
	return runtime.GOOS
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	err = os.Remove(dir)
	if err != nil {
		return err
	}
	return nil
}

func DeleleLineContainSubstring(file string, substring string) string {
	re := regexp.MustCompile("(?m)[\r\n]+^.*" + substring + ".*$")
	res := re.ReplaceAllString(file, "")
	return res
}
