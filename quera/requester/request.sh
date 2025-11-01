#!/bin/bash
set -e

# === Configuration ===
DIR="/home/manouchehrrasouli/Downloads/tabriz/tabriz/jsons"  # directory containing .json files
OUTPUT_FILE="./update_queries.sql"                           # output SQL file
ENDPOINT="http://localhost:1552/radar/backoffice/polygon/doctor/express"
AUTH="Bearer eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoiZXlKMWMyVnlTV1FpT2pFd01ESTVMQ0oxYzJWeWJtRnRaU0k2SW0wdWNtRnpiM1ZzYVNKOSIsImV4cCI6MTc2MTk4NjYyMH0.UZx1keYD6mBlkiA-HPLBFBTuzzMg5Dm2ZxPHpqDP9_WGuvLEEZspwX6fBdOXW97s-sJ0V-zItw2-C9sFg3VlCQ"

# === Initialize output ===
echo "-- Auto-generated update queries" > "$OUTPUT_FILE"
echo "-- $(date)" >> "$OUTPUT_FILE"
echo "" >> "$OUTPUT_FILE"

# === Track failures ===
failed_vendors=()

# === Loop through JSON files ===
for file in "$DIR"/*.json; do
  [ -e "$file" ] || continue  # skip if no files

  filename=$(basename "$file")
  vendor_id="${filename%.json}"  # remove extension (e.g. 56540.json -> 56540)

  echo "Processing vendor_id=$vendor_id ..."

  # Upload the JSON file
  if ! curl -s --location "$ENDPOINT" \
      --header "Authorization: $AUTH" \
      --form "file=@\"$file\"" \
      --fail; then
    echo "❌ Failed to upload $file"
    failed_vendors+=("$vendor_id")
    continue
  fi

  # Generate the SQL update line
  echo "update delivery_methods set pricing_plan_id = 374, resolution_type = 'polygon' where vendor_id = $vendor_id and type = 'express';" >> "$OUTPUT_FILE"
done

echo ""
echo "✅ Done. SQL queries written to: $OUTPUT_FILE"

# === Print failed vendors ===
if [ ${#failed_vendors[@]} -gt 0 ]; then
  echo ""
  echo "❌ Upload failed for the following vendors:"
  printf '(%s)\n' "$(IFS=,; echo "${failed_vendors[*]}")"
else
  echo "🎉 All uploads succeeded!"
fi
