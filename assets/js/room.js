let app = new Vue({
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

			<div class="row">
				<div class="col">
				</div>
			</div>
		</div>
    `,

	props: [
		'name'
	],

	data: () => ({
		socket: null
	}),

	computed: {},

	mounted() {

	},

	methods: {}
});
