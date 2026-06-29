package elasticsearch

var DbusUnitMapping = `
	{
		"properties": {
			"name": {"type": "keyword"},
			"description": {"type": "text"},
			"load_state": {"type": "keyword"},
			"active_state": {"type": "keyword"},
			"sub_state": {"type": "keyword"},
			"followed": {"type": "keyword"},
			"path": {"type": "text"},
			"job_id": {"type": "keyword"},
			"job_type": {"type": "keyword"},
			"job_path": {"type": "text"},
			"@timestamp": {"type": "date"},
			"agent_Id": {"type": "keyword"}
		}
	}
`
var ExecMapping = `
	{
		"properties": {
			"filename": {"type": "keyword"},
			"pid": {"type": "unsigned_long"},
			"old_pid": {"type": "unsigned_long"},
			"uid": {"type": "integer"},
			"@timestamp": {"type": "date"},
			"agent_id": {"type": "keyword"}
		}
	}
`
var ExecveMapping = `
	{
		"properties": {
			"pid": {"type": "unsigned_long"},
			"p_pid": {"type": "unsigned_long"},
			"name": {"type": "keyword"},
			"comm": {"type": "text"},
			"@timestamp": {"type": "date"},
			"agent_id": {"type": "keyword"}
		}
	}
`
var FanotifyMapping = `
	{
		"properties": {
			"pid": {"type": "unsigned_long"},
			"name": {"type": "keyword"},
			"comm": {"type": "text"},
			"p_pid": {"type": "unsigned_long"},
			"path": {"type": "text"},
			"events": {"type": "text"},
			"@timestamp": {"type": "date"},
			"agent_id": {"type": "keyword"}
		}
	}
`
var ISSSMapping = `
	{
		"properties": {
			"old_state": {"type": "keyword"},
			"new_state": {"type": "keyword"},
			"src_port": {"type": "integer"},
			"dest_port": {"type": "integer"},
			"family": {"type": "keyword"},
			"protocol": {"type": "keyword"},
			"src_ip": {"type": "ip"},
			"dest_ip": {"type": "ip"},
			"pid": {"type": "unsigned_long"},
			"name": {"type": "keyword"},
			"comm": {"type": "text"},
			"p_pid": {"type": "unsigned_long"},
			"@timestamp": {"type": "date"},
			"agent_id": {"type": "keyword"}
		}
	}
`
var Connect4Mapping = `
	{
		"properties": {
			"user_family": {"type": "keyword"},
			"dest_ip": {"type": "keyword"},
			"dest_port": {"type": "keyword"},
			"family": {"type": "keyword"},
			"type": {"type": "keyword"},
			"protocol": {"type": "keyword"},
			"uid": {"type": "keyword"},
			"pid": {"type": "keyword"},
			"name": {"type": "keyword"},
			"comm": {"type": "keyword"},
			"p_pid": {"type": "keyword"},
			"@timestamp": {"type": "date"},
			"agent_id": {"type": "keyword"}
		}
	}
`
var Bind4Mapping = `
	{
		"properties": {
			"user_family": {"type": "keyword"},
			"dest_ip": {"type": "keyword"},
			"dest_port": {"type": "keyword"},
			"family": {"type": "keyword"},
			"type": {"type": "keyword"},
			"protocol": {"type": "keyword"},
			"uid": {"type": "keyword"},
			"pid": {"type": "keyword"},
			"name": {"type": "keyword"},
			"comm": {"type": "keyword"},
			"p_pid": {"type": "keyword"},
			"@timestamp": {"type": "date"},
			"agent_id": {"type": "keyword"}
		}
	}
`
