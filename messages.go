package main

type Msg struct {
	Cmd       string    `json:"cmd"`
	User      string    `json:"user,omitempty"`
	Vote      string    `json:"vote,omitempty"`
	StoryDesc string    `json:"storyDesc,omitempty"`
	UserList  []string  `json:"userList,omitempty"`
	VoteList  map[string]string `json:"voteList,omitempty"`
}
