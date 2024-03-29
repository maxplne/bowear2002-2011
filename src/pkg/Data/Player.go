package Data

import . "launchpad.net/gobson/bson"

type DType byte

const (
	Infantry = DType(0)
	Mobile   = DType(1)
	Aviation = DType(2)
	Organic  = DType(3)
	Other    = DType(4)
)

type Player struct {
	ID     string "_id"
	UserID string
	Name   string
 
	Faction int32
  
	Avatar    byte
	Tactics   byte
	Clout     byte
	Education byte
	MechApt   byte

	Points    uint16
	Divisions []Division
	
	RecordWon uint16
	RecordLost uint16
	Prestige uint32
	
	Honor uint32 //Level = Honor / 100
	
	Money int32
	Ore int32
	Silicon int32
	Uranium int32
	Sulfur byte
	
	Reincatnation byte

	Map  int16
	X, Y int16
}

type Division struct {
	Type  DType
	Level byte
	Rank  string
	XP    uint32
}
 
func (d *Division) Influence(p *Player) byte {
	return (p.Clout / 2) + d.Level
}

func (d *Division) TotalXP() uint32{
	level3 := uint32(d.Level) * uint32(d.Level) * uint32(d.Level);
	if (d.Level >= 49) {
		return uint32((float32(1.5) * float32(level3)) * float32(d.Level - 45))
	} else {
		return 6 * level3
	}
	return 0
}

func (p *Player) MaxUnits() byte{
	return 48 + (p.Tactics/20)
}

func (d *Player) TotalHonor() uint32{
	return 200
}

func NewPlayer() *Player {
	p := new(Player)

	p.Divisions = make([]Division, 4)
	p.Divisions[Infantry] = Division{Infantry, 1, "", 0}
	p.Divisions[Mobile] = Division{Mobile, 1, "" , 0}
	p.Divisions[Aviation] = Division{Aviation, 1, "", 0}
	p.Divisions[Organic] = Division{Organic, 1, "", 0}

	return p
} 
 
func (p *Player) SetDefaultStats() {
	p.Points = 0
	p.Money = 300000
}

func RegisterPlayer(plyaer *Player) {
	e := CPlayers.Insert(plyaer)
	if e != nil {
		panic(e)
	}
}

func GetPlayerByUserID(id string) *Player {
	p := new(Player)
	e := CPlayers.Find(M{"userid": id}).One(p)
	if e != nil {
		return nil
	}
	return p
}

func SavePlayer(Player *Player) {
	e := CPlayers.Update(M{"_id": Player.ID}, Player)
	if e != nil {
		panic(e)
	}
}
