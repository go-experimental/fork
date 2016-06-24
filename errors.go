package spoon

import "fmt"

// FileDescriptorError contains a file descriptor error
type FileDescriptorError struct {
	innerError error
}

// Error returns the slave start error
func (f *FileDescriptorError) Error() string {
	return fmt.Sprint("FileDescriptor Error:" + f.innerError.Error())
}

var _ error = new(FileDescriptorError)

// ChildStartError contains the slave start error
type ChildStartError struct {
	innerError error
}

// Error returns the slave start error
func (s *ChildStartError) Error() string {
	return fmt.Sprint("Child Process Failed to Start:" + s.innerError.Error())
}

var _ error = new(ChildStartError)

// ChildShutdownError contains the slave shutdown error
type ChildShutdownError struct {
	innerError error
}

// Error returns the slave start error
func (s *ChildShutdownError) Error() string {
	return fmt.Sprint("Slave Shutdown Error:" + s.innerError.Error() + "\n\nNOTE: could just be the force termination because timeout reached")
}

var _ error = new(ChildShutdownError)

// SignalParentError contains an error regarding signaling the parent
type SignalParentError struct {
	innerError error
}

// Error returns the slave start error
func (s *SignalParentError) Error() string {
	return fmt.Sprint("Error Signaling parent:" + s.innerError.Error())
}

var _ error = new(SignalParentError)
