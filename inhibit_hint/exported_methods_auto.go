// Code generated by "dbusutil-gen em -type Object"; DO NOT EDIT.

package inhibit_hint

import (
	"pkg.deepin.io/lib/dbusutil"
)

func (v *Object) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name:    "Get",
			Fn:      v.Get,
			InArgs:  []string{"locale", "why"},
			OutArgs: []string{"hint"},
		},
	}
}