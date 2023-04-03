package types

const (
	QrInsertLog = `INSERT INTO logging (
									request_payload, response_payload, url
								) VALUES (
									:requestPayload, :responsePayload, :url
								)`
)
