# Go Data Intensive

The purpose of this project is to study in practice some of the key attributes that make 
data-intensive APIs

### Some attributes of data intensive

- High volume of traffic, thousands requests per second
- High volume of data being written or read
- High volume of data processing like indexing
- High volume of data being frequently updated

### Common functional requirements

- Consistency
- Latency
- Scalability

- Check more ilities in Neil book:
> (Ilities)

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

- Atomicity and Eventual Consistency
  > - When you have to update some data in multiple layers of your system, eg:
  > cache, database, search index, you them all to succeed or fail, so that you don't have
  > inconsistency between the layers.
  > - Deal with updates as messages, send the messages to a "queue" and have consumers for each layer to
  > read the messages and process, if it fails, the message will still be there.
  > - Eventually, all layers will be updated, distributed.
  > - Flag messages with UUIDs to track them in the flow, set order and handle duplication.

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
  > - "Graceful Shutdown / Restart" to avoid abrupt exiting to cause inconsistency (data left unprocessed, pending)

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

- Auto Scaling (Horizontal)
  > - By observing the system metrics, specially  under stress with load testing,
  > it will be possible to establish the thresholds for auto scaling:
  > - Eg.: Add 2 units when CPU > 100% for 30 seconds.
  > - Eg.: Decrease 2 units when CPU is 50% idle for 30 seconds.
  > - It's important to fine tune these numbers and consider the uptime required to add units,
  > in case of a bursts of requests, the system might become unavailable before the up-scaling finishes

- Flexible architecture
  > - Strong contracts to easily replace less performant pieces and dependencies by more optimized

- [WIP...]
