# Logging Implementation

## Overview
The terminal-hack application now includes comprehensive file-based logging for debugging and traceability.

## Features

### Automatic Log File Creation
- Logs are automatically created in the `./logs/` directory
- Each run creates a new timestamped log file: `terminal_hack_YYYY-MM-DD_HH-MM-SS.log`
- Log files include microsecond precision timestamps for detailed tracing

### Log Levels
- **INFO**: Application lifecycle, game state changes, user actions
- **DEBUG**: Detailed operation traces (cursor movements, container operations)
- **WARN**: Unusual conditions that don't prevent operation
- **ERROR**: Failures and error conditions

### What Gets Logged

#### Application Lifecycle
- Application start/stop
- Screen initialization
- Game setup parameters (word count, length, dimensions)

#### User Actions
- Cursor movements (up, down, left, right)
- Symbol selections
- ESC key press (game exit)

#### Game Logic
- Word list loading and shuffling
- Winning word selection
- Each guess validation
- Match results and remaining lives
- Game end conditions (win/loss)

#### Container Operations
- Container creation with dimensions
- Word insertion operations
- Errors during word placement

#### Cursor Operations
- Cursor initialization
- Position changes
- Boundary checks
- Symbol transitions

## Usage

The logging system is automatically initialized when the application starts:

```go
logger.Init()        // Called at application start
defer logger.Close() // Ensures logs are flushed on exit
```

### Adding Custom Logs

To add logging to new code:

```go
import "terminal_hack/internal/logger"

// Info level with key-value pairs
logger.Info("Operation completed", "key", value, "count", 42)

// Debug level for detailed tracing
logger.Debug("Processing item", "id", itemId, "name", name)

// Error level for failures
logger.Error("Operation failed", "error", err, "operation", "load")

// Warning level for concerning but non-fatal issues
logger.Warn("Unexpected condition", "value", val)
```

## Log Format

Each log entry contains:
- Date and time (with microseconds)
- Source file and line number
- Log level
- Message
- Key-value pairs

Example:
```
2026/01/18 18:30:45.123456 main.go:35: [INFO] Application starting
2026/01/18 18:30:45.234567 utilities.go:18: [DEBUG] Loading word list count=25 length=5
2026/01/18 18:30:45.345678 validator.go:23: [INFO] Validator initialized winningWord=HELLO lives=4
2026/01/18 18:30:50.456789 main.go:155: [INFO] User selected symbol symbol=WORLD symbolId=abc123
2026/01/18 18:30:50.567890 validator.go:35: [INFO] Incorrect guess guessed=WORLD match=2/5 remainingLives=3
```

## Benefits

1. **Debugging**: Trace exactly what happened during a game session
2. **User Support**: Investigate reported issues with detailed logs
3. **Performance Analysis**: Identify slow operations or bottlenecks
4. **Game Balance**: Analyze win/loss patterns and guess distributions
5. **Development**: Understand code execution flow during feature development

## Log File Management

- Log files are never automatically deleted
- Each run creates a new file (no overwriting)
- Consider setting up log rotation for long-running deployments
- Log files can be safely deleted manually when no longer needed

## Thread Safety

The logging system is thread-safe and can be called from multiple goroutines simultaneously (e.g., cursor blink goroutine, main game loop).
