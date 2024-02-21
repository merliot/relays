import { WebSocketController, ViewMode } from './common.js'

export function run(prefix, url, viewMode) {
	const relays = new Relays(prefix, url, viewMode)
}

class Relays extends WebSocketController {

	open() {
		super.open()
		if (this.state.DeployParams !== "") {
			this.showRelays()
		}
	}

	handle(msg) {
		switch(msg.Path) {
		case "click":
			this.saveClick(msg)
			break
		}
	}

	showRelays() {
		var havesome = false
		for (let i = 0; i < 4; i++) {
			let div = document.getElementById("relay" + i)
			var relay = this.state.Relays[i]
			if (relay.Gpio === "") {
				div.classList.replace("visibleFlex", "hidden")
			} else {
				this.setMouse(i)
				this.setRelayImg(relay, i)
				this.setGpio(relay, i)
				this.setRelayName(relay, i)
				div.classList.replace("hidden", "visibleFlex")
				havesome = true
			}
		}
		if (!havesome) {
			document.getElementById("nodef").classList.replace("hidden", "visible")
		}
	}

	setMouse(i) {
		if (this.viewMode === ViewMode.ViewFull) {
			let div = document.getElementById("relay" + i)
			div.onclick = () => {
				this.relayClick(i)
			}
		}
	}

	setRelayImg(relay, i) {
		let image = document.getElementById("relay" + i + "-img")
		if (relay.State) {
			image.src = "images/relay-on.png"
		} else {
			image.src = "images/relay-off.png"
		}
	}

	setGpio(relay, i) {
		if (this.viewMode === ViewMode.ViewFull) {
			let gpio = document.getElementById("gpio" + i)
			gpio.textContent = relay.Gpio
		}
	}

	setRelayName(relay, i) {
		let label = document.getElementById("relay" + i + "-name")
		label.textContent = relay.Name
	}

	saveClick(msg) {
		var relay = this.state.Relays[msg.Relay]
		relay.State = msg.State
		this.setRelayImg(relay, msg.Relay)
	}

	relayClick(index) {
		var relay = this.state.Relays[index]
		relay.State = !relay.State
		this.setRelayImg(relay, index)
		this.webSocket.send(JSON.stringify({Path: "click", Relay: index, State: relay.State}))
	}
}
