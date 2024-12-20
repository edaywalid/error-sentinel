
# Error Sentinel

## Overview
Error Sentinel is a Go-based web server designed to gracefully handle errors and simplify debugging during development. The server features middleware for error recovery, automatic retries, a health check endpoint, and detailed debugging tools for stack trace analysis.

---

## Features

1. **Error Recovery**:
   - Middleware captures unhandled exceptions (panics) and logs them.
   - Automatically recovers from panics to ensure continued operation.
   - Different error responses based on environment:
     - **Production**: Displays a generic message: "Something went wrong."
     - **Development**: Displays a stack trace with clickable links to source files.

2. **Source File Viewer**:
   - Highlights the line of code causing the panic.
   - Automatically scrolls to the highlighted line when the page loads.

3. **Custom Routes**:
   - `/panic`: Triggers a runtime panic for testing.
   - `/error`: Simulates an HTTP 500 error.
   - `/retry`: Implements automatic retry logic (e.g., retries 3 times before failing).

4. **Health Check Endpoint**:
   - `/health`: Reports the server's operational status.

---

## File Structure

```
error-sentinel/
├── cmd/
│   └── main.go                # Entry point for the server
├── internal/
│   ├── handlers/
│   │   ├── error.go         # Handlers for /panic and /error
│   │   ├── health.go        # Handler for /health
│   │   └── source.go        # Handler for viewing source files
│   ├── middleware/
│   │   └── recovery.go     # Middleware for panic recovery
│   ├── routes/
│   │   └── router.go       # Router configuration
│   ├── utils/
│   │   └── stacktrace.go   # Formatting stack traces
├── go.mod                       # Go module file
├── go.sum                       # Go module dependencies
└── README.md                   # Project documentation
```

---

## Getting Started

### Prerequisites
- **Go 1.23+** installed on your system

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/error-sentinel.git
   cd error-sentinel
   ```

2. Add .env file based on .env.example

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the server:
   ```bash
    make run 
   ```

---

## Usage

### Endpoints

#### 1. `/panic`
- **Description**: Triggers a runtime panic.
- **Purpose**: Test panic recovery and stack trace logging.

#### 2. `/error`
- **Description**: Simulates a server error (HTTP 500).
- **Purpose**: Test custom error handling.

#### 3. `/retry`
- **Description**: Implements automatic retry logic (e.g., retries 3 times before returning a failure).
- **Purpose**: Demonstrate retry middleware.

#### 4. `/health`
- **Description**: Reports the server's health status.
- **Response**: text indicating lifetime :
  ```txt
  OK - Uptime: 9.6042804s
  ```

#### 5. `/source`
- **Description**: Displays source files with syntax highlighting and line numbers.
- **Parameters**:
  - `file`: Path to the source file.
  - `line`: Line number to highlight.
- **Example**:
  ```
  /source?file=internal/handlers/source.go&line=42
  ```

---

## Technical Details

### Error Recovery Middleware
- Captures panics and logs stack traces.
- Switches behavior based on the environment (production vs. development).

### Stack Trace Formatting
- Converts stack traces into clickable links for easier debugging.
- Highlights the specific line causing the error.
- Removes instruction offsets (e.g., `+0x4b4`).

### Source File Viewer
- Highlights the error line in yellow.
- Scrolls to the highlighted line using `scrollIntoView` in JavaScript.


