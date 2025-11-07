package main

import (
	"encoding/json"
	"os"
)

func savePathJSON(p Path, filename string) error {
	// convert []byte{'R','D','L'} -> []string{"R","D","L"}
	strMoves := make([]string, len(p))
	for i, mv := range p {
		strMoves[i] = string([]byte{mv})
	}

	out := map[string]interface{}{
		"path":      strMoves,
		"start":     startPoint,
		"end":       endPoint,
		"maze_rows": len(maze),
		"maze_cols": len(maze[0]),
	}

	bytes, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, bytes, 0644)
}
