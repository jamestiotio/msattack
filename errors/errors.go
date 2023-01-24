package errors

type Error int

// Define all known error codes
const (
	SUCCESS     Error = 0
	MAINTENANCE Error = 99903004
)
