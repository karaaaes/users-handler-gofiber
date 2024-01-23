package request

type ComicRequest struct {
	SeriesName string `json:"series_name"`
	Author     string `json:"author"`
	Cover      string `json:"cover"`
	CoverName  string `json:"cover_name"`
}
