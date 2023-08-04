package main

import (
	"blog/internal/database"
	"docs"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"objects"
	"strconv"
	"time"
	"utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
)

// v1 handlers

// Likes a post
func (dbconfig *dbConfig) LikePostHandle(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	params, err := DecodeLikedRequestBody(r)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if LikedValidation(params) != nil {
		RespondWithError(w, http.StatusBadRequest, "data validation error")
		return
	}
	postID, err := uuid.Parse(params.PostID)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "could not decode post_id: "+params.PostID)
		return
	}
	_, err = dbconfig.DB.GetLikedByPostID(r.Context(), postID)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	feed_followed, err := dbconfig.DB.CreateLike(r.Context(), database.CreateLikeParams{
		ID:        uuid.New(),
		PostID:    postID,
		UserID:    dbUser.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusCreated, feed_followed)
}

// Gets all liked posts from the database
func (dbconfig *dbConfig) GetAllLikedHandle(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
		log.Println("Error decoding page param:", err)
	}
	liked, err := SortingQueryLiked(r, *dbconfig, page, dbUser)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(liked) == 0 {
		RespondWithJSON(w, http.StatusOK, []objects.ResponseBodyLiked{})
		return
	}
	RespondWithJSON(w, http.StatusOK, DatabaseToLiked(liked))
}

// Bookmarks a post
func (dbconfig *dbConfig) BookmarkPostHandle(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	params, err := DecodeBookmarkRequestBody(r)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if BookmarkValidation(params) != nil {
		RespondWithError(w, http.StatusBadRequest, "data validation error")
		return
	}
	postID, err := uuid.Parse(params.PostID)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "could not decode feed_id: "+params.PostID)
		return
	}
	_, err = dbconfig.DB.GetBookmarkByPostID(r.Context(), postID)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	feed_followed, err := dbconfig.DB.CreateBookmark(r.Context(), database.CreateBookmarkParams{
		ID:        uuid.New(),
		PostID:    postID,
		UserID:    dbUser.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusCreated, feed_followed)
}

// Gets all liked posts from the database
func (dbconfig *dbConfig) GetAllBookmarkHandle(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
		log.Println("Error decoding page param:", err)
	}
	bookmarks, err := SortingQueryBookmarks(r, *dbconfig, page, dbUser)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(bookmarks) == 0 {
		RespondWithJSON(w, http.StatusOK, []objects.ResponseBodyBookmark{})
		return
	}
	RespondWithJSON(w, http.StatusOK, DatabaseToBookmark(bookmarks))
}

// Unlikes a post
func (dbconfig *dbConfig) UnLikePostHandle(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	postPID := chi.URLParam(r, "postID")
	postID, err := uuid.Parse(postPID)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "could not decode postID param")
		return
	}
	post, err := dbconfig.DB.RemoveLikedByPostID(r.Context(), postID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			RespondWithError(w, http.StatusNotFound, "could not find record to remove")
			return
		}
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, post)
}

// Removes a bookmark
func (dbconfig *dbConfig) UnBookmarkPostHandle(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	postPID := chi.URLParam(r, "postID")
	postID, err := uuid.Parse(postPID)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "could not decode postID param")
		return
	}
	post, err := dbconfig.DB.RemoveBookmarkByPostID(r.Context(), postID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			RespondWithError(w, http.StatusNotFound, "could not find record to remove")
			return
		}
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, post)
}

// Searches for a liked post
func (dbconfig *dbConfig) SearchLikedHandle(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("q")
	liked, err := dbconfig.DB.SearchLikedByPostTitle(r.Context(), utils.AddSearchChar(search))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			RespondWithJSON(w, http.StatusInternalServerError, []string{})
			return
		}
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(liked) == 0 {
		RespondWithJSON(w, http.StatusOK, []database.Liked{})
		return
	}
	RespondWithJSON(w, http.StatusOK, liked)
}

// searches for a bookmark
func (dbconfig *dbConfig) SearchBookmarkHandle(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("q")
	bookmarks, err := dbconfig.DB.SearchBookmarkByPostTitle(r.Context(), utils.AddSearchChar(search))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			RespondWithJSON(w, http.StatusInternalServerError, []string{})
			return
		}
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(bookmarks) == 0 {
		RespondWithJSON(w, http.StatusOK, []database.Bookmark{})
		return
	}
	RespondWithJSON(w, http.StatusOK, bookmarks)
}

// Searches the database for simular feeds
func (dbconfig *dbConfig) SearchFeedHandle(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("q")
	feeds, err := dbconfig.DB.GetFeedSearchName(r.Context(), utils.AddSearchChar(search))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			RespondWithJSON(w, http.StatusInternalServerError, []string{})
			return
		}
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(feeds) == 0 {
		RespondWithJSON(w, http.StatusOK, []database.Feed{})
		return
	}
	RespondWithJSON(w, http.StatusOK, feeds)
}

