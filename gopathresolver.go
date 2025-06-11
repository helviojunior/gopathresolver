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
	"os"
	"path/filepath"
	"errors"
	"os/user"
    "runtime"
    "strings"
)

var PathSeparator = string(os.PathSeparator)

func ResolveFullPath(file_path string) (string, error) {

	if file_path == "" {
		return "", errors.New("File path cannot be empty!") 
	}

	if runtime.GOOS == "windows" {
		file_path = strings.Replace(file_path, "/", PathSeparator, -1)
	} else {
		file_path = strings.Replace(file_path, "\\", PathSeparator, -1)
	}

	fc := file_path[0:1]
	if fc == "~" {
		usr, err := user.Current()
	    if err != nil {
	       return "", err
	    }

		file_path, err = filepath.Abs(strings.Replace(file_path, "~", usr.HomeDir, 1))
		if err != nil {
	       return "", err
	    }
		if !IsValid(file_path) {
			return "", errors.New("File path '"+ file_path + "' is not a valid path") 
		}

		return file_path, nil
	}

	currentPath, err := os.Getwd()
    if err != nil {
       return "", err
    }

	if fc == "." {
		doAction := false
		if len(file_path) < 3 {
			doAction = true
		}else if file_path[0:3] != "../" && file_path[0:3] != "..\\" {
			doAction = true
		}
		if doAction {
			file_path, err = filepath.Abs(strings.Replace(file_path, ".", currentPath, 1))
			if err != nil {
		       return "", err
		    }
			if !IsValid(file_path) {
				return "", errors.New("File path '"+ file_path + "' is not a valid path") 
			}

			return file_path, nil
		}
	}

	if runtime.GOOS == "windows" {
	    pc := file_path[1:3]
	    if pc == ":\\" {
	    	file_path, err = filepath.Abs(file_path)
	    	if err != nil {
		       return "", err
		    }
			if !IsValid(file_path) {
				return "", errors.New("File path '"+ file_path + "' is not a valid path") 
			}

			return file_path, nil
	    }

	} else {
		if fc == "/" {
			file_path, err = filepath.Abs(file_path)
			if err != nil {
		       return "", err
		    }
			if !IsValid(file_path) {
				return "", errors.New("File path '"+ file_path + "' is not a valid path") 
			}

			return file_path, nil
		}
	}

	file_name := filepath.Base(file_path)

	if file_name != file_path && strings.Contains(file_path, PathSeparator) {
    	file_path, err = filepath.Abs(filepath.Join(currentPath, file_path))
    	if err != nil {
	       return "", err
	    }
		if !IsValid(file_path) {
			return "", errors.New("File path '"+ file_path + "' is not a valid path") 
		}

		return file_path, nil
    }

	
	if file_name == file_path {
		file_path, err = filepath.Abs(filepath.Join(currentPath, file_path))
		if err != nil {
	       return "", err
	    }
		if !IsValid(file_path) {
			return "", errors.New("File path '"+ file_path + "' is not a valid path") 
		}
	}

	return file_path, nil
}

func ResolveRelativePath(base_path string, full_path string) (string, error) {
	var replacer string
	replacer = "/"
	if runtime.GOOS == "windows" {
		replacer = "\\"
	}

	p1, err := ResolveFullPath(base_path)
    if err != nil {
        return full_path, err
    }

    p2, err := ResolveFullPath(full_path)
    if err != nil {
        return full_path, err
    }

    new_path := strings.Replace(p2, p1, "", 1)
    if len(new_path) > 0 && new_path[0:1] == replacer {
    	new_path = new_path[1:]
    }
    new_path = "." + replacer + new_path
    if !IsValid(new_path) {
		return full_path, errors.New("File path '"+ new_path + "' is not a valid path") 
	}

    return new_path, nil
}

func IsValid(fp string) bool {
  // Check if file already exists
  if _, err := os.Stat(fp); err == nil {
    return true
  }

  // Attempt to create it
  var d []byte
  if err := os.WriteFile(fp, d, 0644); err == nil {
    os.Remove(fp) // And delete it
    return true
  }

  return false
}

func IsValidAndNotExists(fp string) (bool, error) {

	_, err := os.Stat(fp)
	if err == nil { // File exists
		return false, errors.New("File path '"+ fp +"' already exists")
	}else if !errors.Is(err, os.ErrNotExist) {
		return false, err
	}

	// Attempt to create it
	var d []byte
	err = os.WriteFile(fp, d, 0644)
	if err != nil {
		return false, err	
	}
	
	os.Remove(fp) // And delete it
	return true, nil
}
