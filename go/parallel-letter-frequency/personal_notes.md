# Results

### Before Concurrency

```
BenchmarkSequentialFrequency-16             3775            317140 ns/op
BenchmarkConcurrentFrequency-16             3396            349705 ns/op
PASS
ok      letter  2.464s
```

### After Concurrency

```
BenchmarkSequentialFrequency-16             3762            323061 ns/op
BenchmarkConcurrentFrequency-16             5503            211829 ns/op
PASS
ok      letter  2.446s
```

### Difference

`137876 ns/op`