// Searches the database for simular posts
func (dbconfig *dbConfig) SearchPostHandle(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("q")
	posts, err := dbconfig.DB.GetPostSearchTitle(r.Context(), utils.AddSearchChar(search))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			RespondWithJSON(w, http.StatusInternalServerError, []string{})
			return
		}
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(posts) == 0 {
		RespondWithJSON(w, http.StatusOK, []database.Post{})
		return
	}
	RespondWithJSON(w, http.StatusOK, DatabaseToPosts(posts))
}

// returns all posts followed by a user
func (dbconfig *dbConfig) GetPostFollowHandle(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
		log.Println("Error decoding page param")
	}
	posts, err := SortingQueryPost(r, *dbconfig, page, dbUser)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(posts) == 0 {
		RespondWithJSON(w, http.StatusOK, []database.Post{})
		return
	}
	RespondWithJSON(w, http.StatusOK, DatabaseToPosts(posts))
}

// returns all feeds followed by a user
func (dbconfig *dbConfig) GetFeedFollowHandle(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	feeds, err := dbconfig.DB.GetUserFeedFollows(r.Context(), dbUser.ID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(feeds) == 0 {
		RespondWithJSON(w, http.StatusOK, []database.Feed{})
		return
	}
	RespondWithJSON(w, http.StatusOK, feeds)
}

// unfollows a feed
func (dbconfig *dbConfig) UnFollowFeedHandle(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	feedFID := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFID)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "could not decode feedFollowID param")
		return
	}
	feed, err := dbconfig.DB.DeleteFeedByID(r.Context(), database.DeleteFeedByIDParams{
		ID:     feedFollowID,
		UserID: dbUser.ID,
	})
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			RespondWithError(w, http.StatusNotFound, "could not find record to remove")
			return
		}
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, feed)
}

// follows a feed
func (dbconfig *dbConfig) CreateFeedFollowHandle(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	params, err := DecodeFeedFollowRequestBody(r)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if FeedFollowValidation(params) != nil {
		RespondWithError(w, http.StatusBadRequest, "data validation error")
		return
	}
	_, err = dbconfig.CheckFeedIDExist(params.FeedID, r)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	_, err = dbconfig.CheckFeedFollowExist(params.FeedID, dbUser, r)
	if err != nil {
		RespondWithError(w, http.StatusMethodNotAllowed, err.Error())
		return
	}
	feedId, err := uuid.Parse(params.FeedID)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "could not decode feed_id: "+params.FeedID)
		return
	}
	feed_followed, err := dbconfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    feedId,
		UserID:    dbUser.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusCreated, feed_followed)
}

// GETs all feeds from the database
func (dbconfig *dbConfig) GetAllFeedsHandle(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
		log.Println("Error decoding page param:", err)
	}
	feeds, err := SortingQueryFeed(r, *dbconfig, page)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(feeds) == 0 {
		RespondWithJSON(w, http.StatusOK, []objects.ResponseBodyFeed{})
		return
	}
	RespondWithJSON(w, http.StatusOK, DatabaseToFeeds(feeds))
}

// creates a feed for a specific user
func (dbconfig *dbConfig) CreateFeedHandler(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	params, err := DecodeFeedRequestBody(r)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if FeedValidation(params) != nil {
		RespondWithError(w, http.StatusBadRequest, "data validation error")
		return
	}
	_, err = dbconfig.CheckFeedURLExist(params.URL, r)
	if err != nil {
		RespondWithError(w, http.StatusConflict, err.Error())
		return
	}
	feed, err := dbconfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Url:       params.URL,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    dbUser.ID,
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	feed_followed, err := dbconfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    feed.ID,
		UserID:    dbUser.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	response := struct {
		Feed       objects.ResponseBodyFeed `json:"feed"`
		FeedFollow database.FeedFollow      `json:"feed_follow"`
	}{
		Feed:       DatabaseToFeed(feed),
		FeedFollow: feed_followed,
	}
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusCreated, response)
}

// returns a user with the specific ApiKey
func (dbconfig *dbConfig) GetUserByApiKeyHandle(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	RespondWithJSON(w, http.StatusAccepted, dbUser)
}

// creates a new user
func (dbconfig *dbConfig) CreateUserHandle(w http.ResponseWriter, r *http.Request) {
	params, err := DecodeUserRequestBody(r)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if UserValidation(params) != nil {
		RespondWithError(w, http.StatusBadRequest, "data validation error")
		return
	}
	_, err = dbconfig.CheckUserExist(params.Name, r)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	dbUser, err := dbconfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusCreated, dbUser)
}

// Displays available endpoints in json format
func (dbconfig *dbConfig) Endpoints(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, http.StatusOK, docs.Endpoints())
}

// test for RespondWithJSON
func ReadiHandle(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, 200, objects.ReadyHandle{
		Status: "ok",
	})
}

// error handler
func ErrHandle(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, 200, "Internal Server Error")
}

// middleware that determines which headers, http methods and orgins are allowed
func MiddleWare() cors.Options {
	return cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	}
}

// helper function
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}

// helper function
func RespondWithError(w http.ResponseWriter, code int, msg string) error {
	return RespondWithJSON(w, code, map[string]string{"error": msg})
}

