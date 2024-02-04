# Blogsite9000
Personal blogging site template. Written in Go & HTMX.

If the site is down, it's not actively being pitched to anyone.

[The site](https://alepablog.com)

# Purpose
To:
- learn the fundementals of how the backend of a website operates
- learn basic interactivity through HTMX
- learn fundimental techniques used in webserver development
- Backend deployment on bare-metal Linux

## Golang is overkill for a blog...
This is true, but as a learning experience creating any full-stack project is a great learning process.

# Usage
This is a personal project, so a major fork is required for deployment for other people.

- Create a MariaDB (MySQL) database with name "Blogsite9000", with appropriate permissions for the user you wish to use
- Source the `create-posts.sql` and `create-comments.sql` in the cli of the user you wish to use
- Set environmental variables for the DB with: `USERNAME="[username]"; PASSWORD="[password]";`

## Build
`go build`
## Run
`./main`
## Creating posts

- Create a post with:

   `INSERT INTO posts (Author, Title, LinkToPost) VALUES ("[author name]", "[post title]", "[path to html formatted blog post]");`

# Roadmap
- Maybe accounts, but since this is a toy-project, it would also be just an excercise.
