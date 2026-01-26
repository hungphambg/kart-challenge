# Story 8.4: Implement coupon code identification

**Description:**
As a developer, I want to implement logic to scan the streamed content and identify potential coupon codes, so that only relevant strings are passed to the Kafka topic.

**Acceptance Criteria:**
- [x] A regular expression or equivalent string matching logic is used to identify sequences of 8 to 10 alphanumeric characters.
- [x] Each identified sequence is considered a potential coupon code.
- [x] The identification logic is applied to each line of the decompressed content.
- [x] Leading/trailing whitespace is trimmed from potential coupon codes.

**Tasks/Subtasks:**
- [x] Modify `pipeline/producer/main.go` to use a regular expression to identify coupon codes.

**Dev Agent Record:**
- **Implementation Plan:**
  - Modified `pipeline/producer/main.go` to include the `regexp` package.
  - Defined a regular expression `\b[a-zA-Z0-9]{8,10}\b` to match alphanumeric sequences of 8 to 10 characters.
  - Applied the regex to each line of the decompressed content using `FindAllString` and trimmed whitespace from matches.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - Coupon codes are now identified from the decompressed file content.

**File List:**
- `pipeline/producer/main.go`

**Change Log:**
- Implemented coupon code identification.

**Status:**
- review