package server

import (
    "time"
    "net/http"
    "net/url"
    "fmt"
    "github.com/Agoudjiliss/IT-challenge/internal/configs"
)

func Run() error {
    config, err := configs.NewConfiguration()
    if err != nil {
        return fmt.Errorf("could not load configuration: %v", err)
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/ping", ping)

    cache := NewCache(30 * time.Second) // Cache expiration time (30 seconds)

    for _, resource := range config.Resources {
        targetURL, err := url.Parse(resource.Destination_URL)
        if err != nil {
            return fmt.Errorf("invalid URL %s: %v", resource.Destination_URL, err)
        }
        proxy := NewProxy(targetURL)
        mux.HandleFunc(resource.Endpoint, ProxyRequestHandler(proxy, targetURL, resource.Endpoint, cache))
        fmt.Printf("Configured proxy for endpoint %s to %s\n", resource.Endpoint, resource.Destination_URL) // Debugging line
    }

    if config.Server.CertFile != "" && config.Server.KeyFile != "" {
        return http.ListenAndServeTLS(config.Server.Host+":"+config.Server.Listen_port, config.Server.CertFile, config.Server.KeyFile, mux)
    } else {
        return http.ListenAndServe(config.Server.Host+":"+config.Server.Listen_port, mux)
    }
}
