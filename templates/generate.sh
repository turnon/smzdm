#!/bin/sh

ls templates/*.go | grep -v templates.go | xargs rm 2>/dev/null

ls templates/*.html |
    while read f; do
        name=$(echo $f | sed 's/\.html//')
        cat <<EOF > "$name.go"
package templates

func init() {
	Templates["${name}"] = \`$(cat $f)\`
}
EOF
    done
