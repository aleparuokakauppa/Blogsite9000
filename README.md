# Blogsite9000
Personal blogging site template. Written in `Go` & `HTMX`.

If [the site](https://alepablog.com) is down, it's not actively being pitched to anyone.

# Purpose
To:
- Learn the fundementals of how the backend of a website operates
- Learn basic interactivity through HTMX
- Learn fundimental techniques used in webserver development
- Backend deployment on bare-metal Linux
# Roadmap
- Clean the inline CSS
- Refactoring

# Usage
This is a personal project, so a major fork is required for usage for other people.

- Create a MariaDB (MySQL) database with name `Blogsite9000`, with appropriate permissions for the user you wish to use
- Source the `create-posts.sql` and `create-comments.sql` in the cli of the user you wish to use
- Set environmental variables for the DB with:
```bash
USERNAME="{username}"; PASSWORD="{password}";
```

## Build
```bash
go build
```

## Run
Run the built executable
```bash
./main
```

## Creating posts
Using the MariaDB cli on the user added in the environmental variables
```SQL
INSERT INTO posts (Author, Title, LinkToPost) VALUES ("{author name}", "{post title}", "{path to html formatted blog post}");
```
