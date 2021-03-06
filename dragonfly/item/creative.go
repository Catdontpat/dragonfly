package item

import (
	"encoding/base64"
	"encoding/json"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"strings"
	_ "unsafe" // Imported for compiler directives.
)

// CreativeItems returns a list with all items that have been registered as a creative item. These items will
// be accessible by players in-game who have creative mode enabled.
func CreativeItems() []Stack {
	return creativeItemStacks
}

// RegisterCreativeItem registers an item as a creative item, exposing it in the creative inventory.
func RegisterCreativeItem(item Stack) {
	creativeItemStacks = append(creativeItemStacks, item)
}

// creativeItemStacks holds a list of all item stacks that were registered to the creative inventory using
// RegisterCreativeItem.
var creativeItemStacks []Stack

//lint:ignore U1000 Type is used using compiler directives.
type creativeItemEntry struct {
	ID   int32
	Meta int16
	NBT  string
}

// registerVanillaCreativeItems initialises the creative items, registering all creative items that have also
// been registered as normal items and are present in vanilla.
//lint:ignore U1000 Function is used using compiler directives.
//noinspection GoUnusedFunction
func registerVanillaCreativeItems() {
	var temp map[string]interface{}

	var m []creativeItemEntry
	if err := json.Unmarshal([]byte(strings.Replace(creativeItems, "\n", "", -1)), &m); err != nil {
		panic(err)
	}
	for _, data := range m {
		it, found := world_itemByID(data.ID, data.Meta)
		if !found {
			// The item wasn't registered, so don't register it as a creative item.
			continue
		}
		_, resultingMeta := it.EncodeItem()
		if resultingMeta != data.Meta {
			// We found an item registered with that ID and a meta of 0, but we only need items with strictly
			// the same meta here.
			continue
		}
		//noinspection ALL
		if nbter, ok := it.(world.NBTer); ok {
			nbtData, _ := base64.StdEncoding.DecodeString(data.NBT)
			if err := nbt.Unmarshal(nbtData, &temp); err != nil {
				panic(err)
			}
			if len(temp) != 0 {
				it = nbter.DecodeNBT(temp).(world.Item)
			}
		}
		RegisterCreativeItem(NewStack(it, 1))
	}
}

//go:linkname world_itemByID github.com/df-mc/dragonfly/dragonfly/world.itemByID
//noinspection ALL
func world_itemByID(id int32, meta int16) (world.Item, bool)

