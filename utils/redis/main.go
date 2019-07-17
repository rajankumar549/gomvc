package redis

import (
	"sync"
	"time"

	redigo "github.com/garyburd/redigo/redis"
)

// NetworkTCP for tcp
const NetworkTCP = "tcp"

// Options for redis
type Options struct {
	MaxIdle   int
	MaxActive int
	Timeout   time.Duration
	Wait      bool
}

// Redis struct
type Redis struct {
	Pool  *redigo.Pool
	mutex sync.Mutex
}

// New redis connection
func New(address, network string, opts ...Options) *Redis {
	opt := Options{
		MaxIdle:   100,
		MaxActive: 100,
		Timeout:   100,
		Wait:      true,
	}
	if len(opts) > 0 {
		opt = opts[0]
	}

	return &Redis{
		Pool: &redigo.Pool{
			MaxIdle:     opt.MaxIdle,
			MaxActive:   opt.MaxActive,
			IdleTimeout: opt.Timeout * time.Second,
			Dial: func() (redigo.Conn, error) {
				return redigo.Dial(network, address)
			},
		},
	}
}

// Del key and value
func (r *Redis) Del(key string) (int64, error) {
	conn := r.Pool.Get()
	defer conn.Close()
	return redigo.Int64(conn.Do("DEL", key))
}

// Get string value
func (r *Redis) Get(key string) (string, error) {
	conn := r.Pool.Get()
	defer conn.Close()
	return redigo.String(conn.Do("GET", key))
}

// Set key and value
func (r *Redis) Set(key, value string) (string, error) {
	conn := r.Pool.Get()
	defer conn.Close()
	return redigo.String(conn.Do("SET", key, value))
}

// SetWithNX with NX params
func (r *Redis) SetWithNX(key, value string, expire int) (string, error) {
	conn := r.Pool.Get()
	defer conn.Close()
	return redigo.String(conn.Do("SET", key, value, "NX", "EX", expire))
}

// SetNX key and value
func (r *Redis) SetNX(key, value string, expire int) (int, error) {
	conn := r.Pool.Get()
	defer conn.Close()
	return redigo.Int(conn.Do("SETNX", key, value))
}

// SetEX key and value
func (r *Redis) SetEX(key, value string, expire int) (string, error) {
	conn := r.Pool.Get()
	defer conn.Close()
	return redigo.String(conn.Do("SETEX", key, expire, value))
}

// IncrBY key and value
func (r *Redis) IncrBY(key, value string) (int64, error) {
	conn := r.Pool.Get()
	defer conn.Close()
	return redigo.Int64(conn.Do("INCRBY", key, value))
}
