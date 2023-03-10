package client

import "gRMS-client/modals"

func (c *MyClient) GetUser(userId uint64) error {
	var req = modals.Req{
		GetUser: userId,
	}
	return c.Conn.WriteJSON(req)
}

func (c *MyClient) ChangePassword() {

}

func (c *MyClient) ChangeUsername() {

}
