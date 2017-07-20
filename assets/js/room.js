(function (window, document, _, undefined) {
	"use strict";

	new Vue({
		el: "#main",
		template: `
		<div class="container-fluid">
			<div class="row">
				<div class="col">
					<nav class="navbar bg-faded">
						<h1 class="navbar-brand">Room {{ name }}</h1>
					</nav>
				</div>
			</div>

			<form>
				<div class="form-group">
					<label for="story-desc"></label>
					<textarea id="story-desc" v-model="storyDesc"></textarea>
				</div>
				
				<div class="form-group">
					<div class="input-group">
						<div v-for="v in availableVotes" class="input-group-btn">
							<button class="btn btn-outline-primary" :class="{ active: v == vote }" type="button" @click.prevent="setVote(v)" style="margin-left: 0;">{{ v }}</button>
      			</div>
				</div>
			</form>
		</div>
    `,

		props: {
			availableVotes: {
				default: [0, 1, 2, 3, 5, 8, 13, 21, '?']
			}
		},

		data: () => ({
			name: null,
			socket: null,

			storyDesc: '',
			vote: null,
			votes: {},
		}),

		watch: {
			storyDesc: _.debounce(function() {
				this.sendStoryDesc()
			}, 500)
		},

		mounted() {
			this.name = window.app.config.roomName;
			this.socketConnect();
		},

		methods: {
			onSocketClose() {
				console.log('socket closed', arguments);
				this.socket = null;
			},

			onSocketMsg(msg) {
				console.log('socket message', msg);
			},

			socketConnect() {
				let wsUrl = window.location.hostname + ':' + window.location.port + '/msg/' + this.name;
				let socket = new WebSocket('ws://' + wsUrl);

				socket.onclose = this.onSocketClose;
				socket.onmessage = this.onSocketMsg;

				this.socket = socket;
				console.log('socket opened', this.socket);
			},

			socketSend(msg) {
				if (this.socket == null) {
					console.log('socket was closed');
					this.socketConnect();
				}
				console.log('sending', msg);
				this.socket.send(JSON.stringify(msg));
			},

			sendStoryDesc() {
				let msg = {
					cmd: 'storyDesc',
					storyDesc: this.storyDesc
				};

				this.socketSend(msg);
			},

			sendVote() {
				let msg = {
					cmd: 'vote',
					vote: this.vote
				};

				this.socketSend(msg);
			},

			setVote(v) {
				this.vote = v;
				this.sendVote();
			}

		}
	});
})(window, document, _, undefined);
