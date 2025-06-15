package wsdl_to_dto

// Dereference a string pointer and handle null strings
func deref(s *string) string {
	if s == nil {
		return "<nil>"
	}
	return *s
}
