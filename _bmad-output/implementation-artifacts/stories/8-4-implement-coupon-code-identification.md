# Story 8.4: Implement coupon code identification

**Epic:** Coupon Producer Service

**Description:**
As a developer, I want to implement logic to scan the streamed content and identify potential coupon codes, so that only relevant strings are passed to the Kafka topic.

**Acceptance Criteria:**
- A regular expression or equivalent string matching logic is used to identify sequences of 8 to 10 alphanumeric characters.
- Each identified sequence is considered a potential coupon code.
- The identification logic is applied to each line of the decompressed content.
- Leading/trailing whitespace is trimmed from potential coupon codes.
