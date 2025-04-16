package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"time"

	"golang.org/x/sys/unix"
)

func (m *model) populateFromInfo(info os.FileInfo) {
	var stat unix.Stat_t
	if err := unix.Stat(m.path, &stat); err != nil {
		m.error = err.Error()
		return
	}
	m.filename = info.Name()
	m.size = info.Size()
	birth := time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	m.created = birth.Format("2006-01-02 15:04")
	m.lastModified = info.ModTime().Format("2006-01-02 15:04")
	m.permission = info.Mode().String()
	u, err := user.LookupId(fmt.Sprintf("%d", stat.Uid))
	if err != nil {
		m.error = err.Error()
		return
	}
	m.owner = u.Username
}

func (m *model) getDescription(path string) {
	out, err := exec.Command("file", "-b", "--separator=,", path).Output()
	if err != nil {
		m.error = err.Error()
		return
	}

	desc := strings.TrimSpace(string(out))
	parts := strings.Split(desc, ",")

	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	m.description = strings.Join(parts, "\n")
}
