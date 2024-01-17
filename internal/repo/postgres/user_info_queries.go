package postgres

const (
	createUserInfo=`
INSERT INTO 
user_info (user_id, name, weigh, height, age, waist, created_at, gender) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
RETURNING id
`
	
	
)