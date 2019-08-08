package entry

// ConfigCloth  ...
type ConfigCloth struct {
	Id    int32
	Name  string
	Icon  string
	Cost  string
	Level int32
	Type  int32
	Exp   int32
	Star  int32
}

// ConfigScene  ...
type ConfigScene struct {
	Id    int32
	Name  string
	Icon  string
	Level int32
	Star  int32
}

// ConfigLevel  ...
type ConfigLevel struct {
	Id    int32
	Name  string
	Icon  string
	Level int32
	Star  int32
}

// ConfigSign  ...
type ConfigSign struct {
	Id  int32
	Day int32
	Num int32
}

// ConfigGift  ...
type ConfigGift struct {
	Id      int32
	Name    string
	Icon    string
	Diamond int32
	Favour  int32
	Reward  int32
}
