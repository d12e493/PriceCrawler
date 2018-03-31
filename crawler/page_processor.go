package crawler

import "io"

type PageProcessor interface {
	FindProduct(body io.ReadCloser)
}
