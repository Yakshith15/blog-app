package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type BlogClient struct {
	baseURL       string
	internalToken string
	httpClient    *http.Client
}

func NewBlogClient(baseURL string, internalToken string, timeout time.Duration) *BlogClient {
	return &BlogClient{
		baseURL:       baseURL,
		internalToken: internalToken,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// BlogExists checks whether a blog exists by calling Blog Service
func (c *BlogClient) BlogExists(blogID uuid.UUID) (bool, error) {

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/internal/blogs/%s", c.baseURL, blogID.String()),
		nil,
	)
	if err != nil {
		return false, err
	}

	// Internal service auth
	req.Header.Set("X-Internal-Token", c.internalToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return true, nil
	case http.StatusNotFound:
		return false, nil
	case http.StatusUnauthorized:
		return false, fmt.Errorf("unauthorized to call blog service")
	default:
		return false, fmt.Errorf("blog service error: %d", resp.StatusCode)
	}
}
