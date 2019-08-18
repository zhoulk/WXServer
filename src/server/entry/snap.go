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

// BarrageReport ...
type BarrageReport struct {
	From string
	To   string
	Msg  string
}

// ExtraMoney ...
type ExtraMoney struct {
	Uid     string
	LvChao  string
	Diamond int32
	Reason  int32
}

// OpenFrom ...
type OpenFrom struct {
	Uid     string
	FromUid string
	Type    int32
}
