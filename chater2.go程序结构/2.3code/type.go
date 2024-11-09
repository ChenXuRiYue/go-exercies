package main

// 定义一个别名集合

type (
	Second int
	Minute int
	Hour   int
)

func SecondToMinute(s Second) Minute {
	return Minute(s / 60)
}

func HourToSecond(h Hour) Second {
	return Second(h * 3600)
}

// ...
