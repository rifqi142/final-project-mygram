# ⚙️ Final Project - MyGramApp API ⚙️

### ⚡️Scalable Web Services with Go - Digitalent ✕ Hacktiv8 ⚡️


MyGram is a photo-sharing application built using the Go programming language. It allows users to freely share, view, and comment on photos posted by others. Registration is open to anyone, and individuals can create an account by providing an email address and selecting a username.


## 🧑🏻‍💻 Author

- [@Muhammad Rifqi Setiawan](https://github.com/rifqi142)


## 📚 Tech Backend

- Golang
- Gin Gonic
- JWT-GO
- Gorm


##  📲End Point

### User
- POST /users/register: Register a new user.
- POST /users/login: Log in an existing user.
- PUT /users/: Update user data. Requires authentication.
- DELETE /users/: Delete user account. Requires authentication.


### Photos
- POST /photos/: Upload a new photo.
- GET /photos/: Get all photos.
- GET /photos/:PhotoId: Get a photo by ID.
- PUT /photos/:PhotoId: Update photo data.
- DELETE /photos/:PhotoId: Delete a photo.

### Comments
- POST /comments/: Create a new comment.
- GET /comments/: Get all comments.
- GET /comments/:CommentId: Get a comment by ID.
- PUT /comments/:CommentId: Update comment data.
- DELETE /comments/:CommentId: Delete a comment.

### Social Medias
- POST /socialmedias/: Create a new social media entry.
- GET /socialmedias/: Get all social media entries.
- GET /socialmedias/:SocialMediaId: Get a social media entry by ID.
- PUT /socialmedias/:SocialMediaId: Update social media entry data.
- DELETE /socialmedias/:SocialMediaId: Delete a social media entry.