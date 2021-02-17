# Go Data Intensive

The purpose of this project is to study in practice some of the key attributes that make 
data-intensive APIs

### Attributes

- High volume of traffic, thousands requests per second
- High volume of data being written
- High volume of data processing/indexing
- High volume of data changing/updated

### Notes

- Caching:
  > - Cache everything immutable, and establish long TTLs for data 
  > that changes less frequently

- Performance optimization:
  > - Use appropriate data structures 
  > - Reduce cyclomatic complexity (Big O), avoid expensive operations 
  > - Minimize IO read / write and network calls (latency)
  > - Do things asynchronously and concurrently
  > - Set timeouts when talking to external dependencies

- Expensive operations in background:
  > - Delegate expensive operations, or, part of expensive operations 
  > to be executed and cached in background, without impacting response times

- Reduce input:
  > - Reduce the input by filtering / ignoring inputs that has been invalidated

- Idempotent logic
  > - Enqueued messages consumed concurrently might be read by multiple consumers, 
  > there must be mechanisms in place to guarantee consistency, regardless of the executions
  > - Use flags, timestamps, acknowledge messages ensure consistency

- Load balancing:
  > - One consumer cannot have it's performance degraded while another is idle
  > - Messages should be distributed evenly between queue groups / streams
  > - Cache keys should be distributed evenly between nodes in a cluster (consistent hashing)
  > - Data should be distributed evenly between database nodes/replicas

- Failure / Reliability:
  > - Clear separation of domain and infrastructure errors
  > - Log and capture relevant errors with all the context needed for debugging (where?, what?, why?)
  > - Retry with exponential backoffs when dependencies performance are degraded
  > - Fallback mechanisms when dependencies are unavailable, to avoid interruption of service
  > - Circuit breakers can be implemented to automatically identify dependencies that are unhealthy
  > and toggle fallbacks, they can redo when the dependencies are healthy again (auto-healing)
  > - Send messages that cannot be processed to a dead letter queue to avoid data-loss

- Observability:
  > - Track metrics to understand how the system is behaving
  > - End 2 End tracings to discover bottlenecks and spots to improve
  
- Performance Testing:
  > - How many requests per second a single unit can handle
  > - How much is scales, doubling the units can respond twice the requests?
  > - How dependencies behave under stress (database, cache..)
  > - How effective is you caching, hits vs misses
  > - Tweaking configs like, increasing number of workers and decreasing the timing they run
  > - Available memory and garbage collection effectiveness (cleaning more space than it's being written)
  > - How complex queries in DB behave during stress, if writing under stress impacts reading

- Flexible architecture
  > - Strong contracts to easily replace less performant pieces and dependencies by more optimized

- [WIP...]
