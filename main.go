package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Comment struct {
    ID int64
    Author string
    Text string
    Target int64
    Time string
}

type Post struct {
    ID int64
    Author string
    Title string
    LinkToPost string
    Comments []Comment
}

// Pings the DB to see if DB is alive
func connectDB() error {
    // Capture connection properties.
    cfg := mysql.Config{
        User: os.Getenv("$USER"),
        Passwd: os.Getenv("$PASSWD"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "blogsite",
        AllowNativePasswords: true,
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        return fmt.Errorf("connectDB: %v", pingErr)
    }
    fmt.Println("DB Connected!")
    return nil
}

func main() {
    println("Trying to access database...")
    if err := connectDB(); err != nil {
        panic(err.Error())
    }
    println("Starting server.")

    http.HandleFunc("/", serveMain)
    http.HandleFunc("/postComment", handlePostComment)
    http.HandleFunc("/getPostWithID", handleGetPostWithID)
    http.HandleFunc("/getPosts", handleGetPosts)
    http.HandleFunc("/image/alepa", serveLogo)
    http.HandleFunc("/gpg", serveGPG)

    hostPort := ":8000"
    fmt.Printf("Listening on port %s", hostPort)
    log.Print(http.ListenAndServe(hostPort, nil))
    println("Bye")
}

func serveMain(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "src/index.html")
}

func handlePostComment(w http.ResponseWriter, r *http.Request) {
    // Get the comment TargetID
    postID := r.URL.Query().Get("ID")
    IntPostID, err := strconv.Atoi(postID)
    if err != nil {
        panic(err.Error())
    }
    var comment Comment
    comment.Author = r.PostFormValue("comment-author")
    comment.Text = r.PostFormValue("comment-text")
    // If user didn't input an author
    if comment.Author == "" {
        // TODO, DOESN'T WORK YET
        alertTemplate := template.Must(template.ParseFiles("src/notification.html"))
        message := "No Author given. Try again."
        print(message)
        alertTemplate.ExecuteTemplate(w, "notification-element", message)
    } else {
        if err := insertComment(IntPostID, comment); err != nil {
            panic(err.Error())
        }
        tmpl := template.Must(template.ParseFiles("src/post.html"))
        tmpl.ExecuteTemplate(w, "comment-list-element", Comment{Author: comment.Author, Text: comment.Text})
    }
}

func handleGetPostWithID(w http.ResponseWriter, r *http.Request) {
    // Get the post ID from the query
    postID := r.URL.Query().Get("ID")
    IntPostID, err := strconv.Atoi(postID)
    if err != nil {
        panic(err.Error())
    }
    // Get the post data with the ID
    postWithID, err := getPostWithID(IntPostID)
    if err != nil {
        panic(err.Error())
    }
    // Get the post's comments
    postWithID.Comments, err = getComments(IntPostID)
    if err != nil {
        panic(err.Error())
    }
    // Get the main template
    mainTmpl, err := template.ParseFiles("src/post.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        println(err.Error())
        return
    }
    // Get the post's content as HTML
    contentTmpl, err := template.ParseFiles(postWithID.LinkToPost)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        println(err.Error())
        return
    }
    // Associate the content template with the main template
    mainTmpl, err = mainTmpl.AddParseTree("content", contentTmpl.Tree)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        println(err.Error())
        return
    }
    // Render the main template
    //err = mainTmpl.ExecuteTemplate(w, "post.html", nil)
    err = mainTmpl.ExecuteTemplate(w, "post.html", postWithID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        println(err.Error())
        return
    }
}

func handleGetPosts(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("src/postsList.html"))
    postsFromDB, err := getPosts()
    if err != nil {
        panic(fmt.Errorf("handleGetPosts err: %v", err.Error()))
    }
    posts := map[string][]Post{
        "Posts": postsFromDB,
    }
    tmpl.Execute(w, posts)

}

func serveGPG(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "src/gpg.html")
}

func serveLogo(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "image/jpeg")
    imagePath := "src/alepa.jpeg"
    http.ServeFile(w, r, imagePath)
}

func getPosts() ([]Post, error) {
    var posts []Post
    rows, err := db.Query("SELECT * FROM posts;")
    if err != nil {
        return nil, fmt.Errorf("getPosts: %v", err)
    }
    for rows.Next() {
        var post Post
        if err = rows.Scan(&post.ID, &post.Author, &post.Title, &post.LinkToPost); err != nil {
            return nil, fmt.Errorf("getPosts: %v", err)
        }
        posts = append(posts, post)
    }
    return posts, nil
}

func getComments(postID int) ([]Comment, error) {
    var comments []Comment
    getCommentQuery := "SELECT * FROM comments WHERE Target=(?);"
    rows, err := db.Query(getCommentQuery, postID)
    if err != nil {
        return nil, fmt.Errorf("getComments: %v", err)
    }
    for rows.Next() {
        var comment Comment
        if err = rows.Scan(&comment.ID, &comment.Author, &comment.Text, &comment.Target, &comment.Time); err != nil {
            return nil, fmt.Errorf("getComments: %v", err)
        }
        comments = append(comments, comment)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("getComments: %v", err)
    }
    return comments, nil
}

func getPostWithID(ID int) (Post, error) {
    var post Post
    getPostQuery := "SELECT * FROM posts WHERE ID=(?);"
    row, err := db.Query(getPostQuery, ID)
    if err != nil {
        return post, fmt.Errorf("getPostWithID: %v", err)
    }
    row.Next()
    if err = row.Scan(&post.ID, &post.Author, &post.Title, &post.LinkToPost); err != nil {
        return post, fmt.Errorf("getPostWithID: %v", err)
    }
    return post, nil
}

func insertComment(targetID int, comment Comment) (error) {
    currentTime := time.Now()
    currentTimeString := currentTime.Format("01-02 15:04")
    insertquery := "INSERT INTO comments (Author, Text, Target, Time) VALUES (?, ?, ?, ?)"
    if _, err := db.Exec(insertquery, comment.Author, comment.Text, targetID, currentTimeString); err != nil {
        return fmt.Errorf("insertComment: %v", err)
    }
    return nil
}
