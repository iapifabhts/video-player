const itemsNode = document.querySelector(".items")

itemsNode.onclick = event => {
	if (event.target.classList.contains("items__image")) {
		play(event.target.dataset.video)
	}
}

const getItems = () =>
	axios.get("http://localhost:80/items")
		.then(({ data }) => drawItems(data))

getItems()

const drawItems = items => {
	itemsNode.innerHTML = null
	for (const item of items) {
		const itemNode = document.createElement("div")
		const itemImg = document.createElement("img")
		const iconNode = document.createElement("span")
		itemNode.classList.add("items__item")
		itemImg.src = `http://localhost:80/files${item.poster}`
		itemImg.dataset.video = item.video
		itemImg.classList.add("items__image")
		iconNode.classList.add("material-symbols-outlined")
		iconNode.innerText = "play_arrow"
		itemNode.append(itemImg)
		itemNode.append(iconNode)
		itemsNode.append(itemNode)
	}
}
