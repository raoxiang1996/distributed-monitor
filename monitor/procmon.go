package monitor

type ProcessData struct {
	// int pid;
	state   rune  `json:"state"` //单字符1一个字节
	ppid    int32 `json:"ppid"`
	pgrp    int32 `json:"pgrp"`
	session int32 `json:"session"`
	tty_nr  int32 `json:"tty_nr"`
	tpgid   int32 `json:"tpgid"`
	flags   int32 `json:"flags"`

	minflt  int64 `json:"minflt"`
	cminflt int64 `json:"cminflt"`
	majflt  int64 `json:"majflt"`
	cmajflt int64 `json:"cmajflt"`

	utime  int64 `json:"utime"`
	stime  int64 `json:"stime"`
	cutime int64 `json:"cutime"`
	cstime int64 `json:"cstime"`

	priority    int64 `json:"priority"`
	nice        int64 `json:"nice"`
	num_threads int64 `json:"num_threads"`
	itrealvalue int64 `json:"itrealvalue"`
	starttime   int64 `json:"starttime"`

	vsizeKb int64 `json:"vsize_kb"`
	rssKb   int64 `json:"rss_kb"`
	rsslim  int64 `json:"rsslim"`
}

func (this *ProcessData) Parse(startdata string) {

}
