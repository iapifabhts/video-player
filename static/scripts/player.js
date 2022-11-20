const player = document.querySelector(".player")
const hls = new Hls()

const play = path => {
	hls.loadSource(`http://localhost:80/files${path}`)
	hls.attachMedia(player)
}
