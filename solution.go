package solution

import "github.com/kyokomi/emoji"


func GetMessage() string {
	hello_world := emoji.Sprint("Hello :world_map:!")
	return hello_world
}
