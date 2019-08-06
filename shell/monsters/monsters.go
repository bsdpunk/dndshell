package monsters

type Monsters struct {
	Monsters []Monster `json:"monsters"`
	//Monster  Monster    `json:"monster"`
}

type Monster struct {
	Name        string `json:"name"`
	Id          int    `json:"id"`
	Age         string `json:"HITPOINTS"`
	AC          int    `json:"AC"`
	CR          int    `json:"CR"`
	XP          string `json:"XP"`
	Description string `json:"Descriptions"`
	//squad
}
