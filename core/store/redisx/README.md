## Redis

### Redis Lock
```
firstLock := NewRedisLock(client, key)
firstLock.SetExpire(30)
firstAcquire, err := firstLock.Acquire()
assert.Nil(t, err)
assert.True(t, firstAcquire)

release, err := firstLock.Release()
assert.Nil(t, err)
assert.True(t, release)
```