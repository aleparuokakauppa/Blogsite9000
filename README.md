# Blogsite9000
Personal blogging site template. Written in go + htmx.

# Purpose
To learn fundimental techniques used in webserver development and basic interactivity through htmx.
To learn how a website is developed from the backend, database, frontend, and deployment on Linux.

## Golang is overkill for a blog...
This is true, but as a learning experience creating any full-stack project incorporates many aspects of the development process which creates a great opportunity for learning

# Usage
This is a personal project, so change source code as needed
## Build
`go build`
## Run
`./main`
## Creating posts
- Create a MariaDB (MySQL) database, with appropriate permissions for the user you wish to use
- Source the `create-posts.sql` and `create-comments.sql` in the cli
- Create a post with:
  
   `INSERT INTO posts (Author, Title, LinkToPost) VALUES ("[author name]", "[post title]", "[path to html formatted blog post]");`
- Set environmental variables for the DB with: `USERNAME="[username]"; PASSWORD="[password]";`

# Roadmap
- Cleaning up
- Maybe a frontend for posting blog posts (currently done manually in mysql...)
