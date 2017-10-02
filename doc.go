// Package servicenow currently provides primarily low-level abstractions for
// communicating with the ServiceNow API, and data types useful for
// unmarshaling the nonstandard JSON types that ServiceNow returns (e.g.
// non-ISO datetime format). As needs arise, higher-level abstractions may be
// added that provide easy access to standard structured objects in the
// ServiceNow API, for example a method to easily access structs representing
// change requests.
package servicenow
