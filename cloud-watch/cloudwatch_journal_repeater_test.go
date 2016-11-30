package cloud_watch

import (
	"testing"
	"time"
)

func TestRepeater(t *testing.T) {

	config_data := `
log_group="dcos-logstream-test"
state_file="/var/lib/journald-cloudwatch-logs/state-test"
log_priority=3
debug=true
local=true
log_stream="test-stream"
log_group="test-group"
	`

	config, _ := LoadConfigFromString(config_data, nil)
	session := NewAWSSession(config)
	repeater, err := NewCloudWatchJournalRepeater(session, nil, config)

	if err != nil {
		t.Errorf("Unable to created new repeater %s", err)
		t.Fail()
	}

	if repeater == nil {
		t.Error("Repeater nil")
		t.Fail()
	}

	records := []Record{
		{Message:"Hello mom", TimeUsec:time.Now().Unix() * 1000},
		{Message:"Hello dad", TimeUsec:time.Now().Unix() * 1000},
	}
	err = repeater.WriteBatch(records)

	if err != nil {
		t.Errorf("Unable to write batch %s", err)
		t.Fail()
	}

}