//noinspection SpellCheckingInspection
const creativeItems = `[{"id":5,"meta":0,"nbt":"CgAA"},{"id":5,"meta":1,"nbt":"CgAA"},{"id":5,"meta":2,"nbt":"CgAA"},{"id":5,"meta":3,"nbt":"CgAA"},{"id":5,"meta":4,"nbt":"CgAA"},{"id":5,"meta":5,"nbt":"CgAA"},{"id":-242,"meta":0
,"nbt":"CgAA"},{"id":-243,"meta":0,"nbt":"CgAA"},{"id":139,"meta":0,"nbt":"CgAA"},{"id":139,"meta":1,"nbt":"CgAA"},{"id":139,"meta":2,"nbt":"CgAA"},{"id":139,"meta":3,"nbt":"CgAA"},{"id":139,"meta":4,"nbt":
"CgAA"},{"id":139,"meta":5,"nbt":"CgAA"},{"id":139,"meta":12,"nbt":"CgAA"},{"id":139,"meta":7,"nbt":"CgAA"},{"id":139,"meta":8,"nbt":"CgAA"},{"id":139,"meta":6,"nbt":"CgAA"},{"id":139,"meta":9,"nbt":"CgAA"}
,{"id":139,"meta":13,"nbt":"CgAA"},{"id":139,"meta":10,"nbt":"CgAA"},{"id":139,"meta":11,"nbt":"CgAA"},{"id":-277,"meta":0,"nbt":"CgAA"},{"id":-297,"meta":0,"nbt":"CgAA"},{"id":-278,"meta":0,"nbt":"CgAA"},{
"id":85,"meta":0,"nbt":"CgAA"},{"id":85,"meta":1,"nbt":"CgAA"},{"id":85,"meta":2,"nbt":"CgAA"},{"id":85,"meta":3,"nbt":"CgAA"},{"id":85,"meta":4,"nbt":"CgAA"},{"id":85,"meta":5,"nbt":"CgAA"},{"id":113,"meta
":0,"nbt":"CgAA"},{"id":-256,"meta":0,"nbt":"CgAA"},{"id":-257,"meta":0,"nbt":"CgAA"},{"id":107,"meta":0,"nbt":"CgAA"},{"id":183,"meta":0,"nbt":"CgAA"},{"id":184,"meta":0,"nbt":"CgAA"},{"id":185,"meta":0,"n
bt":"CgAA"},{"id":187,"meta":0,"nbt":"CgAA"},{"id":186,"meta":0,"nbt":"CgAA"},{"id":-258,"meta":0,"nbt":"CgAA"},{"id":-259,"meta":0,"nbt":"CgAA"},{"id":-180,"meta":0,"nbt":"CgAA"},{"id":67,"meta":0,"nbt":"C
gAA"},{"id":-179,"meta":0,"nbt":"CgAA"},{"id":53,"meta":0,"nbt":"CgAA"},{"id":134,"meta":0,"nbt":"CgAA"},{"id":135,"meta":0,"nbt":"CgAA"},{"id":136,"meta":0,"nbt":"CgAA"},{"id":163,"meta":0,"nbt":"CgAA"},{"
id":164,"meta":0,"nbt":"CgAA"},{"id":109,"meta":0,"nbt":"CgAA"},{"id":-175,"meta":0,"nbt":"CgAA"},{"id":128,"meta":0,"nbt":"CgAA"},{"id":-177,"meta":0,"nbt":"CgAA"},{"id":180,"meta":0,"nbt":"CgAA"},{"id":-1
76,"meta":0,"nbt":"CgAA"},{"id":-169,"meta":0,"nbt":"CgAA"},{"id":-172,"meta":0,"nbt":"CgAA"},{"id":-170,"meta":0,"nbt":"CgAA"},{"id":-173,"meta":0,"nbt":"CgAA"},{"id":-171,"meta":0,"nbt":"CgAA"},{"id":-174
,"meta":0,"nbt":"CgAA"},{"id":108,"meta":0,"nbt":"CgAA"},{"id":114,"meta":0,"nbt":"CgAA"},{"id":-184,"meta":0,"nbt":"CgAA"},{"id":-178,"meta":0,"nbt":"CgAA"},{"id":156,"meta":0,"nbt":"CgAA"},{"id":-185,"met
a":0,"nbt":"CgAA"},{"id":203,"meta":0,"nbt":"CgAA"},{"id":-2,"meta":0,"nbt":"CgAA"},{"id":-3,"meta":0,"nbt":"CgAA"},{"id":-4,"meta":0,"nbt":"CgAA"},{"id":-254,"meta":0,"nbt":"CgAA"},{"id":-255,"meta":0,"nbt
":"CgAA"},{"id":-276,"meta":0,"nbt":"CgAA"},{"id":-292,"meta":0,"nbt":"CgAA"},{"id":-275,"meta":0,"nbt":"CgAA"},{"id":324,"meta":0,"nbt":"CgAA"},{"id":427,"meta":0,"nbt":"CgAA"},{"id":428,"meta":0,"nbt":"Cg
AA"},{"id":429,"meta":0,"nbt":"CgAA"},{"id":430,"meta":0,"nbt":"CgAA"},{"id":431,"meta":0,"nbt":"CgAA"},{"id":330,"meta":0,"nbt":"CgAA"},{"id":755,"meta":0,"nbt":"CgAA"},{"id":756,"meta":0,"nbt":"CgAA"},{"i
d":96,"meta":0,"nbt":"CgAA"},{"id":-149,"meta":0,"nbt":"CgAA"},{"id":-146,"meta":0,"nbt":"CgAA"},{"id":-148,"meta":0,"nbt":"CgAA"},{"id":-145,"meta":0,"nbt":"CgAA"},{"id":-147,"meta":0,"nbt":"CgAA"},{"id":1
67,"meta":0,"nbt":"CgAA"},{"id":-246,"meta":0,"nbt":"CgAA"},{"id":-247,"meta":0,"nbt":"CgAA"},{"id":101,"meta":0,"nbt":"CgAA"},{"id":758,"meta":0,"nbt":"CgAA"},{"id":20,"meta":0,"nbt":"CgAA"},{"id":241,"met
a":0,"nbt":"CgAA"},{"id":241,"meta":8,"nbt":"CgAA"},{"id":241,"meta":7,"nbt":"CgAA"},{"id":241,"meta":15,"nbt":"CgAA"},{"id":241,"meta":12,"nbt":"CgAA"},{"id":241,"meta":14,"nbt":"CgAA"},{"id":241,"meta":1,
"nbt":"CgAA"},{"id":241,"meta":4,"nbt":"CgAA"},{"id":241,"meta":5,"nbt":"CgAA"},{"id":241,"meta":13,"nbt":"CgAA"},{"id":241,"meta":9,"nbt":"CgAA"},{"id":241,"meta":3,"nbt":"CgAA"},{"id":241,"meta":11,"nbt":
"CgAA"},{"id":241,"meta":10,"nbt":"CgAA"},{"id":241,"meta":2,"nbt":"CgAA"},{"id":241,"meta":6,"nbt":"CgAA"},{"id":102,"meta":0,"nbt":"CgAA"},{"id":160,"meta":0,"nbt":"CgAA"},{"id":160,"meta":8,"nbt":"CgAA"}
,{"id":160,"meta":7,"nbt":"CgAA"},{"id":160,"meta":15,"nbt":"CgAA"},{"id":160,"meta":12,"nbt":"CgAA"},{"id":160,"meta":14,"nbt":"CgAA"},{"id":160,"meta":1,"nbt":"CgAA"},{"id":160,"meta":4,"nbt":"CgAA"},{"id
":160,"meta":5,"nbt":"CgAA"},{"id":160,"meta":13,"nbt":"CgAA"},{"id":160,"meta":9,"nbt":"CgAA"},{"id":160,"meta":3,"nbt":"CgAA"},{"id":160,"meta":11,"nbt":"CgAA"},{"id":160,"meta":10,"nbt":"CgAA"},{"id":160
,"meta":2,"nbt":"CgAA"},{"id":160,"meta":6,"nbt":"CgAA"},{"id":65,"meta":0,"nbt":"CgAA"},{"id":-165,"meta":0,"nbt":"CgAA"},{"id":44,"meta":0,"nbt":"CgAA"},{"id":-166,"meta":2,"nbt":"CgAA"},{"id":44,"meta":3
,"nbt":"CgAA"},{"id":182,"meta":5,"nbt":"CgAA"},{"id":158,"meta":0,"nbt":"CgAA"},{"id":158,"meta":1,"nbt":"CgAA"},{"id":158,"meta":2,"nbt":"CgAA"},{"id":158,"meta":3,"nbt":"CgAA"},{"id":158,"meta":4,"nbt":"
CgAA"},{"id":158,"meta":5,"nbt":"CgAA"},{"id":44,"meta":5,"nbt":"CgAA"},{"id":-166,"meta":0,"nbt":"CgAA"},{"id":44,"meta":1,"nbt":"CgAA"},{"id":-166,"meta":3,"nbt":"CgAA"},{"id":182,"meta":6,"nbt":"CgAA"},{
"id":182,"meta":0,"nbt":"CgAA"},{"id":-166,"meta":4,"nbt":"CgAA"},{"id":-162,"meta":1,"nbt":"CgAA"},{"id":-162,"meta":6,"nbt":"CgAA"},{"id":-162,"meta":7,"nbt":"CgAA"},{"id":-162,"meta":4,"nbt":"CgAA"},{"id
":-162,"meta":5,"nbt":"CgAA"},{"id":-162,"meta":3,"nbt":"CgAA"},{"id":-162,"meta":2,"nbt":"CgAA"},{"id":44,"meta":4,"nbt":"CgAA"},{"id":44,"meta":7,"nbt":"CgAA"},{"id":182,"meta":7,"nbt":"CgAA"},{"id":-162,
"meta":0,"nbt":"CgAA"},{"id":44,"meta":6,"nbt":"CgAA"},{"id":-166,"meta":1,"nbt":"CgAA"},{"id":182,"meta":1,"nbt":"CgAA"},{"id":182,"meta":2,"nbt":"CgAA"},{"id":182,"meta":3,"nbt":"CgAA"},{"id":182,"meta":4
,"nbt":"CgAA"},{"id":-264,"meta":0,"nbt":"CgAA"},{"id":-265,"meta":0,"nbt":"CgAA"},{"id":-282,"meta":0,"nbt":"CgAA"},{"id":-293,"meta":0,"nbt":"CgAA"},{"id":-284,"meta":0,"nbt":"CgAA"},{"id":45,"meta":0,"nb
t":"CgAA"},{"id":-302,"meta":0,"nbt":"CgAA"},{"id":-303,"meta":0,"nbt":"CgAA"},{"id":-304,"meta":0,"nbt":"CgAA"},{"id":98,"meta":0,"nbt":"CgAA"},{"id":98,"meta":1,"nbt":"CgAA"},{"id":98,"meta":2,"nbt":"CgAA
"},{"id":98,"meta":3,"nbt":"CgAA"},{"id":206,"meta":0,"nbt":"CgAA"},{"id":168,"meta":2,"nbt":"CgAA"},{"id":-274,"meta":0,"nbt":"CgAA"},{"id":-280,"meta":0,"nbt":"CgAA"},{"id":-281,"meta":0,"nbt":"CgAA"},{"i
d":-279,"meta":0,"nbt":"CgAA"},{"id":4,"meta":0,"nbt":"CgAA"},{"id":48,"meta":0,"nbt":"CgAA"},{"id":-183,"meta":0,"nbt":"CgAA"},{"id":24,"meta":0,"nbt":"CgAA"},{"id":24,"meta":1,"nbt":"CgAA"},{"id":24,"meta
":2,"nbt":"CgAA"},{"id":24,"meta":3,"nbt":"CgAA"},{"id":179,"meta":0,"nbt":"CgAA"},{"id":179,"meta":1,"nbt":"CgAA"},{"id":179,"meta":2,"nbt":"CgAA"},{"id":179,"meta":3,"nbt":"CgAA"},{"id":173,"meta":0,"nbt"
:"CgAA"},{"id":-139,"meta":0,"nbt":"CgAA"},{"id":41,"meta":0,"nbt":"CgAA"},{"id":42,"meta":0,"nbt":"CgAA"},{"id":133,"meta":0,"nbt":"CgAA"},{"id":57,"meta":0,"nbt":"CgAA"},{"id":22,"meta":0,"nbt":"CgAA"},{"
id":155,"meta":0,"nbt":"CgAA"},{"id":155,"meta":2,"nbt":"CgAA"},{"id":155,"meta":1,"nbt":"CgAA"},{"id":155,"meta":3,"nbt":"CgAA"},{"id":168,"meta":0,"nbt":"CgAA"},{"id":168,"meta":1,"nbt":"CgAA"},{"id":165,
"meta":0,"nbt":"CgAA"},{"id":-220,"meta":0,"nbt":"CgAA"},{"id":-221,"meta":0,"nbt":"CgAA"},{"id":170,"meta":0,"nbt":"CgAA"},{"id":-239,"meta":0,"nbt":"CgAA"},{"id":216,"meta":0,"nbt":"CgAA"},{"id":214,"meta
":0,"nbt":"CgAA"},{"id":-227,"meta":0,"nbt":"CgAA"},{"id":112,"meta":0,"nbt":"CgAA"},{"id":215,"meta":0,"nbt":"CgAA"},{"id":-225,"meta":0,"nbt":"CgAA"},{"id":-226,"meta":0,"nbt":"CgAA"},{"id":-240,"meta":0,
"nbt":"CgAA"},{"id":-241,"meta":0,"nbt":"CgAA"},{"id":-299,"meta":0,"nbt":"CgAA"},{"id":-298,"meta":0,"nbt":"CgAA"},{"id":-300,"meta":0,"nbt":"CgAA"},{"id":-301,"meta":0,"nbt":"CgAA"},{"id":-230,"meta":0,"n
bt":"CgAA"},{"id":-232,"meta":0,"nbt":"CgAA"},{"id":-233,"meta":0,"nbt":"CgAA"},{"id":-234,"meta":0,"nbt":"CgAA"},{"id":-235,"meta":0,"nbt":"CgAA"},{"id":-236,"meta":0,"nbt":"CgAA"},{"id":-270,"meta":0,"nbt
":"CgAA"},{"id":-222,"meta":0,"nbt":"CgAA"},{"id":35,"meta":0,"nbt":"CgAA"},{"id":35,"meta":8,"nbt":"CgAA"},{"id":35,"meta":7,"nbt":"CgAA"},{"id":35,"meta":15,"nbt":"CgAA"},{"id":35,"meta":12,"nbt":"CgAA"},
{"id":35,"meta":14,"nbt":"CgAA"},{"id":35,"meta":1,"nbt":"CgAA"},{"id":35,"meta":4,"nbt":"CgAA"},{"id":35,"meta":5,"nbt":"CgAA"},{"id":35,"meta":13,"nbt":"CgAA"},{"id":35,"meta":9,"nbt":"CgAA"},{"id":35,"me
ta":3,"nbt":"CgAA"},{"id":35,"meta":11,"nbt":"CgAA"},{"id":35,"meta":10,"nbt":"CgAA"},{"id":35,"meta":2,"nbt":"CgAA"},{"id":35,"meta":6,"nbt":"CgAA"},{"id":171,"meta":0,"nbt":"CgAA"},{"id":171,"meta":8,"nbt
":"CgAA"},{"id":171,"meta":7,"nbt":"CgAA"},{"id":171,"meta":15,"nbt":"CgAA"},{"id":171,"meta":12,"nbt":"CgAA"},{"id":171,"meta":14,"nbt":"CgAA"},{"id":171,"meta":1,"nbt":"CgAA"},{"id":171,"meta":4,"nbt":"Cg
AA"},{"id":171,"meta":5,"nbt":"CgAA"},{"id":171,"meta":13,"nbt":"CgAA"},{"id":171,"meta":9,"nbt":"CgAA"},{"id":171,"meta":3,"nbt":"CgAA"},{"id":171,"meta":11,"nbt":"CgAA"},{"id":171,"meta":10,"nbt":"CgAA"},
{"id":171,"meta":2,"nbt":"CgAA"},{"id":171,"meta":6,"nbt":"CgAA"},{"id":237,"meta":0,"nbt":"CgAA"},{"id":237,"meta":8,"nbt":"CgAA"},{"id":237,"meta":7,"nbt":"CgAA"},{"id":237,"meta":15,"nbt":"CgAA"},{"id":2
37,"meta":12,"nbt":"CgAA"},{"id":237,"meta":14,"nbt":"CgAA"},{"id":237,"meta":1,"nbt":"CgAA"},{"id":237,"meta":4,"nbt":"CgAA"},{"id":237,"meta":5,"nbt":"CgAA"},{"id":237,"meta":13,"nbt":"CgAA"},{"id":237,"m
eta":9,"nbt":"CgAA"},{"id":237,"meta":3,"nbt":"CgAA"},{"id":237,"meta":11,"nbt":"CgAA"},{"id":237,"meta":10,"nbt":"CgAA"},{"id":237,"meta":2,"nbt":"CgAA"},{"id":237,"meta":6,"nbt":"CgAA"},{"id":236,"meta":0
,"nbt":"CgAA"},{"id":236,"meta":8,"nbt":"CgAA"},{"id":236,"meta":7,"nbt":"CgAA"},{"id":236,"meta":15,"nbt":"CgAA"},{"id":236,"meta":12,"nbt":"CgAA"},{"id":236,"meta":14,"nbt":"CgAA"},{"id":236,"meta":1,"nbt
":"CgAA"},{"id":236,"meta":4,"nbt":"CgAA"},{"id":236,"meta":5,"nbt":"CgAA"},{"id":236,"meta":13,"nbt":"CgAA"},{"id":236,"meta":9,"nbt":"CgAA"},{"id":236,"meta":3,"nbt":"CgAA"},{"id":236,"meta":11,"nbt":"CgA
A"},{"id":236,"meta":10,"nbt":"CgAA"},{"id":236,"meta":2,"nbt":"CgAA"},{"id":236,"meta":6,"nbt":"CgAA"},{"id":82,"meta":0,"nbt":"CgAA"},{"id":172,"meta":0,"nbt":"CgAA"},{"id":159,"meta":0,"nbt":"CgAA"},{"id
":159,"meta":8,"nbt":"CgAA"},{"id":159,"meta":7,"nbt":"CgAA"},{"id":159,"meta":15,"nbt":"CgAA"},{"id":159,"meta":12,"nbt":"CgAA"},{"id":159,"meta":14,"nbt":"CgAA"},{"id":159,"meta":1,"nbt":"CgAA"},{"id":159
,"meta":4,"nbt":"CgAA"},{"id":159,"meta":5,"nbt":"CgAA"},{"id":159,"meta":13,"nbt":"CgAA"},{"id":159,"meta":9,"nbt":"CgAA"},{"id":159,"meta":3,"nbt":"CgAA"},{"id":159,"meta":11,"nbt":"CgAA"},{"id":159,"meta
":10,"nbt":"CgAA"},{"id":159,"meta":2,"nbt":"CgAA"},{"id":159,"meta":6,"nbt":"CgAA"},{"id":220,"meta":0,"nbt":"CgAA"},{"id":228,"meta":0,"nbt":"CgAA"},{"id":227,"meta":0,"nbt":"CgAA"},{"id":235,"meta":0,"nb
t":"CgAA"},{"id":232,"meta":0,"nbt":"CgAA"},{"id":234,"meta":0,"nbt":"CgAA"},{"id":221,"meta":0,"nbt":"CgAA"},{"id":224,"meta":0,"nbt":"CgAA"},{"id":225,"meta":0,"nbt":"CgAA"},{"id":233,"meta":0,"nbt":"CgAA
"},{"id":229,"meta":0,"nbt":"CgAA"},{"id":223,"meta":0,"nbt":"CgAA"},{"id":231,"meta":0,"nbt":"CgAA"},{"id":219,"meta":0,"nbt":"CgAA"},{"id":222,"meta":0,"nbt":"CgAA"},{"id":226,"meta":0,"nbt":"CgAA"},{"id"
:201,"meta":0,"nbt":"CgAA"},{"id":201,"meta":2,"nbt":"CgAA"},{"id":3,"meta":0,"nbt":"CgAA"},{"id":3,"meta":1,"nbt":"CgAA"},{"id":2,"meta":0,"nbt":"CgAA"},{"id":198,"meta":0,"nbt":"CgAA"},{"id":243,"meta":0,
"nbt":"CgAA"},{"id":110,"meta":0,"nbt":"CgAA"},{"id":1,"meta":0,"nbt":"CgAA"},{"id":15,"meta":0,"nbt":"CgAA"},{"id":14,"meta":0,"nbt":"CgAA"},{"id":56,"meta":0,"nbt":"CgAA"},{"id":21,"meta":0,"nbt":"CgAA"},
{"id":73,"meta":0,"nbt":"CgAA"},{"id":16,"meta":0,"nbt":"CgAA"},{"id":129,"meta":0,"nbt":"CgAA"},{"id":153,"meta":0,"nbt":"CgAA"},{"id":-288,"meta":0,"nbt":"CgAA"},{"id":-271,"meta":0,"nbt":"CgAA"},{"id":13
,"meta":0,"nbt":"CgAA"},{"id":1,"meta":1,"nbt":"CgAA"},{"id":1,"meta":3,"nbt":"CgAA"},{"id":1,"meta":5,"nbt":"CgAA"},{"id":-273,"meta":0,"nbt":"CgAA"},{"id":1,"meta":2,"nbt":"CgAA"},{"id":1,"meta":4,"nbt":"
CgAA"},{"id":1,"meta":6,"nbt":"CgAA"},{"id":-291,"meta":0,"nbt":"CgAA"},{"id":12,"meta":0,"nbt":"CgAA"},{"id":12,"meta":1,"nbt":"CgAA"},{"id":81,"meta":0,"nbt":"CgAA"},{"id":17,"meta":0,"nbt":"CgAA"},{"id":
-10,"meta":0,"nbt":"CgAA"},{"id":17,"meta":1,"nbt":"CgAA"},{"id":-5,"meta":0,"nbt":"CgAA"},{"id":17,"meta":2,"nbt":"CgAA"},{"id":-6,"meta":0,"nbt":"CgAA"},{"id":17,"meta":3,"nbt":"CgAA"},{"id":-7,"meta":0,"
nbt":"CgAA"},{"id":162,"meta":0,"nbt":"CgAA"},{"id":-8,"meta":0,"nbt":"CgAA"},{"id":162,"meta":1,"nbt":"CgAA"},{"id":-9,"meta":0,"nbt":"CgAA"},{"id":-212,"meta":0,"nbt":"CgAA"},{"id":-212,"meta":8,"nbt":"Cg
AA"},{"id":-212,"meta":1,"nbt":"CgAA"},{"id":-212,"meta":9,"nbt":"CgAA"},{"id":-212,"meta":2,"nbt":"CgAA"},{"id":-212,"meta":10,"nbt":"CgAA"},{"id":-212,"meta":3,"nbt":"CgAA"},{"id":-212,"meta":11,"nbt":"Cg
AA"},{"id":-212,"meta":4,"nbt":"CgAA"},{"id":-212,"meta":12,"nbt":"CgAA"},{"id":-212,"meta":5,"nbt":"CgAA"},{"id":-212,"meta":13,"nbt":"CgAA"},{"id":18,"meta":0,"nbt":"CgAA"},{"id":18,"meta":1,"nbt":"CgAA"}
,{"id":18,"meta":2,"nbt":"CgAA"},{"id":18,"meta":3,"nbt":"CgAA"},{"id":161,"meta":0,"nbt":"CgAA"},{"id":161,"meta":1,"nbt":"CgAA"},{"id":6,"meta":0,"nbt":"CgAA"},{"id":6,"meta":1,"nbt":"CgAA"},{"id":6,"meta
":2,"nbt":"CgAA"},{"id":6,"meta":3,"nbt":"CgAA"},{"id":6,"meta":4,"nbt":"CgAA"},{"id":6,"meta":5,"nbt":"CgAA"},{"id":-218,"meta":3,"nbt":"CgAA"},{"id":295,"meta":0,"nbt":"CgAA"},{"id":361,"meta":0,"nbt":"Cg
AA"},{"id":362,"meta":0,"nbt":"CgAA"},{"id":458,"meta":0,"nbt":"CgAA"},{"id":296,"meta":0,"nbt":"CgAA"},{"id":457,"meta":0,"nbt":"CgAA"},{"id":392,"meta":0,"nbt":"CgAA"},{"id":394,"meta":0,"nbt":"CgAA"},{"i
d":391,"meta":0,"nbt":"CgAA"},{"id":396,"meta":0,"nbt":"CgAA"},{"id":260,"meta":0,"nbt":"CgAA"},{"id":322,"meta":0,"nbt":"CgAA"},{"id":466,"meta":0,"nbt":"CgAA"},{"id":103,"meta":0,"nbt":"CgAA"},{"id":360,"
meta":0,"nbt":"CgAA"},{"id":382,"meta":0,"nbt":"CgAA"},{"id":477,"meta":0,"nbt":"CgAA"},{"id":86,"meta":0,"nbt":"CgAA"},{"id":-155,"meta":0,"nbt":"CgAA"},{"id":91,"meta":0,"nbt":"CgAA"},{"id":736,"meta":0,"
nbt":"CgAA"},{"id":31,"meta":2,"nbt":"CgAA"},{"id":175,"meta":3,"nbt":"CgAA"},{"id":31,"meta":1,"nbt":"CgAA"},{"id":175,"meta":2,"nbt":"CgAA"},{"id":760,"meta":0,"nbt":"CgAA"},{"id":-131,"meta":3,"nbt":"CgA
A"},{"id":-131,"meta":1,"nbt":"CgAA"},{"id":-131,"meta":2,"nbt":"CgAA"},{"id":-131,"meta":0,"nbt":"CgAA"},{"id":-131,"meta":4,"nbt":"CgAA"},{"id":-131,"meta":11,"nbt":"CgAA"},{"id":-131,"meta":9,"nbt":"CgAA
"},{"id":-131,"meta":10,"nbt":"CgAA"},{"id":-131,"meta":8,"nbt":"CgAA"},{"id":-131,"meta":12,"nbt":"CgAA"},{"id":-133,"meta":3,"nbt":"CgAA"},{"id":-133,"meta":1,"nbt":"CgAA"},{"id":-133,"meta":2,"nbt":"CgAA
"},{"id":-133,"meta":0,"nbt":"CgAA"},{"id":-133,"meta":4,"nbt":"CgAA"},{"id":-134,"meta":3,"nbt":"CgAA"},{"id":-134,"meta":1,"nbt":"CgAA"},{"id":-134,"meta":2,"nbt":"CgAA"},{"id":-134,"meta":0,"nbt":"CgAA"}
,{"id":-134,"meta":4,"nbt":"CgAA"},{"id":335,"meta":0,"nbt":"CgAA"},{"id":-130,"meta":0,"nbt":"CgAA"},{"id":-223,"meta":0,"nbt":"CgAA"},{"id":-224,"meta":0,"nbt":"CgAA"},{"id":37,"meta":0,"nbt":"CgAA"},{"id
":38,"meta":0,"nbt":"CgAA"},{"id":38,"meta":1,"nbt":"CgAA"},{"id":38,"meta":2,"nbt":"CgAA"},{"id":38,"meta":3,"nbt":"CgAA"},{"id":38,"meta":4,"nbt":"CgAA"},{"id":38,"meta":5,"nbt":"CgAA"},{"id":38,"meta":6,
"nbt":"CgAA"},{"id":38,"meta":7,"nbt":"CgAA"},{"id":38,"meta":8,"nbt":"CgAA"},{"id":38,"meta":9,"nbt":"CgAA"},{"id":38,"meta":10,"nbt":"CgAA"},{"id":175,"meta":0,"nbt":"CgAA"},{"id":175,"meta":1,"nbt":"CgAA
"},{"id":175,"meta":4,"nbt":"CgAA"},{"id":175,"meta":5,"nbt":"CgAA"},{"id":-216,"meta":0,"nbt":"CgAA"},{"id":351,"meta":19,"nbt":"CgAA"},{"id":351,"meta":7,"nbt":"CgAA"},{"id":351,"meta":8,"nbt":"CgAA"},{"i
d":351,"meta":16,"nbt":"CgAA"},{"id":351,"meta":17,"nbt":"CgAA"},{"id":351,"meta":1,"nbt":"CgAA"},{"id":351,"meta":14,"nbt":"CgAA"},{"id":351,"meta":11,"nbt":"CgAA"},{"id":351,"meta":10,"nbt":"CgAA"},{"id":
351,"meta":2,"nbt":"CgAA"},{"id":351,"meta":6,"nbt":"CgAA"},{"id":351,"meta":12,"nbt":"CgAA"},{"id":351,"meta":18,"nbt":"CgAA"},{"id":351,"meta":5,"nbt":"CgAA"},{"id":351,"meta":13,"nbt":"CgAA"},{"id":351,"
meta":9,"nbt":"CgAA"},{"id":351,"meta":0,"nbt":"CgAA"},{"id":351,"meta":3,"nbt":"CgAA"},{"id":351,"meta":4,"nbt":"CgAA"},{"id":351,"meta":15,"nbt":"CgAA"},{"id":106,"meta":0,"nbt":"CgAA"},{"id":-231,"meta":
0,"nbt":"CgAA"},{"id":-287,"meta":0,"nbt":"CgAA"},{"id":111,"meta":0,"nbt":"CgAA"},{"id":32,"meta":0,"nbt":"CgAA"},{"id":-163,"meta":0,"nbt":"CgAA"},{"id":80,"meta":0,"nbt":"CgAA"},{"id":79,"meta":0,"nbt":"
CgAA"},{"id":174,"meta":0,"nbt":"CgAA"},{"id":-11,"meta":0,"nbt":"CgAA"},{"id":78,"meta":0,"nbt":"CgAA"},{"id":365,"meta":0,"nbt":"CgAA"},{"id":319,"meta":0,"nbt":"CgAA"},{"id":363,"meta":0,"nbt":"CgAA"},{"
id":423,"meta":0,"nbt":"CgAA"},{"id":411,"meta":0,"nbt":"CgAA"},{"id":349,"meta":0,"nbt":"CgAA"},{"id":460,"meta":0,"nbt":"CgAA"},{"id":461,"meta":0,"nbt":"CgAA"},{"id":462,"meta":0,"nbt":"CgAA"},{"id":39,"
meta":0,"nbt":"CgAA"},{"id":40,"meta":0,"nbt":"CgAA"},{"id":-228,"meta":0,"nbt":"CgAA"},{"id":-229,"meta":0,"nbt":"CgAA"},{"id":99,"meta":14,"nbt":"CgAA"},{"id":100,"meta":14,"nbt":"CgAA"},{"id":99,"meta":1
5,"nbt":"CgAA"},{"id":99,"meta":0,"nbt":"CgAA"},{"id":344,"meta":0,"nbt":"CgAA"},{"id":338,"meta":0,"nbt":"CgAA"},{"id":353,"meta":0,"nbt":"CgAA"},{"id":367,"meta":0,"nbt":"CgAA"},{"id":352,"meta":0,"nbt":"
CgAA"},{"id":30,"meta":0,"nbt":"CgAA"},{"id":375,"meta":0,"nbt":"CgAA"},{"id":52,"meta":0,"nbt":"CgAA"},{"id":97,"meta":0,"nbt":"CgAA"},{"id":97,"meta":1,"nbt":"CgAA"},{"id":97,"meta":2,"nbt":"CgAA"},{"id":
97,"meta":3,"nbt":"CgAA"},{"id":97,"meta":4,"nbt":"CgAA"},{"id":97,"meta":5,"nbt":"CgAA"},{"id":122,"meta":0,"nbt":"CgAA"},{"id":-159,"meta":0,"nbt":"CgAA"},{"id":383,"meta":10,"nbt":"CgAA"},{"id":383,"meta
":122,"nbt":"CgAA"},{"id":383,"meta":11,"nbt":"CgAA"},{"id":383,"meta":12,"nbt":"CgAA"},{"id":383,"meta":13,"nbt":"CgAA"},{"id":383,"meta":14,"nbt":"CgAA"},{"id":383,"meta":28,"nbt":"CgAA"},{"id":383,"meta"
:22,"nbt":"CgAA"},{"id":383,"meta":75,"nbt":"CgAA"},{"id":383,"meta":16,"nbt":"CgAA"},{"id":383,"meta":19,"nbt":"CgAA"},{"id":383,"meta":30,"nbt":"CgAA"},{"id":383,"meta":18,"nbt":"CgAA"},{"id":383,"meta":2
9,"nbt":"CgAA"},{"id":383,"meta":23,"nbt":"CgAA"},{"id":383,"meta":24,"nbt":"CgAA"},{"id":383,"meta":25,"nbt":"CgAA"},{"id":383,"meta":26,"nbt":"CgAA"},{"id":383,"meta":27,"nbt":"CgAA"},{"id":383,"meta":111
,"nbt":"CgAA"},{"id":383,"meta":112,"nbt":"CgAA"},{"id":383,"meta":108,"nbt":"CgAA"},{"id":383,"meta":109,"nbt":"CgAA"},{"id":383,"meta":31,"nbt":"CgAA"},{"id":383,"meta":74,"nbt":"CgAA"},{"id":383,"meta":1
13,"nbt":"CgAA"},{"id":383,"meta":121,"nbt":"CgAA"},{"id":383,"meta":33,"nbt":"CgAA"},{"id":383,"meta":38,"nbt":"CgAA"},{"id":383,"meta":39,"nbt":"CgAA"},{"id":383,"meta":34,"nbt":"CgAA"},{"id":383,"meta":4
8,"nbt":"CgAA"},{"id":383,"meta":46,"nbt":"CgAA"},{"id":383,"meta":37,"nbt":"CgAA"},{"id":383,"meta":35,"nbt":"CgAA"},{"id":383,"meta":32,"nbt":"CgAA"},{"id":383,"meta":36,"nbt":"CgAA"},{"id":383,"meta":47,
"nbt":"CgAA"},{"id":383,"meta":110,"nbt":"CgAA"},{"id":383,"meta":17,"nbt":"CgAA"},{"id":383,"meta":40,"nbt":"CgAA"},{"id":383,"meta":45,"nbt":"CgAA"},{"id":383,"meta":49,"nbt":"CgAA"},{"id":383,"meta":50,"
nbt":"CgAA"},{"id":383,"meta":55,"nbt":"CgAA"},{"id":383,"meta":42,"nbt":"CgAA"},{"id":383,"meta":125,"nbt":"CgAA"},{"id":383,"meta":124,"nbt":"CgAA"},{"id":383,"meta":123,"nbt":"CgAA"},{"id":383,"meta":126
,"nbt":"CgAA"},{"id":383,"meta":41,"nbt":"CgAA"},{"id":383,"meta":43,"nbt":"CgAA"},{"id":383,"meta":54,"nbt":"CgAA"},{"id":383,"meta":57,"nbt":"CgAA"},{"id":383,"meta":104,"nbt":"CgAA"},{"id":383,"meta":105
,"nbt":"CgAA"},{"id":383,"meta":115,"nbt":"CgAA"},{"id":383,"meta":118,"nbt":"CgAA"},{"id":383,"meta":116,"nbt":"CgAA"},{"id":383,"meta":58,"nbt":"CgAA"},{"id":383,"meta":114,"nbt":"CgAA"},{"id":383,"meta":
59,"nbt":"CgAA"},{"id":49,"meta":0,"nbt":"CgAA"},{"id":-289,"meta":0,"nbt":"CgAA"},{"id":7,"meta":0,"nbt":"CgAA"},{"id":88,"meta":0,"nbt":"CgAA"},{"id":87,"meta":0,"nbt":"CgAA"},{"id":213,"meta":0,"nbt":"Cg
AA"},{"id":372,"meta":0,"nbt":"CgAA"},{"id":121,"meta":0,"nbt":"CgAA"},{"id":200,"meta":0,"nbt":"CgAA"},{"id":240,"meta":0,"nbt":"CgAA"},{"id":432,"meta":0,"nbt":"CgAA"},{"id":433,"meta":0,"nbt":"CgAA"},{"i
d":19,"meta":0,"nbt":"CgAA"},{"id":19,"meta":1,"nbt":"CgAA"},{"id":-132,"meta":0,"nbt":"CgAA"},{"id":-132,"meta":1,"nbt":"CgAA"},{"id":-132,"meta":2,"nbt":"CgAA"},{"id":-132,"meta":3,"nbt":"CgAA"},{"id":-13
2,"meta":4,"nbt":"CgAA"},{"id":-132,"meta":8,"nbt":"CgAA"},{"id":-132,"meta":9,"nbt":"CgAA"},{"id":-132,"meta":10,"nbt":"CgAA"},{"id":-132,"meta":11,"nbt":"CgAA"},{"id":-132,"meta":12,"nbt":"CgAA"},{"id":29
8,"meta":0,"nbt":"CgAA"},{"id":302,"meta":0,"nbt":"CgAA"},{"id":306,"meta":0,"nbt":"CgAA"},{"id":314,"meta":0,"nbt":"CgAA"},{"id":310,"meta":0,"nbt":"CgAA"},{"id":748,"meta":0,"nbt":"CgAA"},{"id":299,"meta"
:0,"nbt":"CgAA"},{"id":303,"meta":0,"nbt":"CgAA"},{"id":307,"meta":0,"nbt":"CgAA"},{"id":315,"meta":0,"nbt":"CgAA"},{"id":311,"meta":0,"nbt":"CgAA"},{"id":749,"meta":0,"nbt":"CgAA"},{"id":300,"meta":0,"nbt"
:"CgAA"},{"id":304,"meta":0,"nbt":"CgAA"},{"id":308,"meta":0,"nbt":"CgAA"},{"id":316,"meta":0,"nbt":"CgAA"},{"id":312,"meta":0,"nbt":"CgAA"},{"id":750,"meta":0,"nbt":"CgAA"},{"id":301,"meta":0,"nbt":"CgAA"}
,{"id":305,"meta":0,"nbt":"CgAA"},{"id":309,"meta":0,"nbt":"CgAA"},{"id":317,"meta":0,"nbt":"CgAA"},{"id":313,"meta":0,"nbt":"CgAA"},{"id":751,"meta":0,"nbt":"CgAA"},{"id":268,"meta":0,"nbt":"CgAA"},{"id":2
72,"meta":0,"nbt":"CgAA"},{"id":267,"meta":0,"nbt":"CgAA"},{"id":283,"meta":0,"nbt":"CgAA"},{"id":276,"meta":0,"nbt":"CgAA"},{"id":743,"meta":0,"nbt":"CgAA"},{"id":271,"meta":0,"nbt":"CgAA"},{"id":275,"meta
":0,"nbt":"CgAA"},{"id":258,"meta":0,"nbt":"CgAA"},{"id":286,"meta":0,"nbt":"CgAA"},{"id":279,"meta":0,"nbt":"CgAA"},{"id":746,"meta":0,"nbt":"CgAA"},{"id":270,"meta":0,"nbt":"CgAA"},{"id":274,"meta":0,"nbt
":"CgAA"},{"id":257,"meta":0,"nbt":"CgAA"},{"id":285,"meta":0,"nbt":"CgAA"},{"id":278,"meta":0,"nbt":"CgAA"},{"id":745,"meta":0,"nbt":"CgAA"},{"id":269,"meta":0,"nbt":"CgAA"},{"id":273,"meta":0,"nbt":"CgAA"
},{"id":256,"meta":0,"nbt":"CgAA"},{"id":284,"meta":0,"nbt":"CgAA"},{"id":277,"meta":0,"nbt":"CgAA"},{"id":744,"meta":0,"nbt":"CgAA"},{"id":290,"meta":0,"nbt":"CgAA"},{"id":291,"meta":0,"nbt":"CgAA"},{"id":
292,"meta":0,"nbt":"CgAA"},{"id":294,"meta":0,"nbt":"CgAA"},{"id":293,"meta":0,"nbt":"CgAA"},{"id":747,"meta":0,"nbt":"CgAA"},{"id":261,"meta":0,"nbt":"CgAA"},{"id":471,"meta":0,"nbt":"CgAA"},{"id":262,"met
a":0,"nbt":"CgAA"},{"id":262,"meta":6,"nbt":"CgAA"},{"id":262,"meta":7,"nbt":"CgAA"},{"id":262,"meta":8,"nbt":"CgAA"},{"id":262,"meta":9,"nbt":"CgAA"},{"id":262,"meta":10,"nbt":"CgAA"},{"id":262,"meta":11,"
nbt":"CgAA"},{"id":262,"meta":12,"nbt":"CgAA"},{"id":262,"meta":13,"nbt":"CgAA"},{"id":262,"meta":14,"nbt":"CgAA"},{"id":262,"meta":15,"nbt":"CgAA"},{"id":262,"meta":16,"nbt":"CgAA"},{"id":262,"meta":17,"nb
t":"CgAA"},{"id":262,"meta":18,"nbt":"CgAA"},{"id":262,"meta":19,"nbt":"CgAA"},{"id":262,"meta":20,"nbt":"CgAA"},{"id":262,"meta":21,"nbt":"CgAA"},{"id":262,"meta":22,"nbt":"CgAA"},{"id":262,"meta":23,"nbt"
:"CgAA"},{"id":262,"meta":24,"nbt":"CgAA"},{"id":262,"meta":25,"nbt":"CgAA"},{"id":262,"meta":26,"nbt":"CgAA"},{"id":262,"meta":27,"nbt":"CgAA"},{"id":262,"meta":28,"nbt":"CgAA"},{"id":262,"meta":29,"nbt":"
CgAA"},{"id":262,"meta":30,"nbt":"CgAA"},{"id":262,"meta":31,"nbt":"CgAA"},{"id":262,"meta":32,"nbt":"CgAA"},{"id":262,"meta":33,"nbt":"CgAA"},{"id":262,"meta":34,"nbt":"CgAA"},{"id":262,"meta":35,"nbt":"Cg
AA"},{"id":262,"meta":36,"nbt":"CgAA"},{"id":262,"meta":37,"nbt":"CgAA"},{"id":262,"meta":38,"nbt":"CgAA"},{"id":262,"meta":39,"nbt":"CgAA"},{"id":262,"meta":40,"nbt":"CgAA"},{"id":262,"meta":41,"nbt":"CgAA
"},{"id":262,"meta":42,"nbt":"CgAA"},{"id":262,"meta":43,"nbt":"CgAA"},{"id":513,"meta":0,"nbt":"CgAA"},{"id":366,"meta":0,"nbt":"CgAA"},{"id":320,"meta":0,"nbt":"CgAA"},{"id":364,"meta":0,"nbt":"CgAA"},{"i
d":424,"meta":0,"nbt":"CgAA"},{"id":412,"meta":0,"nbt":"CgAA"},{"id":350,"meta":0,"nbt":"CgAA"},{"id":463,"meta":0,"nbt":"CgAA"},{"id":297,"meta":0,"nbt":"CgAA"},{"id":282,"meta":0,"nbt":"CgAA"},{"id":459,"
meta":0,"nbt":"CgAA"},{"id":413,"meta":0,"nbt":"CgAA"},{"id":393,"meta":0,"nbt":"CgAA"},{"id":357,"meta":0,"nbt":"CgAA"},{"id":400,"meta":0,"nbt":"CgAA"},{"id":354,"meta":0,"nbt":"CgAA"},{"id":464,"meta":0,
"nbt":"CgAA"},{"id":346,"meta":0,"nbt":"CgAA"},{"id":398,"meta":0,"nbt":"CgAA"},{"id":757,"meta":0,"nbt":"CgAA"},{"id":332,"meta":0,"nbt":"CgAA"},{"id":359,"meta":0,"nbt":"CgAA"},{"id":259,"meta":0,"nbt":"C
gAA"},{"id":420,"meta":0,"nbt":"CgAA"},{"id":347,"meta":0,"nbt":"CgAA"},{"id":345,"meta":0,"nbt":"CgAA"},{"id":395,"meta":0,"nbt":"CgAA"},{"id":395,"meta":2,"nbt":"CgAA"},{"id":329,"meta":0,"nbt":"CgAA"},{"
id":416,"meta":0,"nbt":"CgAA"},{"id":417,"meta":0,"nbt":"CgAA"},{"id":418,"meta":0,"nbt":"CgAA"},{"id":419,"meta":0,"nbt":"CgAA"},{"id":455,"meta":0,"nbt":"CgAA"},{"id":469,"meta":0,"nbt":"CgAA"},{"id":444,
"meta":0,"nbt":"CgAA"},{"id":450,"meta":0,"nbt":"CgAA"},{"id":374,"meta":0,"nbt":"CgAA"},{"id":384,"meta":0,"nbt":"CgAA"},{"id":373,"meta":0,"nbt":"CgAA"},{"id":373,"meta":1,"nbt":"CgAA"},{"id":373,"meta":2
,"nbt":"CgAA"},{"id":373,"meta":3,"nbt":"CgAA"},{"id":373,"meta":4,"nbt":"CgAA"},{"id":373,"meta":5,"nbt":"CgAA"},{"id":373,"meta":6,"nbt":"CgAA"},{"id":373,"meta":7,"nbt":"CgAA"},{"id":373,"meta":8,"nbt":"
CgAA"},{"id":373,"meta":9,"nbt":"CgAA"},{"id":373,"meta":10,"nbt":"CgAA"},{"id":373,"meta":11,"nbt":"CgAA"},{"id":373,"meta":12,"nbt":"CgAA"},{"id":373,"meta":13,"nbt":"CgAA"},{"id":373,"meta":14,"nbt":"CgA
A"},{"id":373,"meta":15,"nbt":"CgAA"},{"id":373,"meta":16,"nbt":"CgAA"},{"id":373,"meta":17,"nbt":"CgAA"},{"id":373,"meta":18,"nbt":"CgAA"},{"id":373,"meta":19,"nbt":"CgAA"},{"id":373,"meta":20,"nbt":"CgAA"
},{"id":373,"meta":21,"nbt":"CgAA"},{"id":373,"meta":22,"nbt":"CgAA"},{"id":373,"meta":23,"nbt":"CgAA"},{"id":373,"meta":24,"nbt":"CgAA"},{"id":373,"meta":25,"nbt":"CgAA"},{"id":373,"meta":26,"nbt":"CgAA"},
{"id":373,"meta":27,"nbt":"CgAA"},{"id":373,"meta":28,"nbt":"CgAA"},{"id":373,"meta":29,"nbt":"CgAA"},{"id":373,"meta":30,"nbt":"CgAA"},{"id":373,"meta":31,"nbt":"CgAA"},{"id":373,"meta":32,"nbt":"CgAA"},{"
id":373,"meta":33,"nbt":"CgAA"},{"id":373,"meta":34,"nbt":"CgAA"},{"id":373,"meta":35,"nbt":"CgAA"},{"id":373,"meta":36,"nbt":"CgAA"},{"id":373,"meta":37,"nbt":"CgAA"},{"id":373,"meta":38,"nbt":"CgAA"},{"id
":373,"meta":39,"nbt":"CgAA"},{"id":373,"meta":40,"nbt":"CgAA"},{"id":373,"meta":41,"nbt":"CgAA"},{"id":373,"meta":42,"nbt":"CgAA"},{"id":438,"meta":0,"nbt":"CgAA"},{"id":438,"meta":1,"nbt":"CgAA"},{"id":43
8,"meta":2,"nbt":"CgAA"},{"id":438,"meta":3,"nbt":"CgAA"},{"id":438,"meta":4,"nbt":"CgAA"},{"id":438,"meta":5,"nbt":"CgAA"},{"id":438,"meta":6,"nbt":"CgAA"},{"id":438,"meta":7,"nbt":"CgAA"},{"id":438,"meta"
:8,"nbt":"CgAA"},{"id":438,"meta":9,"nbt":"CgAA"},{"id":438,"meta":10,"nbt":"CgAA"},{"id":438,"meta":11,"nbt":"CgAA"},{"id":438,"meta":12,"nbt":"CgAA"},{"id":438,"meta":13,"nbt":"CgAA"},{"id":438,"meta":14,
"nbt":"CgAA"},{"id":438,"meta":15,"nbt":"CgAA"},{"id":438,"meta":16,"nbt":"CgAA"},{"id":438,"meta":17,"nbt":"CgAA"},{"id":438,"meta":18,"nbt":"CgAA"},{"id":438,"meta":19,"nbt":"CgAA"},{"id":438,"meta":20,"n
bt":"CgAA"},{"id":438,"meta":21,"nbt":"CgAA"},{"id":438,"meta":22,"nbt":"CgAA"},{"id":438,"meta":23,"nbt":"CgAA"},{"id":438,"meta":24,"nbt":"CgAA"},{"id":438,"meta":25,"nbt":"CgAA"},{"id":438,"meta":26,"nbt
":"CgAA"},{"id":438,"meta":27,"nbt":"CgAA"},{"id":438,"meta":28,"nbt":"CgAA"},{"id":438,"meta":29,"nbt":"CgAA"},{"id":438,"meta":30,"nbt":"CgAA"},{"id":438,"meta":31,"nbt":"CgAA"},{"id":438,"meta":32,"nbt":
"CgAA"},{"id":438,"meta":33,"nbt":"CgAA"},{"id":438,"meta":34,"nbt":"CgAA"},{"id":438,"meta":35,"nbt":"CgAA"},{"id":438,"meta":36,"nbt":"CgAA"},{"id":438,"meta":37,"nbt":"CgAA"},{"id":438,"meta":38,"nbt":"C
gAA"},{"id":438,"meta":39,"nbt":"CgAA"},{"id":438,"meta":40,"nbt":"CgAA"},{"id":438,"meta":41,"nbt":"CgAA"},{"id":438,"meta":42,"nbt":"CgAA"},{"id":441,"meta":0,"nbt":"CgAA"},{"id":441,"meta":1,"nbt":"CgAA"
},{"id":441,"meta":2,"nbt":"CgAA"},{"id":441,"meta":3,"nbt":"CgAA"},{"id":441,"meta":4,"nbt":"CgAA"},{"id":441,"meta":5,"nbt":"CgAA"},{"id":441,"meta":6,"nbt":"CgAA"},{"id":441,"meta":7,"nbt":"CgAA"},{"id":
441,"meta":8,"nbt":"CgAA"},{"id":441,"meta":9,"nbt":"CgAA"},{"id":441,"meta":10,"nbt":"CgAA"},{"id":441,"meta":11,"nbt":"CgAA"},{"id":441,"meta":12,"nbt":"CgAA"},{"id":441,"meta":13,"nbt":"CgAA"},{"id":441,
"meta":14,"nbt":"CgAA"},{"id":441,"meta":15,"nbt":"CgAA"},{"id":441,"meta":16,"nbt":"CgAA"},{"id":441,"meta":17,"nbt":"CgAA"},{"id":441,"meta":18,"nbt":"CgAA"},{"id":441,"meta":19,"nbt":"CgAA"},{"id":441,"m
eta":20,"nbt":"CgAA"},{"id":441,"meta":21,"nbt":"CgAA"},{"id":441,"meta":22,"nbt":"CgAA"},{"id":441,"meta":23,"nbt":"CgAA"},{"id":441,"meta":24,"nbt":"CgAA"},{"id":441,"meta":25,"nbt":"CgAA"},{"id":441,"met
a":26,"nbt":"CgAA"},{"id":441,"meta":27,"nbt":"CgAA"},{"id":441,"meta":28,"nbt":"CgAA"},{"id":441,"meta":29,"nbt":"CgAA"},{"id":441,"meta":30,"nbt":"CgAA"},{"id":441,"meta":31,"nbt":"CgAA"},{"id":441,"meta"
:32,"nbt":"CgAA"},{"id":441,"meta":33,"nbt":"CgAA"},{"id":441,"meta":34,"nbt":"CgAA"},{"id":441,"meta":35,"nbt":"CgAA"},{"id":441,"meta":36,"nbt":"CgAA"},{"id":441,"meta":37,"nbt":"CgAA"},{"id":441,"meta":3
8,"nbt":"CgAA"},{"id":441,"meta":39,"nbt":"CgAA"},{"id":441,"meta":40,"nbt":"CgAA"},{"id":441,"meta":41,"nbt":"CgAA"},{"id":441,"meta":42,"nbt":"CgAA"},{"id":280,"meta":0,"nbt":"CgAA"},{"id":355,"meta":0,"n
bt":"CgAA"},{"id":355,"meta":8,"nbt":"CgAA"},{"id":355,"meta":7,"nbt":"CgAA"},{"id":355,"meta":15,"nbt":"CgAA"},{"id":355,"meta":12,"nbt":"CgAA"},{"id":355,"meta":14,"nbt":"CgAA"},{"id":355,"meta":1,"nbt":"
CgAA"},{"id":355,"meta":4,"nbt":"CgAA"},{"id":355,"meta":5,"nbt":"CgAA"},{"id":355,"meta":13,"nbt":"CgAA"},{"id":355,"meta":9,"nbt":"CgAA"},{"id":355,"meta":3,"nbt":"CgAA"},{"id":355,"meta":11,"nbt":"CgAA"}
,{"id":355,"meta":10,"nbt":"CgAA"},{"id":355,"meta":2,"nbt":"CgAA"},{"id":355,"meta":6,"nbt":"CgAA"},{"id":50,"meta":0,"nbt":"CgAA"},{"id":-268,"meta":0,"nbt":"CgAA"},{"id":-156,"meta":0,"nbt":"CgAA"},{"id"
:-208,"meta":0,"nbt":"CgAA"},{"id":-269,"meta":0,"nbt":"CgAA"},{"id":58,"meta":0,"nbt":"CgAA"},{"id":-200,"meta":0,"nbt":"CgAA"},{"id":-201,"meta":0,"nbt":"CgAA"},{"id":-202,"meta":0,"nbt":"CgAA"},{"id":-21
9,"meta":3,"nbt":"CgAA"},{"id":720,"meta":0,"nbt":"CgAA"},{"id":801,"meta":0,"nbt":"CgAA"},{"id":61,"meta":0,"nbt":"CgAA"},{"id":-196,"meta":0,"nbt":"CgAA"},{"id":-198,"meta":0,"nbt":"CgAA"},{"id":-272,"met
a":0,"nbt":"CgAA"},{"id":379,"meta":0,"nbt":"CgAA"},{"id":145,"meta":0,"nbt":"CgAA"},{"id":145,"meta":4,"nbt":"CgAA"},{"id":145,"meta":8,"nbt":"CgAA"},{"id":-195,"meta":0,"nbt":"CgAA"},{"id":116,"meta":0,"n
bt":"CgAA"},{"id":47,"meta":0,"nbt":"CgAA"},{"id":-194,"meta":0,"nbt":"CgAA"},{"id":380,"meta":0,"nbt":"CgAA"},{"id":-213,"meta":0,"nbt":"CgAA"},{"id":54,"meta":0,"nbt":"CgAA"},{"id":146,"meta":0,"nbt":"CgA
A"},{"id":130,"meta":0,"nbt":"CgAA"},{"id":-203,"meta":0,"nbt":"CgAA"},{"id":205,"meta":0,"nbt":"CgAA"},{"id":218,"meta":0,"nbt":"CgAA"},{"id":218,"meta":8,"nbt":"CgAA"},{"id":218,"meta":7,"nbt":"CgAA"},{"i
d":218,"meta":15,"nbt":"CgAA"},{"id":218,"meta":12,"nbt":"CgAA"},{"id":218,"meta":14,"nbt":"CgAA"},{"id":218,"meta":1,"nbt":"CgAA"},{"id":218,"meta":4,"nbt":"CgAA"},{"id":218,"meta":5,"nbt":"CgAA"},{"id":21
8,"meta":13,"nbt":"CgAA"},{"id":218,"meta":9,"nbt":"CgAA"},{"id":218,"meta":3,"nbt":"CgAA"},{"id":218,"meta":11,"nbt":"CgAA"},{"id":218,"meta":10,"nbt":"CgAA"},{"id":218,"meta":2,"nbt":"CgAA"},{"id":218,"me
ta":6,"nbt":"CgAA"},{"id":425,"meta":0,"nbt":"CgAA"},{"id":25,"meta":0,"nbt":"CgAA"},{"id":84,"meta":0,"nbt":"CgAA"},{"id":500,"meta":0,"nbt":"CgAA"},{"id":501,"meta":0,"nbt":"CgAA"},{"id":502,"meta":0,"nbt
":"CgAA"},{"id":503,"meta":0,"nbt":"CgAA"},{"id":504,"meta":0,"nbt":"CgAA"},{"id":505,"meta":0,"nbt":"CgAA"},{"id":506,"meta":0,"nbt":"CgAA"},{"id":507,"meta":0,"nbt":"CgAA"},{"id":508,"meta":0,"nbt":"CgAA"
},{"id":509,"meta":0,"nbt":"CgAA"},{"id":510,"meta":0,"nbt":"CgAA"},{"id":511,"meta":0,"nbt":"CgAA"},{"id":759,"meta":0,"nbt":"CgAA"},{"id":348,"meta":0,"nbt":"CgAA"},{"id":89,"meta":0,"nbt":"CgAA"},{"id":1
23,"meta":0,"nbt":"CgAA"},{"id":169,"meta":0,"nbt":"CgAA"},{"id":323,"meta":0,"nbt":"CgAA"},{"id":472,"meta":0,"nbt":"CgAA"},{"id":473,"meta":0,"nbt":"CgAA"},{"id":474,"meta":0,"nbt":"CgAA"},{"id":475,"meta
":0,"nbt":"CgAA"},{"id":476,"meta":0,"nbt":"CgAA"},{"id":753,"meta":0,"nbt":"CgAA"},{"id":754,"meta":0,"nbt":"CgAA"},{"id":321,"meta":0,"nbt":"CgAA"},{"id":389,"meta":0,"nbt":"CgAA"},{"id":737,"meta":0,"nbt
":"CgAA"},{"id":390,"meta":0,"nbt":"CgAA"},{"id":281,"meta":0,"nbt":"CgAA"},{"id":325,"meta":0,"nbt":"CgAA"},{"id":325,"meta":1,"nbt":"CgAA"},{"id":325,"meta":8,"nbt":"CgAA"},{"id":325,"meta":10,"nbt":"CgAA
"},{"id":325,"meta":2,"nbt":"CgAA"},{"id":325,"meta":3,"nbt":"CgAA"},{"id":325,"meta":4,"nbt":"CgAA"},{"id":325,"meta":5,"nbt":"CgAA"},{"id":397,"meta":3,"nbt":"CgAA"},{"id":397,"meta":2,"nbt":"CgAA"},{"id"
:397,"meta":4,"nbt":"CgAA"},{"id":397,"meta":5,"nbt":"CgAA"},{"id":397,"meta":0,"nbt":"CgAA"},{"id":397,"meta":1,"nbt":"CgAA"},{"id":138,"meta":0,"nbt":"CgAA"},{"id":-206,"meta":0,"nbt":"CgAA"},{"id":-157,"
meta":0,"nbt":"CgAA"},{"id":-197,"meta":0,"nbt":"CgAA"},{"id":120,"meta":0,"nbt":"CgAA"},{"id":263,"meta":0,"nbt":"CgAA"},{"id":263,"meta":1,"nbt":"CgAA"},{"id":264,"meta":0,"nbt":"CgAA"},{"id":452,"meta":0
,"nbt":"CgAA"},{"id":265,"meta":0,"nbt":"CgAA"},{"id":752,"meta":0,"nbt":"CgAA"},{"id":742,"meta":0,"nbt":"CgAA"},{"id":371,"meta":0,"nbt":"CgAA"},{"id":266,"meta":0,"nbt":"CgAA"},{"id":388,"meta":0,"nbt":"
CgAA"},{"id":406,"meta":0,"nbt":"CgAA"},{"id":337,"meta":0,"nbt":"CgAA"},{"id":336,"meta":0,"nbt":"CgAA"},{"id":405,"meta":0,"nbt":"CgAA"},{"id":409,"meta":0,"nbt":"CgAA"},{"id":422,"meta":0,"nbt":"CgAA"},{
"id":465,"meta":0,"nbt":"CgAA"},{"id":467,"meta":0,"nbt":"CgAA"},{"id":468,"meta":0,"nbt":"CgAA"},{"id":470,"meta":0,"nbt":"CgAA"},{"id":287,"meta":0,"nbt":"CgAA"},{"id":288,"meta":0,"nbt":"CgAA"},{"id":318
,"meta":0,"nbt":"CgAA"},{"id":289,"meta":0,"nbt":"CgAA"},{"id":334,"meta":0,"nbt":"CgAA"},{"id":415,"meta":0,"nbt":"CgAA"},{"id":414,"meta":0,"nbt":"CgAA"},{"id":385,"meta":0,"nbt":"CgAA"},{"id":369,"meta":
0,"nbt":"CgAA"},{"id":377,"meta":0,"nbt":"CgAA"},{"id":378,"meta":0,"nbt":"CgAA"},{"id":376,"meta":0,"nbt":"CgAA"},{"id":437,"meta":0,"nbt":"CgAA"},{"id":445,"meta":0,"nbt":"CgAA"},{"id":370,"meta":0,"nbt":
"CgAA"},{"id":341,"meta":0,"nbt":"CgAA"},{"id":368,"meta":0,"nbt":"CgAA"},{"id":381,"meta":0,"nbt":"CgAA"},{"id":399,"meta":0,"nbt":"CgAA"},{"id":208,"meta":0,"nbt":"CgAA"},{"id":426,"meta":0,"nbt":"CgAA"},
{"id":339,"meta":0,"nbt":"CgAA"},{"id":340,"meta":0,"nbt":"CgAA"},{"id":386,"meta":0,"nbt":"CgAA"},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQAAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgIC
aWQAAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQAAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgIDbHZsBAACAmlkAAAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQBAAIDbHZ
sAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQBAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQBAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQBAAIDbHZsBAAAAA=="}
,{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQCAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQCAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgIDbHZsAwACAmlkAgAAAA=="},{"id":403,
"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQCAAIDbHZsBAAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQDAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQDAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"n
bt":"CgAJBGVuY2gKAgICaWQDAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQDAAIDbHZsBAAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQEAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBG
VuY2gKAgICaWQEAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQEAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQEAAIDbHZsBAAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICa
WQFAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQFAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQFAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQGAAIDbHZs
AQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgIDbHZsAgACAmlkBgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQGAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgIDbHZsAQACAmlkBwAAAA=="},
{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQHAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQHAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQIAAIDbHZsAQAAAA=="},{"id":403,"
meta":0,"nbt":"CgAJBGVuY2gKAgICaWQJAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQJAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQJAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nb
t":"CgAJBGVuY2gKAgICaWQJAAIDbHZsBAAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQJAAIDbHZsBQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQKAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGV
uY2gKAgIDbHZsAgACAmlkCgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQKAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQKAAIDbHZsBAAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgIDbH
ZsBQACAmlkCgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQLAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQLAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQLAAIDbHZsA
wAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQLAAIDbHZsBAAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQLAAIDbHZsBQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQMAAIDbHZsAQAAAA=="},{
"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQMAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQNAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQNAAIDbHZsAgAAAA=="},{"id":403,"m
eta":0,"nbt":"CgAJBGVuY2gKAgICaWQOAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQOAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQOAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nbt
":"CgAJBGVuY2gKAgICaWQPAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQPAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgIDbHZsAwACAmlkDwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVu
Y2gKAgICaWQPAAIDbHZsBAAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQPAAIDbHZsBQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQQAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgIDbHZ
sAQACAmlkEQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQRAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQRAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQSAAIDbHZsAQ
AAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgIDbHZsAgACAmlkEgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQSAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQTAAIDbHZsAQAAAA=="},{"
id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQTAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQTAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQTAAIDbHZsBAAAAA=="},{"id":403,"me
ta":0,"nbt":"CgAJBGVuY2gKAgICaWQTAAIDbHZsBQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQUAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgIDbHZsAgACAmlkFAAAAA=="},{"id":403,"meta":0,"nbt"
:"CgAJBGVuY2gKAgICaWQVAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQWAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQXAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY
2gKAgICaWQXAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQXAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQYAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQY
AAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQYAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQZAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQZAAIDbHZsAgA
AAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQaAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQbAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgIDbHZsAQACAmlkHAAAAA=="},{"i
d":403,"meta":0,"nbt":"CgAJBGVuY2gKAgIDbHZsAQACAmlkHQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQdAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQdAAIDbHZsAwAAAA=="},{"id":403,"met
a":0,"nbt":"CgAJBGVuY2gKAgICaWQdAAIDbHZsBAAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQdAAIDbHZsBQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQeAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":
"CgAJBGVuY2gKAgICaWQeAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQeAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQfAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2
gKAgIDbHZsAgACAmlkHwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQfAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQgAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQhA
AIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQiAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQiAAIDbHZsAgAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgIDbHZsAwACAmlkIgAA
AA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQiAAIDbHZsBAAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQjAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQjAAIDbHZsAgAAAA=="},{"id
":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQjAAIDbHZsAwAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQkAAIDbHZsAQAAAA=="},{"id":403,"meta":0,"nbt":"CgAJBGVuY2gKAgICaWQkAAIDbHZsAgAAAA=="},{"id":403,"meta
":0,"nbt":"CgAJBGVuY2gKAgICaWQkAAIDbHZsAwAAAA=="},{"id":333,"meta":0,"nbt":"CgAA"},{"id":333,"meta":1,"nbt":"CgAA"},{"id":333,"meta":2,"nbt":"CgAA"},{"id":333,"meta":3,"nbt":"CgAA"},{"id":333,"meta":4,"nbt"
:"CgAA"},{"id":333,"meta":5,"nbt":"CgAA"},{"id":66,"meta":0,"nbt":"CgAA"},{"id":27,"meta":0,"nbt":"CgAA"},{"id":28,"meta":0,"nbt":"CgAA"},{"id":126,"meta":0,"nbt":"CgAA"},{"id":328,"meta":0,"nbt":"CgAA"},{"
id":342,"meta":0,"nbt":"CgAA"},{"id":408,"meta":0,"nbt":"CgAA"},{"id":407,"meta":0,"nbt":"CgAA"},{"id":331,"meta":0,"nbt":"CgAA"},{"id":152,"meta":0,"nbt":"CgAA"},{"id":76,"meta":0,"nbt":"CgAA"},{"id":69,"m
eta":0,"nbt":"CgAA"},{"id":143,"meta":0,"nbt":"CgAA"},{"id":-144,"meta":0,"nbt":"CgAA"},{"id":-141,"meta":0,"nbt":"CgAA"},{"id":-143,"meta":0,"nbt":"CgAA"},{"id":-140,"meta":0,"nbt":"CgAA"},{"id":-142,"meta
":0,"nbt":"CgAA"},{"id":77,"meta":0,"nbt":"CgAA"},{"id":-260,"meta":0,"nbt":"CgAA"},{"id":-261,"meta":0,"nbt":"CgAA"},{"id":-296,"meta":0,"nbt":"CgAA"},{"id":131,"meta":0,"nbt":"CgAA"},{"id":72,"meta":0,"nb
t":"CgAA"},{"id":-154,"meta":0,"nbt":"CgAA"},{"id":-151,"meta":0,"nbt":"CgAA"},{"id":-153,"meta":0,"nbt":"CgAA"},{"id":-150,"meta":0,"nbt":"CgAA"},{"id":-152,"meta":0,"nbt":"CgAA"},{"id":-262,"meta":0,"nbt"
:"CgAA"},{"id":-263,"meta":0,"nbt":"CgAA"},{"id":70,"meta":0,"nbt":"CgAA"},{"id":147,"meta":0,"nbt":"CgAA"},{"id":148,"meta":0,"nbt":"CgAA"},{"id":-295,"meta":0,"nbt":"CgAA"},{"id":251,"meta":0,"nbt":"CgAA"
},{"id":151,"meta":0,"nbt":"CgAA"},{"id":356,"meta":0,"nbt":"CgAA"},{"id":404,"meta":0,"nbt":"CgAA"},{"id":410,"meta":0,"nbt":"CgAA"},{"id":125,"meta":3,"nbt":"CgAA"},{"id":23,"meta":3,"nbt":"CgAA"},{"id":3
3,"meta":1,"nbt":"CgAA"},{"id":29,"meta":1,"nbt":"CgAA"},{"id":46,"meta":0,"nbt":"CgAA"},{"id":421,"meta":0,"nbt":"CgAA"},{"id":-204,"meta":0,"nbt":"CgAA"},{"id":446,"meta":0,"nbt":"CgAA"},{"id":446,"meta":
8,"nbt":"CgAA"},{"id":446,"meta":7,"nbt":"CgAA"},{"id":446,"meta":15,"nbt":"CgAA"},{"id":446,"meta":12,"nbt":"CgAA"},{"id":446,"meta":14,"nbt":"CgAA"},{"id":446,"meta":1,"nbt":"CgAA"},{"id":446,"meta":4,"nb
t":"CgAA"},{"id":446,"meta":5,"nbt":"CgAA"},{"id":446,"meta":13,"nbt":"CgAA"},{"id":446,"meta":9,"nbt":"CgAA"},{"id":446,"meta":3,"nbt":"CgAA"},{"id":446,"meta":11,"nbt":"CgAA"},{"id":446,"meta":10,"nbt":"C
gAA"},{"id":446,"meta":2,"nbt":"CgAA"},{"id":446,"meta":6,"nbt":"CgAA"},{"id":446,"meta":15,"nbt":"CgADBFR5cGUCAA=="},{"id":434,"meta":0,"nbt":"CgAA"},{"id":434,"meta":1,"nbt":"CgAA"},{"id":434,"meta":2,"nb
t":"CgAA"},{"id":434,"meta":3,"nbt":"CgAA"},{"id":434,"meta":4,"nbt":"CgAA"},{"id":434,"meta":5,"nbt":"CgAA"},{"id":434,"meta":6,"nbt":"CgAA"},{"id":401,"meta":0,"nbt":"CgAKCUZpcmV3b3JrcwkKRXhwbG9zaW9ucwEAA
QZGbGlnaHQBAAA="},{"id":401,"meta":0,"nbt":"CgAKCUZpcmV3b3JrcwkKRXhwbG9zaW9ucwoCAQ1GaXJld29ya1RyYWlsAAEMRmlyZXdvcmtUeXBlAAcNRmlyZXdvcmtDb2xvcgIABwxGaXJld29ya0ZhZGUAAQ9GaXJld29ya0ZsaWNrZXIAAAEGRmxpZ2h0AQAA"}
,{"id":401,"meta":0,"nbt":"CgAKCUZpcmV3b3JrcwkKRXhwbG9zaW9ucwoCAQxGaXJld29ya1R5cGUABw1GaXJld29ya0NvbG9yAggHDEZpcmV3b3JrRmFkZQABD0ZpcmV3b3JrRmxpY2tlcgABDUZpcmV3b3JrVHJhaWwAAAEGRmxpZ2h0AQAA"},{"id":401,"meta"
:0,"nbt":"CgAKCUZpcmV3b3JrcwkKRXhwbG9zaW9ucwoCBw1GaXJld29ya0NvbG9yAgcHDEZpcmV3b3JrRmFkZQABD0ZpcmV3b3JrRmxpY2tlcgABDUZpcmV3b3JrVHJhaWwAAQxGaXJld29ya1R5cGUAAAEGRmxpZ2h0AQAA"},{"id":401,"meta":0,"nbt":"CgAKCUZ
pcmV3b3JrcwkKRXhwbG9zaW9ucwoCBwxGaXJld29ya0ZhZGUAAQ9GaXJld29ya0ZsaWNrZXIAAQ1GaXJld29ya1RyYWlsAAEMRmlyZXdvcmtUeXBlAAcNRmlyZXdvcmtDb2xvcgIPAAEGRmxpZ2h0AQAA"},{"id":401,"meta":0,"nbt":"CgAKCUZpcmV3b3JrcwkKRXhw
bG9zaW9ucwoCBw1GaXJld29ya0NvbG9yAgwHDEZpcmV3b3JrRmFkZQABD0ZpcmV3b3JrRmxpY2tlcgABDUZpcmV3b3JrVHJhaWwAAQxGaXJld29ya1R5cGUAAAEGRmxpZ2h0AQAA"},{"id":401,"meta":0,"nbt":"CgAKCUZpcmV3b3JrcwkKRXhwbG9zaW9ucwoCAQ9Ga
XJld29ya0ZsaWNrZXIAAQ1GaXJld29ya1RyYWlsAAEMRmlyZXdvcmtUeXBlAAcNRmlyZXdvcmtDb2xvcgIOBwxGaXJld29ya0ZhZGUAAAEGRmxpZ2h0AQAA"},{"id":401,"meta":0,"nbt":"CgAKCUZpcmV3b3JrcwkKRXhwbG9zaW9ucwoCAQ9GaXJld29ya0ZsaWNrZX
IAAQ1GaXJld29ya1RyYWlsAAEMRmlyZXdvcmtUeXBlAAcNRmlyZXdvcmtDb2xvcgIBBwxGaXJld29ya0ZhZGUAAAEGRmxpZ2h0AQAA"},{"id":401,"meta":0,"nbt":"CgAKCUZpcmV3b3JrcwkKRXhwbG9zaW9ucwoCBw1GaXJld29ya0NvbG9yAgQHDEZpcmV3b3JrRmF
kZQABD0ZpcmV3b3JrRmxpY2tlcgABDUZpcmV3b3JrVHJhaWwAAQxGaXJld29ya1R5cGUAAAEGRmxpZ2h0AQAA"},{"id":401,"meta":0,"nbt":"CgAKCUZpcmV3b3JrcwEGRmxpZ2h0AQkKRXhwbG9zaW9ucwoCAQ1GaXJld29ya1RyYWlsAAEMRmlyZXdvcmtUeXBlAAcN
RmlyZXdvcmtDb2xvcgIFBwxGaXJld29ya0ZhZGUAAQ9GaXJld29ya0ZsaWNrZXIAAAAA"},{"id":401,"meta":0,"nbt":"CgAKCUZpcmV3b3JrcwkKRXhwbG9zaW9ucwoCAQ9GaXJld29ya0ZsaWNrZXIAAQ1GaXJld29ya1RyYWlsAAEMRmlyZXdvcmtUeXBlAAcNRmlyZ
XdvcmtDb2xvcgINBwxGaXJld29ya0ZhZGUAAAEGRmxpZ2h0AQAA"},{"id":401,"meta":0,"nbt":"CgAKCUZpcmV3b3JrcwkKRXhwbG9zaW9ucwoCBw1GaXJld29ya0NvbG9yAgkHDEZpcmV3b3JrRmFkZQABD0ZpcmV3b3JrRmxpY2tlcgABDUZpcmV3b3JrVHJhaWwAAQ
xGaXJld29ya1R5cGUAAAEGRmxpZ2h0AQAA"},{"id":401,"meta":0,"nbt":"CgAKCUZpcmV3b3JrcwEGRmxpZ2h0AQkKRXhwbG9zaW9ucwoCAQ9GaXJld29ya0ZsaWNrZXIAAQ1GaXJld29ya1RyYWlsAAEMRmlyZXdvcmtUeXBlAAcNRmlyZXdvcmtDb2xvcgIDBwxGaXJ
ld29ya0ZhZGUAAAAA"},{"id":401,"meta":0,"nbt":"CgAKCUZpcmV3b3JrcwkKRXhwbG9zaW9ucwoCAQxGaXJld29ya1R5cGUABw1GaXJld29ya0NvbG9yAgsHDEZpcmV3b3JrRmFkZQABD0ZpcmV3b3JrRmxpY2tlcgABDUZpcmV3b3JrVHJhaWwAAAEGRmxpZ2h0AQAA
"},{"id":401,"meta":0,"nbt":"CgAKCUZpcmV3b3JrcwkKRXhwbG9zaW9ucwoCBw1GaXJld29ya0NvbG9yAgoHDEZpcmV3b3JrRmFkZQABD0ZpcmV3b3JrRmxpY2tlcgABDUZpcmV3b3JrVHJhaWwAAQxGaXJld29ya1R5cGUAAAEGRmxpZ2h0AQAA"},{"id":401,"met
a":0,"nbt":"CgAKCUZpcmV3b3JrcwkKRXhwbG9zaW9ucwoCBwxGaXJld29ya0ZhZGUAAQ9GaXJld29ya0ZsaWNrZXIAAQ1GaXJld29ya1RyYWlsAAEMRmlyZXdvcmtUeXBlAAcNRmlyZXdvcmtDb2xvcgICAAEGRmxpZ2h0AQAA"},{"id":401,"meta":0,"nbt":"CgAKC
UZpcmV3b3JrcwkKRXhwbG9zaW9ucwoCBw1GaXJld29ya0NvbG9yAgYHDEZpcmV3b3JrRmFkZQABD0ZpcmV3b3JrRmxpY2tlcgABDUZpcmV3b3JrVHJhaWwAAQxGaXJld29ya1R5cGUAAAEGRmxpZ2h0AQAA"},{"id":402,"meta":0,"nbt":"CgAKDUZpcmV3b3Jrc0l0ZW
0HDEZpcmV3b3JrRmFkZQABD0ZpcmV3b3JrRmxpY2tlcgABDUZpcmV3b3JrVHJhaWwAAQxGaXJld29ya1R5cGUABw1GaXJld29ya0NvbG9yAgAAAwtjdXN0b21Db2xvcr2Llw4A"},{"id":402,"meta":8,"nbt":"CgAKDUZpcmV3b3Jrc0l0ZW0BD0ZpcmV3b3JrRmxpY2t
lcgABDUZpcmV3b3JrVHJhaWwAAQxGaXJld29ya1R5cGUABw1GaXJld29ya0NvbG9yAggHDEZpcmV3b3JrRmFkZQAAAwtjdXN0b21Db2xvctvCxQsA"},{"id":402,"meta":7,"nbt":"CgADC2N1c3RvbUNvbG9y0YmTBgoNRmlyZXdvcmtzSXRlbQcNRmlyZXdvcmtDb2xv
cgIHBwxGaXJld29ya0ZhZGUAAQ9GaXJld29ya0ZsaWNrZXIAAQ1GaXJld29ya1RyYWlsAAEMRmlyZXdvcmtUeXBlAAAA"},{"id":402,"meta":15,"nbt":"CgAKDUZpcmV3b3Jrc0l0ZW0HDUZpcmV3b3JrQ29sb3ICDwcMRmlyZXdvcmtGYWRlAAEPRmlyZXdvcmtGbGlj
a2VyAAENRmlyZXdvcmtUcmFpbAABDEZpcmV3b3JrVHlwZQAAAwtjdXN0b21Db2xvcp+8eAA="},{"id":402,"meta":12,"nbt":"CgAKDUZpcmV3b3Jrc0l0ZW0BDUZpcmV3b3JrVHJhaWwAAQxGaXJld29ya1R5cGUABw1GaXJld29ya0NvbG9yAgwHDEZpcmV3b3JrRmFk
ZQABD0ZpcmV3b3JrRmxpY2tlcgAAAwtjdXN0b21Db2xvcsuwqgwA"},{"id":402,"meta":14,"nbt":"CgAKDUZpcmV3b3Jrc0l0ZW0BD0ZpcmV3b3JrRmxpY2tlcgABDUZpcmV3b3JrVHJhaWwAAQxGaXJld29ya1R5cGUABw1GaXJld29ya0NvbG9yAg4HDEZpcmV3b3Jr
RmFkZQAAAwtjdXN0b21Db2xvcsX/MwA="},{"id":402,"meta":1,"nbt":"CgAKDUZpcmV3b3Jrc0l0ZW0HDUZpcmV3b3JrQ29sb3ICAQcMRmlyZXdvcmtGYWRlAAEPRmlyZXdvcmtGbGlja2VyAAENRmlyZXdvcmtUcmFpbAABDEZpcmV3b3JrVHlwZQAAAwtjdXN0b21Db
2xvcrPH/gQA"},{"id":402,"meta":4,"nbt":"CgAKDUZpcmV3b3Jrc0l0ZW0HDUZpcmV3b3JrQ29sb3ICBAcMRmlyZXdvcmtGYWRlAAEPRmlyZXdvcmtGbGlja2VyAAENRmlyZXdvcmtUcmFpbAABDEZpcmV3b3JrVHlwZQAAAwtjdXN0b21Db2xvcqvtnQwA"},{"id":4
02,"meta":5,"nbt":"CgAKDUZpcmV3b3Jrc0l0ZW0HDEZpcmV3b3JrRmFkZQABD0ZpcmV3b3JrRmxpY2tlcgABDUZpcmV3b3JrVHJhaWwAAQxGaXJld29ya1R5cGUABw1GaXJld29ya0NvbG9yAgUAAwtjdXN0b21Db2xvco+1tgcA"},{"id":402,"meta":13,"nbt":"C
gAKDUZpcmV3b3Jrc0l0ZW0HDUZpcmV3b3JrQ29sb3ICDQcMRmlyZXdvcmtGYWRlAAEPRmlyZXdvcmtGbGlja2VyAAENRmlyZXdvcmtUcmFpbAABDEZpcmV3b3JrVHlwZQAAAwtjdXN0b21Db2xvcoXFxQMA"},{"id":402,"meta":9,"nbt":"CgAKDUZpcmV3b3Jrc0l0ZW
0HDUZpcmV3b3JrQ29sb3ICCQcMRmlyZXdvcmtGYWRlAAEPRmlyZXdvcmtGbGlja2VyAAENRmlyZXdvcmtUcmFpbAABDEZpcmV3b3JrVHlwZQAAAwtjdXN0b21Db2xvcqvRYwA="},{"id":402,"meta":3,"nbt":"CgAKDUZpcmV3b3Jrc0l0ZW0BDEZpcmV3b3JrVHlwZQA
HDUZpcmV3b3JrQ29sb3ICAwcMRmlyZXdvcmtGYWRlAAEPRmlyZXdvcmtGbGlja2VyAAENRmlyZXdvcmtUcmFpbAAAAwtjdXN0b21Db2xvcpuv5QcA"},{"id":402,"meta":11,"nbt":"CgAKDUZpcmV3b3Jrc0l0ZW0BD0ZpcmV3b3JrRmxpY2tlcgABDUZpcmV3b3JrVHJ
haWwAAQxGaXJld29ya1R5cGUABw1GaXJld29ya0NvbG9yAgsHDEZpcmV3b3JrRmFkZQAAAwtjdXN0b21Db2xvcoWfCQA="},{"id":402,"meta":10,"nbt":"CgAKDUZpcmV3b3Jrc0l0ZW0HDEZpcmV3b3JrRmFkZQABD0ZpcmV3b3JrRmxpY2tlcgABDUZpcmV3b3JrVHJ
haWwAAQxGaXJld29ya1R5cGUABw1GaXJld29ya0NvbG9yAgoAAwtjdXN0b21Db2xvcsHj+QcA"},{"id":402,"meta":2,"nbt":"CgAKDUZpcmV3b3Jrc0l0ZW0BDEZpcmV3b3JrVHlwZQAHDUZpcmV3b3JrQ29sb3ICAgcMRmlyZXdvcmtGYWRlAAEPRmlyZXdvcmtGbGlj
a2VyAAENRmlyZXdvcmtUcmFpbAAAAwtjdXN0b21Db2xvctOPjAoA"},{"id":402,"meta":6,"nbt":"CgAKDUZpcmV3b3Jrc0l0ZW0BDUZpcmV3b3JrVHJhaWwAAQxGaXJld29ya1R5cGUABw1GaXJld29ya0NvbG9yAgYHDEZpcmV3b3JrRmFkZQABD0ZpcmV3b3JrRmxpY
2tlcgAAAwtjdXN0b21Db2xvcseNyw4A"}]`
