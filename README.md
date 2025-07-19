# md-merger

md-merger is a tool for merging markdown files.


## Examples

Merge markdown files. This will create a root_merged.md inside the root directory. 
Relative markdown links will be updated.

**root.md**
```md
# Root
The following comment is used to import content.
<!-- merge:partials/content.md -->
```

**partials/content.md**
```md
# Content
Lorem Ipsum Dolor
![My Image](.my-image.png)
```

By running the following command the merged file will be created.

```bash
md-merger --input root.md
```

The result will be the following markdown

**root_merged.md**
```md
# Root
The following comment is used to import content.
# Content
Lorem Ipsum Dolor
![My Image](partials/my-image.png)
```

## Development

**Compiling to executbale (Linux)**
```bash
GOOS=linux GOARCH=amd64 go build -o build/linux/md-merger cmd/main.go
```

**Compiling to executbale (Windows)**
```bash
GOOS=windows GOARCH=amd64 go build -o build/win/md-merger.exe cmd/main.go
```