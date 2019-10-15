curl -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq --argjson VAR "$HERO_ID" '.[] | select (.id==$VAR) | .connections.relatives' | tr -d '""'
