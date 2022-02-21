package solution
import 
"github.com/kyokomi/emoji/v2"
func GetMessage() string {
	myMessage := emoji.Sprint("Hello, :world_map:")
	return myMessage
}
