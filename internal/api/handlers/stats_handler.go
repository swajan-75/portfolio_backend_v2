package handlers

import (
	"encoding/json"
	"net/http"

	"firebase.google.com/go/v4/db"
	"github.com/gin-gonic/gin"
)
type StatsHandler struct{
	DB *db.Client
}



func NewStatsHandler(db *db.Client) *StatsHandler{
	return &StatsHandler{DB: db}
}

func (h *StatsHandler) Track_Visitor(c *gin.Context) {
    _, err := c.Cookie("visited")

    ip := c.ClientIP()
    var geo GeoData


    resp, errHttp := http.Get("http://ip-api.com/json/" + ip)
    if errHttp == nil {
        defer resp.Body.Close()
        json.NewDecoder(resp.Body).Decode(&geo)
    }


    if err != nil {
        ctx := c.Request.Context()
        country := "Unknown"
        if geo.Country != "" {
            country = geo.Country
        }


        h.DB.NewRef("stats/unique_visitors").Transaction(ctx, func(node db.TransactionNode) (interface{}, error) {
            var count int64
            node.Unmarshal(&count)
            return count + 1, nil
        })

        countryRef := h.DB.NewRef("stats/countries/" + country)
        countryRef.Transaction(ctx, func(node db.TransactionNode) (interface{}, error) {
            var count int64
            node.Unmarshal(&count)
            return count + 1, nil
        })


        c.SetCookie("visited", "true", 31536000, "/", "", true, true)
    }

    // c.JSON(http.StatusOK, gin.H{
    //     "message": "Visitor tracking processed",
    //     "is_new":  err != nil, // true if it's their first time
    //     "geo":     geo,         // This contains the full ip-api response
    //     "ip":      ip,
    // })

	c.Status(http.StatusOK)
}

func (h *StatsHandler) TrackDownload(c *gin.Context) {
    ref := h.DB.NewRef("stats/cv_downloads")
    ref.Transaction(c.Request.Context(), func(node db.TransactionNode) (interface{}, error) {
        var count int
        node.Unmarshal(&count)
        return count + 1, nil
    })
    c.Status(http.StatusOK)
}

func (h *StatsHandler) GetStats(c *gin.Context) {
    ctx := c.Request.Context()
    
    ref := h.DB.NewRef("stats")
    var stats map[string]interface{}
    
    if err := ref.Get(ctx, &stats); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch statistics"})
        return
    }

    c.JSON(http.StatusOK, stats)
}
type GeoData struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Query       string  `json:"query"` 
}

