package main

import (
	"fmt"

	"github.com/gentee/systray"
	"github.com/gentee/systray/example/icon"
)

type TreeInfo struct {
	Name     string
	Title    string
	Children []TreeInfo
}

var (
	list = []TreeInfo{
		{Name: `welcome`, Title: `Welcome Title`},
		{Name: `russian`, Title: `Русский`},
		{Name: ``, Title: `Папка`, Children: []TreeInfo{
			{Name: `test-first`, Title: `Test First`},
			{Name: `item-2`, Title: `This is an item 2`},
			{Name: `item-3`, Title: `Item 3`},
		}},
		{Name: ``, Title: `Folder №2`, Children: []TreeInfo{
			{Name: `item-4`, Title: `Item 4`},
			{Name: `item-5`, Title: `Item 5`},
			{Name: `item-6`, Title: `Тестовое сообщение`},
		}},
		{Name: `english`, Title: `English`},
	}
	mChan = make(chan *systray.MenuItem)
	menu  = make([]*systray.MenuItem, 16)
)

func onExit() {
	fmt.Println(`exit`)
}

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Gentee systray")

	for _, item := range list {
		m := systray.AddMenuItemChan(item.Title, item.Name, mChan)
		for _, sub := range item.Children {
			m.AddSubMenuItemChan(sub.Title, sub.Name, mChan)
		}
	}
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		var m *systray.MenuItem
		for {
			select {
			case m = <-mChan:
				_, name := m.Name()
				if len(name) > 0 {
					fmt.Println(`Click`, name)
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}
