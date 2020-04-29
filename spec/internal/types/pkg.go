package types

import "fmt"

// pkgMap maps a package path to a package.
var pkgMap = make(map[string]*Pkg)

// NewPkg returns a new Pkg for the given package path and name.
// Unless name is the empty string, if the package exists already,
// the existing package name and the provided name must match.
func NewPkg(path, name string) *Pkg {
	if p := pkgMap[path]; p != nil {
		if name != "" && p.Name != name {
			panic(fmt.Sprintf("conflicting package names %s and %s for path %q", p.Name, name, path))
		}
		return p
	}

	p := &Pkg{
		Path: path,
		Name: name,
		TypeMap: map[string]*Type{},
	}

	pkgMap[path] = p
	return p
}

type Pkg struct {
	Path    string // string literal used in import statement, e.g. "runtime/internal/sys"
	Name    string // package name, e.g. "sys"
	TypeMap map[string]*Type
}

func (s *Pkg) AddType(name string, t *Type) {
	if p := s.TypeMap[name]; p != nil {
		panic(fmt.Sprintf("pkg: conflicting type %s already added to package %s", name, s.Name))
	}
	s.TypeMap[name] = t
}