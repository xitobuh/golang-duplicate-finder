package pkg

import "os"

func CheckPath(path string) bool {
	pathCheck, err := os.Stat(path)
	if err != nil {
		return false
	}
	return pathCheck.Mode().IsRegular()
}
