az account list-locations | jq '[
  .[] |
  select(.metadata.regionType == "Physical") |
  {
    name,
    displayName,
    id,
    regionalDisplayName,
    type,
    shortName: (.displayName | split(" ") | map(.[0:1]) | join("") | ascii_downcase),
    regionCategory: (.metadata.regionCategory),
    pairedRegionNames: (.metadata.pairedRegion // [] | map(.name) | join(","))
  }
]' > resourcelocations.json