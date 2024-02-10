# Blogsite9000
Personal blogging site template. Written in Go & HTMX.

If [the site](https://alepablog.com) is down, it's not actively being pitched to anyone.

# Purpose
To:
- learn the fundementals of how the backend of a website operates
- learn basic interactivity through HTMX
- learn fundimental techniques used in webserver development
- Backend deployment on bare-metal Linux

# Usage
This is a personal project, so a major fork is required for deployment for other people.

- Create a MariaDB (MySQL) database with name "Blogsite9000", with appropriate permissions for the user you wish to use
- Source the `create-posts.sql` and `create-comments.sql` in the cli of the user you wish to use
- Set environmental variables for the DB with: `USERNAME="[username]"; PASSWORD="[password]";`

## Build

```bash
go build
```

## Run

```bash
`./main`
```

## Creating posts

- Create a post with:

   `INSERT INTO posts (Author, Title, LinkToPost) VALUES ("[author name]", "[post title]", "[path to html formatted blog post]");`

# Roadmap
- Back-button functionality for easier navigation
