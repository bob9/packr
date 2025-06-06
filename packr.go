package packr

import (
	"fmt"

	"github.com/gobuffalo/packr/v2/file/resolver"
	"github.com/gobuffalo/packr/v2/jam/parser"
	"github.com/gobuffalo/packr/v2/plog"
	"github.com/markbates/safe"
)

var boxes = &boxMap{}

var _ = safe.Run(func() error {
	p, err := parser.NewFromRoots([]string{}, nil)
	if err != nil {
		plog.Logger.Error(err)
		return err
	}
	boxes, err := p.Run()
	if err != nil {
		plog.Logger.Error(err)
		return err
	}
	for _, box := range boxes {
		b := construct(box.Name, box.AbsPath)
		_, err = placeBox(b)
		if err != nil {
			plog.Logger.Error(err)
			return err
		}
	}
	return nil
})

func findBox(name string) (*Box, error) {
	key := resolver.Key(name)
	plog.Debug("packr", "findBox", "name", name, "key", key)

	b, ok := boxes.Load(key)
	if !ok {
		plog.Debug("packr", "findBox", "name", name, "key", key, "found", ok)
		return nil, fmt.Errorf("could not find box %s", name)
	}

	plog.Debug(b, "found", "box", b)
	return b, nil
}

func placeBox(b *Box) (*Box, error) {
	key := resolver.Key(b.Name)
	eb, _ := boxes.LoadOrStore(key, b)

	plog.Debug("packr", "placeBox", "name", eb.Name, "path", eb.Path, "resolution directory", eb.ResolutionDir)
	return eb, nil
}
