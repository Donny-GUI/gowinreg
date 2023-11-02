package winreg

import (
    "fmt"
    "golang.org/x/sys/windows/registry"
)


// ReadDWordValue reads a DWORD value from the Windows Registry.
func ReadDWordValue(root registry.Key, keyPath, valueName string) (uint32, error) {
    k, err := registry.OpenKey(root, keyPath, registry.QUERY_VALUE)
    if err != nil {
        return 0, err
    }
    defer k.Close()

    value, _, err := k.GetDWordValue(valueName)
    if err != nil {
        return 0, err
    }

    return value, nil
}

// WriteDWordValue writes a DWORD value to the Windows Registry.
func WriteDWordValue(root registry.Key, keyPath, valueName string, data uint32) error {
    k, err := registry.OpenKey(root, keyPath, registry.WRITE)
    if err != nil {
        return err
    }
    defer k.Close()

    if err := k.SetDWordValue(valueName, data); err != nil {
        return err
    }

    return nil
}

// ReadBinaryValue reads a binary value from the Windows Registry.
func ReadBinaryValue(root registry.Key, keyPath, valueName string) ([]byte, error) {
    k, err := registry.OpenKey(root, keyPath, registry.QUERY_VALUE)
    if err != nil {
        return nil, err
    }
    defer k.Close()

    value, _, err := k.GetBinaryValue(valueName)
    if err != nil {
        return nil, err
    }

    return value, nil
}

// WriteBinaryValue writes a binary value to the Windows Registry.
func WriteBinaryValue(root registry.Key, keyPath, valueName string, data []byte) error {
    k, err := registry.OpenKey(root, keyPath, registry.WRITE)
    if err != nil {
        return err
    }
    defer k.Close()

    if err := k.SetBinaryValue(valueName, data); err != nil {
        return err
    }

    return nil
}

// DeleteValue deletes a registry value.
func DeleteValue(root registry.Key, keyPath, valueName string) error {
    k, err := registry.OpenKey(root, keyPath, registry.WRITE)
    if err != nil {
        return err
    }
    defer k.Close()

    if err := k.DeleteValue(valueName); err != nil {
        return err
    }

    return nil
}

// DeleteSubKey deletes a registry subkey and all its subkeys and values.
func DeleteSubKey(root registry.Key, keyPath, subKeyName string) error {
    k, err := registry.OpenKey(root, keyPath, registry.WRITE)
    if err != nil {
        return err
    }
    defer k.Close()

    if err := k.DeleteKey(subKeyName); err != nil {
        return err
    }

    return nil
}

// Check if a registry key exists.
func KeyExists(root registry.Key, keyPath string) bool {
    k, err := registry.OpenKey(root, keyPath, registry.QUERY_VALUE)
    if err != nil {
        return false
    }
    k.Close()
    return true
}

// Check if a registry value exists.
func ValueExists(root registry.Key, keyPath, valueName string) bool {
    k, err := registry.OpenKey(root, keyPath, registry.QUERY_VALUE)
    if err != nil {
        return false
    }
    defer k.Close()

    _, _, err = k.GetStringValue(valueName)
    return err == nil
}

// EnumerateSubKeys returns a list of subkeys under the given key.
func EnumerateSubKeys(root registry.Key, keyPath string) ([]string, error) {
    k, err := registry.OpenKey(root, keyPath, registry.ENUMERATE_SUB_KEYS)
    if err != nil {
        return nil, err
    }
    defer k.Close()

    subKeys, err := k.ReadSubKeyNames(-1)
    if err != nil {
        return nil, err
    }

    return subKeys, nil
}

// EnumerateValues returns a list of value names under the given key.
func EnumerateValues(root registry.Key, keyPath string) ([]string, error) {
    k, err := registry.OpenKey(root, keyPath, registry.QUERY_VALUE)
    if err != nil {
        return nil, err
    }
    defer k.Close()

    valueNames, err := k.ReadValueNames(-1)
    if err != nil {
        return nil, err
    }

    return valueNames, nil
}

// CreateKey creates a new registry key or opens an existing one.
func CreateKey(root registry.Key, keyPath string) (registry.Key, error) {
    k, _, err := registry.CreateKey(root, keyPath, registry.ALL_ACCESS)
    return k, err
}

// ReadStringValueWithDefault reads a string value from the Windows Registry with a default value.
func ReadStringValueWithDefault(root registry.Key, keyPath, valueName, defaultValue string) (string, error) {
    k, err := registry.OpenKey(root, keyPath, registry.QUERY_VALUE)
    if err != nil {
        return defaultValue, nil // Return the default value if the key or value doesn't exist
    }
    defer k.Close()

    value, _, err := k.GetStringValue(valueName)
    if err != nil {
        return defaultValue, nil // Return the default value if the value doesn't exist
    }

    return value, nil
}

