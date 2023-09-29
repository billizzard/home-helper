package menuPage

type MenuPage struct {
	MenuItems []MenuItem
}

type MenuItem struct {
	Name string
	Icon string
	Link string
}

func NewMenuPage() *MenuPage {
	model := MenuPage{}
	model.addItem("Files", "folder", "/files/root")

	return &model
}

func (m *MenuPage) addItem(name string, icon string, link string) {
	m.MenuItems = append(m.MenuItems, MenuItem{
		Name: name,
		Icon: icon,
		Link: link,
	})
}
