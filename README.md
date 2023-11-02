# gowinreg 💨💿⚙

![win](https://github.com/Donny-GUI/gowinreg/assets/108424001/601b9a71-2155-41f0-871b-99e0eacda537)


# Windows Registry Access in Go  📝

`winreg` is a Go package that provides easy access to the Windows Registry, allowing you to perform various operations programmatically. This package simplifies tasks such as reading and writing registry values, deleting keys, checking the existence of keys and values, and more.

## Features

- Read and write string values.
- Read and write DWORD (32-bit integer) values.
- Read and write binary values.
- Read and write expandable string (REG_EXPAND_SZ) values.
- Read and write 32-bit and 64-bit integer values.
- Delete registry values and keys.
- Check the existence of registry keys and values.
- Enumerate subkeys and value names.
- Read multi-string values (REG_MULTI_SZ).

## Installation

To use the `winreg` package, you need to install the "golang.org/x/sys/windows/registry" package first:

```bash
go get golang.org/x/sys/windows/registry

```

After that, you can import winreg into your Go code:
```go
import (
    "golang.org/x/sys/windows/registry"
    "github.com/Donny-GUI/gowinreg"
)
```

## Usage
Here's an example of how to use the winreg package:

```go
package main

import (
    "fmt"
    "golang.org/x/sys/windows/registry"
    "github.com/yourusername/winreg"
)

func main() {
    root := registry.CURRENT_USER
    keyPath := `Software\MyApp`

    // Write a DWORD value
    if err := winreg.WriteDWordValue(root, keyPath, "Setting2", 42); err != nil {
        fmt.Println("Error writing DWORD registry value:", err)
    }

    // Read a DWORD value
    dwordValue, err := winreg.ReadDWordValue(root, keyPath, "Setting2")
    if err != nil {
        fmt.Println("Error reading DWORD registry value:", err)
    } else {
        fmt.Printf("Setting2 = %d\n", dwordValue)
    }

}
```
## Functions


`ReadDWordValue`: Reads a DWORD value from the Windows Registry.


`WriteDWordValue`: Writes a DWORD value to the Windows Registry.


`ReadBinaryValue`: Reads a binary value from the Windows Registry.


`WriteBinaryValue`: Writes a binary value to the Windows Registry.


`DeleteValue`: Deletes a registry value.


`DeleteSubKey`: Deletes a registry subkey and all its subkeys and values.


`KeyExists`: Checks if a registry key exists.


`ValueExists`: Checks if a registry value exists.


`EnumerateSubKeys`: Returns a list of subkeys under the given key.


`EnumerateValues`: Returns a list of value names under the given key.


`CreateKey`: Creates a new registry key or opens an existing one.


`ReadStringValueWithDefault`: Reads a string value from the Windows Registry with a default value.


`ReadMultiStringValue`: Reads a multi-string value from the Windows Registry.


`WriteMultiStringValue`: Writes a multi-string value to the Windows Registry.


`ReadExpandStringValue`: Reads an expandable string value (REG_EXPAND_SZ) from the Windows Registry.


`WriteExpandStringValue`: Writes an expandable string value (REG_EXPAND_SZ) to the Windows Registry.


`ReadInt32Value`: Reads a 32-bit integer value from the Windows Registry.


`WriteInt32Value`: Writes a 32-bit integer value to the Windows Registry.


`ReadInt64Value`: Reads a 64-bit integer value from the Windows Registry.


`WriteInt64Value`: Writes a 64-bit integer value to the Windows Registry.


`DeleteKey`: Deletes a registry key and all its subkeys and values.

## Contributions
Contributions and improvements to this package are welcome. Feel free to submit pull requests, report issues, or provide feedback.

