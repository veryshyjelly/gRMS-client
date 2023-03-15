package client

import "gRMS-client/modals"

func (c *client) GetUser(userId uint64) error {
	var req = modals.Req{
		GetUser: userId,
	}
	return c.Conn.WriteJSON(req)
}

func (c *client) ChangePassword() {

}

func (c *client) ChangeUsername() {

}