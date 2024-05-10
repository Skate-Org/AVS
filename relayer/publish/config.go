package publish

import (
// avsMemcache "github.com/Skate-Org/AVS/relayer/db/avs/mem"
// skateappMemcache "github.com/Skate-Org/AVS/relayer/db/skateapp/mem"
)

var Verbose = true

// NOTE: use shared in-mem cache if facing performance is a bottleneck.
// To be consider in future versions
// taskCache     = skateappMemcache.NewCache(100 * 1024 * 1024) // 100MB
// operatorCache = avsMemcache.NewCache(2 * 1024 * 1024)        // 2MB
