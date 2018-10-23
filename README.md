#API Endpoints

`users`

- `POST /api/users` - for signing up new users

###post-MVP:
- `GET /api/users/:id` - returns the user page, for looking at user's comments, history

`session`

- `POST /api/session` - log in
- `DELETE /api/session` - log out

`songs`

- `POST /api/songs` - add new song to database
- `GET /api/songs` - get new trending songs? pass parameter?

###post-MVP:
- `GET /api/songs/:id` - gets specific song by id

`comments`

- `POST /api/songs/:id/comments` - add a comment to a song
- `DELETE /api/songs/:id/comments/:comment_id` - delete comment from song

###post-MVP:
- `PATCH /api/songs/:id/comments/:comment_id` - edit comment

#Database Schema

##Users

| Column Name     | Data Type | Details                   |
|-----------------|-----------|---------------------------|
| id              | integer   | not null, primary key     |
| username        | string    | not null, indexed, unique |
| password_digest | string    | not null                  |
| session_token   | string    | not null, indexed, unique |
| created_at      | datetime  | not null                  |
| updated_at      | datetime  | not null                  |

##Songs

| Column Name | Data Type | Details               |
|-------------|-----------|-----------------------|
| id          | integer   | not null, primary key |
| name        | string    | not null, indexed     |
| artist_id   | integer   | not null, indexed     |
| artist_name | string    | not null              |
| url         | string    | not null              |
| created_at  | datetime  | not null              |
| updated_at  | datetime  | not null              |

##Artists

| Column Name | Data Type | Details               |
|-------------|-----------|-----------------------|
| id          | integer   | not null, primary key |
| name        | string    | not null, indexed     |
| bio         | string    |                       |
| location    | string    |                       |
| created_at  | datetime  | not null              |
| updated_at  | datetime  | not null              |

##Song Comments

| Column Name | Data Type | Details               |
|-------------|-----------|-----------------------|
| id          | integer   | not null, primary key |
| user_id     | integer   | not null, indexed     |
| song_id     | integer   | not null, indexed     |
| body        | string    | not null              |
| created_at  | datetime  | not null              |
| updated_at  | datetime  | not null              |