// ReadMultiStringValue reads a multi-string value from the Windows Registry.
func ReadMultiStringValue(root registry.Key, keyPath, valueName string) ([]string, error) {
    k, err := registry.OpenKey(root, keyPath, registry.QUERY_VALUE)
    if err != nil {
        return nil, err
    }
    defer k.Close()

    value, _, err := k.GetStringsValue(valueName)
    if err != nil {
        return nil, err
    }

    return value, nil
}

// WriteMultiStringValue writes a multi-string value to the Windows Registry.
func WriteMultiStringValue(root registry.Key, keyPath, valueName string, data []string) error {
    k, err := registry.OpenKey(root, keyPath, registry.WRITE)
    if err != nil {
        return err
    }
    defer k.Close()

    if err := k.SetStringsValue(valueName, data); err != nil {
        return err
    }

    return nil
}

// ReadQWordValue reads a QWORD (64-bit integer) value from the Windows Registry.
func ReadQWordValue(root registry.Key, keyPath, valueName string) (uint64, error) {
    k, err := registry.OpenKey(root, keyPath, registry.QUERY_VALUE)
    if err != nil {
        return 0, err
    }
    defer k.Close()

    value, _, err := k.GetQWordValue(valueName)
    if err != nil {
        return 0, err
    }

    return value, nil
}

// WriteQWordValue writes a QWORD (64-bit integer) value to the Windows Registry.
func WriteQWordValue(root registry.Key, keyPath, valueName string, data uint64) error {
    k, err := registry.OpenKey(root, keyPath, registry.WRITE)
    if err != nil {
        return err
    }
    defer k.Close()

    if err := k.SetQWordValue(valueName, data); err != nil {
        return err
    }

    return nil
}

// ReadExpandStringValue reads an expandable string value (REG_EXPAND_SZ) from the Windows Registry.
func ReadExpandStringValue(root registry.Key, keyPath, valueName string) (string, error) {
    k, err := registry.OpenKey(root, keyPath, registry.QUERY_VALUE)
    if err != nil {
        return "", err
    }
    defer k.Close()

    value, _, err := k.GetExpandStringValue(valueName)
    if err != nil {
        return "", err
    }

    return value, nil
}

// WriteExpandStringValue writes an expandable string value (REG_EXPAND_SZ) to the Windows Registry.
func WriteExpandStringValue(root registry.Key, keyPath, valueName, data string) error {
    k, err := registry.OpenKey(root, keyPath, registry.WRITE)
    if err != nil {
        return err
    }
    defer k.Close()

    if err := k.SetExpandStringValue(valueName, data); err != nil {
        return err
    }

    return nil
}

// ReadInt32Value reads a 32-bit integer value from the Windows Registry.
func ReadInt32Value(root registry.Key, keyPath, valueName string) (int32, error) {
    k, err := registry.OpenKey(root, keyPath, registry.QUERY_VALUE)
    if err != nil {
        return 0, err
    }
    defer k.Close()

    value, _, err := k.GetIntValue(valueName)
    if err != nil {
        return 0, err
    }

    return value, nil
}

// WriteInt32Value writes a 32-bit integer value to the Windows Registry.
func WriteInt32Value(root registry.Key, keyPath, valueName string, data int32) error {
    k, err := registry.OpenKey(root, keyPath, registry.WRITE)
    if err != nil {
        return err
    }
    defer k.Close()

    if err := k.SetIntValue(valueName, data); err != nil {
        return err
    }

    return nil
}

// ReadInt64Value reads a 64-bit integer value from the Windows Registry.
func ReadInt64Value(root registry.Key, keyPath, valueName string) (int64, error) {
    k, err := registry.OpenKey(root, keyPath, registry.QUERY_VALUE)
    if err != nil {
        return 0, err
    }
    defer k.Close()

    value, _, err := k.GetIntValue(valueName)
    if err != nil {
        return 0, err
    }

    return int64(value), nil
}

// WriteInt64Value writes a 64-bit integer value to the Windows Registry.
func WriteInt64Value(root registry.Key, keyPath, valueName string, data int64) error {
    k, err := registry.OpenKey(root, keyPath, registry.WRITE)
    if err != nil {
        return err
    }
    defer k.Close()

    if err := k.SetIntValue(valueName, int32(data)); err != nil {
        return err
    }

    return nil
}

// DeleteKey deletes a registry key and all its subkeys and values.
func DeleteKey(root registry.Key, keyPath string) error {
    return registry.DeleteKey(root, keyPath)
}
