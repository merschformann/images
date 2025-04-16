package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var lock sync.Mutex

func AddItem(c *gin.Context) {
	var newItem Item

	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lock.Lock()
	defer lock.Unlock()
	newItem.ID = len(DB)
	DB = append(DB, newItem)

	c.JSON(http.StatusCreated, newItem)
}

func GetItem(c *gin.Context) {
	id := c.Param("id")
	itemID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	if itemID < 0 || itemID >= len(DB) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	item := DB[itemID]

	c.JSON(http.StatusOK, item)
}

func ListItems(c *gin.Context) {
	offset, count := 0, 10
	pCount := c.Param("count")
	var err error
	if pCount != "" {
		count, err = strconv.Atoi(pCount)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid count parameter"})
			return
		}
	}
	pOffset := c.Param("offset")
	if pOffset != "" {
		offset, err = strconv.Atoi(pOffset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
			return
		}
	}

	if offset < 0 || count < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset or count parameter"})
		return
	}

	if offset >= len(DB) {
		c.JSON(http.StatusOK, []Item{})
		return
	}
	if offset+count > len(DB) {
		count = len(DB) - offset
	}

	items := make([]Item, count)
	copy(items, DB[offset:offset+count])
	c.JSON(http.StatusOK, items)
}
