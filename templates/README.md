# Templates

Templates for missions can go in this folder (or even subfolders) - just be sure to link them in the mBot/mission/templateMap.go file!

## Example Template

Each template should be stored as their own JSON file. The content of a template file should be similar to the following:

```json
{"structuredResponse":"n/a","introduction":"The target is unreachable.","testing_methodology":"The target is unreachable.","conclusion":"The target is unreachable."}
```

## How does it work?

The above content is already stored as a template file - `mBot/templates/template.json` - and linked to the map in `mBot/mission/templateMap.go`, as shown below:

```go
var missionMap = map[string]string{
	"Example Mission Title": "templates/template.json",
}
```

This means if a mission were claimed with the title `Example Mission Title`, each field would be auto-populated with the data in the template file.
