package storage

import (
	"strings"
	"testing"
	"fmt"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	//save and restore original notifyUser.
	saved := notifyUser
	defer func() { notifyUser = saved }()

	//install test's fake notifyUser
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}
	//simulate 980 mb used condition...
	const user = "joe@example.org"
	CheckQuota(user)
	fmt.Printf("notifiedUser: '%q'\nnotifiedMsg: '%q'\n", notifiedUser, notifiedMsg)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+"want substring %q", notifiedMsg, wantSubstring)
	}
}
