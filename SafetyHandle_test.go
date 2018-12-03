package Common

import (
	"testing"
)

func TestSignCheck(t *testing.T) {
	var ReqSecurity HttpRequestSecurity
	ReqSecurity.SignCheck("ask_id=215262670&goal_id=215038067&token=14a10a0770957d5a527c13940d7a362d&v_id=1.4.5")
}
