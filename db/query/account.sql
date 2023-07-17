-- name: CreateAccount :one
INSERT INTO accounts (
  owner, balance,currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccount :many
SELECT * FROM accounts
WHERE owner = $1
ORDER BY id
LIMIT $2 OFFSET $3;

-- name: UpdateAccount :one
UPDATE accounts
set balance = $2
WHERE id = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE accounts
set balance = balance + sqlc.arg(amount)
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;