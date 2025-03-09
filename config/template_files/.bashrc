# Vine stuff
vine_root_dir="$HOME/.vine"
vine_binary=""$vine_root_dir"/vine_aly"
vine_execution_var="VINE_EXECUTE"

execution_dir=$(pwd)

function vine() {
    eval "export $vine_execution_var=bash"

    trap 'unset '"$vine_execution_var" RETURN

    if [[ "$1" == "-d" ]]; then
        output=$("$vine_binary" alias "$2")
        _cd_back
        eval "$output"
    elif [[ "$1" == "-e" ]]; then
        output=$("$vine_binary" alias "$2")
        trap '_cd_back' INT
        cd "$vine_root_dir"
        "$output"
        _cd_back
        trap - INT
    else
        echo "Usage: vine -d/-e <drive_letter/script>"
    fi
}

function _cd_back() {
    cd "$execution_dir"
    # if [[ -n "$OLDPWD" ]]; then
    #     cd - > /dev/null
    # fi
}
# ffuts eniV
