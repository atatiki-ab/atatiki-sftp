# sftp

Upload to a sftp server

	`
	curl -X POST \
	  http://localhost:8080/upload \
	  -H 'Content-Type: application/json' \
	  -d '{
		"pathAndFilename": "./a.txt",
		"content": "Test with åäö ÅÄÖ\nand some new lines\nanother one ;)\n",
		"addWindowsLineEndings": true,
		"convertTo8859": true
	}'`

Download from an sftp server

	`
	curl -X POST \
	  http://localhost:8080/download \
	  -H 'Content-Type: application/json' \
	  -d '{
		"pathAndFilename": "./a.txt",
		"removeWindowsLineEndings": true,
		"convertFrom8859": true
	}'`

Delete a file from an sftp server

	`
	curl -X POST \
	  http://localhost:8080/delete \
	  -H 'Content-Type: application/json' \
	  -d '{
		"pathAndFilename": "./a.txt"
	}'`