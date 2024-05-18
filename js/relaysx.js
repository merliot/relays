class DeviceRelays extends DeviceBase {

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
			if (relay.Name === "" || relay.Gpio === "") {
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
	}

	setMouse(i) {
		if (this.view === 0) {
			let div = document.getElementById("relay" + i)
			div.onclick = () => {
				if (this.state.Online) {
					this.relayClick(i)
				}
			}
		}
	}

	setRelayImg(relay, i) {
		let image = document.getElementById("relay" + i + "-img")
		if (relay.State) {
			image.src = this.assets + "images/relay-on.png"
		} else {
			image.src = this.assets + "images/relay-off.png"
		}
	}

	setGpio(relay, i) {
		if (this.view === 0) {
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
		this.send("click", {Relay: index, State: relay.State})
	}
}
