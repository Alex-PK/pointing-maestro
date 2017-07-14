package main

type MsgGetVote struct {
	user string
	vote string
}

type MsgSendVote struct {
	user string
}

type MsgSendVoteInfo struct {
	user string
	vote string
}

type MsgGetCmd struct {
	cmd string
}

type msgSendCmd struct {
	cmd string
}
