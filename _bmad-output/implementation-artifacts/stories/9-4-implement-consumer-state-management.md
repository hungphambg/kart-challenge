# Story 9.4: Implement consumer state management

**Epic:** Coupon Consumer Service

**Description:**
As a developer, I want to implement (or leverage the stream-processing library's feature for) state management to track which Kafka topics/source files each unique coupon code has been observed in, so that we can determine its validity.

**Acceptance Criteria:**
- The consumer maintains a data structure (e.g., a hash map, or a KTable/GlobalKTable if using a stream processing library) where:
  - The key is the coupon code string.
  - The value tracks the unique source files (e.g., `couponbase1.gz`, `couponbase2.gz`, `couponbase3.gz`) in which the code has been seen. A bitmask or a set of strings could be used for this value.
- This state is updated continuously as messages are consumed from Kafka.
- The state management mechanism is fault-tolerant, ensuring that the state is not lost upon consumer restarts, leveraging the chosen stream-processing library's capabilities.
```

I will now create this file.