/*
 * GOPATHRESOLVER - Resolve file path in Golang
 * Copyright (c) 2025 Helvio Junior (M4v3r1ck) <helvio_junior [at] hotmail [dot] com>
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package gopathresolver

import (
	"fmt"
	"os"
	"testing"
    "runtime"
    "path/filepath"
)

func TestFileNameOnly(t *testing.T) {
	file := "teste.md"
	if _, err := ResolveFullPath(file); err != nil {
		t.Error(err)
	}
}

func TestRelative1(t *testing.T) {

	//Creating test path
	currentPath, err := os.Getwd()
    if err != nil {
       t.Error(err)
    }
	dir := filepath.Join(currentPath, "/mypath_rel1/")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Error(err)
	}

	//Testing
	file := "mypath_rel1/teste.md"
	if _, err := ResolveFullPath(file); err != nil {
		t.Error(err)
	}

	removeFolder(dir)
}

func TestRelative2(t *testing.T) {
	//Creating test path
	currentPath, err := os.Getwd()
    if err != nil {
       t.Error(err)
    }
	dir := filepath.Join(currentPath, "/mypath_rel2/")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Error(err)
	}

	file := "mypath_rel2\\teste.md"
	if _, err := ResolveFullPath(file); err != nil {
		t.Error(err)
	}

	removeFolder(dir)
}

func TestRelative3(t *testing.T) {
	//Creating test path
	currentPath, err := os.Getwd()
    if err != nil {
       t.Error(err)
    }
	dir := filepath.Join(currentPath, "/mypath_rel3/")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Error(err)
	}

	file := "./mypath_rel3/teste.md"
	if _, err := ResolveFullPath(file); err != nil {
		t.Error(err)
	}

	removeFolder(dir)
}

func TestRelative4(t *testing.T) {
	//Creating test path
	currentPath, err := os.Getwd()
    if err != nil {
       t.Error(err)
    }
	dir := filepath.Join(currentPath, "/mypath_rel4/")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Error(err)
	}

	file := ".\\mypath_rel4\\teste.md"
	if _, err := ResolveFullPath(file); err != nil {
		t.Error(err)
	}

	removeFolder(dir)
}


func TestHome1(t *testing.T) {

	//Creating test path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Error(err)
	}
	dir := filepath.Join(homeDir, "/home1/")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Error(err)
	}

	//Testing
	file := "~/home1/teste.md"
	if _, err := ResolveFullPath(file); err != nil {
		t.Error(err)
	}

	removeFolder(dir)
}

func TestHome2(t *testing.T) {
	//Creating test path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Error(err)
	}
	dir := filepath.Join(homeDir, "/home2/")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Error(err)
	}

	file := "~\\home2\\teste.md"
	if _, err := ResolveFullPath(file); err != nil {
		t.Error(err)
	}

	removeFolder(dir)
}

func TestParent(t *testing.T) {
	//Creating test path
	currentPath, err := os.Getwd()
    if err != nil {
       t.Error(err)
    }
	dir, err := filepath.Abs(filepath.Join(currentPath, "../mypath_parent/"))
	if err != nil {
       t.Error(err)
    }
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Error(err)
	}

	//Testing
	file := "../mypath_parent/teste.md"
	if _, err := ResolveFullPath(file); err != nil {
		t.Error(err)
	}

	removeFolder(dir)
}

func TestPathTransversal(t *testing.T) {
	//Creating test path
	currentPath, err := os.Getwd()
    if err != nil {
       t.Error(err)
    }
    d1 := ""
    if runtime.GOOS == "windows" {
		d1 = "..\\..\\..\\..\\..\\..\\..\\..\\..\\..\\..\\..\\windows\\temp\\mypath_transversal"
	}else{
		d1 = "../../../../../../../../../../../../../../tmp/mypath_transversal"
	}
	dir, err := filepath.Abs(filepath.Join(currentPath, d1))
	if err != nil {
       t.Error(err)
    }
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Error(err)
	}

	//Testing
	file := ""
	if runtime.GOOS == "windows" {
		file = "../../../../../../../../../../../../../..\\windows\\temp\\mypath_transversal/teste.md"
	}else{
		file = "../../../../../../../../../../../../../../tmp/mypath_transversal/teste.md"
	}
	if _, err := ResolveFullPath(file); err != nil {
		t.Error(err)
	}

	removeFolder(dir)
}


func TestAbsolute(t *testing.T) {
	file := ""
	if runtime.GOOS == "windows" {
		file = "c:\\windows\\temp\\teste.md"
	}else{
		file = "/tmp/teste.md"
	}
	if _, err := ResolveFullPath(file); err != nil {
		t.Error(err)
	}
}

func TestAbsolutePathTransversal(t *testing.T) {
	file := ""
	if runtime.GOOS == "windows" {
		file = "c:\\..\\..\\windows\\temp\\teste.md"
	}else{
		file = "/tmp/../../../tmp/teste.md"
	}
	if _, err := ResolveFullPath(file); err != nil {
		t.Error(err)
	}
}


func TestRelativeResolver(t *testing.T) {
	base := ""
	file := ""
	if runtime.GOOS == "windows" {
		base = "c:\\..\\..\\windows\\"
		file = "c:\\..\\..\\windows\\temp\\teste.md"
	}else{
		base = "/tmp/"
		file = "/tmp/../../../tmp/teste.md"
	}
	if out, err := ResolveRelativePath(base, file); err != nil {
		t.Error(err)
	}else{
		fmt.Printf("    Relative path: %s\n", out)
	}
}


func removeFolder(path string) {

	fi, err := os.Stat(path)

    if err != nil {
    	return
    }

    if fi.Mode().IsDir() {
    	err = os.RemoveAll(path)
		if err != nil {
			return
		}

    }

}

