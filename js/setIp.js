fetch("/api/ip", { headers: { 'accept': 'application/json' } })
    .then(response => response.json())
    .then(data => {
        document.getElementById("ip").innerHTML = data["ip"]
    }).catch(err => {
        msg = "There was a problem getting your IP address"
        document.getElementById("ip").innerHTML = msg
        console.error(err)
    })
