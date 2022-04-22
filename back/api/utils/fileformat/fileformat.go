package fileformat

import (
	"path"
	"strings"

	"github.com/twinj/uuid"
)

// UniqueFormat : need to make sure that the image added in the user profile has a unique name.
func UniqueFormat(fn string) string {
	//path.Ext() get the extension of the file
	fileName := strings.TrimSuffix(fn, path.Ext(fn))
	extension := path.Ext(fn)
	u := uuid.NewV4()
	newFileName := fileName + "-" + u.String() + extension

	return newFileName

}
