(function (window, document, _, undefined) {
	"use strict";

	new Vue({
		el: "#main",
		template: `
		<div class="container-fluid">
			<div class="row">
				<div class="col">
					<nav class="navbar bg-faded">
						<h1 class="navbar-brand">Room {{ roomName }}</h1>
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
				
				<div class="form-group">
					<button class="btn btn-outline-info" type="button" @click.prevent="sendShowVotes">Show votes</button>
					<button class="btn btn-outline-danger" type="button" @click.prevent="sendClearVotes">Clear votes</button>
			</form>
		</div>
    `,

		props: {
			availableVotes: {
				default: [0, 1, 2, 3, 5, 8, 13, 21, '?']
			}
		},

		data: () => ({
			roomName: null,
			socket: null,

			userName: 'dummy',

			storyDesc: '',
			vote: '',
			votes: {},
		}),

		watch: {
			storyDesc: _.debounce(function() {
				this.sendStoryDesc()
			}, 500)
		},

		mounted() {
			this.roomName = window.app.config.roomName;
			this.socketConnect();
		},

		methods: {
			onSocketClose() {
				console.log('socket closed', arguments);
				this.socket = null;
			},

			onSocketMsg(msg) {
				console.log('socket message', msg.data);
			},

			socketConnect() {
				let wsUrl = window.location.hostname + ':' + window.location.port + '/msg/' + this.roomName;
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
					user: this.userName,
					vote: this.vote
				};

				this.socketSend(msg);
			},

			setVote(v) {
				this.vote = '' + v;
				this.sendVote();
			},

			sendShowVotes() {
				let msg = {
					cmd: 'showVotes'
				};
				this.socketSend(msg);
			},

			sendClearVotes() {
				let msg = {
					cmd: 'clearVotes'
				};
				this.socketSend(msg);
			},

		}
	});
})(window, document, _, undefined);
