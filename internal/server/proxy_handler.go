package server

import (
    "fmt"
    "net/http"
    "net/http/httputil"
    "net/url"
    "strings"
    "time"
)

// NewProxy creates a new reverse proxy for the given target URL.
func NewProxy(target *url.URL) *httputil.ReverseProxy {
    return httputil.NewSingleHostReverseProxy(target)
}

// ProxyRequestHandler processes incoming requests, checks the cache, and proxies requests to the backend server.
func ProxyRequestHandler(proxy *httputil.ReverseProxy, targetURL *url.URL, endpoint string, cache *Cache) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("[ PROXY SERVER ] Request received at %s at %s\n", r.URL, time.Now().UTC())

        // Create a unique cache key based on the request URL
        cacheKey := r.URL.Path + "?" + r.URL.RawQuery

        // Check if response is cached
        if cachedResponse, found := cache.Get(cacheKey); found {
            fmt.Printf("[CACHE HIT] Serving cached response for %s with key %s\n", r.URL, cacheKey)
            w.Write(cachedResponse)
            return
        }

        r.URL.Host = targetURL.Host
        r.URL.Scheme = targetURL.Scheme
        r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
        r.Host = targetURL.Host

        // Trim the reverseProxyRouterPrefix from the request path
        path := r.URL.Path
        r.URL.Path = strings.TrimLeft(path, endpoint)

        // Use a response recorder to capture the backend response
        recorder := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
        
        fmt.Printf("[ PROXY SERVER ] Proxying request to %s at %s\n", r.URL, time.Now().UTC())
        proxy.ServeHTTP(recorder, r)

        // Check for forbidden response
        if recorder.statusCode == http.StatusForbidden {
            fmt.Printf("[ERROR] Received Forbidden response from %s\n", targetURL)
        }

        // Cache the response
        cache.Set(cacheKey, recorder.Body)
        fmt.Printf("[CACHE MISS] Caching response for %s\n", r.URL)
    }
}

// responseRecorder struct to capture the response body
type responseRecorder struct {
    http.ResponseWriter
    statusCode int
    Body       []byte
}

// WriteHeader captures the status code of the response
func (rec *responseRecorder) WriteHeader(code int) {
    rec.statusCode = code
    rec.ResponseWriter.WriteHeader(code)
}

// Write captures the response body for caching
func (rec *responseRecorder) Write(body []byte) (int, error) {
    rec.Body = append(rec.Body, body...)
    return rec.ResponseWriter.Write(body)
}
