package models

// "@timestamp" is for data stream

type DbusUnitRecord struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	LoadState   string `json:"load_state"`
	ActiveState string `json:"active_state"`
	Substate    string `json:"sub_state"`
	Followed    string `json:"followed"`
	Path        string `json:"path"`
	JobId       uint32 `json:"job_id"`
	JobType     string `json:"job_type"`
	JobPath     string `json:"job_path"`
	Timestamp   string `json:"@timestamp"`
	AgentId     string `json:"agent_Id"`
}

type ExecRecord struct {
	Filename  string `json:"filename"`
	Pid       uint32 `json:"pid"`
	OldPid    uint32 `json:"old_pid"`
	Uid       uint32 `json:"uid"`
	Comm      string `json:"comm"`
	Timestamp string `json:"@timestamp"`
	AgentId   string `json:"agent_id"`
}

type ExecveRecord struct {
	Pid       uint32 `json:"pid"`
	Name      string `json:"name"`
	Comm      string `json:"comm"`
	PPid      int64  `json:"p_pid"`
	Timestamp string `json:"@timestamp"`
	AgentId   string `json:"agent_id"`
}

type FanotifyRecord struct {
	Pid       int64    `json:"pid"`
	Name      string   `json:"name"`
	Comm      string   `json:"comm"`
	PPid      int64    `json:"p_pid"`
	Path      string   `json:"path"`
	Events    []string `json:"events"`
	Timestamp string   `json:"@timestamp"`
	AgentId   string   `json:"agent_id"`
}

type ISSSRecord struct {
	OldState  string `json:"old_state"`
	NewState  string `json:"new_state"`
	SPort     uint32 `json:"src_port"`
	DPort     uint32 `json:"dest_port"`
	Family    string `json:"family"`
	Protocol  string `json:"protocol"`
	SAddr     string `json:"src_ip"`
	DAddr     string `json:"dest_ip"`
	Pid       int64  `json:"pid"`
	Name      string `json:"name"`
	Comm      string `json:"comm"`
	PPid      int64  `json:"p_pid"`
	Timestamp string `json:"@timestamp"`
	AgentId   string `json:"agent_id"`
}

type Connect4Record struct {
	UserFamily string `json:"user_family"`
	UserIPv4   string `json:"dest_ip"`
	UserPort   uint32 `json:"dest_port"`
	Family     string `json:"family"`
	Type       string `json:"type"`
	Protocol   string `json:"protocol"`
	Uid        uint32 `json:"uid"`
	Pid        int64  `json:"pid"`
	Name       string `json:"name"`
	Comm       string `json:"comm"`
	PPid       int64  `json:"p_pid"`
	Timestamp  string `json:"@timestamp"`
	AgentId    string `json:"agent_id"`
}

type Bind4Record struct {
	UserFamily string `json:"user_family"`
	UserIPv4   string `json:"dest_ip"`
	UserPort   uint32 `json:"dest_port"`
	Family     string `json:"family"`
	Type       string `json:"type"`
	Protocol   string `json:"protocol"`
	Uid        uint32 `json:"uid"`
	Pid        int64  `json:"pid"`
	Name       string `json:"name"`
	Comm       string `json:"comm"`
	PPid       int64  `json:"p_pid"`
	Timestamp  string `json:"@timestamp"`
	AgentId    string `json:"agent_id"`
}
