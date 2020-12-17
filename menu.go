package systray

import "sync/atomic"

// newMenuItemChan returns a populated MenuItem object
func newMenuItemChan(title string, tooltip string, parent *MenuItem, mchan chan *MenuItem) *MenuItem {
	if mchan == nil {
		mchan = make(chan *MenuItem)
	}
	return &MenuItem{
		ClickedCh:   mchan,
		id:          atomic.AddUint32(&currentID, 1),
		title:       title,
		tooltip:     tooltip,
		disabled:    false,
		checked:     false,
		isCheckable: false,
		parent:      parent,
	}
}

func AddMenuItemChan(title string, tooltip string, mchan chan *MenuItem) *MenuItem {
	item := newMenuItemChan(title, tooltip, nil, mchan)
	item.update()
	return item
}

func (item *MenuItem) AddSubMenuItemChan(title string, tooltip string, mchan chan *MenuItem) *MenuItem {
	child := newMenuItemChan(title, tooltip, item, mchan)
	child.update()
	return child
}

func (item *MenuItem) Name() (string, string) {
	return item.title, item.tooltip
}
