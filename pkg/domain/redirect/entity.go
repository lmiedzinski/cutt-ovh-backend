package redirect

type redirect struct {
	Slug string `json:"slug"`
	Url  string `json:"url"`
}

type postRedirectRequest struct {
	Url string `json:"url"`
}

type postRedirectResponse struct {
	Url  string `json:"url"`
	Slug string `json:"slug"`
}

type getRedirectBySlugResponse struct {
	Url  string `json:"url"`
	Slug string `json:"slug"`
}

type redirectErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}
