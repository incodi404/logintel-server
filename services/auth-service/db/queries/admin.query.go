package queries

var CHECK_IF_ADMIN_EXISTS = `
SELECT EXISTS (
	SELECT 1 FROM admins WHERE email=$1
)
`
var ADD_NEW_ADMIN = `INSERT INTO admins (email, password_hash) VALUES ($1, $2) RETURNING id;`
