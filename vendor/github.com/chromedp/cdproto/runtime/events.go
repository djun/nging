package runtime

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"github.com/mailru/easyjson"
)

// EventBindingCalled notification is issued every time when binding is
// called.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#event-bindingCalled
type EventBindingCalled struct {
	Name               string             `json:"name"`
	Payload            string             `json:"payload"`
	ExecutionContextID ExecutionContextID `json:"executionContextId"` // Identifier of the context where the call was made.
}

// EventConsoleAPICalled issued when console API was called.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#event-consoleAPICalled
type EventConsoleAPICalled struct {
	Type               APIType            `json:"type"`                 // Type of the call.
	Args               []*RemoteObject    `json:"args"`                 // Call arguments.
	ExecutionContextID ExecutionContextID `json:"executionContextId"`   // Identifier of the context where the call was made.
	Timestamp          *Timestamp         `json:"timestamp"`            // Call timestamp.
	StackTrace         *StackTrace        `json:"stackTrace,omitempty"` // Stack trace captured when the call was made. The async stack chain is automatically reported for the following call types: assert, error, trace, warning. For other types the async call chain can be retrieved using Debugger.getStackTrace and stackTrace.parentId field.
	Context            string             `json:"context,omitempty"`    // Console context descriptor for calls on non-default console context (not console.*): 'anonymous#unique-logger-id' for call on unnamed context, 'name#unique-logger-id' for call on named context.
}

// EventExceptionRevoked issued when unhandled exception was revoked.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#event-exceptionRevoked
type EventExceptionRevoked struct {
	Reason      string `json:"reason"`      // Reason describing why exception was revoked.
	ExceptionID int64  `json:"exceptionId"` // The id of revoked exception, as reported in exceptionThrown.
}

// EventExceptionThrown issued when exception was thrown and unhandled.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#event-exceptionThrown
type EventExceptionThrown struct {
	Timestamp        *Timestamp        `json:"timestamp"` // Timestamp of the exception.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails"`
}

// EventExecutionContextCreated issued when new execution context is created.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#event-executionContextCreated
type EventExecutionContextCreated struct {
	Context *ExecutionContextDescription `json:"context"` // A newly created execution context.
}

// EventExecutionContextDestroyed issued when execution context is destroyed.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#event-executionContextDestroyed
type EventExecutionContextDestroyed struct {
	ExecutionContextID ExecutionContextID `json:"executionContextId"` // Id of the destroyed context
}

// EventExecutionContextsCleared issued when all executionContexts were
// cleared in browser.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#event-executionContextsCleared
type EventExecutionContextsCleared struct{}

// EventInspectRequested issued when object should be inspected (for example,
// as a result of inspect() command line API call).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#event-inspectRequested
type EventInspectRequested struct {
	Object             *RemoteObject       `json:"object"`
	Hints              easyjson.RawMessage `json:"hints"`
	ExecutionContextID ExecutionContextID  `json:"executionContextId,omitempty"` // Identifier of the context where the call was made.
}
