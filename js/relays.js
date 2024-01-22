import { WebSocketController } from './common.js'

export function run(prefix, url) {
	const relays = new Relays(prefix, url)
}

class Relays extends WebSocketController {

	open() {
		super.open()
		this.showRelays()
	}

	handle(msg) {
		switch(msg.Path) {
		case "click":
			this.saveClick(msg)
			break
		}
	}

	showRelays() {

		let nodef = document.getElementById("nodef")
		nodef.classList.replace("visible", "hidden")

		var undef = true
		for (let i = 0; i < 4; i++) {
			let div = document.getElementById("relay" + i)
			let label = document.getElementById("relay" + i + "-name")
			let image = document.getElementById("relay" + i + "-img")
			let gpio = document.getElementById("gpio" + i)
			var relay = this.state.Relays[i]
			if (relay.Gpio === "") {
				div.classList.replace("visibleFlex", "hidden")
			} else {
				undef = false
				gpio.textContent = relay.Gpio
				label.textContent = relay.Name
				div.classList.replace("hidden", "visibleFlex")
				div.onclick = () => {
					this.relayClick(image, i)
				}
				this.setRelayImg(relay, image)
			}
		}

		if (undef) {
			nodef.classList.replace("hidden", "visible")
		}
	}

	setRelayImg(relay, image) {
		image.disabled = false
		if (relay.State) {
			image.src = "images/relay-on.png"
		} else {
			image.src = "images/relay-off.png"
		}
	}

	saveClick(msg) {
		var image = document.getElementById("relay" + msg.Relay + "-img")
		var relay = this.state.Relays[msg.Relay]
		relay.State = msg.State
		this.setRelayImg(relay, image)
	}

	relayClick(image, index) {
		var relay = this.state.Relays[index]
		relay.State = !relay.State
		this.setRelayImg(relay, image)
		this.webSocket.send(JSON.stringify({Path: "click", Relay: index, State: relay.State}))
	}
}
