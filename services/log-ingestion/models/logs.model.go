package models

import "google.golang.org/protobuf/types/known/timestamppb"

type DbusUnitRecord struct {
	Name        string
	Description string
	LoadState   string
	ActiveState string
	Substate    string
	Followed    string
	Path        string
	JobId       uint32
	JobType     string
	JobPath     string
	Timestamp   *timestamppb.Timestamp
	AgentId     string
}

type ExecRecord struct {
	Filename  string
	Pid       uint32
	OldPid    uint32
	Uid       uint32
	Comm      string
	Timestamp *timestamppb.Timestamp
	AgentId   string
}

type ExecveRecord struct {
	Pid       uint32
	Name      string
	Comm      string
	PPid      int64
	Timestamp *timestamppb.Timestamp
	AgentId   string
}

type FanotifyRecord struct {
	Pid       int64
	Name      string
	Comm      string
	PPid      int64
	Path      string
	Events    []string
	Timestamp *timestamppb.Timestamp
	AgentId   string
}

type ISSSRecord struct {
	OldState  string
	NewState  string
	SPort     uint32
	DPort     uint32
	Family    string
	Protocol  string
	SAddr     string
	DAddr     string
	Pid       int64
	Name      string
	Comm      string
	PPid      int64
	Timestamp *timestamppb.Timestamp
	AgentId   string
}

type Connect4Record struct {
	UserFamily string
	UserIPv4   string
	UserPort   uint32
	Family     string
	Type       string
	Protocol   string
	Uid        uint32
	Pid        int64
	Name       string
	Comm       string
	PPid       int64
	Timestamp  *timestamppb.Timestamp
	AgentId    string
}

type Bind4Record struct {
	UserFamily string
	UserIPv4   string
	UserPort   uint32
	Family     string
	Type       string
	Protocol   string
	Uid        uint32
	Pid        int64
	Name       string
	Comm       string
	PPid       int64
	Timestamp  *timestamppb.Timestamp
	AgentId    string
}
