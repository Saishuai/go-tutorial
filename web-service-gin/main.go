package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type albumResponse struct {
	ErrorCode int32   `json:"error_code"`
	Albums    []album `json:"albums"`
}

type getAlbumResponse struct {
	ErrorCode int32  `json:"error_code"`
	Album     *album `json:"album"`
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Geryy Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vauhan and Cliffford Brown", Artist: "Sarah Vaughan", Price: 17.99},
}

func (a *albumResponse) successResp() {
	a.ErrorCode = 0
	a.Albums = append(a.Albums, albums...)
}

func albumListTransferToMap() map[string]album {
	albumMap := make(map[string]album, 16)
	for _, v := range albums {
		if _, ok := albumMap[v.ID]; !ok {
			albumMap[v.ID] = v
		}
	}
	return albumMap
}

func getAlbumById(id string) (*album, error) {
	albumMap := albumListTransferToMap()
	if albumVal, exist := albumMap[id]; exist {
		return &albumVal, nil
	}
	return nil, errors.New("this id not exist")

}

func getAlbumRequest(c *gin.Context) {
	resp := &getAlbumResponse{
		ErrorCode: 101,
		Album:     nil,
	}

	id := c.Param("id")
	albumVal, err := getAlbumById(id)
	if err != nil {
		c.IndentedJSON(http.StatusOK, resp)
		log.Panic(err)
	} else {
		resp.Album = albumVal
		resp.ErrorCode = 0
		c.IndentedJSON(http.StatusOK, resp)
	}
}

func getAlbums(c *gin.Context) {
	resp := &albumResponse{
		ErrorCode: 101,
		Albums:    make([]album, 0),
	}
	resp.successResp()
	c.IndentedJSON(http.StatusOK, resp)
}

func postAlbums(c *gin.Context) {
	resp := &albumResponse{
		ErrorCode: 101,
		Albums:    make([]album, 0),
	}
	var newAlbums album
	// bind new albums in request to new ablum struct
	if err := c.BindJSON(&newAlbums); err != nil {
		c.IndentedJSON(http.StatusOK, resp)
		log.Panic("bind error, please check struct for new album")
	} else {
		// getAlbumById() return album by id, set unique id with this judge
		if exist, _ := getAlbumById(newAlbums.ID); exist != nil {
			c.IndentedJSON(http.StatusOK, resp)
			log.Panic("new album id has been used, pls change to another one")
		} else {
			albums = append(albums, newAlbums)
			resp.successResp()
			c.IndentedJSON(http.StatusOK, resp)
		}
	}
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumRequest)
	router.POST("/add-albums", postAlbums)
	// router.POST("/get-album-by-id", getAlbumRequest)

	router.Run("localhost:8080")
}