// helper function
// checks if a username already exists inside database
func (dbconfig *dbConfig) CheckUserExist(name string, r *http.Request) (bool, error) {
	username, err := dbconfig.DB.GetUserByName(r.Context(), name)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return false, err
		}
	}
	if username.Name == name {
		return true, errors.New("name already exists")
	}
	return false, nil
}

// helper function
// checks if a feed with the respective URL already exists
func (dbconfig *dbConfig) CheckFeedURLExist(feedUrl string, r *http.Request) (bool, error) {
	feed, err := dbconfig.DB.GetFeedByURL(r.Context(), feedUrl)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return false, err
		}
	}
	if feed.Url == feedUrl {
		return true, errors.New("feed id '" + feedUrl + "' already exists found")
	}
	return false, nil
}

// helper function
// checks if a feed with id exists
func (dbconfig *dbConfig) CheckFeedIDExist(feedID string, r *http.Request) (bool, error) {
	feedId, err := uuid.Parse(feedID)
	if err != nil {
		return false, err
	}
	feed, err := dbconfig.DB.GetFeed(r.Context(), feedId)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return false, err
		}
	}
	if feed.ID == feedId {
		return true, nil
	}
	return false, errors.New("feed url '" + feedID + "' not found")
}

// helper function
// checks if a user is already following the feed
func (dbconfig *dbConfig) CheckFeedFollowExist(feedID string, dbUser database.User, r *http.Request) (bool, error) {
	feedId, err := uuid.Parse(feedID)
	if err != nil {
		return false, err
	}
	queryParams := database.GetFeedFollowParams{
		FeedID: feedId,
		UserID: dbUser.ID,
	}
	feed_followed, err := dbconfig.DB.GetFeedFollow(r.Context(), queryParams)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return false, err
		}
	}
	if feed_followed.UserID == dbUser.ID &&
		feed_followed.FeedID == feedId {
		return true, errors.New("user cannot follow a feed more than once")
	}
	return false, nil
}

// helper function
// returns a separate error message instead of the SQL version
func SQLErrorWrapper(err error) error {
	if err.Error() == "sql: no rows in result set" {
		return errors.New("could not find record")
	}
	return err
}

// helper function
func SortingQueryFeed(r *http.Request, dbconfig dbConfig, page int) ([]database.Feed, error) {
	sort := r.URL.Query().Get("sort")
	if utils.FindSortParam(sort) == "acs" {
		feeds, err := dbconfig.DB.GetFeedsAsc(r.Context(), database.GetFeedsAscParams{
			Limit:  10,
			Offset: int32((page - 1) * 10),
		})
		return feeds, err
	}
	feeds, err := dbconfig.DB.GetFeedsDesc(r.Context(), database.GetFeedsDescParams{
		Limit:  10,
		Offset: int32((page - 1) * 10),
	})
	return feeds, err
}

// helper function
func SortingQueryPost(r *http.Request, dbconfig dbConfig, page int, dbUser database.User) ([]database.Post, error) {
	sort := r.URL.Query().Get("sort")
	if utils.FindSortParam(sort) == "asc" {
		posts, err := dbconfig.DB.GetPostsByUserAsc(r.Context(), database.GetPostsByUserAscParams{
			Limit:  10,
			Offset: int32((page - 1) * 10),
			UserID: dbUser.ID,
		})
		return posts, err
	}
	posts, err := dbconfig.DB.GetPostsByUserDesc(r.Context(), database.GetPostsByUserDescParams{
		Limit:  10,
		Offset: int32((page - 1) * 10),
		UserID: dbUser.ID,
	})
	return posts, err
}

// helper function
func SortingQueryLiked(r *http.Request, dbconfig dbConfig, page int, dbUser database.User) ([]database.Liked, error) {
	sort := r.URL.Query().Get("sort")
	if utils.FindSortParam(sort) == "asc" {
		posts, err := dbconfig.DB.GetLikedPostByUserAsc(r.Context(), database.GetLikedPostByUserAscParams{
			Limit:  10,
			Offset: int32((page - 1) * 10),
			UserID: dbUser.ID,
		})
		return posts, err
	}
	posts, err := dbconfig.DB.GetLikedPostByUserDesc(r.Context(), database.GetLikedPostByUserDescParams{
		Limit:  10,
		Offset: int32((page - 1) * 10),
		UserID: dbUser.ID,
	})
	return posts, err
}

// helper function
func SortingQueryBookmarks(r *http.Request, dbconfig dbConfig, page int, dbUser database.User) ([]database.Bookmark, error) {
	sort := r.URL.Query().Get("sort")
	if utils.FindSortParam(sort) == "asc" {
		posts, err := dbconfig.DB.GetBookmarkPostByUserAsc(r.Context(), database.GetBookmarkPostByUserAscParams{
			Limit:  10,
			Offset: int32((page - 1) * 10),
			UserID: dbUser.ID,
		})
		return posts, err
	}
	posts, err := dbconfig.DB.GetBookmarkPostByUserDesc(r.Context(), database.GetBookmarkPostByUserDescParams{
		Limit:  10,
		Offset: int32((page - 1) * 10),
		UserID: dbUser.ID,
	})
	return posts, err
}
