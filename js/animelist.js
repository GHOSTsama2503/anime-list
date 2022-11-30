function createList(animeDB) {
    let table = document.getElementById("animeList")
    let tbody = document.createElement("tbody")

    let listLen = animeDB.length

    for (let i = 0; i < listLen; i++) {
        element = animeDB[i]
        elementId = "a" + element["id"]

        let row = document.createElement("tr")

        // status
        let status = document.createElement("td")

        if (element["status"] == true) {
            status.setAttribute("class", "statusTrue")
        }
        else {
            status.setAttribute("class", "statusFalse")
        }

        row.appendChild(status)

        // name
        let name = document.createElement("td")
        name.appendChild(document.createTextNode(element["name"]))
        name.setAttribute("id", elementId)
        name.setAttribute("onclick", `copyName(${elementId})`)
        row.appendChild(name)

        // info
        let info = document.createElement("td")
        let infoLink = document.createElement("a")

        infoLink.href = "https://anilist.co/anime/" + elementId
        infoLink.appendChild(document.createTextNode("Info"))

        info.appendChild(infoLink)
        row.appendChild(info)

        // episodes
        let episodes = document.createElement("td")
        episodes.appendChild(document.createTextNode(element["episodes"]))
        row.appendChild(episodes)

        // format
        let format = document.createElement("td")
        format.appendChild(document.createTextNode(element["format"]))
        row.appendChild(format)

        // row color
        if (i % 2 == 0) {
            row.setAttribute("class", "evenRow")
        }
        else {
            row.setAttribute("class", "oddRow")
        }

        tbody.appendChild(row)
    }

    // insert row
    table.appendChild(tbody)
}


function getAnimeDB(callback) {
    fetch("db.json")
        .then(response => response.text())
        .then(content => {
            callback(JSON.parse(content))
    })
}

getAnimeDB(createList)