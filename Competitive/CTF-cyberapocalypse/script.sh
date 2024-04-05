#!/bin/bash

# Connect to the game server using nc
nc 94.237.54.245 50787 << EOF
y

# Function to send the appropriate response based on the scenario
send_response() {
    local response=""
    for scenario in "$@"; do
        case $scenario in
            GORGE)
                response+="STOP-"
                ;;
            PHREAK)
                response+="DROP-"
                ;;
            FIRE)
                response+="ROLL-"
                ;;
            *)
                echo "UNKNOWN SCENARIO"
                ;;
        esac
    done
    # Remove the trailing dash, if any
    response=${response%-}
    echo "$response"
}

# Read the game prompts and send responses
while IFS= read -r line; do
    # Check if the line is the last line
    if [[ $line == "What do you do?"* ]]; then
        # Extract the scenarios from the previous line
        scenarios=$(echo "$previous_line" | cut -d ' ' -f2-)
        IFS=', ' read -r -a scenario_array <<< "$scenarios"
        send_response "${scenario_array[@]}"
    fi
    # Store the current line as the previous line for the next iteration
    previous_line="$line"
done

EOF
