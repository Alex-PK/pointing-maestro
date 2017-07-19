let app = new Vue({
	el: "#main",
	template: `
		<div class="container-fluid">
			<div class="row">
				<div class="col">
					<nav class="navbar bg-faded">
						<h1 class="navbar-brand">Home</h1>
					</nav>
				</div>
			</div>

			<div class="row">
				<div class="col">
					<form @submit.prevent="openRoom">
						<div class="form-group">
							<label for="roomname">Create room:</label>
							<input class="form-control" id="roomname" type="text" name="room" v-model="name" @keyup="checkName" />
							<small class="form-text text-danger">{{ error }}</small>
						</div>
						
						<button class="btn btn-primary" type="submit" :disabled="inError">Create</button>
					</form>
				</div>
			</div>
		</div>
    `,

	data: () => ({
		name: '',
		error: ''
	}),

	computed: {
		inError() {
			return this.error.length > 0;
		}
	},

	methods: {
		checkName() {
			if (!/^[a-z0-9_-]+$/.test(this.name)) {
				this.error = 'Please provide a valid room name';
				return false;
			}
			this.error = '';
			return true;
		},

		openRoom() {
			if (!this.checkName()) {
				return;
			}
			window.location.href = '/room/' + this.name;
		}
	}
});
