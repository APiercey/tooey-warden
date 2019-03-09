import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
)

func getItems() map[string]string {
	dict := make(map[string]string)
	dict["Gitlab"] = "asd"
	dict["Github"] = "ghj"
	dict["bitbucket"] = "uuu"

	return dict
}

func readBitWardenItems() []byte {
	cmd := exec.Command("bw", "list", "items")
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		os.Exit(1)
	}

	return out
}

// BwLogin Bitwarden login information
type BwLogin struct {
	Username string
	Password string
}

// BwItem struct of a Bitwarden item
type BwItem struct {
	ID    string
	Name  string
	Login BwLogin
}

// BwItemCollection a collection of BwItems
type BwItemCollection struct {
	Collection []BwItem
}

func getBitWardenItems() []BwItem {
	data := readBitWardenItems()

	bwItems := make([]BwItem, 0)
	json.Unmarshal(data, &bwItems)

	return bwItems
}
