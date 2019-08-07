package entry

// Snap ...
type Snap struct {
	Uid    string
	LvChao string
}

// FavourLog ...
type FavourLog struct {
	Uid   string
	ToUID string
	Day   string
	Num   int32
}

// FavourReport ...
type FavourReport struct {
	From string
	To   string
	Num  int32
}
