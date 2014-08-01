package error

import (
    "strings"
)

// cast error
type CastError struct {
    Type, Value, Path string
}

func (c *CastError) Error() string {
    return fmt.Sprintf("CastError: Cast to %s failed for value %s at path %s", c.Type, c.Value, c.Path)
}

// divergent array error
type DivergentArrayError struct {
    Paths []string
}

func (d *DivergentArrayError) Error() string {
    msg := "DivergentArrayError: For your own good, using `document.save()` to update an array "
    msg += "which was selected using an $elemMatch projection OR "
    msg += "populated using skip, limit, query conditions, or exclusion of "
    msg += "the _id field when the operation results in a $pop or $set of "
    msg += "the entire array is not supported. The following "
    msg += "path(s) would have been modified unsafely:\n"
    msg += " %s \n"
    msg += "Use Model.update() to update these arrays instead."

    return fmt.Sprintf(msg, strings.Join(d.Paths, "\n "))
}

// version error
type VersionError struct {
}

func (v *VersionError) Error() string {
    return fmt.Sprintf("VersionError: No matching document found.")
}

// missing schema error
type MissingSchemaError struct {
    Name string
}

func (m *MissingSchemaError) Error() string {
    msg := "Schema hasn't been registered for model %s.\n"
    msg += "Use mongoose.model(name, schema)";
    return fmt.Sprintf(msg, m.Name)
}

// overwrite model error
type OverwriteModelError struct {
    Name string
}

func (o *OverwriteModelError) Error() string {
    return fmt.Sprintf("OverwriteModelError: Cannot overwrite %s model once compiled.", o.Name)
}
