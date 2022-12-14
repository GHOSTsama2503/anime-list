function createList(animeDB) {
    let table = document.getElementById("animeList")
    let tbody = document.createElement("tbody")

    let listLen = animeDB.length
    let statusFalse = 0

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
            statusFalse += 1
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

        infoLink.href = "https://anilist.co/anime/" + elementId.split("a")[1]
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

    // add footer info
    footer = document.getElementsByTagName("footer")[0]

    total = document.createElement("div")
    total.setAttribute("class", "fItem")
    total.appendChild(document.createTextNode("Total: " + listLen))
    footer.appendChild(total)

    completed = document.createElement("div")
    completed.setAttribute("class", "fItem")
    completed.appendChild(document.createTextNode("Completed: " + (listLen - statusFalse)))
    footer.appendChild(completed)

    incompleted = document.createElement("div")
    incompleted.setAttribute("class", "fItem")
    incompleted.appendChild(document.createTextNode("Incompleted: " + statusFalse))
    footer.appendChild(incompleted)

}


function getAnimeDB(callback) {
    fetch("db.json")
        .then(response => response.text())
        .then(content => {
            callback(JSON.parse(content))
    })
}

getAnimeDB(createList)