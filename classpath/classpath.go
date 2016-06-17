package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExt(jreOption)
	cp.parseUser(cpOption)
	return cp
}

func (c *Classpath) parseBootAndExt(jreOption string) {
	jreDir := jre(jreOption)
	c.bootClasspath = newEntry(filepath.Join(jreDir, "lib", "*"))
	c.extClasspath = newEntry(filepath.Join(jreDir, "lib", "ext", "*"))
}

func jre(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return ".jre"
	}
	if javaHome := os.Getenv("JAVA_HOME"); javaHome != "" {
		return filepath.Join(javaHome, "jre")
	}
	panic("Can not find jre")
}

func (c *Classpath) parseUser(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	c.userClasspath = newEntry(cpOption)
}

func (c *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	classFile := className + ".class"
	if data, entry, err := c.bootClasspath.readClass(classFile); err == nil {
		return data, entry, err
	}
	if data, entry, err := c.extClasspath.readClass(classFile); err == nil {
		return data, entry, err
	}
	return c.userClasspath.readClass(classFile)
}

func (c *Classpath) String() string {
	return c.userClasspath.String()
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
