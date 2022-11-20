const dialog = document.querySelector(".dialog")
const form = document.querySelector(".dialog__form")
const posterInput = document.querySelector(".form__input-poster input")
const videoInput = document.querySelector(".form__input-video input")
const button = document.querySelector(".form__button")
const headerButton = document.querySelector(".header__button")

headerButton.onclick = () => dialog.style.display = "flex"

form.onsubmit = event => event.preventDefault()

dialog.onclick = event =>
	event.target.classList.contains("dialog") ?
		event.target.style.display = "none" : null

button.onclick = () => {
	let posterId, videoId
	uploadFile(posterInput.files[0])
		.then(({ data }) => {
			posterId = data
			uploadFile(videoInput.files[0])
				.then(({ data }) => {
					videoId = data
					uploadItem(posterId, videoId)
				})
		})
}

const uploadFile = file => {
	const formData = new FormData()
	formData.append("file", file)
	return axios.post("http://localhost:80/files", formData)
}

const uploadItem = (poster, video) => {
	axios.post("http://localhost:80/items", { poster, video })
		.then(() => {
			getItems()
			dialog.style.display = "none"
		})
}
