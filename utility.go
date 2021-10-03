package main

import "time"

type Utility struct {
	Log LogManager
}

func (utility *Utility) ConvertStringToTime(dateTime string, messageErr string) (time.Time, bool) {
	date, errConvertStart := time.Parse(time.RFC3339, dateTime)
	if errConvertStart != nil {
		utility.Log.WriteErrorLog(messageErr)
		return time.Time{}, false
	} else {
		return date, true
	}
}
