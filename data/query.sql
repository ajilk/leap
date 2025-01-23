-- name: GetBookmark :one
select *
from bookmarks
where id = ?
limit 1
;

-- name: ListBookmarks :many
select *
from bookmarks
order by key
;

-- name: CreateBookmark :one
insert into
  bookmarks (key, value)
values
  (?, ?) returning *;

-- name: UpdateBookmark :exec
update bookmarks
set
  key = ?,
  value = ?
where
  id = ? returning *;

-- name: DeleteBookmark :exec
delete from bookmarks
where id = ?
;

-- name: DeleteAllBookmarks :exec
delete from bookmarks
;

