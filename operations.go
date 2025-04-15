package main

import (
	"os"
	"os/user"
	"os/exec"
	"syscall"
	"fmt"
)

func (m *model) populateFromInfo(info os.FileInfo) {
	stat := info.Sys().(*syscall.Stat_t)

	m.filename = info.Name()
	m.size = info.Size()
	m.lastModified = info.ModTime().String()
	m.permission = info.Mode().String()

	u, err := user.LookupId(fmt.Sprintf("%d", stat.Uid))
	if err != nil {
		m.error = err.Error()
		return
	}
	m.owner = u.Username
}

func (m *model) getDescription(path string) {
	output, err := exec.Command("file", path).Output()
	if err != nil {
		m.error = err.Error()
		return
	}
	m.description = string(output)
}

func (m *model) String() string {
	return fmt.Sprintf(`
Path:          %s
Owner:         %s
Size:          %d bytes
Permissions:   %s
Created:       %s
Last Modified: %s
File Type:     %s
`,
		m.path,
		m.owner,
		m.size,
		m.permission,
		m.created,
		m.lastModified,
		m.description,
	)
}
