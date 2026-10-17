package cache

import "sync"

// Cache — потокобезопасный кеш строковых ключей и значений.
//
// Все методы обязаны быть безопасны при одновременном вызове из
// нескольких горутин: никаких гонок данных и потерянных обновлений.
type Cache struct {
	mu    sync.RWMutex
	items map[string]string
}

// New возвращает пустой кеш, готовый к работе.
func New() *Cache {
	return &Cache{items: make(map[string]string)}
}

// Get возвращает значение по ключу и true, либо пустую строку и false,
// если ключа нет.
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.items[key]
	return v, ok
}

// Set сохраняет значение по ключу. Повторный Set по тому же ключу
// заменяет значение, но не увеличивает размер кеша.
func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = value
}

// Len возвращает число различных ключей в кеше.
func (c *Cache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}
