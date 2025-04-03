package main

import (
	"fmt"
	"time"
)

type Database interface {
	Query(string) string
}

type RealDatabase struct {
	table []string
}

func (db *RealDatabase) Query(query string) string {
	// Simulate slow database operation
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("Results for: %s", query)
}

// ---- Proxy ----
type DatabaseProxy struct {
	db      *RealDatabase
	cache   map[string]string
	logging bool
}

func NewDatabaseProxy(logging bool) *DatabaseProxy {
	return &DatabaseProxy{
		cache:   make(map[string]string),
		logging: logging,
	}
}

func (p *DatabaseProxy) Query(query string) string {
	// Check cache
	// If cache hit return value from cache
	// If cache miss then initialization db object
	// then forward the query to db
	// finally update the cache

	// check cache
	if value, isPresent := p.cache[query]; isPresent {
		if p.logging {
			fmt.Printf("Cache hit for: %s\n", query)
		}
		return value
	}

	// cache miss then Lazy initialization
	if p.db == nil {
		p.db = &RealDatabase{}
		if p.logging {
			fmt.Println("[PROXY] Initialized database")
		}
	}

	// forward query to real database
	if p.logging {
		fmt.Printf("[PROXY] Forwarding query: %s\n", query)
	}
	results := p.db.Query(query)

	// update the cache
	p.cache[query] = results
	return results
}

func main() {
	// Client code works with Database interface
	var db Database = NewDatabaseProxy(true)

	fmt.Println(db.Query("SELECT * FROM users")) // First time -> slow
	fmt.Println(db.Query("SELECT * FROM users")) // Cache hit -> fast
}
