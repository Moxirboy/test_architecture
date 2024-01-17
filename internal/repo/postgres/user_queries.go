package postgres

const(
	createUser=`
insert into 
users (email,password,role,created_at,updated_at,deleted_at,is_email_verified) 
values($1,$2,$3,$4,$5,$6,$7) 
returning id
`
	getUser=`
SELECT is_active
FROM users
WHERE email = $1 AND is_active = true
`
		getUserByEmail=	`
SELECT id ,password FROM users WHERE email=$1
	`
		deleteUserById=`
UPDATE users SET is_active=false ,deleted_at=current_timestamp WHERE id=$1
`
		updateIsVerified=`
UPDATE users SET is_email_verified=true WHERE id=$1
`

)