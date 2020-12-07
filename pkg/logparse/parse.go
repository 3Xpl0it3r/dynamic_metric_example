package logparse



import (
"regexp"
"strconv"
"time"
)

const (
	TimeLayout = "02/Jan/2006:15:04:04 -0700"
)

// LogDataGram is spec of record nginx log
type LogDataGram struct {
	// remote ip address
	RemoteAddr string
	// remote user
	RemoteUser string
	// the time request come in
	TimeLocal string
	// request url, the format is "method url http/version"
	Request string
	Status string
	BodyByteSent string
	HttpRefer string
	HttpUserAgent string
	HttpXForwardFor string
	RequestTime string
	CookieUserTag *TidEntry
	// extra payload
	CurrentTime int64
}

// TidEntry is spec of tid information
type TidEntry struct {
	Login bool
	UserId string
	TimeStamp int64
	SeqNumber string
}


// nginx log parse
func LogParse(body string)*LogDataGram{
	currentTimeStamp := time.Now().UTC().UnixNano()/1e6
	pattern := `(?P<remote_addr>\S*)\s-\s(?P<remote_user>\S*)\s\[(?P<time_local>.*?)\]\s\"(?P<request>.*?)\"\s(?P<status>\d{3})\s(?P<body_bytes_sent>\S*)\s\"(?P<http_refer>[^\"]*)\"\s\"(?P<http_user_agent>.*?)\"\s\"(?P<http_x_forwarded_for>\S*)\"\s(?P<request_time>\S*)\s-\s(?P<cookie_userTag>\S*)`
	regex := regexp.MustCompile(pattern)
	fields := regex.FindStringSubmatch(body)
	dataGram := &LogDataGram{
		RemoteAddr:      fields[1],
		RemoteUser:      fields[2],

		Request:         fields[4],
		Status:          fields[5],
		BodyByteSent:    fields[6],
		HttpRefer:       fields[7],
		HttpUserAgent:   fields[8],
		HttpXForwardFor: fields[9],
		RequestTime:     fields[10],
		// extra time
		CurrentTime: currentTimeStamp,
	}
	tidPattern := `(?P<login>\d+)\.(?P<userId>\d{11,13})\.(?P<timestamp>\d{13})(?P<seq>\d{4})`
	tidGex := regexp.MustCompile(tidPattern)
	tidFields := tidGex.FindStringSubmatch(fields[11])


	if len(tidFields) == 0{
		// not match tid format, set cookieUserTag = nil, then return
		dataGram.CookieUserTag = nil
		return dataGram
	}
	print(currentTimeStamp)

	login := tidFields[1]
	userId := tidFields[2]
	timestamp ,_ := strconv.Atoi(tidFields[3])
	seqNumber := tidFields[4]
	if len(login) > 1{
		dataGram.CookieUserTag = &TidEntry{
			Login:     false,
			UserId:    "",
			TimeStamp: int64(timestamp),
			SeqNumber: seqNumber,
		}
	}else {
		dataGram.CookieUserTag = &TidEntry{
			Login:     true,
			UserId:    userId[:5],
			TimeStamp: int64(timestamp),
			SeqNumber: seqNumber,
		}
	}
	return dataGram
}
