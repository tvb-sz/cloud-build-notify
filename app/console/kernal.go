package console

import "fmt"

func Help() {
	fmt.Println("Usage of notify:")
	fmt.Println("\tnotify dingTalk --token=*** --secret=*** --title=*** --content='markDown text'")
	fmt.Println("\tnotify feiShu --token=*** --secret=*** --title=*** --content='markDown text'")
}
